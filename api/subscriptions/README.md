# Subscriptions API

Manage recurring billing subscriptions including creation, management, and lifecycle operations.

## Available Operations

- **Create** - Create a new subscription
- **List** - List subscriptions with filtering
- **Fetch** - Get subscription by ID or code
- **Enable** - Reactivate a subscription
- **Disable** - Deactivate a subscription
- **Generate Update Link** - Create link for customer to update payment method

## Quick Examples

### Create Subscription

```go
import "github.com/huysamen/paystack-go/api/subscriptions"

request := subscriptions.NewCreateRequestBuilder("CUS_customer_code", "PLN_plan_code").
    Authorization("AUTH_authorization_code").  // Optional: for immediate charge
    StartDate(time.Now().Add(7*24*time.Hour)). // Optional: future start
    Build()

result, err := client.Subscriptions.Create(ctx, *request)
if err != nil {
    return fmt.Errorf("subscription creation failed: %w", err)
}

if err := result.Err(); err != nil {
    return fmt.Errorf("subscription error: %w", err)
}

subscription := result.Data
fmt.Printf("Subscription created: %s (Status: %s)\n",
    subscription.SubscriptionCode.String(),
    subscription.Status.String())
```

### List Subscriptions

```go
query := subscriptions.NewListRequestBuilder().
    PerPage(50).
    Page(1).
    Customer(12345).
    Plan(67890).
    Build()

result, err := client.Subscriptions.List(ctx, *query)
if err != nil {
    return err
}

for _, sub := range result.Data {
    fmt.Printf("Subscription: %s - %s (Next: %s)\n",
        sub.SubscriptionCode.String(),
        sub.Status.String(),
        sub.NextPaymentDate.Time().Format("2006-01-02"))
}
```

### Enable/Disable Subscription

```go
// Disable subscription
disableReq := subscriptions.NewDisableRequestBuilder().
    Code("SUB_subscription_code").
    Token("email_token").  // From subscription object
    Build()

result, err := client.Subscriptions.Disable(ctx, *disableReq)

// Enable subscription
enableReq := subscriptions.NewEnableRequestBuilder().
    Code("SUB_subscription_code").
    Token("email_token").
    Build()

result, err := client.Subscriptions.Enable(ctx, *enableReq)
```

### Generate Update Link

```go
request := subscriptions.NewGenerateUpdateLinkRequestBuilder().
    Code("SUB_subscription_code").
    Build()

result, err := client.Subscriptions.GenerateUpdateLink(ctx, *request)
if err != nil {
    return err
}

if err := result.Err(); err != nil {
    return err
}

// Send link to customer
updateLink := result.Data.Link.String()
fmt.Printf("Update link: %s\n", updateLink)
```

## Advanced Usage

### Subscription with Custom Start Date

```go
// Start subscription next month
nextMonth := time.Now().AddDate(0, 1, 0)

request := subscriptions.NewCreateRequestBuilder("CUS_xxx", "PLN_xxx").
    StartDate(nextMonth).
    Build()

result, err := client.Subscriptions.Create(ctx, *request)
```

### Create with Immediate Authorization

```go
// Create subscription and charge immediately using saved card
request := subscriptions.NewCreateRequestBuilder("CUS_xxx", "PLN_xxx").
    Authorization("AUTH_saved_card").
    Build()

result, err := client.Subscriptions.Create(ctx, *request)
if err != nil {
    return err
}

// Check if first charge was successful
if result.Data.Status.String() == "active" {
    fmt.Println("Subscription created and first payment successful")
} else {
    fmt.Printf("Subscription created but payment status: %s\n", 
        result.Data.Status.String())
}
```

## Subscription Lifecycle

### Status Flow

```
active → non-renewing → cancelled
  ↓           ↑
attention → active (on successful payment)
```

### Status Meanings

- **active**: Subscription is active and will renew
- **non-renewing**: Active but won't auto-renew
- **attention**: Payment failed, needs customer action
- **cancelled**: Subscription is cancelled

### Handling Different States

```go
func handleSubscriptionStatus(subscription *types.Subscription) error {
    switch subscription.Status.String() {
    case "active":
        return enableUserFeatures(subscription.Customer.Email.String())
        
    case "non-renewing":
        return notifySubscriptionEnding(subscription)
        
    case "attention":
        return requestPaymentUpdate(subscription)
        
    case "cancelled":
        return disableUserFeatures(subscription.Customer.Email.String())
        
    default:
        return fmt.Errorf("unknown subscription status: %s", 
            subscription.Status.String())
    }
}
```

## Request Builders

### Create Subscription Builder

```go
request := subscriptions.NewCreateRequestBuilder(customerCode, planCode).
    // Optional parameters
    Authorization("AUTH_xxx").      // Use saved payment method
    StartDate(futureDate).          // Delay start
    Build()
```

### List Subscriptions Builder

```go
query := subscriptions.NewListRequestBuilder().
    PerPage(100).
    Page(1).
    Customer(customerID).           // Filter by customer
    Plan(planID).                   // Filter by plan
    Build()
```

### Enable/Disable Builders

```go
// Both use same pattern
request := subscriptions.NewDisableRequestBuilder().
    Code("SUB_xxx").
    Token("email_token").
    Build()

request := subscriptions.NewEnableRequestBuilder().
    Code("SUB_xxx").
    Token("email_token").
    Build()
```

