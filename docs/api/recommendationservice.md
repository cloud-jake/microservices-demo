# Recommendation Service API Reference

## Overview

The Recommendation Service provides product recommendations based on the items in a user's cart.

**Service Name:** `hipstershop.RecommendationService`
**Default Port:** `8080`

---

## Methods

### `ListRecommendations`

Returns a list of recommended product IDs.

**Request:** `hipstershop.ListRecommendationsRequest`

```json
{
  "user_id": "c2a8-4328-82e6-056c51b6a2f3",
  "product_ids": [
    "OLJCESPC7Z",
    "6E92ZMYYFZ"
  ]
}
```

**Response:** `hipstershop.ListRecommendationsResponse`

```json
{
  "product_ids": [
    "1YMWWN1N4O",
    "L9ECAV7KIM"
  ]
}
```