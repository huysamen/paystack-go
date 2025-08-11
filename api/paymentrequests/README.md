# Payment Requests API

Send payment requests via email or SMS to collect payments from customers.

## Available Operations

- **Create** - Create and send a payment request
- **List** - List payment requests
- **Fetch** - Get payment request details
- **Update** - Update pending requests
- **Verify** - Verify payment request status
- **Send Notification** - Resend request notification
- **Archive** - Archive completed requests
- **Finalize** - Finalize draft requests

## Quick Examples

### Create Payment Request

```go
import "github.com/huysamen/paystack-go/api/paymentrequests"

request := paymentrequests.NewCreateRequestBuilder().
    Customer("customer@example.com").
    Amount(250000).  // 2500 NGN
    Description("Invoice #INV-001 payment").
    DueDate(time.Now().Add(7*24*time.Hour)).  // Due in 7 days
    SendNotification(true).
    Draft(false).  // Send immediately
    Build()

result, err := client.PaymentRequests.Create(ctx, *request)
if err != nil {
    return err
}

fmt.Printf("Payment request sent: %s\n", result.Data.RequestCode.String())
```

### List Requests

```go
query := paymentrequests.NewListRequestBuilder().
    PerPage(50).
    Page(1).
    Customer(12345).
    Status("pending").
    Build()

result, err := client.PaymentRequests.List(ctx, *query)
```

### Send Notification

```go
request := paymentrequests.NewSendNotificationRequestBuilder().
    RequestCode("PRQ_request_code").
    Build()

result, err := client.PaymentRequests.SendNotification(ctx, *request)
```

## Use Cases

### Invoice Payments

```go
request := paymentrequests.NewCreateRequestBuilder().
    Customer("client@company.com").
    Amount(500000).  // 5000 NGN
    Description("Website Development - Invoice #2024-001").
    DueDate(time.Now().Add(30*24*time.Hour)).
    LineItems([]map[string]any{
        {
            "name": "Website Design",
            "amount": 300000,
        },
        {
            "name": "Development",
            "amount": 200000,
        },
    }).
    Tax([]map[string]any{
        {
            "name": "VAT",
            "amount": 37500,  // 7.5% VAT
        },
    }).
    Build()
```

### Service Payments

```go
request := paymentrequests.NewCreateRequestBuilder().
    Customer("+2348012345678").  // SMS notification
    Amount(150000).
    Description("Plumbing service payment").
    SendNotification(true).
    Build()
```

## Best Practices

### 1. Clear Descriptions

```go
description := fmt.Sprintf("Payment for Order #%s - %s", 
    orderNumber, 
    productName)

request := paymentrequests.NewCreateRequestBuilder().
    Description(description).
    Build()
```

### 2. Reasonable Due Dates

```go
// For invoices: 30 days
dueDate := time.Now().Add(30*24*time.Hour)

// For services: 7 days  
dueDate := time.Now().Add(7*24*time.Hour)

request := paymentrequests.NewCreateRequestBuilder().
    DueDate(dueDate).
    Build()
```