## Webhook Integration

Handle subscription events for real-time updates:

```go
func handleSubscriptionWebhook(event *webhook.Event) error {
    switch event.Event {
    case webhook.EventSubscriptionCreate:
        data, _ := event.AsSubscriptionCreate()
        return activateUserSubscription(
            data.Customer.Email.String(),
            data.SubscriptionCode.String(),
        )
        
    case webhook.EventSubscriptionDisable:
        data, _ := event.AsSubscriptionDisable()
        return deactivateUserSubscription(
            data.SubscriptionCode.String(),
        )
        
    case webhook.EventSubscriptionNotRenew:
        data, _ := event.AsSubscriptionNotRenew()
        return notifySubscriptionEnding(data)
        
    case webhook.EventSubscriptionExpiringCards:
        return handleExpiringCards(event)
    }
    return nil
}

func handleExpiringCards(event *webhook.Event) error {
    data, err := event.AsSubscriptionExpiringCards()
    if err != nil {
        return err
    }
    
    // Process expiring cards data
    if !data.Entries.Valid {
        return nil
    }
    
    entries, ok := data.Entries.Metadata["entries"].([]interface{})
    if !ok {
        return fmt.Errorf("invalid expiring cards format")
    }
    
    for _, entry := range entries {
        cardInfo, ok := entry.(map[string]interface{})
        if !ok {
            continue
        }
        
        customerEmail, _ := cardInfo["customer_email"].(string)
        subscriptionCode, _ := cardInfo["subscription_code"].(string)
        
        // Notify customer to update payment method
        if err := notifyCardExpiry(customerEmail, subscriptionCode); err != nil {
            log.Printf("Failed to notify %s: %v", customerEmail, err)
        }
    }
    
    return nil
}
```

## Best Practices

### 1. Customer Communication

Always provide clear communication about subscription status:

```go
func notifySubscriptionChange(customerEmail, subscriptionCode, newStatus string) error {
    messages := map[string]string{
        "active":      "Your subscription is now active!",
        "non-renewing": "Your subscription will not auto-renew.",
        "attention":   "Please update your payment method.",
        "cancelled":   "Your subscription has been cancelled.",
    }
    
    message, exists := messages[newStatus]
    if !exists {
        message = fmt.Sprintf("Subscription status changed to: %s", newStatus)
    }
    
    return sendEmail(customerEmail, "Subscription Update", message)
}
```

### 2. Graceful Failures

Handle payment failures gracefully:

```go
func handlePaymentFailure(subscriptionCode string) error {
    // Generate update link for customer
    request := subscriptions.NewGenerateUpdateLinkRequestBuilder().
        Code(subscriptionCode).
        Build()
        
    result, err := client.Subscriptions.GenerateUpdateLink(ctx, *request)
    if err != nil {
        return err
    }
    
    if err := result.Err(); err != nil {
        return err
    }
    
    updateLink := result.Data.Link.String()
    
    // Send to customer
    return sendPaymentUpdateLink(subscriptionCode, updateLink)
}
```

### 3. Subscription Analytics

Track subscription metrics:

```go
type SubscriptionMetrics struct {
    Active       int
    NonRenewing  int
    Attention    int
    Cancelled    int
    TotalRevenue int64
}

func getSubscriptionMetrics() (*SubscriptionMetrics, error) {
    query := subscriptions.NewListRequestBuilder().
        PerPage(100).
        Page(1).
        Build()
    
    metrics := &SubscriptionMetrics{}
    page := 1
    
    for {
        query.Page = page
        result, err := client.Subscriptions.List(ctx, *query)
        if err != nil {
            return nil, err
        }
        
        if err := result.Err(); err != nil {
            return nil, err
        }
        
        for _, sub := range result.Data {
            switch sub.Status.String() {
            case "active":
                metrics.Active++
                metrics.TotalRevenue += sub.Amount.Int64()
            case "non-renewing":
                metrics.NonRenewing++
            case "attention":
                metrics.Attention++
            case "cancelled":
                metrics.Cancelled++
            }
        }
        
        if len(result.Data) < 100 {
            break
        }
        page++
    }
    
    return metrics, nil
}
```

### 4. Proactive Card Management

Monitor and handle expiring cards:

```go
func checkExpiringCards() error {
    // This would typically be triggered by webhook
    // but you might also want to run periodic checks
    
    query := subscriptions.NewListRequestBuilder().
        PerPage(100).
        Build()
    
    result, err := client.Subscriptions.List(ctx, *query)
    if err != nil {
        return err
    }
    
    for _, sub := range result.Data {
        // Check if card expires soon (within 30 days)
        if isCardExpiringSoon(sub.Authorization) {
            updateLink, err := generateUpdateLink(sub.SubscriptionCode.String())
            if err != nil {
                log.Printf("Failed to generate update link for %s: %v", 
                    sub.SubscriptionCode.String(), err)
                continue
            }
            
            err = notifyCardExpiry(
                sub.Customer.Email.String(),
                sub.SubscriptionCode.String(),
                updateLink,
            )
            if err != nil {
                log.Printf("Failed to notify customer: %v", err)
            }
        }
    }
    
    return nil
}
```

## Testing

See `api/subscriptions/*_test.go` for comprehensive examples using real JSON fixtures from `resources/examples/responses/subscriptions/`.
