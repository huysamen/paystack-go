# Webhooks API

Secure webhook signature validation and typed event parsing for all Paystack webhook events.

## Overview

Paystack sends webhooks to notify your application about events that happen in your account. This package provides:

- **Signature Validation** - Verify webhook authenticity using HMAC-SHA512
- **Event Parsing** - Strongly typed event data structures
- **Helper Methods** - Easy conversion to specific event types

## Quick Start

### Basic Webhook Handler

```go
import (
    "net/http"
    "github.com/huysamen/paystack-go/api/webhook"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
    validator := webhook.NewValidator("sk_live_your_secret_key")
    
    event, err := validator.ValidateRequest(r)
    if err != nil {
        http.Error(w, "Invalid signature", http.StatusBadRequest)
        return
    }
    
    // Process the event
    if err := handleEvent(event); err != nil {
        http.Error(w, "Processing failed", http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusOK)
}
```

### Event Processing

```go
func handleEvent(event *webhook.Event) error {
    switch event.Event {
    case webhook.EventChargeSuccess:
        return handleChargeSuccess(event)
    case webhook.EventTransferSuccess:
        return handleTransferSuccess(event)
    case webhook.EventSubscriptionCreate:
        return handleSubscriptionCreate(event)
    default:
        log.Printf("Unhandled event: %s", event.Event)
        return nil
    }
}

func handleChargeSuccess(event *webhook.Event) error {
    data, err := event.AsChargeSuccess()
    if err != nil {
        return fmt.Errorf("failed to parse charge.success: %w", err)
    }
    
    // Verify the transaction for security
    result, err := client.Transactions.Verify(ctx, data.Reference.String())
    if err != nil {
        return fmt.Errorf("verification failed: %w", err)
    }
    
    if err := result.Err(); err != nil {
        return fmt.Errorf("transaction verification failed: %w", err)
    }
    
    // Process successful payment
    return processPayment(result.Data)
}
```

## Supported Events

### Payment Events

| Event | Description | Helper Method |
|-------|-------------|---------------|
| `charge.success` | Successful payment | `AsChargeSuccess()` |
| `charge.dispute.create` | Chargeback initiated | `AsChargeDisputeCreate()` |
| `charge.dispute.remind` | Dispute reminder | `AsChargeDisputeRemind()` |
| `charge.dispute.resolve` | Dispute resolved | `AsChargeDisputeResolve()` |

### Customer Events

| Event | Description | Helper Method |
|-------|-------------|---------------|
| `customeridentification.failed` | Customer ID verification failed | `AsCustomerIdentificationFailed()` |
| `customeridentification.success` | Customer ID verification succeeded | `AsCustomerIdentificationSuccess()` |

### Account Events

| Event | Description | Helper Method |
|-------|-------------|---------------|
| `dedicatedaccount.assign.failed` | Virtual account assignment failed | `AsDedicatedAccountAssignFailed()` |
| `dedicatedaccount.assign.success` | Virtual account assigned | `AsDedicatedAccountAssignSuccess()` |

### Invoice Events

| Event | Description | Helper Method |
|-------|-------------|---------------|
| `invoice.create` | Invoice created | `AsInvoiceCreate()` |
| `invoice.update` | Invoice updated | `AsInvoiceUpdate()` |
| `invoice.payment_failed` | Invoice payment failed | `AsInvoicePaymentFailed()` |

### Payment Request Events

| Event | Description | Helper Method |
|-------|-------------|---------------|
| `paymentrequest.pending` | Payment request created | `AsPaymentRequestPending()` |
| `paymentrequest.success` | Payment request paid | `AsPaymentRequestSuccess()` |

### Refund Events

| Event | Description | Helper Method |
|-------|-------------|---------------|
| `refund.failed` | Refund failed | `AsRefundFailed()` |
| `refund.pending` | Refund pending | `AsRefundPending()` |
| `refund.processed` | Refund processed | `AsRefundProcessed()` |
| `refund.processing` | Refund processing | `AsRefundProcessing()` |

### Subscription Events

| Event | Description | Helper Method |
|-------|-------------|---------------|
| `subscription.create` | Subscription created | `AsSubscriptionCreate()` |
| `subscription.disable` | Subscription disabled | `AsSubscriptionDisable()` |
| `subscription.not_renew` | Subscription won't renew | `AsSubscriptionNotRenew()` |
| `subscription.expiring_cards` | Cards expiring soon | `AsSubscriptionExpiringCards()` |

