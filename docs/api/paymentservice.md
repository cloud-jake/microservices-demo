# Payment Service API Reference

## Overview

The Payment Service handles credit card transactions. It takes payment details and an amount, and simulates a charge.

**Service Name:** `hipstershop.PaymentService`
**Default Port:** `50051`

---

## Methods

### `Charge`

Charges the provided credit card with the given amount.

**Request:** `hipstershop.ChargeRequest`

```json
{
  "amount": {
    "currency_code": "USD",
    "units": 117,
    "nanos": 980000000
  },
  "credit_card": {
    "credit_card_number": "4444-4444-4444-4444",
    "credit_card_cvv": 123,
    "credit_card_expiration_year": 2030,
    "credit_card_expiration_month": 10
  }
}
```

**Response:** `hipstershop.ChargeResponse`

```json
{
  "transaction_id": "a27b-4d8f-81e6-056c51b6a2f3"
}
```