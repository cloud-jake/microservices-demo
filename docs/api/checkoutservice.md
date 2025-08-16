# Checkout Service API Reference

## Overview

The Checkout Service handles the order placement process. It orchestrates calls to other services to charge the user's credit card, ship the items, and send a confirmation email.

**Service Name:** `hipstershop.CheckoutService`
**Default Port:** `5050`

---

## Methods

### `PlaceOrder`

Places an order for the items in a user's cart.

**Request:** `hipstershop.PlaceOrderRequest`

```json
{
  "user_id": "c2a8-4328-82e6-056c51b6a2f3",
  "user_currency": "USD",
  "address": {
    "street_address": "1600 Amphitheatre Parkway",
    "city": "Mountain View",
    "state": "CA",
    "country": "USA",
    "zip_code": 94043
  },
  "email": "jake.holmquist@example.com",
  "credit_card": {
    "credit_card_number": "4444-4444-4444-4444",
    "credit_card_cvv": 123,
    "credit_card_expiration_year": 2030,
    "credit_card_expiration_month": 10
  }
}
```

**Response:** `hipstershop.PlaceOrderResponse`

```json
{
  "order": {
    "order_id": "d8f2-4c86-80e6-756c21b6a2f1",
    "shipping_tracking_id": "W3S4-50P6-82R1-756C",
    "shipping_cost": { "currencyCode": "USD", "units": 10, "nanos": 0 },
    "shipping_address": {
      "street_address": "1600 Amphitheatre Parkway",
      "city": "Mountain View",
      "state": "CA",
      "country": "USA",
      "zip_code": 94043
    },
    "items": [
      {
        "item": { "product_id": "6E92ZMYYFZ", "quantity": 2 },
        "cost": { "currencyCode": "USD", "units": 17, "nanos": 980000000 }
      }
    ]
  }
}
```

---

## Message Types

### `Address`

| Field            | Type     | Description                               |
|------------------|----------|-------------------------------------------|
| `street_address` | `string` | The street address.                       |
| `city`           | `string` | The city name.                            |
| `state`          | `string` | The state or province.                    |
| `country`        | `string` | The country name.                         |
| `zip_code`       | `int32`  | The postal code.                          |

### `CreditCardInfo`

| Field                          | Type    | Description                               |
|--------------------------------|---------|-------------------------------------------|
| `credit_card_number`           | `string`| The 16-digit credit card number.        |
| `credit_card_cvv`              | `int32` | The 3 or 4-digit CVV code.            |
| `credit_card_expiration_year`  | `int32` | The 4-digit expiration year.            |
| `credit_card_expiration_month` | `int32` | The 2-digit expiration month.           |