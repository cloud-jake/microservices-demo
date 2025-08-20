// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	"cloud.google.com/go/alloydbconn"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	retail "cloud.google.com/go/retail/apiv2"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	pb "github.com/GoogleCloudPlatform/microservices-demo/src/newproductcatalogservice/genproto"
	"github.com/golang/protobuf/jsonpb"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/api/iterator"
	retailpb "google.golang.org/genproto/googleapis/cloud/retail/v2"
)

func loadCatalog(catalog *pb.ListProductsResponse) error {
	catalogMutex.Lock()
	defer catalogMutex.Unlock()

	// Attempt to load from an external source if configured
	if os.Getenv("ALLOYDB_CLUSTER_NAME") != "" {
		err := loadCatalogFromAlloyDB(catalog)
		if err == nil {
			return nil
		}
		log.Warnf("failed to load catalog from AlloyDB: %v. Falling back to local file.", err)
	} else if os.Getenv("VERTEX_AI_SEARCH_ENABLED") == "true" {
		err := loadCatalogFromVertex(catalog)
		if err == nil {
			return nil
		}
		log.Warnf("failed to load catalog from Vertex AI Search: %v. Falling back to local file.", err)
	}

	return loadCatalogFromLocalFile(catalog)
}

func loadCatalogFromLocalFile(catalog *pb.ListProductsResponse) error {
	log.Info("loading catalog from local products.json file...")

	catalogJSON, err := os.ReadFile("products.json")
	if err != nil {
		log.Warnf("failed to open product catalog json file: %v", err)
		return err
	}

	if err := jsonpb.Unmarshal(bytes.NewReader(catalogJSON), catalog); err != nil {
		log.Warnf("failed to parse the catalog JSON: %v", err)
		return err
	}

	log.Info("successfully parsed product catalog json")
	return nil
}

func getSecretPayload(project, secret, version string) (string, error) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Warnf("failed to create SecretManager client: %v", err)
		return "", err
	}
	defer client.Close()

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/%s", project, secret, version),
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Warnf("failed to access SecretVersion: %v", err)
		return "", err
	}

	return string(result.Payload.Data), nil
}

func loadCatalogFromAlloyDB(catalog *pb.ListProductsResponse) error {
	log.Info("loading catalog from AlloyDB...")

	projectID := os.Getenv("PROJECT_ID")
	region := os.Getenv("REGION")
	pgClusterName := os.Getenv("ALLOYDB_CLUSTER_NAME")
	pgInstanceName := os.Getenv("ALLOYDB_INSTANCE_NAME")
	pgDatabaseName := os.Getenv("ALLOYDB_DATABASE_NAME")
	pgTableName := os.Getenv("ALLOYDB_TABLE_NAME")
	pgSecretName := os.Getenv("ALLOYDB_SECRET_NAME")

	pgPassword, err := getSecretPayload(projectID, pgSecretName, "latest")
	if err != nil {
		return err
	}

	dialer, err := alloydbconn.NewDialer(context.Background())
	if err != nil {
		log.Warnf("failed to set-up dialer connection: %v", err)
		return err
	}
	cleanup := func() error { return dialer.Close() }
	defer cleanup()

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		"postgres", pgPassword, pgDatabaseName,
	)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Warnf("failed to parse DSN config: %v", err)
		return err
	}

	pgInstanceURI := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/instances/%s", projectID, region, pgClusterName, pgInstanceName)
	config.ConnConfig.DialFunc = func(ctx context.Context, _ string, _ string) (net.Conn, error) {
		return dialer.Dial(ctx, pgInstanceURI)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Warnf("failed to set-up pgx pool: %v", err)
		return err
	}
	defer pool.Close()

	query := "SELECT id, name, description, picture, price_usd_currency_code, price_usd_units, price_usd_nanos, categories FROM " + pgTableName
	rows, err := pool.Query(context.Background(), query)
	if err != nil {
		log.Warnf("failed to query database: %v", err)
		return err
	}
	defer rows.Close()

	catalog.Products = catalog.Products[:0]
	for rows.Next() {
		product := &pb.Product{}
		product.PriceUsd = &pb.Money{}

		var categories string
		err = rows.Scan(&product.Id, &product.Name, &product.Description,
			&product.Picture, &product.PriceUsd.CurrencyCode, &product.PriceUsd.Units,
			&product.PriceUsd.Nanos, &categories)
		if err != nil {
			log.Warnf("failed to scan query result row: %v", err)
			return err
		}
		categories = strings.ToLower(categories)
		product.Categories = strings.Split(categories, ",")

		catalog.Products = append(catalog.Products, product)
	}

	log.Info("successfully parsed product catalog from AlloyDB")
	return nil
}

func loadCatalogFromVertex(catalog *pb.ListProductsResponse) error {
	log.Info("loading catalog from Vertex AI Search...")
	ctx := context.Background()

	projectID := os.Getenv("PROJECT_ID")
	location := os.Getenv("LOCATION")
	apiCatalog := os.Getenv("CATALOG")
	branch := os.Getenv("BRANCH")

	if projectID == "" || location == "" || apiCatalog == "" || branch == "" {
		return fmt.Errorf("missing required environment variables for Vertex AI Search integration")
	}

	c, err := retail.NewProductClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create retail client: %w", err)
	}
	defer c.Close()

	parent := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/branches/%s", projectID, location, apiCatalog, branch)

	req := &retailpb.ListProductsRequest{
		Parent: parent,
	}
	it := c.ListProducts(ctx, req)
	catalog.Products = catalog.Products[:0]

	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to list products: %w", err)
		}

		// Map the Vertex AI Product to our internal Product protobuf
		product := &pb.Product{
			Id:          resp.Id,
			Name:        resp.Title,
			Description: resp.Description,
			Categories:  resp.Categories,
		}

		if len(resp.Images) > 0 {
			product.Picture = resp.Images[0].Uri
		}

		if resp.PriceInfo != nil {
			product.PriceUsd = &pb.Money{
				CurrencyCode: resp.PriceInfo.CurrencyCode,
				Units:        int64(resp.PriceInfo.Price),
			}
		}
		catalog.Products = append(catalog.Products, product)
	}

	log.Infof("successfully loaded %d products from Vertex AI Search", len(catalog.Products))
	return nil
}
