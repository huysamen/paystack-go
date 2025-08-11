# Transactions API

Complete transaction management including initialization, verification, listing, exporting, and specialized operations like partial debits.

## Available Operations

- **Initialize** - Create a new transaction
- **Verify** - Verify transaction by reference
- **Fetch** - Get transaction by ID
- **List** - List transactions with filtering and pagination
- **Export** - Export transactions to CSV/PDF
- **Partial Debit** - Charge a portion of an authorized amount
- **Timeline** - Get transaction timeline/history
- **Totals** - Get transaction totals and statistics

## Quick Examples

### Verify a Transaction

```go
import "github.com/huysamen/paystack-go/api/transactions"

result, err := client.Transactions.Verify(ctx, "transaction_reference")
if err != nil {
    return fmt.Errorf("verification failed: %w", err)
}

if err := result.Err(); err != nil {
    return fmt.Errorf("transaction failed: %w", err)
}

txn := result.Data
fmt.Printf("Status: %s, Amount: %d kobo\n", 
    txn.Status.String(), 
    txn.Amount.Int64())
```

### List Transactions with Filters

```go
query := transactions.NewListRequestBuilder().
    PerPage(50).
    Page(1).
    Status("success").
    Customer(customerID).
    From(time.Now().Add(-30*24*time.Hour)).
    To(time.Now()).
    Amount(10000).  // Exact amount
    Build()

result, err := client.Transactions.List(ctx, *query)
if err != nil {
    return err
}

for _, txn := range result.Data {
    fmt.Printf("Ref: %s, Amount: %d, Status: %s\n",
        txn.Reference.String(),
        txn.Amount.Int64(),
        txn.Status.String())
}

// Handle pagination
if result.Meta.Next != nil {
    nextQuery := query
    nextQuery.Page = *result.Meta.Next
    // Fetch next page...
}
```

### Initialize a Transaction

```go
request := transactions.NewInitializeRequestBuilder().
    Email("customer@example.com").
    Amount(500000).  // 5000 NGN in kobo
    Currency(enums.NGN).
    Reference("unique_ref_" + time.Now().Format("20060102150405")).
    CallbackURL("https://yoursite.com/payment/callback").
    Build()

result, err := client.Transactions.Initialize(ctx, *request)
if err != nil {
    return err
}

if err := result.Err(); err != nil {
    return err
}

// Redirect customer to payment page
paymentURL := result.Data.AuthorizationURL.String()
fmt.Printf("Payment URL: %s\n", paymentURL)
```

### Partial Debit

Charge a portion of a previously authorized amount:

```go
request := transactions.NewPartialDebitRequestBuilder().
    AuthorizationCode("AUTH_authorization_code").
    Currency(enums.NGN).
    Amount(200000).  // 2000 NGN in kobo
    Email("customer@example.com").
    Reference("partial_" + time.Now().Format("20060102150405")).
    AtLeast("100000").  // Minimum amount to charge if full amount fails
    Build()

result, err := client.Transactions.PartialDebit(ctx, *request)
if err != nil {
    return err
}

if err := result.Err(); err != nil {
    return fmt.Errorf("partial debit failed: %w", err)
}

fmt.Printf("Charged: %d kobo\n", result.Data.Amount.Int64())
```

### Export Transactions

```go
request := transactions.NewExportRequestBuilder().
    From(time.Now().Add(-7*24*time.Hour)).
    To(time.Now()).
    Status("success").
    Currency(enums.NGN).
    Amount(50000).
    Settled(true).
    Settlement(12345).
    PaymentPage(67890).
    Build()

result, err := client.Transactions.Export(ctx, *request)
if err != nil {
    return err
}

if err := result.Err(); err != nil {
    return err
}

fmt.Printf("Export URL: %s\n", result.Data.Path.String())
```

### Get Transaction Timeline

```go
result, err := client.Transactions.Timeline(ctx, "transaction_id_or_reference")
if err != nil {
    return err
}

if err := result.Err(); err != nil {
    return err
}

timeline := result.Data
fmt.Printf("Timeline has %d entries\n", len(timeline.History))
```

### Get Transaction Totals

```go
result, err := client.Transactions.Totals(ctx)
if err != nil {
    return err
}

if err := result.Err(); err != nil {
    return err
}

totals := result.Data
fmt.Printf("Total transactions: %d\n", totals.TotalTransactions.Int64())
fmt.Printf("Total volume: %d kobo\n", totals.TotalVolume.Int64())
```

## Advanced Filtering

### List Builder Options

```go
query := transactions.NewListRequestBuilder().
    // Pagination
    PerPage(100).                    // Results per page (max 100)
    Page(1).                         // Page number
    
    // Filters
    Customer(12345).                 // Filter by customer ID
    Status("success").               // success, failed, abandoned
    From(startDate).                 // Start date
    To(endDate).                     // End date
    Amount(100000).                  // Exact amount in kobo
    
    // Advanced filters
    Currency(enums.NGN).             // Currency filter
    Settled(true).                   // Only settled transactions
    Settlement(settlementID).        // Specific settlement
    PaymentPage(pageID).             // Specific payment page
    
    Build()
```

