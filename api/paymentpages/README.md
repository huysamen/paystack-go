# Payment Pages API

Create hosted payment pages for easy online payments without building your own checkout.

## Available Operations

- **Create** - Create a new payment page
- **List** - List payment pages
- **Fetch** - Get payment page by ID or slug
- **Update** - Update payment page details
- **Add Products** - Add products to payment page

## Quick Examples

### Create Payment Page

```go
import "github.com/huysamen/paystack-go/api/paymentpages"

request := paymentpages.NewCreateRequestBuilder().
    Name("Product Purchase").
    Description("Buy our amazing products").
    Amount(100000).  // Fixed amount: 1000 NGN
    Slug("buy-products").  // Custom URL slug
    RedirectURL("https://yoursite.com/thank-you").
    Build()

result, err := client.PaymentPages.Create(ctx, *request)
if err != nil {
    return err
}

page := result.Data
fmt.Printf("Payment page created: %s\n", page.Slug.String())
fmt.Printf("URL: https://paystack.com/pay/%s\n", page.Slug.String())
```

### List Payment Pages

```go
query := paymentpages.NewListRequestBuilder().
    PerPage(50).
    Page(1).
    Build()

result, err := client.PaymentPages.List(ctx, *query)

for _, page := range result.Data {
    fmt.Printf("Page: %s (Active: %t)\n", 
        page.Name.String(), 
        page.Active.Bool())
}
```

### Update Payment Page

```go
request := paymentpages.NewUpdateRequestBuilder(pageID).
    Name("Updated Product Page").
    Description("New and improved product offerings").
    Active(true).
    Build()

result, err := client.PaymentPages.Update(ctx, *request)
```

## Use Cases

### E-commerce Store

```go
// Variable amount page for general purchases
request := paymentpages.NewCreateRequestBuilder().
    Name("Store Checkout").
    Description("Complete your purchase").
    // No amount = customer enters amount
    Slug("store-checkout").
    RedirectURL("https://store.com/success").
    CustomFields([]map[string]any{
        {"display_name": "Product SKU", "variable_name": "sku"},
        {"display_name": "Quantity", "variable_name": "qty"},
    }).
    Build()
```

### Donation Page

```go
request := paymentpages.NewCreateRequestBuilder().
    Name("Support Our Cause").
    Description("Help us make a difference").
    Slug("donate").
    MinAmount(100000).  // Minimum 1000 NGN
    RedirectURL("https://charity.org/thank-you").
    Build()
```

### Event Tickets

```go
request := paymentpages.NewCreateRequestBuilder().
    Name("Conference 2024 Tickets").
    Amount(5000000).  // Fixed price: 50,000 NGN
    Slug("conference-2024").
    RedirectURL("https://event.com/ticket-confirmation").
    CustomFields([]map[string]any{
        {"display_name": "Attendee Name", "variable_name": "attendee_name"},
        {"display_name": "Dietary Requirements", "variable_name": "dietary"},
    }).
    Build()
```

## Advanced Features

### Custom Fields

Collect additional information from customers:

```go
customFields := []map[string]any{
    {
        "display_name": "Full Name",
        "variable_name": "full_name",
    },
    {
        "display_name": "Phone Number", 
        "variable_name": "phone",
    },
    {
        "display_name": "Company",
        "variable_name": "company",
    },
}

request := paymentpages.NewCreateRequestBuilder().
    Name("Business Service").
    Amount(25000000).
    CustomFields(customFields).
    Build()
```

### Split Payments

Share revenue with subaccounts:

```go
splitConfig := map[string]any{
    "type": "percentage",
    "subaccounts": []map[string]any{
        {
            "subaccount": "ACCT_subaccount_code",
            "share": 20,  // 20% to subaccount
        },
    },
}

request := paymentpages.NewCreateRequestBuilder().
    Name("Marketplace Product").
    Amount(10000000).
    Split(splitConfig).
    Build()
```

## Best Practices

### 1. SEO-Friendly Slugs

```go
slug := strings.ToLower(strings.ReplaceAll(productName, " ", "-"))
slug = regexp.MustCompile(`[^a-z0-9-]`).ReplaceAllString(slug, "")

request := paymentpages.NewCreateRequestBuilder().
    Name(productName).
    Slug(slug).
    Build()
```

### 2. Clear Descriptions

```go
request := paymentpages.NewCreateRequestBuilder().
    Name("Premium Subscription").
    Description("Monthly subscription with premium features including priority support, advanced analytics, and unlimited usage.").
    Build()
```

### 3. Success Page Integration

```go
request := paymentpages.NewCreateRequestBuilder().
    Name("Product Purchase").
    RedirectURL("https://yoursite.com/payment/success?reference={{reference}}").
    Build()
```

The `{{reference}}` placeholder will be replaced with the actual transaction reference.
