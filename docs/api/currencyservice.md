# Currency Service API Reference

## Overview

The Currency Service provides currency conversion and lists supported currencies. It uses real exchange rates from the European Central Bank.

**Service Name:** `hipstershop.CurrencyService`
**Default Port:** `7000`

---

## Methods

### `GetSupportedCurrencies`

Retrieves a list of all supported currency codes.

**Request:** `hipstershop.Empty`

An empty message.

**Response:** `hipstershop.GetSupportedCurrenciesResponse`

```json
{
  "currency_codes": ["USD", "EUR", "GBP", "JPY", "CAD"]
}
```

---

### `Convert`

Converts a given amount of money from one currency to another.

**Request:** `hipstershop.CurrencyConversionRequest`

```json
{
  "from": {
    "currency_code": "USD",
    "units": 100,
    "nanos": 0
  },
  "to_code": "EUR"
}
```

**Response:** `hipstershop.Money`

```json
{
  "currency_code": "EUR",
  "units": 93,
  "nanos": 150000000
}
```