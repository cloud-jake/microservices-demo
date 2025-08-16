# Testing Services with grpcurl

While you can use `curl` to interact with gRPC services (often requiring a gRPC-web proxy), a more direct and convenient tool for command-line interaction is grpcurl. This guide provides instructions on how to use `grpcurl` to test the microservices in this application.

## 1. Installation

First, you need to install `grpcurl`.

**macOS (using Homebrew):**
```sh
brew install grpcurl
```

**Other platforms:**
You can download pre-compiled binaries from the grpcurl releases page.

## 2. Port-Forwarding the Service

To access a service from your local machine, you need to forward its port from the Kubernetes cluster. Open a new terminal window and run the `kubectl port-forward` command.

**For `productcatalogservice`:**
```sh
kubectl port-forward deployment/productcatalogservice 3550:3550
```

**For `newproductcatalogservice`:**
```sh
kubectl port-forward deployment/newproductcatalogservice 3550:3550
```

Leave this terminal window running. All `grpcurl` commands should be run from a different terminal.

## 3. Testing API Methods

The following commands demonstrate how to call each method of the `ProductCatalogService`. The `-plaintext` flag is used because the services do not use TLS encryption within the cluster.

### ListProducts

This method retrieves all products and takes no parameters.

**Command:**
```sh
grpcurl -plaintext localhost:3550 hipstershop.ProductCatalogService/ListProducts
```

**Example Output:**
```json
{
  "products": [
    {
      "id": "OLJCESPC7Z",
      "name": "Sunglasses",
      // ... other products
    }
  ]
}
```

### GetProduct

This method retrieves a single product by its ID.

**Command:**
```sh
grpcurl -plaintext -d '{"id": "6E92ZMYYFZ"}' localhost:3550 hipstershop.ProductCatalogService/GetProduct
```

### SearchProducts

This method searches for products matching a query string.

**Command:**
```sh
grpcurl -plaintext -d '{"query": "mug"}' localhost:3550 hipstershop.ProductCatalogService/SearchProducts
```

You can use these patterns to test any of the gRPC-based services in the application by changing the port, service name, method, and request data (`-d` flag) as needed.