### Transfer Events

| Event | Description | Helper Method |
|-------|-------------|---------------|
| `transfer.failed` | Transfer failed | `AsTransferFailed()` |
| `transfer.reversed` | Transfer reversed | `AsTransferReversed()` |
| `transfer.success` | Transfer successful | `AsTransferSuccess()` |

## Detailed Examples

### Processing Payment Success

```go
func handleChargeSuccess(event *webhook.Event) error {
    data, err := event.AsChargeSuccess()
    if err != nil {
        return err
    }
    
    // Extract payment details
    reference := data.Reference.String()
    amount := data.Amount.Int64()
    currency := data.Currency.String()
    customerEmail := data.Customer.Email.String()
    
    log.Printf("Payment successful: %s paid %d %s (ref: %s)",
        customerEmail, amount, currency, reference)
    
    // Update your database
    return updateOrderStatus(reference, "paid", amount)
}
```

### Handling Transfer Events

```go
func handleTransferSuccess(event *webhook.Event) error {
    data, err := event.AsTransferSuccess()
    if err != nil {
        return err
    }
    
    transferCode := data.TransferCode.String()
    amount := data.Amount.Int64()
    recipientName := data.Recipient.Name.String()
    
    log.Printf("Transfer successful: %d kobo to %s (code: %s)",
        amount, recipientName, transferCode)
    
    // Update payout status
    return updatePayoutStatus(transferCode, "completed")
}
```

### Managing Subscription Events

```go
func handleSubscriptionCreate(event *webhook.Event) error {
    data, err := event.AsSubscriptionCreate()
    if err != nil {
        return err
    }
    
    subscriptionCode := data.SubscriptionCode.String()
    customerEmail := data.Customer.Email.String()
    planName := data.Plan.Name.String()
    
    log.Printf("New subscription: %s subscribed to %s (code: %s)",
        customerEmail, planName, subscriptionCode)
    
    // Activate user features
    return activateSubscription(customerEmail, planName, subscriptionCode)
}

func handleSubscriptionDisable(event *webhook.Event) error {
    data, err := event.AsSubscriptionDisable()
    if err != nil {
        return err
    }
    
    subscriptionCode := data.SubscriptionCode.String()
    customerEmail := data.Customer.Email.String()
    
    log.Printf("Subscription disabled: %s (code: %s)",
        customerEmail, subscriptionCode)
    
    // Deactivate user features
    return deactivateSubscription(subscriptionCode)
}
```

### Handling Disputes

```go
func handleChargeDisputeCreate(event *webhook.Event) error {
    data, err := event.AsChargeDisputeCreate()
    if err != nil {
        return err
    }
    
    disputeID := data.ID.Uint64()
    transactionRef := data.Transaction.Reference.String()
    reason := data.Reason.String()
    amount := data.Transaction.Amount.Int64()
    
    log.Printf("Dispute created: ID %d for transaction %s (reason: %s, amount: %d)",
        disputeID, transactionRef, reason, amount)
    
    // Notify relevant team
    return notifyDisputeTeam(disputeID, transactionRef, reason)
}
```

### Subscription Card Expiry

```go
func handleSubscriptionExpiringCards(event *webhook.Event) error {
    data, err := event.AsSubscriptionExpiringCards()
    if err != nil {
        return err
    }
    
    // This event contains an array of expiring cards
    if !data.Entries.Valid {
        return fmt.Errorf("invalid expiring cards data")
    }
    
    entries, ok := data.Entries.Metadata["entries"].([]interface{})
    if !ok {
        return fmt.Errorf("unexpected expiring cards format")
    }
    
    for _, entry := range entries {
        cardData, ok := entry.(map[string]interface{})
        if !ok {
            continue
        }
        
        customerEmail, _ := cardData["customer_email"].(string)
        cardLast4, _ := cardData["last4"].(string)
        expMonth, _ := cardData["exp_month"].(string)
        expYear, _ := cardData["exp_year"].(string)
        
        log.Printf("Card expiring: %s's card ending in %s expires %s/%s",
            customerEmail, cardLast4, expMonth, expYear)
        
        // Notify customer
        if err := notifyCardExpiry(customerEmail, cardLast4, expMonth, expYear); err != nil {
            log.Printf("Failed to notify customer %s: %v", customerEmail, err)
        }
    }
    
    return nil
}
```

