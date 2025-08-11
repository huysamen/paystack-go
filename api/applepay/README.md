# Apple Pay API

Manage Apple Pay integration and domain registration.

## Available Operations

- **Register Domain** - Register domain for Apple Pay
- **List Domains** - List registered domains
- **Unregister Domain** - Remove domain registration

## Quick Examples

### Register Domain

```go
import "github.com/huysamen/paystack-go/api/applepay"

request := applepay.NewRegisterDomainRequestBuilder().
    DomainName("shop.yourstore.com").
    Build()

result, err := client.ApplePay.RegisterDomain(ctx, *request)
if err != nil {
    return err
}

if err := result.Err(); err != nil {
    return fmt.Errorf("domain registration failed: %w", err)
}

fmt.Printf("Domain registered successfully\n")
```

### List Domains

```go
query := applepay.NewListDomainsRequestBuilder().
    UseCursor(false).
    Build()

result, err := client.ApplePay.ListDomains(ctx, *query)
if err != nil {
    return err
}

for _, domain := range result.Data {
    fmt.Printf("Domain: %s\n", domain.DomainName.String())
}
```

### Unregister Domain

```go
request := applepay.NewUnregisterDomainRequestBuilder().
    DomainName("old.yourstore.com").
    Build()

result, err := client.ApplePay.UnregisterDomain(ctx, *request)
```

## Use Cases

Required for integrating Apple Pay on your website. Register all domains where Apple Pay will be used.
