# Refunds API

Process refunds for completed transactions.

## Available Operations

- **Create** - Process a refund
- **List** - List refunds with filtering
- **Fetch** - Get refund details

## Quick Examples

### Create Refund

```go
import "github.com/huysamen/paystack-go/api/refunds"

request := refunds.NewCreateRequestBuilder().
    Transaction("transaction_reference").
    Amount(50000).  // Partial refund: 500 NGN
    Currency(enums.NGN).
    CustomerNote("Requested by customer").
    MerchantNote("Product returned in good condition").
    Build()

result, err := client.Refunds.Create(ctx, *request)
```

### List Refunds

```go
query := refunds.NewListRequestBuilder().
    PerPage(50).
    Page(1).
    Reference("transaction_ref").
    Currency(enums.NGN).
    Build()

result, err := client.Refunds.List(ctx, *query)
```

## Use Cases

### Full Refund

```go
request := refunds.NewCreateRequestBuilder().
    Transaction("TXN_123").
    // No amount = full refund
    CustomerNote("Order cancelled").
    Build()
```

### Partial Refund

```go
request := refunds.NewCreateRequestBuilder().
    Transaction("TXN_123").
    Amount(25000).  // Partial amount
    CustomerNote("Shipping fee refund").
    Build()
```
