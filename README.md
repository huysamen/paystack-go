# paystack-go

A comprehensive Go client library for the [Paystack API](https://paystack.com/docs/api/).

[![Go Reference](https://pkg.go.dev/badge/github.com/huysamen/paystack-go.svg)](https://pkg.go.dev/github.com/huysamen/paystack-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/huysamen/paystack-go)](https://goreportcard.com/report/github.com/huysamen/paystack-go)

## Features

- ✅ **Transactions**: Initialize, verify, charge authorization, list (with advanced filtering), fetch, export, partial debit, view timeline, totals
- ✅ **Plans**: Create, list, fetch, update plans by ID or code
- ✅ **Products**: Create, list, fetch, update products with inventory management and metadata support
- ✅ **Customers**: Create, list, fetch, update, validate identity, whitelist/blacklist, authorization management, direct debit
- ✅ **Subscriptions**: Create, list, fetch, enable, disable, generate update links, send management emails
- ✅ **Transfers**: Initiate, finalize, bulk transfers, list, fetch, verify transfer status
- ✅ **Transfer Recipients**: Create, bulk create, list, fetch, update, delete recipient accounts
- ✅ **Subaccounts**: Create, list, fetch, update subaccounts for revenue splitting
- ✅ **Transaction Splits**: Create, list, fetch, update splits, add/remove subaccounts from splits
- ✅ **Settlements**: List settlements, filter by status/date/subaccount, list settlement transactions
- ✅ **Terminal**: Send events, check status, list/fetch/update terminals, commission/decommission devices
- ✅ **Virtual Terminal**: Create virtual terminals, manage settings, assign destinations, handle splits
- ✅ **Direct Debit**: Manage mandate authorizations for direct debit payments
- ✅ **Dedicated Virtual Account**: Create and manage dedicated virtual accounts for unique customer payments
- ✅ **Apple Pay**: Register and manage domains for Apple Pay integration
- ✅ **Integration**: Manage payment session timeout settings and integration configuration
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

### Products

- **Create Product**: Create products with inventory management, pricing, and metadata
- **List Products**: List all products with pagination and date filtering
- **Fetch Product**: Get detailed product information by ID or product code
- **Update Product**: Update product details, pricing, and inventory levels

### Payment Pages

- **Create Payment Page**: Create secure payment pages for one-time payments, subscriptions, or products
- **List Payment Pages**: List all payment pages with pagination and date filtering
- **Fetch Payment Page**: Get detailed payment page information by ID or slug
- **Update Payment Page**: Update page details, pricing, and activation status
- **Check Slug Availability**: Verify if a custom slug is available for use
- **Add Products**: Associate existing products with payment pages for product showcases

### Payment Requests

- **Create Payment Request**: Create payment requests with line items, taxes, and due dates
- **List Payment Requests**: List all payment requests with status and customer filtering
- **Fetch Payment Request**: Get detailed payment request information by ID or code
- **Verify Payment Request**: Verify payment request status and payment details
- **Update Payment Request**: Update request details, line items, and due dates
- **Finalize Payment Request**: Convert draft payment requests to active invoices
- **Archive Payment Request**: Archive payment requests to hide from lists
- **Send Notification**: Send email notifications to customers about payment requests
- **Get Totals**: Retrieve payment request analytics and totals by status and currency

### Transfer Control

- **Check Balance**: Fetch available balance on your integration by currency
- **Fetch Balance Ledger**: Get detailed ledger of all pay-ins and pay-outs with reasons
- **Resend OTP**: Generate and resend OTP for transfer verification
- **Disable OTP**: Initiate process to disable OTP requirement for transfers
- **Finalize Disable OTP**: Complete OTP disabling with verification code
- **Enable OTP**: Re-enable OTP requirement for transfer security

### Bulk Charges

- **Initiate Bulk Charge**: Process multiple charges in a single batch with authorization codes
- **List Bulk Charge Batches**: Get all bulk charge batches with filtering and pagination
- **Fetch Bulk Charge Batch**: Retrieve specific batch details including progress status
- **Fetch Charges in a Batch**: Get detailed charge information within a specific batch
- **Pause Bulk Charge Batch**: Temporarily halt processing of an active batch
- **Resume Bulk Charge Batch**: Resume processing of a paused batch

### Integration

- **Fetch Timeout**: Retrieve the current payment session timeout setting on your integration
- **Update Timeout**: Configure payment session timeout duration (0 to disable timeouts)

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

### Direct Debit

- **List Mandate Authorizations**: Get all direct debit mandate authorizations with status filtering
- **Trigger Activation Charge**: Trigger activation charges on pending mandates for multiple customers

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

### Dedicated Virtual Account

- **Create Dedicated Virtual Account**: Create dedicated virtual accounts for existing customers
- **Assign Dedicated Virtual Account**: Create customer and assign dedicated virtual account in one call
- **List Dedicated Virtual Accounts**: Get all dedicated virtual accounts with filtering by status, currency, provider
- **Fetch Dedicated Virtual Account**: Retrieve detailed dedicated virtual account information by ID
- **Requery Dedicated Account**: Check account status and transaction updates from providers
- **Deactivate Dedicated Virtual Account**: Deactivate a dedicated virtual account
- **Split Transaction**: Add transaction splits to dedicated virtual accounts for automatic fund distribution
- **Remove Split**: Remove splits from dedicated virtual accounts
- **Fetch Bank Providers**: Get available bank providers for dedicated virtual account creation

### Apple Pay

- **Register Domain**: Register top-level domains or subdomains for Apple Pay integration
- **List Domains**: Get all registered Apple Pay domains with optional cursor pagination
- **Unregister Domain**: Remove previously registered domains from Apple Pay integration

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

### Direct Debit

The direct debit API allows you to manage authorization on your customer's bank accounts.

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/huysamen/paystack-go"
    "github.com/huysamen/paystack-go/api/direct-debit"
)

func main() {
    client := paystack.DefaultClient("sk_test_your_secret_key_here")
    ctx := context.Background()

    // List active mandate authorizations
    listReq := &directdebit.ListMandateAuthorizationsRequest{
        Status:  directdebit.MandateAuthorizationStatusActive,
        PerPage: 10,
    }

    mandates, err := client.DirectDebit.ListMandateAuthorizations(ctx, listReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d active mandate authorizations\n", len(mandates.Data))
    for _, mandate := range mandates.Data {
        fmt.Printf("- %s: %s (%s)\n", 
            mandate.Customer.Email, mandate.BankName, mandate.Status)
    }

    // List pending mandates
    pendingReq := &directdebit.ListMandateAuthorizationsRequest{
        Status: directdebit.MandateAuthorizationStatusPending,
    }

    pendingMandates, err := client.DirectDebit.ListMandateAuthorizations(ctx, pendingReq)
    if err != nil {
        log.Fatal(err)
    }

    // Trigger activation charges for pending mandates
    if len(pendingMandates.Data) > 0 {
        var customerIDs []int
        for _, mandate := range pendingMandates.Data {
            customerIDs = append(customerIDs, mandate.Customer.ID)
        }

        activationReq := &directdebit.TriggerActivationChargeRequest{
            CustomerIDs: customerIDs,
        }

        activationResp, err := client.DirectDebit.TriggerActivationCharge(ctx, activationReq)
        if err != nil {
            log.Printf("Activation error: %v\n", err)
        } else {
            fmt.Printf("Activation triggered: %s\n", activationResp.Message)
        }
    }
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

### Dedicated Virtual Account Management

```go
// Fetch available bank providers
providers, err := client.DedicatedVirtualAccount.FetchBankProviders(context.Background())
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Available bank providers: %d\n", len(providers.Data))
preferredBank := providers.Data[0].ProviderSlug // Use first provider

// Create dedicated virtual account for existing customer
createReq := &dedicatedvirtualaccount.CreateDedicatedVirtualAccountRequest{
    Customer:      "CUS_xnxdt6s1zg5f4nx", // existing customer code
    PreferredBank: preferredBank,
    FirstName:     "John",
    LastName:      "Doe",
    Phone:         "+2348100000000",
}

account, err := client.DedicatedVirtualAccount.Create(context.Background(), createReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Created account: %s (%s) - %s\n", 
    account.AccountNumber, account.AccountName, account.Bank.Name)

// List active dedicated virtual accounts
active := true
listReq := &dedicatedvirtualaccount.ListDedicatedVirtualAccountsRequest{
    Active:   &active,
    Currency: "NGN",
}

accounts, err := client.DedicatedVirtualAccount.List(context.Background(), listReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Active accounts: %d\n", len(accounts.Data))

// Add split to account for automatic fund distribution
if len(accounts.Data) > 0 {
    splitReq := &dedicatedvirtualaccount.SplitDedicatedAccountTransactionRequest{
        Customer:      accounts.Data[0].Customer.CustomerCode,
        SplitCode:     "SPL_98WF13Zu8w5", // your split code
        PreferredBank: preferredBank,
    }

    splitAccount, err := client.DedicatedVirtualAccount.SplitTransaction(context.Background(), splitReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Split added to account: %s\n", splitAccount.AccountNumber)
}
```

### Products Management

```go
// Create a physical product with inventory
unlimited := false
quantity := 100
createReq := &products.CreateProductRequest{
    Name:        "Wireless Headphones",
    Description: "High-quality wireless headphones with noise cancellation",
    Price:       25000, // ₦250.00 in kobo
    Currency:    "NGN",
    Unlimited:   &unlimited,
    Quantity:    &quantity,
}

product, err := client.Products.Create(context.Background(), createReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Product created: %s (Code: %s)\n", product.Name, product.ProductCode)

// List all products with pagination
perPage := 10
listReq := &products.ListProductsRequest{
    PerPage: &perPage,
}

productsResp, err := client.Products.List(context.Background(), listReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Found %d products\n", len(productsResp.Data))
for _, prod := range productsResp.Data {
    fmt.Printf("  - %s: ₦%.2f\n", prod.Name, float64(prod.Price)/100)
}

// Update product pricing and details
newPrice := 30000 // ₦300.00
updateReq := &products.UpdateProductRequest{
    Price: &newPrice,
}

updatedProduct, err := client.Products.Update(context.Background(), product.ProductCode, updateReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Updated price: ₦%.2f\n", float64(updatedProduct.Price)/100)
```

### Payment Pages Management

```go
// Create a payment page for course enrollment
fixedAmount := true
collectPhone := false
createReq := &payment_pages.CreatePaymentPageRequest{
    Name:         "Premium Course Access",
    Description:  "One-time payment for premium course access with lifetime updates",
    Amount:       func() *int { amount := 25000; return &amount }(), // ₦250.00
    Currency:     "NGN",
    Type:         "payment",
    FixedAmount:  &fixedAmount,
    CollectPhone: &collectPhone,
    CustomFields: []payment_pages.CustomField{
        {
            DisplayName:  "Full Name",
            VariableName: "student_name",
            Required:     true,
        },
        {
            DisplayName:  "Phone Number",
            VariableName: "phone_number",
            Required:     true,
        },
    },
    Metadata: &types.Metadata{
        "course_id":   "premium-001",
        "instructor":  "John Doe",
        "duration":    "6 months",
    },
    RedirectURL:    "https://myapp.com/course-access",
    SuccessMessage: "Welcome! You now have access to the premium course.",
}

page, err := client.PaymentPages.Create(context.Background(), createReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Payment page created: %s\n", page.Name)
fmt.Printf("Payment URL: https://paystack.com/pay/%s\n", page.Slug)

// Check if a custom slug is available
slug := "premium-course-2024"
slugResp, err := client.PaymentPages.CheckSlugAvailability(context.Background(), slug)
if err != nil {
    log.Fatal(err)
}

if slugResp.Status {
    fmt.Printf("Slug '%s' is available\n", slug)
} else {
    fmt.Printf("Slug '%s' is not available\n", slug)
}

// List all payment pages with pagination
listReq := &payment_pages.ListPaymentPagesRequest{
    PerPage: 10,
    Page:    1,
}

pagesResp, err := client.PaymentPages.List(context.Background(), listReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Found %d payment pages:\n", len(pagesResp.Data))
for _, p := range pagesResp.Data {
    amountStr := "Variable amount"
    if p.Amount != nil {
        amountStr = fmt.Sprintf("₦%.2f", float64(*p.Amount)/100)
    }
    fmt.Printf("  - %s: %s (Active: %t)\n", p.Name, amountStr, p.Active)
}

// Update payment page with early bird pricing
updateReq := &payment_pages.UpdatePaymentPageRequest{
    Name:        "Premium Course Access - Early Bird Special",
    Description: "🐦 Early Bird: Save 20%! Premium course access with lifetime updates",
    Amount:      func() *int { amount := 20000; return &amount }(), // ₦200.00 (discounted)
}

updatedPage, err := client.PaymentPages.Update(context.Background(), page.Slug, updateReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Updated page: %s - ₦%.2f\n", updatedPage.Name, float64(*updatedPage.Amount)/100)
```

### Payment Requests Management

```go
// Create a payment request with line items
createReq := &payment_requests.CreatePaymentRequestRequest{
    Customer:    "CUS_customer_code_here",
    Description: "Invoice for professional services",
    DueDate:     time.Now().AddDate(0, 0, 14).Format("2006-01-02"), // Due in 14 days
    Currency:    "NGN",
    LineItems: []payment_requests.LineItem{
        {
            Name:     "Website Development",
            Amount:   150000, // ₦1,500.00
            Quantity: 1,
        },
        {
            Name:     "SEO Optimization",
            Amount:   75000, // ₦750.00
            Quantity: 1,
        },
        {
            Name:     "Monthly Maintenance",
            Amount:   25000, // ₦250.00
            Quantity: 3,
        },
    },
    Tax: []payment_requests.Tax{
        {
            Name:   "VAT (7.5%)",
            Amount: 22500, // ₦225.00
        },
    },
}

paymentRequest, err := client.PaymentRequests.Create(context.Background(), createReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Payment Request Created: %s\n", paymentRequest.RequestCode)
fmt.Printf("Total Amount: ₦%.2f\n", float64(paymentRequest.Amount)/100)

// List payment requests with filtering
listReq := &payment_requests.ListPaymentRequestsRequest{
    PerPage:  10,
    Page:     1,
    Status:   "pending",
    Currency: "NGN",
}

requestsResp, err := client.PaymentRequests.List(context.Background(), listReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Found %d payment requests\n", len(requestsResp.Data))
for _, req := range requestsResp.Data {
    fmt.Printf("  - %s: ₦%.2f (%s)\n", req.Description, float64(req.Amount)/100, req.Status)
}

// Update payment request with new line items
updateReq := &payment_requests.UpdatePaymentRequestRequest{
    Description: "Updated: Invoice for professional services (Rush Order)",
    DueDate:     time.Now().AddDate(0, 0, 3).Format("2006-01-02"), // Rush - due in 3 days
    LineItems: []payment_requests.LineItem{
        {
            Name:     "Website Development (Express)",
            Amount:   180000, // ₦1,800.00 (rush fee)
            Quantity: 1,
        },
        {
            Name:     "SEO Optimization",
            Amount:   75000,
            Quantity: 1,
        },
    },
}

updatedRequest, err := client.PaymentRequests.Update(context.Background(), paymentRequest.RequestCode, updateReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Updated Amount: ₦%.2f\n", float64(updatedRequest.Amount)/100)

// Get payment request analytics
totals, err := client.PaymentRequests.GetTotals(context.Background())
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Payment Request Analytics:\n")
for _, pending := range totals.Pending {
    fmt.Printf("  Pending %s: ₦%.2f\n", pending.Currency, float64(pending.Amount)/100)
}
```

### Apple Pay Domain Management

```go
// Register domain for Apple Pay integration
registerReq := &applepay.RegisterDomainRequest{
    DomainName: "checkout.mystore.com",
}

registerResp, err := client.ApplePay.RegisterDomain(context.Background(), registerReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Domain registered: %s\n", registerResp.Message)

// List all registered domains
domainsResp, err := client.ApplePay.ListDomains(context.Background(), &applepay.ListDomainsRequest{})
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Registered domains: %d\n", len(domainsResp.Data.DomainNames))
for _, domain := range domainsResp.Data.DomainNames {
    fmt.Printf("  - %s\n", domain)
}

// Unregister domain when no longer needed
unregisterReq := &applepay.UnregisterDomainRequest{
    DomainName: "old-checkout.mystore.com",
}

unregisterResp, err := client.ApplePay.UnregisterDomain(context.Background(), unregisterReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Domain unregistered: %s\n", unregisterResp.Message)
```

### Transfer Control and Balance Management

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/huysamen/paystack-go"
    "github.com/huysamen/paystack-go/api/transfer-control"
)

func main() {
    client := paystack.DefaultClient("sk_test_your_secret_key_here")
    ctx := context.Background()

    // Check current balance across all currencies
    balance, err := client.TransferControl.CheckBalance(ctx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Account Balances:")
    for _, bal := range balance.Data {
        majorAmount := float64(bal.Balance) / 100
        fmt.Printf("  %s: %.2f\n", bal.Currency, majorAmount)
    }

    // Fetch balance ledger for transaction history
    ledger, err := client.TransferControl.FetchBalanceLedger(ctx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Balance Ledger Entries: %d\n", len(ledger.Data))
    if len(ledger.Data) > 0 {
        latest := ledger.Data[0]
        fmt.Printf("Latest transaction: %s (₦%.2f change)\n", 
            latest.Reason, float64(latest.Difference)/100)
    }

    // Enable OTP requirement for transfers
    enableResp, err := client.TransferControl.EnableOTP(ctx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("OTP Status: %s\n", enableResp.Message)

    // Resend OTP for specific transfer (requires existing transfer code)
    resendReq := &transfercontrol.ResendOTPRequest{
        TransferCode: "TRF_your_transfer_code",
        Reason:       "resend_otp",
    }

    resendResp, err := client.TransferControl.ResendOTP(ctx, resendReq)
    if err != nil {
        log.Printf("Resend OTP error: %v", err)
    } else {
        fmt.Printf("OTP Resent: %s\n", resendResp.Message)
    }
}
```

### Integration Management

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/huysamen/paystack-go"
    "github.com/huysamen/paystack-go/api/integration"
)

func main() {
    client := paystack.DefaultClient("sk_test_your_secret_key_here")
    ctx := context.Background()

    // Fetch current payment session timeout
    timeout, err := client.Integration.FetchTimeout(ctx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Current payment session timeout: %d seconds\n", timeout.Data.PaymentSessionTimeout)

    // Configure timeout for different checkout scenarios
    scenarios := []struct {
        name    string
        timeout int
        useCase string
    }{
        {"quick-checkout", 15, "Simple payment forms"},
        {"standard-checkout", 30, "Regular transactions"},
        {"complex-checkout", 60, "Multi-step forms"},
        {"unlimited", 0, "Disable session timeout"},
    }

    for _, scenario := range scenarios {
        fmt.Printf("\nConfiguring %s (%s)...\n", scenario.name, scenario.useCase)
        
        updateReq := &integration.UpdateTimeoutRequest{
            Timeout: scenario.timeout,
        }

        updatedTimeout, err := client.Integration.UpdateTimeout(ctx, updateReq)
        if err != nil {
            log.Printf("Failed to update timeout: %v", err)
            continue
        }

        timeoutText := fmt.Sprintf("%d seconds", updatedTimeout.Data.PaymentSessionTimeout)
        if updatedTimeout.Data.PaymentSessionTimeout == 0 {
            timeoutText = "unlimited"
        }
        fmt.Printf("✅ Updated to: %s\n", timeoutText)
    }

    // Best practices
    fmt.Println("\nBest Practices:")
    fmt.Println("• 15-30s: Ideal for simple forms")
    fmt.Println("• 30-60s: Good for multi-field forms")
    fmt.Println("• 60s+: For complex checkout flows")
    fmt.Println("• 0 (unlimited): Use with caution")
}
```

### Bulk Charges for Mass Payment Processing

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/huysamen/paystack-go"
    "github.com/huysamen/paystack-go/api/bulk-charges"
)

func main() {
    client := paystack.DefaultClient("sk_test_your_secret_key_here")
    ctx := context.Background()

    // Prepare bulk charges (e.g., salary payments)
    charges := bulkcharges.InitiateBulkChargeRequest{
        {
            Authorization: "AUTH_employee_001",
            Amount:        250000, // ₦2,500.00 salary
            Reference:     "salary-jan-2024-001",
        },
        {
            Authorization: "AUTH_employee_002", 
            Amount:        180000, // ₦1,800.00 salary
            Reference:     "salary-jan-2024-002",
        },
        {
            Authorization: "AUTH_employee_003",
            Amount:        320000, // ₦3,200.00 salary  
            Reference:     "salary-jan-2024-003",
        },
    }

    // Initiate bulk charge batch
    batch, err := client.BulkCharges.Initiate(ctx, charges)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Bulk charge initiated: %s\n", batch.Data.BatchCode)
    fmt.Printf("Total charges: %d\n", batch.Data.TotalCharges)
    fmt.Printf("Status: %s\n", batch.Data.Status)

    // Monitor batch progress
    fetchedBatch, err := client.BulkCharges.Fetch(ctx, batch.Data.BatchCode)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Pending charges: %d/%d\n", 
        fetchedBatch.Data.PendingCharges, 
        fetchedBatch.Data.TotalCharges)

    // Get detailed charge information
    chargesReq := &bulkcharges.FetchChargesInBatchRequest{
        PerPage: &[]int{50}[0],
        Status:  &[]string{"success"}[0], // Filter successful charges
    }

    chargeDetails, err := client.BulkCharges.FetchChargesInBatch(
        ctx, batch.Data.BatchCode, chargesReq)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Successful charges: %d\n", len(chargeDetails.Data))
    for _, charge := range chargeDetails.Data {
        fmt.Printf("  - %s: ₦%.2f (%s)\n", 
            charge.Reference, 
            float64(charge.Amount)/100,
            charge.Customer.Email)
    }

    // Pause batch if needed (for large batches)
    if fetchedBatch.Data.Status == "active" {
        pauseResp, err := client.BulkCharges.Pause(ctx, batch.Data.BatchCode)
        if err != nil {
            log.Printf("Pause error: %v", err)
        } else {
            fmt.Printf("Batch paused: %s\n", pauseResp.Message)
        }

        // Resume when ready
        resumeResp, err := client.BulkCharges.Resume(ctx, batch.Data.BatchCode)
        if err != nil {
            log.Printf("Resume error: %v", err)
        } else {
            fmt.Printf("Batch resumed: %s\n", resumeResp.Message)
        }
    }
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For issues related to this library, please create an issue on GitHub.
For Paystack API support, please contact [Paystack Support](https://paystack.com/support).
