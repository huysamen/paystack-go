# Transfer Recipients API

Manage bank account recipients for money transfers and payouts.

## Available Operations

- **Create** - Add a new transfer recipient
- **List** - List recipients with filtering
- **Fetch** - Get recipient details
- **Update** - Update recipient information  
- **Delete** - Remove a recipient

## Quick Examples

### Create Recipient

```go
import "github.com/huysamen/paystack-go/api/transferrecipients"

request := transferrecipients.NewCreateRequestBuilder().
    Type(enums.Nuban).  // Nigerian bank account
    Name("John Doe").
    AccountNumber("0123456789").
    BankCode("044").  // Access Bank
    Currency(enums.NGN).
    Description("Freelancer payment").
    Build()

result, err := client.TransferRecipients.Create(ctx, *request)
if err != nil {
    return err
}

recipient := result.Data
fmt.Printf("Recipient created: %s (Code: %s)\n", 
    recipient.Name.String(), 
    recipient.RecipientCode.String())
```

### List Recipients

```go
query := transferrecipients.NewListRequestBuilder().
    PerPage(50).
    Page(1).
    Type(enums.Nuban).
    Build()

result, err := client.TransferRecipients.List(ctx, *query)
```

### Update Recipient

```go
request := transferrecipients.NewUpdateRequestBuilder("RCP_recipient_code").
    Name("John Smith").  // Updated name
    Email("john.smith@example.com").
    Build()

result, err := client.TransferRecipients.Update(ctx, *request)
```

## Recipient Types

- **nuban** - Nigerian bank accounts
- **mobile_money** - Mobile money wallets
- **basa** - BASA (Bank Account and Salary Account)

## Best Practices

### Bank Account Validation

```go
// Verify account details before creating recipient
verifyRequest := verification.NewResolveAccountRequestBuilder().
    AccountNumber("0123456789").
    BankCode("044").
    Build()

verifyResult, err := client.Verification.ResolveAccount(ctx, *verifyRequest)
if err != nil {
    return fmt.Errorf("account verification failed: %w", err)
}

// Create recipient with verified details
request := transferrecipients.NewCreateRequestBuilder().
    Name(verifyResult.Data.AccountName.String()).
    AccountNumber("0123456789").
    BankCode("044").
    Build()
```

### Recipient Management

```go
// Store recipient codes for future transfers
type RecipientInfo struct {
    Code        string
    Name        string
    AccountNum  string
    BankCode    string
    CreatedAt   time.Time
}

func saveRecipient(recipient *types.Recipient) error {
    info := RecipientInfo{
        Code:       recipient.RecipientCode.String(),
        Name:       recipient.Name.String(),
        AccountNum: recipient.Details.AccountNumber.String(),
        BankCode:   recipient.Details.BankCode.String(),
        CreatedAt:  time.Now(),
    }
    
    return database.SaveRecipient(info)
}
```
