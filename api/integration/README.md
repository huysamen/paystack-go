# Integration API

Manage your Paystack integration settings and retrieve account information.

## Available Operations

- **Fetch Payment Session Timeout** - Get session timeout settings
- **Update Payment Session Timeout** - Update timeout settings

## Quick Examples

### Fetch Timeout Settings

```go
import "github.com/huysamen/paystack-go/api/integration"

result, err := client.Integration.FetchPaymentSessionTimeout(ctx)
if err != nil {
    return err
}

timeout := result.Data
fmt.Printf("Payment session timeout: %d seconds\n", 
    timeout.PaymentSessionTimeout.Int64())
```

### Update Timeout

```go
request := integration.NewUpdatePaymentSessionTimeoutRequestBuilder().
    TimeoutInSeconds(1800).  // 30 minutes
    Build()

result, err := client.Integration.UpdatePaymentSessionTimeout(ctx, *request)
```

## Use Cases

Configure payment session timeouts for your integration to balance security and user experience.
