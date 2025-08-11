# Transfer Control API

Manage transfer settings, balances, and controls for your account.

## Available Operations

- **Check Balance** - Get account balance
- **Fetch Balance History** - Get balance ledger
- **Resend OTP** - Resend transfer OTP
- **Disable OTP** - Disable OTP requirement
- **Finalize Disable OTP** - Complete OTP disabling
- **Enable OTP** - Re-enable OTP requirement

## Quick Examples

### Check Balance

```go
import "github.com/huysamen/paystack-go/api/transferscontrol"

result, err := client.TransfersControl.CheckBalance(ctx)
if err != nil {
    return err
}

for _, balance := range result.Data {
    fmt.Printf("Currency: %s, Balance: %d kobo\n", 
        balance.Currency.String(), 
        balance.Balance.Int64())
}
```

### Balance History

```go
result, err := client.TransfersControl.FetchBalanceHistory(ctx)
if err != nil {
    return err
}

for _, entry := range result.Data {
    fmt.Printf("Date: %s, Amount: %d, Reason: %s\n",
        entry.CreatedAt.Time().Format("2006-01-02"),
        entry.Difference.Int64(),
        entry.Reason.String())
}
```

### OTP Management

```go
// Disable OTP (requires OTP)
disableReq := transferscontrol.NewDisableOTPRequestBuilder().Build()
result, err := client.TransfersControl.DisableOTP(ctx, *disableReq)

// Finalize OTP disabling with OTP code
finalizeReq := transferscontrol.NewFinalizeDisableOTPRequestBuilder().
    OTP("123456").
    Build()
result, err := client.TransfersControl.FinalizeDisableOTP(ctx, *finalizeReq)
```

## Use Cases

### Balance Monitoring

```go
func checkSufficientBalance(requiredAmount int64) error {
    result, err := client.TransfersControl.CheckBalance(ctx)
    if err != nil {
        return err
    }
    
    for _, balance := range result.Data {
        if balance.Currency.String() == "NGN" {
            if balance.Balance.Int64() < requiredAmount {
                return fmt.Errorf("insufficient balance: have %d, need %d",
                    balance.Balance.Int64(), requiredAmount)
            }
            break
        }
    }
    
    return nil
}
```
