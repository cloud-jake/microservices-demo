# Shipping Service API Reference

## Overview

The Shipping Service calculates shipping costs and "ships" orders by providing a tracking ID.

**Service Name:** `hipstershop.ShippingService`
**Default Port:** `50051`

---

## Methods

### `GetQuote`

Calculates the shipping cost for a list of items to a given address.

**Request:** `hipstershop.GetQuoteRequest`

```json
{
  "address": {
    "street_address": "1600 Amphitheatre Parkway",
    "city": "Mountain View",
    "state": "CA",
    "country": "USA",
    "zip_code": 94043
  },
  "items": [
    { "product_id": "6E92ZMYYFZ", "quantity": 2 }
  ]
}
```

**Response:** `hipstershop.GetQuoteResponse`

```json
{
  "cost_usd": {
    "currency_code": "USD",
    "units": 10,
    "nanos": 0
  }
}
```

---

### `ShipOrder`

Simulates shipping an order and returns a tracking ID.

**Request:** `hipstershop.ShipOrderRequest` (Same structure as `GetQuoteRequest`)

**Response:** `hipstershop.ShipOrderResponse`

```json
{
  "tracking_id": "W3S4-50P6-82R1-756C"
}
```