## Security Best Practices

### 1. Always Validate Signatures

```go
func webhookHandler(w http.ResponseWriter, r *http.Request) {
    // NEVER process webhooks without signature validation
    validator := webhook.NewValidator("sk_live_your_secret_key")
    
    event, err := validator.ValidateRequest(r)
    if err != nil {
        log.Printf("Invalid webhook signature: %v", err)
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    
    // Now safe to process
    handleEvent(event)
}
```

### 2. Verify Critical Events via API

For important events like successful payments, always verify via API:

```go
func handleChargeSuccess(event *webhook.Event) error {
    data, err := event.AsChargeSuccess()
    if err != nil {
        return err
    }
    
    // Verify via API for additional security
    result, err := client.Transactions.Verify(ctx, data.Reference.String())
    if err != nil {
        return fmt.Errorf("verification failed: %w", err)
    }
    
    if err := result.Err(); err != nil {
        return fmt.Errorf("transaction not successful: %w", err)
    }
    
    // Verify amounts match
    if result.Data.Amount.Int64() != data.Amount.Int64() {
        return fmt.Errorf("amount mismatch: webhook=%d, api=%d",
            data.Amount.Int64(), result.Data.Amount.Int64())
    }
    
    // Now safe to process
    return processPayment(result.Data)
}
```

### 3. Idempotency

Handle duplicate webhooks gracefully:

```go
func processPayment(transaction types.Transaction) error {
    reference := transaction.Reference.String()
    
    // Check if already processed
    if isPaymentProcessed(reference) {
        log.Printf("Payment %s already processed, skipping", reference)
        return nil
    }
    
    // Process payment
    err := updateOrderStatus(reference, "paid", transaction.Amount.Int64())
    if err != nil {
        return err
    }
    
    // Mark as processed
    return markPaymentProcessed(reference)
}
```

### 4. Error Handling and Retries

Paystack will retry failed webhooks, so handle errors appropriately:

```go
func handleEvent(event *webhook.Event) error {
    switch event.Event {
    case webhook.EventChargeSuccess:
        if err := handleChargeSuccess(event); err != nil {
            // Log error but return nil to prevent retries for business logic errors
            if isBusinesLogicError(err) {
                log.Printf("Business logic error for %s: %v", event.Event, err)
                return nil
            }
            // Return error for temporary failures (DB down, etc.)
            return err
        }
    }
    return nil
}

func isBusinesLogicError(err error) bool {
    // Determine if error is due to business logic vs infrastructure
    return strings.Contains(err.Error(), "order not found") ||
           strings.Contains(err.Error(), "already processed")
}
```

## Testing

### Testing with Real Fixtures

The SDK includes real webhook JSON fixtures:

```go
func TestWebhookParsing(t *testing.T) {
    // Load fixture
    data, err := os.ReadFile("../../resources/examples/webhook/charge.success.json")
    require.NoError(t, err)
    
    // Parse event
    var event webhook.Event
    err = json.Unmarshal(data, &event)
    require.NoError(t, err)
    
    // Test specific event type
    chargeData, err := event.AsChargeSuccess()
    require.NoError(t, err)
    
    // Verify fields
    assert.Equal(t, "charge.success", event.Event)
    assert.True(t, chargeData.Amount.Int64() > 0)
    assert.NotEmpty(t, chargeData.Reference.String())
}
```

### Mock Webhook Testing

```go
func TestWebhookHandler(t *testing.T) {
    // Create test event
    eventData := map[string]interface{}{
        "event": "charge.success",
        "data": map[string]interface{}{
            "reference": "test_ref_123",
            "amount": 100000,
            "currency": "NGN",
            "status": "success",
        },
    }
    
    body, _ := json.Marshal(eventData)
    
    // Create signature
    validator := webhook.NewValidator("test_secret")
    signature := validator.GenerateSignature(body) // You'll need to implement this for testing
    
    // Create request
    req := httptest.NewRequest("POST", "/webhook", bytes.NewReader(body))
    req.Header.Set("X-Paystack-Signature", signature)
    
    // Test handler
    recorder := httptest.NewRecorder()
    webhookHandler(recorder, req)
    
    assert.Equal(t, http.StatusOK, recorder.Code)
}
```

See `api/webhook/webhook_test.go` for comprehensive examples using real JSON fixtures from `resources/examples/webhook/`.
