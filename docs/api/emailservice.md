# Email Service API Reference

## Overview

The Email Service is responsible for sending order confirmation emails to users. In this demo, it simulates sending an email.

**Service Name:** `hipstershop.EmailService`
**Default Port:** `5000`

---

## Methods

### `SendOrderConfirmation`

Sends an order confirmation email to the specified address with the order details.

**Request:** `hipstershop.SendOrderConfirmationRequest`

```json
{
  "email": "jake.holmquist@example.com",
  "order": { ... }
}
```

**Response:** `hipstershop.Empty`

An empty message, indicating success.