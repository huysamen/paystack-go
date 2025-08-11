# Bulk Charges API

Process multiple charges in a single batch operation.

## Available Operations

- **Initiate** - Start a bulk charge batch
- **List Batches** - List bulk charge batches
- **Fetch Batch** - Get batch details
- **Fetch Batch Charges** - Get charges in a batch
- **Pause Batch** - Pause processing
- **Resume Batch** - Resume processing

## Quick Examples

### Initiate Bulk Charges

```go
import "github.com/huysamen/paystack-go/api/bulkcharges"

charges := []map[string]any{
    {
        "authorization": "AUTH_code_1", 
        "amount": 100000,
        "reference": "ref_1",
    },
    {
        "authorization": "AUTH_code_2",
        "amount": 150000, 
        "reference": "ref_2",
    },
}

request := bulkcharges.NewInitiateRequestBuilder().
    Charges(charges).
    Build()

result, err := client.BulkCharges.Initiate(ctx, *request)
```

### List Batches

```go
query := bulkcharges.NewListBatchesRequestBuilder().
    PerPage(50).
    Page(1).
    Build()

result, err := client.BulkCharges.ListBatches(ctx, *query)
```

## Use Cases

Perfect for subscription renewals, membership fees, or any scenario requiring multiple charges at once.
