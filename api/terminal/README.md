# Terminal API

Manage physical payment terminals for in-person transactions.

## Available Operations

- **List** - List terminals
- **Fetch** - Get terminal details
- **Commission** - Commission a terminal
- **Decommission** - Decommission a terminal
- **Update** - Update terminal settings
- **Send Event** - Send events to terminal

## Quick Examples

### List Terminals

```go
import "github.com/huysamen/paystack-go/api/terminal"

query := terminal.NewListRequestBuilder().
    PerPage(50).
    Page(1).
    Build()

result, err := client.Terminal.List(ctx, *query)
```

### Commission Terminal

```go
request := terminal.NewCommissionRequestBuilder().
    SerialNumber("1234567890").
    Build()

result, err := client.Terminal.Commission(ctx, *request)
```

### Send Event

```go
request := terminal.NewSendEventRequestBuilder().
    TerminalID("terminal_id").
    Type(enums.TerminalEventTypeInvoice).
    Action(enums.TerminalActionProcess).
    Data(map[string]any{
        "id": 12345,
        "reference": "INV_001",
    }).
    Build()

result, err := client.Terminal.SendEvent(ctx, *request)
```

## Use Cases

Perfect for retail stores, restaurants, and physical businesses that need in-person payment processing.
