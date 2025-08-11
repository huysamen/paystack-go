# Dedicated Virtual Accounts API

Create and manage dedicated virtual account numbers for customers.

## Available Operations

- **Create** - Create a dedicated virtual account
- **List** - List virtual accounts
- **Fetch** - Get virtual account details
- **Requery** - Requery account transactions
- **Deactivate** - Deactivate an account
- **Split** - Configure split payments

## Quick Examples

### Create Virtual Account

```go
import "github.com/huysamen/paystack-go/api/dedicatedvirtualaccounts"

request := dedicatedvirtualaccounts.NewCreateRequestBuilder().
    Customer("customer@example.com").
    PreferredBank("wema-bank").  // Optional bank preference
    Build()

result, err := client.DedicatedVirtualAccounts.Create(ctx, *request)
if err != nil {
    return err
}

account := result.Data
fmt.Printf("Virtual Account: %s - %s\n", 
    account.AccountNumber.String(), 
    account.BankName.String())
```

### List Virtual Accounts

```go
query := dedicatedvirtualaccounts.NewListRequestBuilder().
    PerPage(50).
    Page(1).
    Active(true).
    Currency(enums.NGN).
    Build()

result, err := client.DedicatedVirtualAccounts.List(ctx, *query)
```

### Deactivate Account

```go
request := dedicatedvirtualaccounts.NewDeactivateRequestBuilder().
    DedicatedAccountID(12345).
    Build()

result, err := client.DedicatedVirtualAccounts.Deactivate(ctx, *request)
```

## Use Cases

### Customer Wallet System

```go
// Create virtual account for each customer
func createCustomerWallet(customerEmail string) (*types.DedicatedAccount, error) {
    request := dedicatedvirtualaccounts.NewCreateRequestBuilder().
        Customer(customerEmail).
        Build()
        
    result, err := client.DedicatedVirtualAccounts.Create(ctx, *request)
    if err != nil {
        return nil, err
    }
    
    if err := result.Err(); err != nil {
        return nil, err
    }
    
    return &result.Data, nil
}
```

### Split Payments

```go
request := dedicatedvirtualaccounts.NewSplitRequestBuilder().
    DedicatedAccountID(12345).
    SubAccount("ACCT_subaccount_code").
    SplitCode("SPL_split_code").
    Build()

result, err := client.DedicatedVirtualAccounts.Split(ctx, *request)
```

Perfect for creating unique bank account numbers for each customer, enabling easy bank transfer payments.
