# Disputes API

Handle transaction disputes and chargebacks.

## Available Operations

- **List** - List disputes with filtering
- **Fetch** - Get dispute details
- **List Evidence** - Get dispute evidence
- **Get Upload URL** - Get evidence upload URL
- **Add Evidence** - Submit dispute evidence
- **Resolve** - Resolve a dispute
- **Export** - Export disputes data

## Quick Examples

### List Disputes

```go
import "github.com/huysamen/paystack-go/api/disputes"

query := disputes.NewListRequestBuilder().
    PerPage(50).
    Page(1).
    Status("awaiting-merchant-feedback").
    Transaction("TXN_123").
    Build()

result, err := client.Disputes.List(ctx, *query)
```

### Add Evidence

```go
request := disputes.NewAddEvidenceRequestBuilder().
    DisputeID(12345).
    CustomerEmail("customer@example.com").
    CustomerName("John Doe").
    CustomerPhone("+2348012345678").
    ServiceDetails("Premium subscription service").
    DeliveryAddress("123 Lagos Street, Lagos").
    Build()

result, err := client.Disputes.AddEvidence(ctx, *request)
```

### Resolve Dispute

```go
request := disputes.NewResolveRequestBuilder().
    DisputeID(12345).
    Resolution("merchant-accepted").
    Message("Customer concern addressed").
    RefundAmount(50000).  // Partial refund
    UploadFilename("evidence.pdf").
    Evidence(disputeEvidenceID).
    Build()

result, err := client.Disputes.Resolve(ctx, *request)
```

## Dispute Status Flow

- **awaiting-merchant-feedback** - Merchant needs to respond
- **awaiting-bank-feedback** - Under bank review
- **pending** - Processing
- **resolved** - Dispute resolved

## Best Practices

### Quick Response

Respond to disputes within 24-48 hours:

```go
func handleNewDispute(disputeID uint64) error {
    // Gather evidence immediately
    evidence := gatherDisputeEvidence(disputeID)
    
    request := disputes.NewAddEvidenceRequestBuilder().
        DisputeID(disputeID).
        CustomerEmail(evidence.CustomerEmail).
        ServiceDetails(evidence.ServiceDescription).
        Build()
        
    return client.Disputes.AddEvidence(ctx, *request)
}
```
