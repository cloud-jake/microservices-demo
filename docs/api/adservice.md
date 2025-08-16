# Ad Service API Reference

## Overview

The Ad Service provides contextual advertisements based on a list of keywords.

**Service Name:** `hipstershop.AdService`
**Default Port:** `9555`

---

## Methods

### `GetAds`

Retrieves a list of advertisements based on the provided context keywords.

**Request:** `hipstershop.AdRequest`

```json
{
  "context_keys": ["clothing", "accessories"]
}
```

**Response:** `hipstershop.AdResponse`

```json
{
  "ads": [
    {
      "redirect_url": "/product/L9ECAV7KIM",
      "text": "Stylish loafers for just $89.99"
    }
  ]
}
```

---

## Message Types

### `Ad`

| Field          | Type     | Description                               |
|----------------|----------|-------------------------------------------|
| `redirect_url` | `string` | The URL to redirect to when the ad is clicked. |
| `text`         | `string` | The text content of the advertisement.    |