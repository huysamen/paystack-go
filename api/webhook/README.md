# Webhook Package

This package provides utilities for validating and handling Paystack webhooks.

## Package Structure

- `webhook.go` - Package documentation and overview
- `types.go` - Event types and data parsing utilities  
- `events.go` - Specific event data structures
- `validator.go` - Webhook signature validation and request handling

## Usage

```go
package main

import (
    "log"
    "net/http"
    
    "github.com/huysamen/paystack-go/api/webhook"
    "github.com/huysamen/paystack-go/types"
)

func main() {
    // Create webhook validator
    validator := webhook.NewValidator("your_secret_key_here")
    
    // Set up webhook handler
    http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
        // Validate the webhook
        event, err := validator.ValidateRequest(r)
        if err != nil {
            log.Printf("Invalid webhook: %v", err)
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        
        // Handle different event types
        switch event.Event {
        case webhook.EventChargeSuccess:
            // Parse charge success event using convenience method
            chargeEvent, err := event.AsChargeSuccess()
            if err != nil {
                log.Printf("Error parsing charge event: %v", err)
                return
            }
            
            log.Printf("Payment successful: %s - ₦%.2f", 
                chargeEvent.Reference, float64(chargeEvent.Amount)/100)
            
        case webhook.EventTransferSuccess:
            log.Printf("Transfer completed for event: %s", event.Event)
            
        case webhook.EventChargeDisputeCreate:
            log.Printf("Dispute created for event: %s", event.Event)
            
        default:
            log.Printf("Unhandled event: %s", event.Event)
        }
        
        w.WriteHeader(http.StatusOK)
    })
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Types

- `Event`: Represents a Paystack webhook event (defined in `types.go`)
- `Validator`: Provides webhook validation functionality (defined in `validator.go`)

## Constants

The package provides constants for all Paystack webhook event names (defined in `types.go`):

### Charge Events
- `EventChargeSuccess` - Successful charge

### Dispute Events
- `EventChargeDisputeCreate` - A dispute was logged against your business
- `EventChargeDisputeRemind` - A logged dispute has not been resolved
- `EventChargeDisputeResolve` - A dispute has been resolved

### Customer Identification Events  
- `EventCustomerIdentificationSuccess` - Customer identification successful
- `EventCustomerIdentificationFailed` - Customer identification failed

### Dedicated Virtual Account Events
- `EventDedicatedAccountAssignSuccess` - DVA successfully created and assigned
- `EventDedicatedAccountAssignFailed` - DVA couldn't be created and assigned

### Transfer Events
- `EventTransferSuccess` - Successful transfer
- `EventTransferFailed` - Failed transfer
- `EventTransferReversed` - Reversed transfer

### Subscription Events
- `EventSubscriptionCreate` - Subscription created
- `EventSubscriptionDisable` - Subscription disabled
- `EventSubscriptionNotRenew` - Subscription not renewing
- `EventSubscriptionExpiringCards` - Subscriptions with expiring cards

### Invoice Events
- `EventInvoiceCreate` - Invoice created
- `EventInvoicePaymentFailed` - Invoice payment failed
- `EventInvoiceUpdate` - Invoice updated

### Payment Request Events
- `EventPaymentRequestPending` - Payment request pending
- `EventPaymentRequestSuccess` - Payment request successful

### Refund Events
- `EventRefundProcessed` - Refund successfully processed
- `EventRefundFailed` - Refund failed
- `EventRefundPending` - Refund pending
- `EventRefundProcessing` - Refund processing

For the complete list with exact event strings, see `types.go`.

## Functions

### From `validator.go`:
- `NewValidator(secretKey string)`: Creates a new webhook validator
- `ValidateSignature(payload []byte, signature string)`: Validates webhook signature
- `ValidateRequest(r *http.Request)`: Validates complete webhook request

### From `types.go`:
- `ParseEventData[T any](event *Event)`: Parses event data into specified type

### Convenience Methods (from `types.go`):
- `Event.AsChargeSuccess()`: Parses as ChargeSuccessEvent
- `Event.AsCustomerIdentificationFailed()`: Parses as CustomerIdentificationFailedEvent  
- `Event.AsCustomerIdentificationSuccess()`: Parses as CustomerIdentificationSuccessEvent
- `Event.AsTransferSuccess()`: Parses as TransferSuccessEvent
- `Event.AsTransferFailed()`: Parses as TransferFailedEvent
- `Event.AsTransferReversed()`: Parses as TransferReversedEvent
- `Event.AsSubscriptionCreate()`: Parses as SubscriptionCreateEvent
- `Event.AsInvoiceCreate()`: Parses as InvoiceCreateEvent
- `Event.AsRefundProcessed()`: Parses as RefundProcessedEvent

## Event Data Structures

The package provides specific event structures (defined in `events.go`) for type-safe parsing:

- `ChargeSuccessEvent` - Data for successful charge events
- `CustomerIdentificationFailedEvent` - Data for failed customer identification
- `CustomerIdentificationSuccessEvent` - Data for successful customer identification  
- `TransferSuccessEvent` - Data for successful transfer events
- `TransferFailedEvent` - Data for failed transfer events
- `TransferReversedEvent` - Data for reversed transfer events
- `SubscriptionCreateEvent` - Data for subscription creation events
- `InvoiceCreateEvent` - Data for invoice creation events
- `RefundProcessedEvent` - Data for processed refund events

These structures reuse types from the main `types` package where applicable (Authorization, Customer, Plan, etc.) to maintain consistency across the SDK.
