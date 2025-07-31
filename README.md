# paystack-go

A comprehensive Go client library for the [Paystack API](https://paystack.com/docs/api/).

[![Go Reference](https://pkg.go.dev/badge/github.com/huysamen/paystack-go.svg)](https://pkg.go.dev/github.com/huysamen/paystack-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/huysamen/paystack-go)](https://goreportcard.com/report/github.com/huysamen/paystack-go)

## Features

- ✅ **Transactions**: Initialize, verify, charge authorization, list (with advanced filtering), fetch, export, partial debit, view timeline, totals
- ✅ **Plans**: Create, list, fetch, update plans by ID or code
- ✅ **Customers**: Create, list, fetch, update, validate identity, whitelist/blacklist, authorization management, direct debit
- ✅ **Subscriptions**: Create, list, fetch, enable, disable, generate update links, send management emails
- ✅ **Transfers**: Initiate, finalize, bulk transfers, list, fetch, verify transfer status
- ✅ **Transfer Recipients**: Create, bulk create, list, fetch, update, delete recipient accounts
- ✅ **Subaccounts**: Create, list, fetch, update subaccounts for revenue splitting
- ✅ **Transaction Splits**: Create, list, fetch, update splits, add/remove subaccounts from splits
- ✅ **Settlements**: List settlements, filter by status/date/subaccount, list settlement transactions
- ✅ **Terminal**: Send events, check status, list/fetch/update terminals, commission/decommission devices
- ✅ **Verification**: Resolve account numbers, validate accounts, resolve card BINs
- ✅ **Miscellaneous**: List banks/countries/states for address verification and geographic support
- ✅ **Type Safety**: Strongly typed request/response structures
- ✅ **Error Handling**: Comprehensive error handling with Paystack-specific error types
- ✅ **Configuration**: Support for different environments and custom HTTP clients
- ✅ **Validation**: Request validation for required fields

## Installation

```bash
go get github.com/huysamen/paystack-go
```

## Quick Start

### Basic Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/huysamen/paystack-go"
    "github.com/huysamen/paystack-go/api/transactions"
    "github.com/huysamen/paystack-go/types"
)

func main() {
    // Create a client with your secret key
    client := paystack.DefaultClient("sk_test_your_secret_key_here")

    // Initialize a transaction
    req := &transactions.TransactionInitializeRequest{
        Amount:   50000, // Amount in kobo (500.00 NGN)
        Email:    "customer@example.com",
        Currency: types.CurrencyNGN,
    }

    resp, err := client.Transactions.Initialize(context.Background(), req)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Authorization URL: %s\n", resp.Data.AuthorizationURL)
    fmt.Printf("Reference: %s\n", resp.Data.Reference)
}
```

### Advanced Configuration

```go
package main

import (
    "time"

    "github.com/huysamen/paystack-go"
)

func main() {
    // Create a custom configuration
    config := paystack.NewConfig("sk_test_your_secret_key_here").
        WithTimeout(30 * time.Second).
        WithEnvironment(paystack.EnvironmentProduction)

    // Create client with custom config
    client := paystack.NewClient(config)

    // Use the client...
}
```

### With Custom HTTP Client

```go
package main

import (
    "net/http"
    "time"

    "github.com/huysamen/paystack-go"
)

func main() {
    // Create custom HTTP client
    httpClient := &http.Client{
        Timeout: 15 * time.Second,
        // Add your custom transport, retry logic, etc.
    }

    config := paystack.NewConfig("sk_test_your_secret_key_here").
        WithHTTPClient(httpClient)

    client := paystack.NewClient(config)
}
```

## API Coverage

### Transactions

- **Initialize Transaction**: Create a new transaction with validation
- **Verify Transaction**: Verify a transaction by reference
- **Charge Authorization**: Charge a customer's authorization
- **List Transactions**: List all transactions with advanced filtering (customer, status, date range, amount, etc.)
- **Fetch Transaction**: Fetch a single transaction by ID
- **Export Transactions**: Export transactions to CSV
- **Partial Debit**: Perform partial debit on a customer's account
- **View Timeline**: View transaction timeline by ID or reference
- **Transaction Totals**: Get transaction totals and volume statistics

### Plans

- **Create Plan**: Create a new subscription plan
- **List Plans**: List all plans with filtering options
- **Fetch Plan**: Fetch plan by ID or plan code
- **Update Plan**: Update an existing plan

### Customers

- **Create Customer**: Create a new customer with email and optional details
- **List Customers**: List all customers with date filtering and pagination
- **Fetch Customer**: Get customer details including transactions and authorizations
- **Update Customer**: Update customer information
- **Validate Customer**: Validate customer identity with BVN and bank details
- **Whitelist/Blacklist**: Set customer risk action (allow/deny/default)
- **Initialize Authorization**: Start authorization process for direct debit
- **Verify Authorization**: Check authorization status by reference
- **Initialize Direct Debit**: Link bank account for direct debit transactions
- **Direct Debit Activation**: Trigger activation charge for inactive mandates
- **Fetch Mandate Authorizations**: Get all direct debit mandates for a customer
- **Deactivate Authorization**: Deactivate payment authorization codes

### Subscriptions

- **Create Subscription**: Create recurring payment subscriptions for customers
- **List Subscriptions**: List all subscriptions with customer and plan filtering
- **Fetch Subscription**: Get subscription details including invoices and payment history
- **Enable Subscription**: Reactivate a disabled subscription with email token
- **Disable Subscription**: Temporarily disable a subscription with email token
- **Generate Update Link**: Create secure link for customers to update payment methods
- **Send Update Link**: Email management link directly to customer

### Transfers

- **Initiate Transfer**: Send money to customers from your balance
- **Finalize Transfer**: Complete OTP-protected transfers with verification code
- **Bulk Transfer**: Process multiple transfers in a single batch operation
- **List Transfers**: Get all transfers with recipient and date filtering
- **Fetch Transfer**: Retrieve detailed transfer information by ID or code
- **Verify Transfer**: Check transfer status and details using reference

### Transfer Recipients

- **Create Transfer Recipient**: Create recipients for bank transfers with account validation
- **Bulk Create Recipients**: Create multiple recipients in a single batch operation  
- **List Transfer Recipients**: Get all recipients with date filtering and pagination
- **Fetch Transfer Recipient**: Retrieve detailed recipient information by ID or code
- **Update Transfer Recipient**: Modify recipient name and email address
- **Delete Transfer Recipient**: Deactivate recipient accounts (sets inactive status)

### Subaccounts

- **Create Subaccount**: Create subaccounts for revenue splitting with commission settings
- **List Subaccounts**: Get all subaccounts with date filtering and pagination
- **Fetch Subaccount**: Retrieve detailed subaccount information by ID or code
- **Update Subaccount**: Modify subaccount details, commission rates, and settlement schedules

### Transaction Splits

- **Create Split**: Create transaction splits for automatic revenue distribution
- **List Splits**: Get all splits with filtering by name, active status, and date range
- **Fetch Split**: Retrieve detailed split information by ID or split code
- **Update Split**: Modify split name, active status, and bearer type settings
- **Add Subaccount**: Add or update subaccount shares within a split
- **Remove Subaccount**: Remove subaccounts from existing splits

### Settlements

- **List Settlements**: Get all settlements with status, date, and subaccount filtering
- **List Settlement Transactions**: Retrieve transactions within a specific settlement

### Terminal

- **Send Event**: Send invoice or transaction events to terminal devices
- **Fetch Event Status**: Check delivery status of events sent to terminals
- **Fetch Terminal Status**: Check online and availability status of terminals
- **List Terminals**: Get all terminals available on your integration
- **Fetch Terminal**: Retrieve detailed terminal information by ID
- **Update Terminal**: Modify terminal name and address details
- **Commission Device**: Activate debug devices by linking to your integration
- **Decommission Device**: Unlink debug devices from your integration

### Virtual Terminal

- **Create Virtual Terminal**: Create virtual terminals for in-person payments without POS devices
- **List Virtual Terminals**: Get all virtual terminals with status and search filtering
- **Fetch Virtual Terminal**: Retrieve detailed virtual terminal information by code
- **Update Virtual Terminal**: Modify virtual terminal name and settings
- **Deactivate Virtual Terminal**: Deactivate a virtual terminal
- **Assign Destination**: Add WhatsApp notification destinations to virtual terminals
- **Unassign Destination**: Remove notification destinations from virtual terminals
- **Add Split Code**: Associate transaction splits with virtual terminals
- **Remove Split Code**: Remove split codes from virtual terminals

### Verification

- **Resolve Account**: Get account details by account number and bank code
- **Validate Account**: Perform enhanced account validation with KYC information
- **Resolve Card BIN**: Get card information from Bank Identification Number (BIN)

### Miscellaneous

- **List Banks**: Get banks by country with payment capability filtering
- **List Countries**: Retrieve all countries supported by Paystack
- **List States**: Get states/provinces for address verification purposes

## Error Handling

The library provides structured error handling with Paystack-specific error types:

```go
resp, err := client.Transactions.Initialize(context.Background(), req)
if err != nil {
    if paystackErr, ok := err.(*net.PaystackError); ok {
        fmt.Printf("Paystack API Error: %s (Status: %d)\n", paystackErr.Message, paystackErr.StatusCode)
    } else {
        fmt.Printf("Other error: %v\n", err)
    }
    return
}
```

## Types and Enums

The library includes comprehensive type definitions for all Paystack API entities:

- **Currency**: NGN, USD, GHS, ZAR, KES
- **Channel**: Card, Bank, USSD, QR, Mobile Money, Bank Transfer, EFT
- **CardBrand**: Visa, MasterCard, Verve
- **Interval**: Hourly, Daily, Weekly, Monthly, Quarterly, Biannually, Annually
- **Bearer**: Account, Subaccount

## Examples

### Initialize and Verify Transaction

```go
// Initialize transaction
initReq := &transactions.TransactionInitializeRequest{
    Amount:      100000, // 1000.00 NGN
    Email:       "customer@example.com",
    Currency:    types.CurrencyNGN,
    CallbackURL: "https://yourapp.com/callback",
}

initResp, err := client.Transactions.Initialize(context.Background(), initReq)
if err != nil {
    log.Fatal(err)
}

// Customer pays using the authorization URL
fmt.Printf("Pay here: %s\n", initResp.Data.AuthorizationURL)

// Later, verify the transaction
verifyResp, err := client.Transactions.Verify(context.Background(), initResp.Data.Reference)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Payment status: %s\n", verifyResp.Data.Status)
```

### List Transactions with Advanced Filtering

```go
// List with customer filter and date range
from := time.Now().AddDate(0, -1, 0) // Last month
to := time.Now()
customer := uint64(12345)
status := "success"

listReq := &transactions.TransactionListRequest{
    PerPage:  &[]int{50}[0],
    Page:     &[]int{1}[0],
    Customer: &customer,
    Status:   &status,
    From:     &from,
    To:       &to,
}

resp, err := client.Transactions.List(listReq)
if err != nil {
    log.Fatal(err)
}

for _, transaction := range resp.Data {
    fmt.Printf("Transaction: %s - %s\n", transaction.Reference, transaction.Status)
}
```

### Get Transaction Totals

```go
// Get transaction totals for the last month
from := time.Now().AddDate(0, -1, 0)
to := time.Now()

totalsReq := &transactions.TransactionTotalsRequest{
    From: &from,
    To:   &to,
}

resp, err := client.Transactions.Totals(totalsReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Total Transactions: %d\n", resp.Data.TotalTransactions)
fmt.Printf("Total Volume: %d\n", resp.Data.TotalVolume)
for _, currency := range resp.Data.TotalVolumeByCurrency {
    fmt.Printf("Currency %s: %d\n", currency.Currency.String(), currency.Amount)
}
```

### Create and Manage Plans

```go
// Create a new plan
createReq := &plans.PlanCreateRequest{
    Name:     "Premium Monthly",
    Amount:   2500000, // 25,000.00 NGN in kobo
    Interval: types.IntervalMonthly,
    Currency: types.CurrencyNGN,
    Description: "Premium monthly subscription",
}

createResp, err := client.Plans.Create(createReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Plan created: %s\n", createResp.Data.PlanCode)

// List plans with filtering
listReq := &plans.PlanListRequest{
    PerPage:  &[]int{20}[0],
    Status:   &[]string{"active"}[0],
    Interval: &types.IntervalMonthly,
}

listResp, err := client.Plans.List(listReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Found %d plans\n", len(listResp.Data))

// Update a plan
updateReq := &plans.PlanUpdateRequest{
    Name:     "Premium Monthly (Updated)",
    Amount:   3000000, // 30,000.00 NGN in kobo
    Interval: types.IntervalMonthly,
    Description: "Updated premium monthly subscription",
}

updateResp, err := client.Plans.Update(createResp.Data.PlanCode, updateReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Plan updated: %s\n", updateResp.Data.Message)
```
```

### Export Transactions

```go
from := time.Now().AddDate(0, -1, 0) // Last month
to := time.Now()

exportReq := &transactions.TransactionExportRequest{
    From:    &from,
    To:      &to,
    PerPage: &[]int{100}[0],
}

resp, err := client.Transactions.Export(exportReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Export file: %s\n", resp.Data.Path)
fmt.Printf("Expires at: %s\n", resp.Data.ExpiresAt.Time)
```

### Customer Management

```go
// Create a new customer
createReq := &customers.CustomerCreateRequest{
    Email:     "customer@example.com",
    FirstName: &[]string{"John"}[0],
    LastName:  &[]string{"Doe"}[0],
    Phone:     &[]string{"+2348123456789"}[0],
    Metadata: map[string]interface{}{
        "user_id": "12345",
        "tier":    "premium",
    },
}

createResp, err := client.Customers.Create(createReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Customer created: %s\n", createResp.Data.CustomerCode)

// Fetch customer with related data
fetchResp, err := client.Customers.Fetch(createResp.Data.CustomerCode)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Customer: %s (%s)\n", fetchResp.Data.Email, fetchResp.Data.CustomerCode)
fmt.Printf("Authorizations: %d\n", len(fetchResp.Data.Authorizations))

// Update customer
updateReq := &customers.CustomerUpdateRequest{
    FirstName: &[]string{"Jane"}[0],
    Metadata: map[string]interface{}{
        "user_id": "12345",
        "tier":    "enterprise",
        "updated": true,
    },
}

updateResp, err := client.Customers.Update(createResp.Data.CustomerCode, updateReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Customer updated: %s\n", *updateResp.Data.FirstName)

// Whitelist customer
riskReq := &customers.CustomerRiskActionRequest{
    Customer:   createResp.Data.CustomerCode,
    RiskAction: &[]customers.RiskAction{customers.RiskActionAllow}[0],
}

riskResp, err := client.Customers.SetRiskAction(riskReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Customer whitelisted: %s\n", riskResp.Data.RiskAction)
```

### Subscription Management

```go
// Create a subscription
createReq := &subscriptions.SubscriptionCreateRequest{
    Customer:      "customer@example.com", // or customer code
    Plan:          "PLN_gx2wn530m0i3w3m",  // plan code
    Authorization: &[]string{"AUTH_6tmt288t0o"}[0], // optional specific authorization
    StartDate:     &[]time.Time{time.Now().AddDate(0, 0, 7)}[0], // start in 1 week
}

createResp, err := client.Subscriptions.Create(createReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Subscription: %s (Status: %s)\n", 
    createResp.Data.SubscriptionCode, 
    createResp.Data.Status)

// List subscriptions with filtering
listReq := &subscriptions.SubscriptionListRequest{
    PerPage:  &[]int{20}[0],
    Customer: &[]int{12345}[0], // filter by customer ID
    Plan:     &[]int{67890}[0], // filter by plan ID
}

listResp, err := client.Subscriptions.List(listReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Found %d subscriptions\n", len(listResp.Data.Data))

// Fetch subscription with invoices
fetchResp, err := client.Subscriptions.Fetch(createResp.Data.SubscriptionCode)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Customer: %s\n", fetchResp.Data.Customer.Email)
fmt.Printf("Plan: %s (₦%.2f %s)\n", 
    fetchResp.Data.Plan.Name,
    float64(fetchResp.Data.Plan.Amount)/100,
    fetchResp.Data.Plan.Interval)
fmt.Printf("Invoices: %d\n", len(fetchResp.Data.Invoices))

// Generate update link for customer
updateLinkResp, err := client.Subscriptions.GenerateUpdateLink(createResp.Data.SubscriptionCode)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Update link: %s\n", updateLinkResp.Data.Link)

// Send update link via email
sendResp, err := client.Subscriptions.SendUpdateLink(createResp.Data.SubscriptionCode)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Email sent: %s\n", sendResp.Data.Message)
```

### Transfer Management

```go
// Initiate a single transfer
transferReq := &transfers.TransferInitiateRequest{
    Source:    "balance",
    Amount:    2000000, // ₦20,000 in kobo
    Recipient: "RCP_gx2wn530m0i3w3m", // recipient code
    Reason:    &[]string{"Payment for services"}[0],
    Currency:  &[]string{"NGN"}[0],
    Reference: &[]string{"my-transfer-ref-001"}[0],
}

transferResp, err := client.Transfers.Initiate(transferReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Transfer: %s (Status: %s)\n", 
    transferResp.Data.TransferCode, 
    transferResp.Data.Status)

// Handle OTP if required
if transferResp.Data.Status == "otp" {
    finalizeReq := &transfers.TransferFinalizeRequest{
        TransferCode: transferResp.Data.TransferCode,
        OTP:          "123456", // OTP from phone
    }
    
    finalizeResp, err := client.Transfers.Finalize(finalizeReq)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Transfer finalized: %s\n", finalizeResp.Data.Status)
}

// Bulk transfer for multiple recipients
bulkReq := &transfers.BulkTransferRequest{
    Source:   "balance",
    Currency: &[]string{"NGN"}[0],
    Transfers: []transfers.BulkTransferItem{
        {
            Amount:    500000, // ₦5,000
            Reference: "bulk-001",
            Reason:    "Salary payment",
            Recipient: "RCP_recipient_1",
        },
        {
            Amount:    750000, // ₦7,500
            Reference: "bulk-002",
            Reason:    "Bonus payment",
            Recipient: "RCP_recipient_2",
        },
    },
}

bulkResp, err := client.Transfers.Bulk(bulkReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Bulk transfer: %d transfers processed\n", len(bulkResp.Data))

// List and verify transfers
listReq := &transfers.TransferListRequest{
    PerPage: &[]int{20}[0],
    Page:    &[]int{1}[0],
}

listResp, err := client.Transfers.List(listReq)
if err != nil {
    log.Fatal(err)
}

for _, transfer := range listResp.Data.Data {
    fmt.Printf("%s: ₦%.2f (%s)\n", 
        transfer.TransferCode,
        float64(transfer.Amount)/100,
        transfer.Status)
}
```

### Transfer Recipients Management

```go
// Create a transfer recipient
recipientReq := &transfer_recipients.TransferRecipientCreateRequest{
    Type:          transfer_recipients.RecipientTypeNuban,
    Name:          "John Doe",
    AccountNumber: "0123456789",
    BankCode:      "058", // GTBank
    Currency:      &[]string{"NGN"}[0],
    Description:   &[]string{"Employee salary account"}[0],
    Metadata: map[string]interface{}{
        "employee_id": "EMP001",
        "department":  "Engineering",
    },
}

recipientResp, err := client.TransferRecipients.Create(recipientReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Recipient: %s (%s)\n", 
    recipientResp.Data.Name, 
    recipientResp.Data.RecipientCode)

// Bulk create multiple recipients
bulkReq := &transfer_recipients.BulkCreateTransferRecipientRequest{
    Batch: []transfer_recipients.BulkRecipientItem{
        {
            Type:          transfer_recipients.RecipientTypeNuban,
            Name:          "Alice Johnson",
            AccountNumber: "9876543210",
            BankCode:      "044", // Access Bank
            Currency:      &[]string{"NGN"}[0],
            Metadata: map[string]interface{}{
                "employee_id": "EMP002",
                "department":  "Marketing",
            },
        },
        {
            Type:          transfer_recipients.RecipientTypeNuban,
            Name:          "Bob Wilson",
            AccountNumber: "1122334455",
            BankCode:      "033", // UBA
            Currency:      &[]string{"NGN"}[0],
            Metadata: map[string]interface{}{
                "employee_id": "EMP003",
                "department":  "Finance",
            },
        },
    },
}

bulkResp, err := client.TransferRecipients.BulkCreate(bulkReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Created: %d recipients\n", len(bulkResp.Data.Success))
fmt.Printf("Failed: %d recipients\n", len(bulkResp.Data.Errors))

// List and filter recipients
listReq := &transfer_recipients.TransferRecipientListRequest{
    PerPage: &[]int{20}[0],
    Page:    &[]int{1}[0],
}

listResp, err := client.TransferRecipients.List(listReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Total recipients: %d\n", len(listResp.Data))
for _, recipient := range listResp.Data {
    fmt.Printf("- %s (%s) - %s\n", 
        recipient.Name, 
        recipient.RecipientCode,
        recipient.Type.String())
}

// Update recipient details
updateReq := &transfer_recipients.TransferRecipientUpdateRequest{
    Name:  "John Smith Doe",
    Email: &[]string{"john.doe@company.com"}[0],
}

updateResp, err := client.TransferRecipients.Update(recipientResp.Data.RecipientCode, updateReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Updated: %s\n", updateResp.Data.Name)

// Delete (deactivate) recipient
deleteResp, err := client.TransferRecipients.Delete(recipientResp.Data.RecipientCode)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Status: %s\n", deleteResp.Message)
```

### Subaccounts Management

```go
// Create a subaccount for revenue splitting
subaccountReq := &subaccounts.SubaccountCreateRequest{
    BusinessName:        "Vendor Store",
    BankCode:            "058", // GTBank
    AccountNumber:       "0123456789",
    PercentageCharge:    20.0, // 20% goes to main account
    Description:         &[]string{"Online electronics store"}[0],
    PrimaryContactName:  &[]string{"John Doe"}[0],
    PrimaryContactEmail: &[]string{"john@vendorstore.com"}[0],
    PrimaryContactPhone: &[]string{"+2348123456789"}[0],
    Metadata: map[string]interface{}{
        "vendor_id": "VND001",
        "category":  "electronics",
        "tier":      "premium",
    },
}

subaccountResp, err := client.Subaccounts.Create(subaccountReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Subaccount: %s (%s)\n", 
    subaccountResp.Data.BusinessName, 
    subaccountResp.Data.SubaccountCode)
fmt.Printf("Commission: %.1f%%\n", subaccountResp.Data.PercentageCharge)

// List subaccounts with filtering
listReq := &subaccounts.SubaccountListRequest{
    PerPage: &[]int{20}[0],
    Page:    &[]int{1}[0],
}

listResp, err := client.Subaccounts.List(listReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Total subaccounts: %d\n", len(listResp.Data))
for _, sub := range listResp.Data {
    status := "Active"
    if !sub.Active {
        status = "Inactive"
    }
    fmt.Printf("- %s (%s) - %.1f%% - %s\n", 
        sub.BusinessName, 
        sub.SubaccountCode,
        sub.PercentageCharge,
        status)
}

// Update subaccount details
updateReq := &subaccounts.SubaccountUpdateRequest{
    BusinessName:     "Updated Vendor Store",
    Description:      "Premium electronics vendor",
    PercentageCharge: &[]float64{15.0}[0], // Reduced commission
    SettlementSchedule: &[]subaccounts.SettlementSchedule{
        subaccounts.SettlementScheduleWeekly,
    }[0],
    Metadata: map[string]interface{}{
        "tier":        "platinum",
        "updated_at":  time.Now().Format(time.RFC3339),
    },
}

updateResp, err := client.Subaccounts.Update(subaccountResp.Data.SubaccountCode, updateReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Updated: %s\n", updateResp.Data.BusinessName)
fmt.Printf("New commission: %.1f%%\n", updateResp.Data.PercentageCharge)

// Deactivate subaccount
active := false
deactivateReq := &subaccounts.SubaccountUpdateRequest{
    BusinessName: updateResp.Data.BusinessName,
    Description:  "Deactivated due to policy violation",
    Active:       &active,
}

_, err = client.Subaccounts.Update(subaccountResp.Data.SubaccountCode, deactivateReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Subaccount deactivated\n")
```

### Transaction Splits

The transaction splits API enables merchants to split settlement for a transaction across their payout account and one or more subaccounts.

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/huysamen/paystack-go"
    transaction_splits "github.com/huysamen/paystack-go/api/transaction-splits"
    "github.com/huysamen/paystack-go/types"
)

func main() {
    client := paystack.DefaultClient("sk_test_your_secret_key_here")

    // Create a transaction split
    createReq := &transaction_splits.TransactionSplitCreateRequest{
        Name:     "Revenue Split",
        Type:     transaction_splits.TransactionSplitTypePercentage,
        Currency: types.CurrencyNGN,
        Subaccounts: []transaction_splits.TransactionSplitSubaccount{
            {
                Subaccount: "ACCT_xxxxxxxxxx",
                Share:      70, // 70% to this subaccount
            },
            {
                Subaccount: "ACCT_yyyyyyyyyy", 
                Share:      30, // 30% to this subaccount
            },
        },
        BearerType: &[]transaction_splits.TransactionSplitBearerType{
            transaction_splits.TransactionSplitBearerTypeAllProportional,
        }[0],
    }

    createResp, err := client.TransactionSplits.Create(createReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Split created: %s\n", createResp.Data.SplitCode)

    // List transaction splits
    listReq := &transaction_splits.TransactionSplitListRequest{
        Active:  &[]bool{true}[0],
        PerPage: &[]int{20}[0],
    }

    listResp, err := client.TransactionSplits.List(listReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d splits\n", len(listResp.Data))

    // Fetch a specific split
    fetchResp, err := client.TransactionSplits.Fetch(createResp.Data.SplitCode)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Split: %s (%s)\n", fetchResp.Data.Name, fetchResp.Data.Type)

    // Update split
    updateReq := &transaction_splits.TransactionSplitUpdateRequest{
        Name: &[]string{"Updated Revenue Split"}[0],
        BearerType: &[]transaction_splits.TransactionSplitBearerType{
            transaction_splits.TransactionSplitBearerTypeAccount,
        }[0],
    }

    updateResp, err := client.TransactionSplits.Update(createResp.Data.SplitCode, updateReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Split updated: %s\n", updateResp.Data.Name)

    // Add subaccount to split
    addReq := &transaction_splits.TransactionSplitSubaccountAddRequest{
        Subaccount: "ACCT_zzzzzzzzzz",
        Share:      15,
    }

    addResp, err := client.TransactionSplits.AddSubaccount(createResp.Data.SplitCode, addReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Subaccount added. Total subaccounts: %d\n", addResp.Data.TotalSubaccounts)

    // Remove subaccount from split
    removeReq := &transaction_splits.TransactionSplitSubaccountRemoveRequest{
        Subaccount: "ACCT_zzzzzzzzzz",
    }

    _, err = client.TransactionSplits.RemoveSubaccount(createResp.Data.SplitCode, removeReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Subaccount removed from split\n")
}
```

### Settlements

The settlements API allows you to track payouts and settlement records.

```go
package main

import (
    "fmt"
    "log"
    "time"
    
    "github.com/huysamen/paystack-go"
    "github.com/huysamen/paystack-go/api/settlements"
)

func main() {
    client := paystack.DefaultClient("sk_test_your_secret_key_here")

    // List all settlements
    perPage := 50
    page := 1
    listReq := &settlements.SettlementListRequest{
        PerPage: &perPage,
        Page:    &page,
    }

    resp, err := client.Settlements.List(listReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d settlements\n", len(resp.Data))
    
    // List settlements with filters
    status := settlements.SettlementStatusSuccess
    thirtyDaysAgo := time.Now().AddDate(0, 0, -30)
    now := time.Now()
    
    filteredReq := &settlements.SettlementListRequest{
        Status:  &status,
        From:    &thirtyDaysAgo,
        To:      &now,
        PerPage: &perPage,
        Page:    &page,
    }

    filteredResp, err := client.Settlements.List(filteredReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d successful settlements in last 30 days\n", len(filteredResp.Data))

    // Get settlement transactions
    if len(resp.Data) > 0 {
        settlementID := fmt.Sprintf("%d", resp.Data[0].ID)
        
        txReq := &settlements.SettlementTransactionListRequest{
            PerPage: &perPage,
            Page:    &page,
        }

        txResp, err := client.Settlements.ListTransactions(settlementID, txReq)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("Found %d transactions in settlement %s\n", len(txResp.Data), settlementID)
    }
}
```

### Terminal

The terminal API allows you to build delightful in-person payment experiences with Paystack terminal devices.

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/huysamen/paystack-go"
    "github.com/huysamen/paystack-go/api/terminal"
)

func main() {
    client := paystack.DefaultClient("sk_test_your_secret_key_here")

    // List terminals
    listReq := &terminal.TerminalListRequest{
        PerPage: &[]int{10}[0],
    }

    listResp, err := client.Terminal.List(listReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d terminals\n", len(listResp.Data))

    if len(listResp.Data) > 0 {
        terminalID := listResp.Data[0].TerminalID

        // Fetch terminal details
        fetchResp, err := client.Terminal.Fetch(terminalID)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("Terminal: %s (Status: %s)\n", fetchResp.Data.Name, fetchResp.Data.Status)

        // Check terminal status
        statusResp, err := client.Terminal.FetchTerminalStatus(terminalID)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("Online: %v, Available: %v\n", statusResp.Data.Online, statusResp.Data.Available)

        // Send invoice event to terminal
        eventReq := &terminal.TerminalSendEventRequest{
            Type:   terminal.TerminalEventTypeInvoice,
            Action: terminal.TerminalEventActionProcess,
            Data: terminal.TerminalEventData{
                "id":        123456,
                "reference": 4634337895939,
            },
        }

        eventResp, err := client.Terminal.SendEvent(terminalID, eventReq)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("Event sent: %s\n", eventResp.Data.ID)

        // Check event delivery status
        eventStatusResp, err := client.Terminal.FetchEventStatus(terminalID, eventResp.Data.ID)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("Event delivered: %v\n", eventStatusResp.Data.Delivered)

        // Update terminal
        updateReq := &terminal.TerminalUpdateRequest{
            Name:    &[]string{"Updated Terminal Name"}[0],
            Address: &[]string{"New Address"}[0],
        }

        _, err = client.Terminal.Update(terminalID, updateReq)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("Terminal updated\n")
    }

    // Commission a device
    commissionReq := &terminal.TerminalCommissionRequest{
        SerialNumber: "1111150412230003899",
    }

    _, err = client.Terminal.CommissionDevice(commissionReq)
    if err != nil {
        log.Printf("Commission error (expected): %v\n", err)
    }
}
```

### Virtual Terminal

The virtual terminal API allows you to accept in-person payments without a POS device using WhatsApp notifications.

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/huysamen/paystack-go"
    "github.com/huysamen/paystack-go/api/virtual-terminal"
    "github.com/huysamen/paystack-go/types"
)

func main() {
    client := paystack.DefaultClient("sk_test_your_secret_key_here")
    ctx := context.Background()

    // Create virtual terminal
    createReq := &virtualterminal.CreateVirtualTerminalRequest{
        Name: "Sales Point #1",
        Destinations: []virtualterminal.VirtualTerminalDestination{
            {
                Target: "+2347012345678",
                Name:   "Sales Rep",
            },
        },
        Currency: "NGN",
        CustomFields: []virtualterminal.CustomField{
            {
                DisplayName:  "Customer ID",
                VariableName: "customer_id",
            },
        },
        Metadata: &types.Metadata{
            "department": "sales",
            "location":   "lagos",
        },
    }

    terminal, err := client.VirtualTerminal.Create(ctx, createReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Created virtual terminal: %s (Code: %s)\n", terminal.Name, terminal.Code)

    // List virtual terminals
    listReq := &virtualterminal.ListVirtualTerminalsRequest{
        Status:  "active",
        PerPage: 10,
    }

    terminals, err := client.VirtualTerminal.List(ctx, listReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d virtual terminals\n", len(terminals.Data))

    // Fetch virtual terminal
    fetchedTerminal, err := client.VirtualTerminal.Fetch(ctx, terminal.Code)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Terminal details: %s (Active: %t)\n", fetchedTerminal.Name, fetchedTerminal.Active)

    // Update virtual terminal
    updateReq := &virtualterminal.UpdateVirtualTerminalRequest{
        Name: "Updated Sales Point #1",
    }

    updatedTerminal, err := client.VirtualTerminal.Update(ctx, terminal.Code, updateReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Updated terminal: %s\n", updatedTerminal.Name)

    // Assign additional destination
    assignReq := &virtualterminal.AssignDestinationRequest{
        Destinations: []virtualterminal.VirtualTerminalDestination{
            {
                Target: "+2348012345678",
                Name:   "Manager",
            },
        },
    }

    destinations, err := client.VirtualTerminal.AssignDestination(ctx, terminal.Code, assignReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Assigned %d destinations\n", len(*destinations))

    // Add split code (optional)
    addSplitReq := &virtualterminal.AddSplitCodeRequest{
        SplitCode: "SPL_98WF13Zu8w5", // Replace with actual split code
    }

    _, err = client.VirtualTerminal.AddSplitCode(ctx, terminal.Code, addSplitReq)
    if err != nil {
        log.Printf("Split code error (expected): %v\n", err)
    }

    // Deactivate virtual terminal
    deactivateResult, err := client.VirtualTerminal.Deactivate(ctx, terminal.Code)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Terminal deactivated: %s\n", deactivateResult.Message)
}
```

### Verification

The verification API provides account validation and card intelligence functionality.

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/huysamen/paystack-go"
    "github.com/huysamen/paystack-go/api/verification"
)

func main() {
    client := paystack.DefaultClient("sk_test_your_secret_key_here")

    // Resolve account number
    resolveReq := &verification.AccountResolveRequest{
        AccountNumber: "0022728151",
        BankCode:      "063",
    }

    resolveResp, err := client.Verification.ResolveAccount(resolveReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Account holder: %s\n", resolveResp.Data.AccountName)

    // Validate account with KYC
    validateReq := &verification.AccountValidateRequest{
        AccountName:    "John Doe",
        AccountNumber:  "0123456789",
        AccountType:    "personal",
        BankCode:       "058",
        CountryCode:    "NG",
        DocumentType:   "bvn",
        DocumentNumber: &[]string{"12345678901"}[0],
    }

    validateResp, err := client.Verification.ValidateAccount(validateReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Verified: %v\n", validateResp.Data.Verified)

    // Resolve card BIN
    binResp, err := client.Verification.ResolveCardBIN("539983")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Card: %s %s from %s\n", 
        binResp.Data.Brand, binResp.Data.CardType, binResp.Data.Bank)
}
```

### Miscellaneous

The miscellaneous API provides supporting functionality like bank listings and geographic information.

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/huysamen/paystack-go"
    "github.com/huysamen/paystack-go/api/miscellaneous"
)

func main() {
    client := paystack.DefaultClient("sk_test_your_secret_key_here")

    // List banks for a specific country
    country := "nigeria"
    perPage := 50
    banksReq := &miscellaneous.BankListRequest{
        Country: &country,
        PerPage: &perPage,
    }

    banksResp, err := client.Miscellaneous.ListBanks(banksReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d banks in %s\n", len(banksResp.Data), country)

    // List supported countries
    countriesResp, err := client.Miscellaneous.ListCountries()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Paystack supports %d countries\n", len(countriesResp.Data))

    // List states for address verification
    statesReq := &miscellaneous.StateListRequest{
        Country: "US",
    }

    statesResp, err := client.Miscellaneous.ListStates(statesReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d states for address verification\n", len(statesResp.Data))
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For issues related to this library, please create an issue on GitHub.
For Paystack API support, please contact [Paystack Support](https://paystack.com/support).
