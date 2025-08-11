# Transfers API

Send money to bank accounts, mobile money wallets, and other recipients with comprehensive transfer management.

## Available Operations

- **Initiate** - Send money to a recipient
- **Finalize** - Complete transfer with OTP
- **Fetch** - Get transfer by ID or code
- **List** - List transfers with filtering
- **Verify** - Verify transfer by reference
- **Bulk** - Send to multiple recipients at once

## Quick Examples

### Initiate Transfer

```go
import "github.com/huysamen/paystack-go/api/transfers"

request := transfers.NewInitiateRequestBuilder("balance", 50000, "RCP_recipient_code").
    Reason("Salary payment").
    Currency(enums.NGN).
    Reference("transfer_" + time.Now().Format("20060102150405")).
    Build()

result, err := client.Transfers.Initiate(ctx, *request)
if err != nil {
    return fmt.Errorf("transfer initiation failed: %w", err)
}

if err := result.Err(); err != nil {
    return fmt.Errorf("transfer failed: %w", err)
}

transfer := result.Data
fmt.Printf("Transfer initiated: %s (Status: %s)\n", 
    transfer.TransferCode.String(), 
    transfer.Status.String())
```

### Finalize Transfer with OTP

```go
request := transfers.FinalizeRequest{
    TransferCode: "TRF_transfer_code",
    OTP:          "123456",
}

result, err := client.Transfers.Finalize(ctx, request)
if err != nil {
    return err
}

if err := result.Err(); err != nil {
    return fmt.Errorf("finalization failed: %w", err)
}

fmt.Printf("Transfer finalized: %s\n", result.Data.Status.String())
```

### Bulk Transfers

```go
bulk := transfers.NewBulkRequestBuilder("balance").
    AddTransfer(transfers.BulkTransferItem{
        Amount:    100000,
        Reference: "bulk_1",
        Reason:    "Commission payment",
        Recipient: "RCP_recipient_1",
    }).
    AddTransfer(transfers.BulkTransferItem{
        Amount:    200000,
        Reference: "bulk_2", 
        Reason:    "Vendor payment",
        Recipient: "RCP_recipient_2",
    }).
    Build()

result, err := client.Transfers.Bulk(ctx, *bulk)
if err != nil {
    return err
}

if err := result.Err(); err != nil {
    return err
}

fmt.Printf("Bulk transfer initiated with %d transfers\n", len(result.Data))
```

### List Transfers

```go
query := transfers.NewListRequestBuilder().
    PerPage(50).
    Page(1).
    Customer(12345).
    Status("success").
    From(time.Now().Add(-30*24*time.Hour)).
    To(time.Now()).
    Build()

result, err := client.Transfers.List(ctx, *query)
if err != nil {
    return err
}

for _, transfer := range result.Data {
    fmt.Printf("Transfer: %s - %d kobo to %s\n",
        transfer.TransferCode.String(),
        transfer.Amount.Int64(),
        transfer.Recipient.Name.String())
}
```

## Advanced Usage

### Transfer with Metadata

```go
request := transfers.NewInitiateRequestBuilder("balance", 75000, "RCP_xyz").
    Reason("Product purchase payout").
    Currency(enums.NGN).
    Reference("order_123_payout").
    Metadata(map[string]any{
        "order_id": "ORDER_123",
        "product": "Premium Plan",
        "commission_rate": 0.15,
    }).
    Build()
```

### Error Handling for Transfer States

```go
result, err := client.Transfers.Initiate(ctx, request)
if err != nil {
    return err
}

if err := result.Err(); err != nil {
    // Handle specific transfer errors
    switch {
    case strings.Contains(err.Error(), "insufficient"):
        return fmt.Errorf("insufficient balance: %w", err)
    case strings.Contains(err.Error(), "recipient"):
        return fmt.Errorf("invalid recipient: %w", err)
    case strings.Contains(err.Error(), "limit"):
        return fmt.Errorf("transfer limit exceeded: %w", err)
    default:
        return fmt.Errorf("transfer failed: %w", err)
    }
}

// Check if OTP is required
transfer := result.Data
if transfer.Status.String() == "otp" {
    fmt.Println("OTP required for this transfer")
    // Collect OTP from user and call Finalize
}
```