### Export Builder Options

```go
request := transactions.NewExportRequestBuilder().
    // Date range (required)
    From(startDate).
    To(endDate).
    
    // Optional filters
    Status("success").
    Currency(enums.NGN).
    Amount(50000).
    Settled(true).
    Settlement(12345).
    PaymentPage(67890).
    Customer(customerId).
    
    Build()
```

## Response Data Types

### Transaction Object

The main transaction object contains:

```go
type Transaction struct {
    ID                 data.Uint           `json:"id"`
    Domain             data.String         `json:"domain"`
    Status             data.String         `json:"status"`
    Reference          data.String         `json:"reference"`
    Amount             data.Int            `json:"amount"`
    Message            data.NullString     `json:"message"`
    GatewayResponse    data.String         `json:"gateway_response"`
    PaidAt             data.NullTime       `json:"paid_at"`
    CreatedAt          data.Time           `json:"created_at"`
    Channel            enums.Channel       `json:"channel"`
    Currency           enums.Currency      `json:"currency"`
    IPAddress          data.String         `json:"ip_address"`
    Metadata           Metadata            `json:"metadata"`
    Log                *TransactionLog     `json:"log"`
    Fees               data.Int            `json:"fees"`
    Customer           Customer            `json:"customer"`
    Authorization      Authorization       `json:"authorization"`
    Plan               *Plan               `json:"plan"`
    // ... additional fields
}
```

### Initialize Response

```go
type InitializeResponseData struct {
    AuthorizationURL data.String `json:"authorization_url"`
    AccessCode      data.String `json:"access_code"`
    Reference       data.String `json:"reference"`
}
```

## Error Handling Patterns

### Common Error Scenarios

```go
result, err := client.Transactions.Verify(ctx, reference)

// Network/HTTP errors
if err != nil {
    if errors.Is(err, context.DeadlineExceeded) {
        return fmt.Errorf("request timeout: %w", err)
    }
    return fmt.Errorf("network error: %w", err)
}

// Paystack API errors
if err := result.Err(); err != nil {
    switch {
    case strings.Contains(err.Error(), "not found"):
        return fmt.Errorf("transaction not found: %s", reference)
    case strings.Contains(err.Error(), "invalid"):
        return fmt.Errorf("invalid reference format: %s", reference)
    default:
        return fmt.Errorf("verification failed: %w", err)
    }
}

// Success
transaction := result.Data
```

## Best Practices

### 1. Reference Generation

Always use unique, traceable references:

```go
import "crypto/rand"
import "encoding/hex"

func generateReference(prefix string) string {
    bytes := make([]byte, 8)
    rand.Read(bytes)
    return fmt.Sprintf("%s_%s_%d", 
        prefix, 
        hex.EncodeToString(bytes), 
        time.Now().Unix())
}

reference := generateReference("order")
```

### 2. Idempotency

Use consistent references for idempotent operations:

```go
// Use order ID or similar unique identifier
reference := fmt.Sprintf("order_%d", orderID)

result, err := client.Transactions.Initialize(ctx, 
    transactions.NewInitializeRequestBuilder().
        Reference(reference).
        // ... other fields
        Build())
```

### 3. Webhook Verification

Always verify webhooks for transaction updates:

```go
func handleTransactionWebhook(event *webhook.Event) error {
    data, err := event.AsChargeSuccess()
    if err != nil {
        return err
    }
    
    // Verify the transaction via API for security
    result, err := client.Transactions.Verify(ctx, data.Reference.String())
    if err != nil {
        return err
    }
    
    if err := result.Err(); err != nil {
        return err
    }
    
    // Process the verified transaction
    return processTransaction(result.Data)
}
```

### 4. Pagination Handling

```go
func fetchAllTransactions(client *paystack.Client, query transactions.ListRequest) ([]types.Transaction, error) {
    var allTransactions []types.Transaction
    page := 1
    
    for {
        query.Page = page
        result, err := client.Transactions.List(ctx, query)
        if err != nil {
            return nil, err
        }
        
        if err := result.Err(); err != nil {
            return nil, err
        }
        
        allTransactions = append(allTransactions, result.Data...)
        
        // Check if there are more pages
        if result.Meta.Next == nil || len(result.Data) == 0 {
            break
        }
        
        page++
    }
    
    return allTransactions, nil
}
```

## Testing

See `api/transactions/*_test.go` for comprehensive examples using real JSON fixtures from `resources/examples/responses/transactions/`.

Each test demonstrates:
- Request builder usage
- Response parsing
- Error handling
- Field validation
