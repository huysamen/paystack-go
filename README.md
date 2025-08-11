# Paystack Go SDK

A comprehensive, production-ready Go SDK for the Paystack API. Built with idiomatic Go patterns, robust error handling, and extensive test coverage.

## Features

- 🚀 **Complete API Coverage** - All Paystack endpoints with typed request/response models
- 🛡️ **Type Safety** - Strongly typed with custom JSON handling for API quirks
- 🔧 **Flexible Configuration** - Custom timeouts, headers, HTTP clients, and base URLs
- 📝 **Fluent Builders** - Chainable request builders for clean, readable code
- 🔐 **Webhook Validation** - HMAC signature validation with typed event parsing
- ✅ **Thoroughly Tested** - Comprehensive tests with real JSON fixtures
- 🌐 **Context Aware** - All operations support context.Context for cancellation
- 📚 **Rich Documentation** - Detailed examples for every package and endpoint

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Configuration](#configuration)
- [Making Requests](#making-requests)
- [Error Handling](#error-handling)
- [Types and JSON Handling](#types-and-json-handling)
- [Webhooks](#webhooks)
- [API Packages](#api-packages)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Installation

```bash
go get github.com/huysamen/paystack-go
```

**Requirements:** Go 1.21+

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    paystack "github.com/huysamen/paystack-go"
    "github.com/huysamen/paystack-go/api/transactions"
)

func main() {
    // Create a client with your secret key
    client := paystack.DefaultClient("sk_test_your_secret_key")
    
    ctx := context.Background()
    
    // Verify a transaction
    result, err := client.Transactions.Verify(ctx, "your_transaction_reference")
    if err != nil {
        log.Fatal("HTTP error:", err)
    }
    
    // Check if the transaction was successful
    if err := result.Err(); err != nil {
        log.Fatal("Paystack error:", err)
    }
    
    fmt.Printf("Transaction Status: %s\n", result.Data.Status.String())
    fmt.Printf("Amount: %d kobo\n", result.Data.Amount.Int64())
}
```

## Configuration

### Basic Configuration

```go
// Simple client with default settings
client := paystack.DefaultClient("sk_test_your_secret_key")
```

### Advanced Configuration

```go
import (
    "time"
    "net/http"
    paystack "github.com/huysamen/paystack-go"
)

cfg := paystack.NewConfig("sk_test_your_secret_key").
    WithTimeout(30*time.Second).
    WithDefaultHeaders(map[string]string{
        "X-Application": "my-app",
        "X-Version": "1.0.0",
    }).
    WithUserAgentSuffix("MyApp/1.0.0")

// Optional: Use your own HTTP client
customClient := &http.Client{
    Timeout: 45 * time.Second,
    // Add custom transport, proxy settings, etc.
}
cfg = cfg.WithHTTPClient(customClient)

client := paystack.NewClient(cfg)
```

### Configuration Options

| Option | Description | Example |
|--------|-------------|---------|
| `WithTimeout` | Request timeout | `30*time.Second` |
| `WithDefaultHeaders` | Headers added to every request | `map[string]string{"X-App": "myapp"}` |
| `WithUserAgentSuffix` | Suffix for User-Agent header | `"MyApp/1.0.0"` |
| `WithHTTPClient` | Custom HTTP client | Custom transport, proxy, etc. |
| `WithBaseURL` | Override API base URL | For testing/staging environments |

## Making Requests

All API calls follow a consistent pattern using fluent request builders:

### List Operations with Pagination

```go
import "github.com/huysamen/paystack-go/api/transactions"

// Build query with filters
query := transactions.NewListRequestBuilder().
    PerPage(50).
    Page(1).
    Status("success").
    From(time.Now().Add(-30*24*time.Hour)).
    To(time.Now()).
    Build()

result, err := client.Transactions.List(ctx, *query)
if err != nil {
    // Handle error
}

for _, txn := range result.Data {
    fmt.Printf("Transaction: %s - %d kobo\n", 
        txn.Reference.String(), 
        txn.Amount.Int64())
}
```

### Create Operations

```go
import "github.com/huysamen/paystack-go/api/customers"

request := customers.NewCreateRequestBuilder().
    Email("customer@example.com").
    FirstName("John").
    LastName("Doe").
    Phone("+2348012345678").
    Build()

result, err := client.Customers.Create(ctx, *request)
```

### Context and Cancellation

All operations accept `context.Context` for timeout and cancellation:

```go
// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

result, err := client.Transactions.Verify(ctx, "ref_123")
```

## Error Handling

The SDK provides two levels of error handling:

### 1. HTTP/Network Errors

```go
result, err := client.Transactions.Verify(ctx, "ref_123")
if err != nil {
    // Network error, timeout, invalid response, etc.
    log.Fatal("Request failed:", err)
}
```

### 2. Paystack API Errors

```go
// Method 1: Check status manually
if !result.Status.Bool() {
    log.Fatal("Paystack error:", result.Message.String())
}

// Method 2: Use convenience method
if err := result.Err(); err != nil {
    log.Fatal("Paystack error:", err)
}
```

### Complete Error Handling Pattern

```go
result, err := client.Transactions.Verify(ctx, "ref_123")
if err != nil {
    return fmt.Errorf("transaction verification failed: %w", err)
}

if err := result.Err(); err != nil {
    return fmt.Errorf("paystack error: %w", err)
}

// Success - use result.Data
transaction := result.Data
```

## Types and JSON Handling

Paystack's API can return inconsistent JSON shapes. The SDK handles this gracefully:

### Custom Data Types

The SDK provides custom types in `types/data` that handle JSON variations:

```go
import "github.com/huysamen/paystack-go/types/data"

// Non-nullable types (convert null to zero value)
var amount data.Int      // Handles: 1000, "1000", null (becomes 0)
var email data.String    // Handles: "test@example.com", null (becomes "")
var active data.Bool     // Handles: true, "true", 1, null (becomes false)

// Nullable types (preserve null state)
var description data.NullString  // Tracks if value was null
var updatedAt data.NullTime      // Can be null or valid time

if description.Valid {
    fmt.Println("Description:", description.String())
} else {
    fmt.Println("No description provided")
}
```

### Metadata Handling

```go
import "github.com/huysamen/paystack-go/types"

// Metadata is flexible and handles various JSON shapes
var meta types.Metadata

// Can handle: null, "", 0, {}, {"key": "value"}
// Only valid objects set Valid=true
if meta.Valid {
    if value, exists := meta.Metadata["custom_field"]; exists {
        fmt.Println("Custom field:", value)
    }
}
```

### Response Variants

Some endpoints return different shapes for the same entity. The SDK handles this with specialized response types:

```go
// When creating a subscription, API returns numeric IDs
createResult, _ := client.Subscriptions.Create(ctx, request)
customerID := createResult.Data.Customer.Int64()  // Numeric ID

// When fetching a subscription, API returns full objects
fetchResult, _ := client.Subscriptions.Fetch(ctx, subscriptionID)
customerEmail := fetchResult.Data.Customer.Email.String()  // Full object
```

## Webhooks

### Signature Validation

```go
import "github.com/huysamen/paystack-go/api/webhook"

func webhookHandler(w http.ResponseWriter, r *http.Request) {
    validator := webhook.NewValidator("sk_live_your_secret_key")
    
    event, err := validator.ValidateRequest(r)
    if err != nil {
        http.Error(w, "Invalid signature", http.StatusBadRequest)
        return
    }
    
    // Process the event
    handleEvent(event)
    
    w.WriteHeader(http.StatusOK)
}
```

### Event Processing

```go
import "github.com/huysamen/paystack-go/api/webhook"

func handleEvent(event *webhook.Event) {
    switch event.Event {
    case webhook.EventChargeSuccess:
        data, err := event.AsChargeSuccess()
        if err != nil {
            log.Printf("Failed to parse charge.success: %v", err)
            return
        }
        
        fmt.Printf("Payment successful: %s for %d kobo\n", 
            data.Reference.String(), 
            data.Amount.Int64())
            
    case webhook.EventTransferSuccess:
        data, err := event.AsTransferSuccess()
        if err != nil {
            log.Printf("Failed to parse transfer.success: %v", err)
            return
        }
        
        fmt.Printf("Transfer completed: %s\n", data.TransferCode.String())
        
    case webhook.EventSubscriptionCreate:
        data, err := event.AsSubscriptionCreate()
        if err != nil {
            log.Printf("Failed to parse subscription.create: %v", err)
            return
        }
        
        fmt.Printf("New subscription: %s\n", data.SubscriptionCode.String())
        
    default:
        log.Printf("Unhandled event: %s", event.Event)
    }
}
```

## API Packages

Each API surface has its own package with detailed documentation and examples:

### Core APIs
- [**Transactions**](api/transactions/README.md) - Initialize, verify, list, export transactions
- [**Customers**](api/customers/README.md) - Create and manage customer records
- [**Plans**](api/plans/README.md) - Subscription billing plans
- [**Subscriptions**](api/subscriptions/README.md) - Recurring billing management

### Payment Methods
- [**Payment Pages**](api/paymentpages/README.md) - Hosted payment pages
- [**Payment Requests**](api/paymentrequests/README.md) - Request payments via email/SMS

### Transfers and Payouts
- [**Transfers**](api/transfers/README.md) - Send money to bank accounts
- [**Transfer Recipients**](api/transferrecipients/README.md) - Manage payout destinations
- [**Transfer Control**](api/transferscontrol/README.md) - Balance and transfer controls

### Products and Inventory
- [**Products**](api/products/README.md) - Product catalog management

### Account Management
- [**Subaccounts**](api/subaccounts/README.md) - Split payments and settlements
- [**Dedicated Virtual Accounts**](api/dedicatedvirtualaccounts/README.md) - Virtual account numbers

### Transaction Management
- [**Refunds**](api/refunds/README.md) - Process refunds
- [**Disputes**](api/disputes/README.md) - Handle chargebacks

### Bulk Operations
- [**Bulk Charges**](api/bulkcharges/README.md) - Batch payment processing

### Point of Sale
- [**Terminal**](api/terminal/README.md) - Physical terminal management
- [**Virtual Terminal**](api/virtualterminal/README.md) - Virtual POS operations

### Integration and Utilities
- [**Integration**](api/integration/README.md) - Account and integration settings
- [**Verification**](api/verification/README.md) - Identity verification services
- [**Miscellaneous**](api/miscellaneous/README.md) - Banks, countries, and other utilities
- [**Apple Pay**](api/applepay/README.md) - Apple Pay integration

### Webhooks
- [**Webhooks**](api/webhook/README.md) - Event validation and processing

## Testing

The SDK includes comprehensive tests with real JSON fixtures:

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests for a specific package
go test ./api/transactions

# Run tests with coverage
go test -cover ./...
```

### Test Structure

- **Unit tests** for every API package and webhook parsing
- **JSON fixtures** in `resources/examples/` directory
- **Builder tests** to verify request construction
- **Response tests** to verify JSON unmarshaling

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Ensure all tests pass
5. Submit a pull request

### Development Setup

```bash
git clone https://github.com/huysamen/paystack-go.git
cd paystack-go
go mod download
go test ./...
```

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

---

**Need help?** Check the individual package READMEs for detailed examples, or review the test files for comprehensive usage patterns.
