# Models Package

This package contains comprehensive type definitions for Paystack API data entities. The models are designed with:

- **JSON Serialization**: All fields have proper JSON tags for marshaling/unmarshaling
- **Enum Integration**: Uses the centralized `enums` package for type safety
- **Flexible Time Handling**: Custom `DateTime` type supporting multiple timestamp formats
- **Comprehensive Field Coverage**: Based on actual API response examples
- **Pure Data Models**: Focuses on entity structures, not request/response types

## Model Files

- **authorization.go** - Payment authorizations and mandate authorizations
- **bank.go** - Bank information and bank providers
- **bulk_charge.go** - Bulk charge operations and batches
- **charge.go** - Individual charge data and responses
- **common.go** - Core response structures, metadata, and common types
- **customer.go** - Customer management and related request types
- **dedicated_virtual_account.go** - Dedicated virtual accounts and bank providers
- **dispute.go** - Dispute handling, evidence, and resolution
- **location.go** - Countries and states for address verification
- **payment_page.go** - Payment pages, custom fields, and filters
- **payment_request.go** - Payment requests, line items, and invoices
- **plan.go** - Subscription plans and subscriptions
- **product.go** - Product catalog and management
- **refund.go** - Refund processing
- **settlement.go** - Settlement records and transactions
- **subaccount.go** - Subaccounts and transaction splits
- **terminal.go** - Terminal devices and event handling
- **time.go** - Enhanced datetime handling with multiple format support
- **transaction.go** - Core transaction model with comprehensive fields
- **transfer.go** - Transfers, recipients, and balance information
- **verification.go** - Account validation and card BIN resolution
- **virtual_terminal.go** - Virtual terminals and custom fields

## Key Features

### Generic Response Wrapper
```go
type Response[T any] struct {
    Status  bool   `json:"status"`
    Message string `json:"message"`
    Data    T      `json:"data"`
    Meta    *Meta  `json:"meta,omitempty"`
}
```

### Flexible DateTime
Supports RFC3339, Unix timestamps, and custom formats from Paystack API.

### Comprehensive Coverage
All models include fields observed in actual API responses, ensuring compatibility with the live Paystack API.

### Type Safety
Leverages the `enums` package for all enumerated values, providing compile-time type checking and validation.

### Clear Separation of Concerns  
Models package contains only data entity definitions. Request and response structures are defined in their respective API operation files where they are used.