## Request Builders

### Initiate Transfer Builder

```go
request := transfers.NewInitiateRequestBuilder(source, amount, recipient).
    // Required parameters in constructor
    
    // Optional parameters
    Reason("Payment description").
    Currency(enums.NGN).
    Reference("unique_reference").
    Metadata(map[string]any{
        "key": "value",
    }).
    Build()
```

### Bulk Transfer Builder

```go
bulk := transfers.NewBulkRequestBuilder("balance").
    Currency(enums.NGN).  // Optional: defaults to NGN
    AddTransfer(transfers.BulkTransferItem{
        Amount:    amount1,
        Reference: "ref1",
        Reason:    "reason1", 
        Recipient: "RCP_xxx",
    }).
    AddTransfer(transfers.BulkTransferItem{
        Amount:    amount2,
        Reference: "ref2",
        Reason:    "reason2",
        Recipient: "RCP_yyy", 
    }).
    Build()
```

### List Transfers Builder

```go
query := transfers.NewListRequestBuilder().
    PerPage(100).
    Page(1).
    Customer(customerID).
    Status("success").      // otp, pending, success, failed, reversed
    From(startDate).
    To(endDate).
    Build()
```

## Transfer Status Flow

```
pending → otp → success
        ↓
      failed
        ↓
     reversed (in some cases)
```

### Status Meanings

- **pending**: Transfer is being processed
- **otp**: OTP required to complete transfer
- **success**: Transfer completed successfully  
- **failed**: Transfer failed
- **reversed**: Transfer was reversed/refunded

## Best Practices

### 1. Reference Generation

Use unique, traceable references:

```go
func generateTransferReference(orderID string) string {
    timestamp := time.Now().Format("20060102150405")
    return fmt.Sprintf("transfer_%s_%s", orderID, timestamp)
}
```

### 2. Balance Checking

Check balance before initiating transfers:

```go
// Get balance first
balanceResult, err := client.Transfers.GetBalance(ctx)
if err != nil {
    return err
}

balance := balanceResult.Data[0].Balance.Int64() // Assuming NGN
if balance < transferAmount {
    return fmt.Errorf("insufficient balance: have %d, need %d", balance, transferAmount)
}

// Proceed with transfer
```

### 3. Webhook Integration

Handle transfer webhooks for real-time updates:

```go
func handleTransferWebhook(event *webhook.Event) error {
    switch event.Event {
    case webhook.EventTransferSuccess:
        data, _ := event.AsTransferSuccess()
        return markTransferCompleted(data.TransferCode.String())
        
    case webhook.EventTransferFailed:
        data, _ := event.AsTransferFailed()
        return markTransferFailed(data.TransferCode.String(), data.Status.String())
        
    case webhook.EventTransferReversed:
        data, _ := event.AsTransferReversed()
        return processTransferReversal(data.TransferCode.String())
    }
    return nil
}
```

### 4. Bulk Transfer Management

```go
func processBulkPayouts(payouts []PayoutRequest) error {
    const maxBulkSize = 100 // Paystack limit
    
    for i := 0; i < len(payouts); i += maxBulkSize {
        end := i + maxBulkSize
        if end > len(payouts) {
            end = len(payouts)
        }
        
        batch := payouts[i:end]
        if err := processBatch(batch); err != nil {
            return fmt.Errorf("batch %d failed: %w", i/maxBulkSize+1, err)
        }
    }
    
    return nil
}

func processBatch(payouts []PayoutRequest) error {
    bulk := transfers.NewBulkRequestBuilder("balance")
    
    for _, payout := range payouts {
        bulk.AddTransfer(transfers.BulkTransferItem{
            Amount:    payout.Amount,
            Reference: payout.Reference,
            Reason:    payout.Reason,
            Recipient: payout.RecipientCode,
        })
    }
    
    result, err := client.Transfers.Bulk(ctx, *bulk.Build())
    if err != nil {
        return err
    }
    
    return result.Err()
}
```

## Testing

See `api/transfers/*_test.go` for comprehensive examples using real JSON fixtures from `resources/examples/responses/transfers/`.
