# Product Catalog Service API Reference

## Overview

The Product Catalog Service provides a gRPC API for retrieving product information from the Online Boutique's catalog. It allows clients to list all available products, get details for a specific product by its ID, and search for products based on a query string.

**Service Name:** `hipstershop.ProductCatalogService`
**Default Port:** `3550`

---

## Methods

### `ListProducts`

Retrieves a list of all products available in the catalog.

**Request:** `hipstershop.Empty`

An empty message.

**Response:** `hipstershop.ListProductsResponse`

```json
{
  "products": [
    {
      "id": "OLJCESPC7Z",
      "name": "Sunglasses",
      "description": "Add a modern touch to your outfits with these sleek aviator sunglasses.",
      "picture": "/static/img/products/sunglasses.jpg",
      "priceUsd": {
        "currencyCode": "USD",
        "units": 19,
        "nanos": 990000000
      },
      "categories": ["accessories"]
    }
  ]
}
```

---

### `GetProduct`

Retrieves a single product by its unique ID. Returns a `NOT_FOUND` error if no product with the given ID exists.

**Request:** `hipstershop.GetProductRequest`

```json
{
  "id": "L9ECAV7KIM"
}
```

**Response:** `hipstershop.Product`

```json
{
  "id": "L9ECAV7KIM",
  "name": "Loafers",
  "description": "A neat addition to your summer wardrobe.",
  "picture": "/static/img/products/loafers.jpg",
  "priceUsd": {
    "currencyCode": "USD",
    "units": 89,
    "nanos": 990000000
  },
  "categories": ["footwear"]
}
```

---

### `SearchProducts`

Searches for products by a query string. The query is matched against product names and descriptions.

**Request:** `hipstershop.SearchProductsRequest`

```json
{
  "query": "mug"
}
```

**Response:** `hipstershop.SearchProductsResponse`

```json
{
  "results": [
    {
      "id": "6E92ZMYYFZ",
      "name": "Mug",
      "description": "A simple mug with a mustard interior.",
      "picture": "/static/img/products/mug.jpg",
      "priceUsd": {
        "currencyCode": "USD",
        "units": 8,
        "nanos": 990000000
      },
      "categories": ["kitchen"]
    }
  ]
}
```

---

## Message Types

### `Product`

| Field         | Type          | Description                                        |
|---------------|---------------|----------------------------------------------------|
| `id`          | `string`      | The unique identifier for the product.             |
| `name`        | `string`      | The name of the product.                           |
| `description` | `string`      | A detailed description of the product.             |
| `picture`     | `string`      | The path to the product's image.                   |
| `priceUsd`    | `Money`       | The price of the product in USD.                   |
| `categories`  | `array<string>` | A list of categories the product belongs to.       |

### `Money`

| Field          | Type     | Description                                                              |
|----------------|----------|--------------------------------------------------------------------------|
| `currencyCode` | `string` | The three-letter currency code (e.g., "USD").                            |
| `units`        | `int64`  | The whole units of the amount.                                           |
| `nanos`        | `int32`  | The nano (10^-9) units of the amount. Must be between -999,999,999 and +999,999,999. |