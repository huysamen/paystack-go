# Products API

Manage your product catalog for e-commerce and digital goods.

## Available Operations

- **Create** - Create a new product
- **List** - List products with filtering
- **Fetch** - Get product by ID
- **Update** - Update product details

## Quick Examples

### Create Product

```go
import "github.com/huysamen/paystack-go/api/products"

request := products.NewCreateRequestBuilder().
    Name("Premium Software License").
    Description("Annual license for premium software features").
    Price(50000000).  // 500,000 NGN in kobo
    Currency(enums.NGN).
    Limited(true).
    Quantity(100).
    Build()

result, err := client.Products.Create(ctx, *request)
```

### List Products

```go
query := products.NewListRequestBuilder().
    PerPage(50).
    Page(1).
    Active(true).
    Build()

result, err := client.Products.List(ctx, *query)

for _, product := range result.Data {
    fmt.Printf("Product: %s - %d kobo\n", 
        product.Name.String(), 
        product.Price.Int64())
}
```

### Update Product

```go
request := products.NewUpdateRequestBuilder(productID).
    Name("Premium Software License v2").
    Price(60000000).  // Price increase
    Description("Enhanced features with AI integration").
    Build()

result, err := client.Products.Update(ctx, *request)
```

## Use Cases

Perfect for:
- Digital products and licenses
- Physical goods inventory
- Service packages
- Subscription add-ons

## Best Practices

```go
// Include detailed metadata for inventory tracking
request := products.NewCreateRequestBuilder().
    Name("MacBook Pro 16-inch").
    Price(250000000).  // 2,500,000 NGN
    Currency(enums.NGN).
    Limited(true).
    Quantity(10).
    Metadata(map[string]any{
        "sku": "MBP16-001",
        "category": "electronics",
        "weight_kg": 2.1,
        "warranty_months": 12,
    }).
    Build()
```
