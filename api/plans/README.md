# Plans API

Create and manage subscription billing plans with flexible pricing and intervals.

## Available Operations

- **Create** - Create a new plan
- **List** - List plans with filtering  
- **Fetch** - Get plan by ID or code
- **Update** - Update plan details

## Quick Examples

### Create Plan

```go
import "github.com/huysamen/paystack-go/api/plans"

request := plans.NewCreateRequestBuilder().
    Name("Premium Plan").
    Amount(500000).  // 5000 NGN in kobo
    Interval(enums.Monthly).
    Currency(enums.NGN).
    Description("Premium features with priority support").
    Build()

result, err := client.Plans.Create(ctx, *request)
if err != nil {
    return err
}

plan := result.Data
fmt.Printf("Plan created: %s (Code: %s)\n", 
    plan.Name.String(), 
    plan.PlanCode.String())
```

### List Plans

```go
query := plans.NewListRequestBuilder().
    PerPage(50).
    Page(1).
    Status("active").
    Amount(500000).
    Interval(enums.Monthly).
    Build()

result, err := client.Plans.List(ctx, *query)
```

### Update Plan

```go
request := plans.NewUpdateRequestBuilder(planID).
    Name("Premium Plan Pro").
    Description("Enhanced premium features").
    Build()

result, err := client.Plans.Update(ctx, *request)
```

## Plan Intervals

- `hourly` - Every hour
- `daily` - Every day  
- `weekly` - Every week
- `monthly` - Every month
- `quarterly` - Every 3 months
- `biannually` - Every 6 months
- `annually` - Every year

## Best Practices

Use descriptive names and include currency information:

```go
request := plans.NewCreateRequestBuilder().
    Name("Pro Plan - Monthly (NGN)").
    Amount(1000000).
    Interval(enums.Monthly).
    Currency(enums.NGN).
    PlanCode("pro-monthly-ngn").  // Custom code for easy reference
    Build()
```
