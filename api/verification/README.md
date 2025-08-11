# Verification API

Verify bank accounts, BVN, and other identity information.

## Available Operations

- **Resolve Account** - Verify bank account details
- **Validate Account** - Validate account ownership
- **Resolve BIN** - Get card BIN information

## Quick Examples

### Resolve Bank Account

```go
import "github.com/huysamen/paystack-go/api/verification"

request := verification.NewResolveAccountRequestBuilder().
    AccountNumber("0123456789").
    BankCode("044").  // Access Bank
    Build()

result, err := client.Verification.ResolveAccount(ctx, *request)
if err != nil {
    return err
}

if err := result.Err(); err != nil {
    return fmt.Errorf("account verification failed: %w", err)
}

account := result.Data
fmt.Printf("Account Name: %s\n", account.AccountName.String())
fmt.Printf("Account Number: %s\n", account.AccountNumber.String())
```

### Validate Account

```go
request := verification.NewValidateAccountRequestBuilder().
    AccountName("John Doe").
    AccountNumber("0123456789").
    AccountType("personal").
    BankCode("044").
    CountryCode("NG").
    DocumentType("identityNumber").
    DocumentNumber("12345678901").
    Build()

result, err := client.Verification.ValidateAccount(ctx, *request)
```

### Resolve Card BIN

```go
result, err := client.Verification.ResolveBIN(ctx, "539983")
if err != nil {
    return err
}

bin := result.Data  
fmt.Printf("Card Type: %s\n", bin.CardType.String())
fmt.Printf("Bank: %s\n", bin.Bank.String())
fmt.Printf("Country: %s\n", bin.CountryName.String())
```

## Use Cases

### Pre-Transfer Verification

```go
func verifyRecipientAccount(accountNumber, bankCode string) error {
    request := verification.NewResolveAccountRequestBuilder().
        AccountNumber(accountNumber).
        BankCode(bankCode).
        Build()
        
    result, err := client.Verification.ResolveAccount(ctx, *request)
    if err != nil {
        return fmt.Errorf("verification failed: %w", err)
    }
    
    if err := result.Err(); err != nil {
        return fmt.Errorf("invalid account: %w", err)
    }
    
    fmt.Printf("✓ Account verified: %s\n", result.Data.AccountName.String())
    return nil
}
```

### Card Validation

```go
func validateCard(bin string) (*types.BINData, error) {
    result, err := client.Verification.ResolveBIN(ctx, bin)
    if err != nil {
        return nil, err
    }
    
    if err := result.Err(); err != nil {
        return nil, err
    }
    
    return &result.Data, nil
}
```

Essential for preventing fraud and ensuring accurate recipient information before processing transfers.
