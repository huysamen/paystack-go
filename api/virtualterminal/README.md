# Virtual Terminal API

Manage virtual payment terminals for phone and online orders.

## Available Operations

- **List** - List virtual terminals
- **Fetch** - Get virtual terminal details
- **Update** - Update terminal settings

## Quick Examples

```go
import "github.com/huysamen/paystack-go/api/virtualterminal"

// List virtual terminals
query := virtualterminal.NewListRequestBuilder().
    PerPage(50).
    Build()

result, err := client.VirtualTerminal.List(ctx, *query)

// Fetch specific terminal
result, err := client.VirtualTerminal.Fetch(ctx, terminalID)

// Update terminal
request := virtualterminal.NewUpdateRequestBuilder(terminalID).
    Name("Updated Terminal Name").
    Build()

result, err := client.VirtualTerminal.Update(ctx, *request)
```

## Use Cases

Ideal for businesses that take payments over the phone or need a simple payment interface for staff.
