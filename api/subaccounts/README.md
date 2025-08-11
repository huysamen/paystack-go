# Subaccounts API

Create and manage subaccounts for split payments and marketplace scenarios.

## Available Operations

- **Create** - Create a new subaccount
- **List** - List subaccounts
- **Fetch** - Get subaccount details
- **Update** - Update subaccount information

## Quick Examples

### Create Subaccount

```go
import "github.com/huysamen/paystack-go/api/subaccounts"

request := subaccounts.NewCreateRequestBuilder().
    BusinessName("Vendor Store").
    SettlementBank("044").  // Access Bank
    AccountNumber("0123456789").
    PercentageCharge(2.5).  // 2.5% commission
    Description("Electronics vendor").
    PrimaryContactEmail("vendor@store.com").
    PrimaryContactName("John Vendor").
    PrimaryContactPhone("+2348012345678").
    Build()

result, err := client.Subaccounts.Create(ctx, *request)
```

### List Subaccounts

```go
query := subaccounts.NewListRequestBuilder().
    PerPage(50).
    Page(1).
    Build()

result, err := client.Subaccounts.List(ctx, *query)
```

## Use Cases

### Marketplace Platform

```go
// Create subaccount for each vendor
request := subaccounts.NewCreateRequestBuilder().
    BusinessName(vendor.BusinessName).
    SettlementBank(vendor.BankCode).
    AccountNumber(vendor.AccountNumber).
    PercentageCharge(platformCommission).
    Description(fmt.Sprintf("Vendor: %s", vendor.Name)).
    Build()
```

### Revenue Sharing

```go
// Service provider gets 80%, platform keeps 20%
request := subaccounts.NewCreateRequestBuilder().
    BusinessName("Service Provider").
    PercentageCharge(20.0).  // Platform commission
    Build()
```
