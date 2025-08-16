# Cart Service API Reference

## Overview

The Cart Service manages user shopping carts. It provides methods to add items to a cart, retrieve the contents of a cart, and empty a cart. Carts are associated with a `user_id`.

**Service Name:** `hipstershop.CartService`
**Default Port:** `7070`

---

## Methods

### `AddItem`

Adds a specified quantity of a product to a user's cart.

**Request:** `hipstershop.AddItemRequest`

```json
{
  "user_id": "c2a8-4328-82e6-056c51b6a2f3",
  "item": {
    "product_id": "6E92ZMYYFZ",
    "quantity": 2
  }
}
```

**Response:** `hipstershop.Empty`

An empty message, indicating success.

---

### `GetCart`

Retrieves the contents of a user's shopping cart.

**Request:** `hipstershop.GetCartRequest`

```json
{
  "user_id": "c2a8-4328-82e6-056c51b6a2f3"
}
```

**Response:** `hipstershop.Cart`

```json
{
  "user_id": "c2a8-4328-82e6-056c51b6a2f3",
  "items": [
    { "product_id": "6E92ZMYYFZ", "quantity": 2 }
  ]
}
```

---

### `EmptyCart`

Removes all items from a user's shopping cart.

**Request:** `hipstershop.EmptyCartRequest`

```json
{
  "user_id": "c2a8-4328-82e6-056c51b6a2f3"
}
```

**Response:** `hipstershop.Empty`

An empty message, indicating success.