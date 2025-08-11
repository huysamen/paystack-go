# Customers API

Comprehensive customer management including creation, updates, authorization management, and risk controls.

## Available Operations

- **Create** - Create a new customer
- **List** - List customers with filtering and pagination
- **Fetch** - Get customer by ID or code
- **Update** - Update customer information
- **Validate** - Validate customer identity
- **Whitelist/Blacklist** - Set customer risk status
- **Deactivate Authorization** - Disable saved payment methods
- **Risk Action** - Set customer risk actions

## Quick Examples

### Create a Customer

```go
import "github.com/huysamen/paystack-go/api/customers"

request := customers.NewCreateRequestBuilder().
    Email("customer@example.com").
    FirstName("John").
    LastName("Doe").
    Phone("+2348012345678").
    Build()

result, err := client.Customers.Create(ctx, *request)
if err != nil {
    return fmt.Errorf("failed to create customer: %w", err)
}

if err := result.Err(); err != nil {
    return fmt.Errorf("customer creation failed: %w", err)
}

customer := result.Data
fmt.Printf("Created customer: %s (ID: %d)\n", 
    customer.Email.String(), 
    customer.ID.Uint64())
```

### List Customers

```go
query := customers.NewListRequestBuilder().
    PerPage(50).
    Page(1).
    Build()

result, err := client.Customers.List(ctx, *query)
if err != nil {
    return err
}

if err := result.Err(); err != nil {
    return err
}

for _, customer := range result.Data {
    fmt.Printf("Customer: %s %s <%s>\n",
        customer.FirstName.String(),
        customer.LastName.String(),
        customer.Email.String())
}
```

### Fetch Customer

```go
// By ID
result, err := client.Customers.Fetch(ctx, 12345)

// By email or customer code
result, err := client.Customers.FetchByCode(ctx, "CUS_abc123xyz")
```

### Update Customer

```go
request := customers.NewUpdateRequestBuilder(12345).
    FirstName("Jane").
    LastName("Smith").
    Phone("+2348087654321").
    Build()

result, err := client.Customers.Update(ctx, *request)
if err != nil {
    return err
}

if err := result.Err(); err != nil {
    return err
}

fmt.Printf("Updated customer: %s\n", result.Data.Email.String())
```

## Advanced Operations

### Customer Validation

Validate customer identity with BVN, account number, etc.:

```go
request := customers.NewValidateRequestBuilder().
    CustomerCode("CUS_abc123xyz").
    FirstName("John").
    LastName("Doe").
    Type("bvn").
    Value("12345678901").  // BVN
    Country("NG").
    Build()

result, err := client.Customers.Validate(ctx, *request)
if err != nil {
    return err
}

if err := result.Err(); err != nil {
    return fmt.Errorf("validation failed: %w", err)
}

fmt.Printf("Validation status: %s\n", result.Data.Status.String())
```

### Risk Management

#### Whitelist a Customer

```go
request := customers.NewWhitelistRequestBuilder().
    CustomerCode("CUS_abc123xyz").
    Build()

result, err := client.Customers.Whitelist(ctx, *request)
```

#### Blacklist a Customer

```go
request := customers.NewBlacklistRequestBuilder().
    CustomerCode("CUS_abc123xyz").
    RiskAction("deny").  // deny, allow
    Build()

result, err := client.Customers.Blacklist(ctx, *request)
```

#### Set Risk Action

```go
request := customers.NewRiskActionRequestBuilder().
    Customer(12345).
    Action(customers.RiskActionDeny).  // or RiskActionAllow
    Build()

result, err := client.Customers.RiskAction(ctx, *request)
```

### Authorization Management

#### Deactivate Authorization

Remove a saved payment method:

```go
request := customers.NewDeactivateAuthorizationRequestBuilder().
    AuthorizationCode("AUTH_abc123xyz").
    Build()

result, err := client.Customers.DeactivateAuthorization(ctx, *request)
if err != nil {
    return err
}

if err := result.Err(); err != nil {
    return err
}

fmt.Printf("Authorization deactivated: %s\n", 
    result.Data.Status.String())
```

## Request Builder Options

### Create Customer Builder

```go
request := customers.NewCreateRequestBuilder().
    // Required
    Email("customer@example.com").
    
    // Optional personal info
    FirstName("John").
    LastName("Doe").
    Phone("+2348012345678").
    
    // Optional metadata
    Metadata(map[string]any{
        "customer_id": "internal_123",
        "source": "website",
    }).
    
    Build()
```

### List Customers Builder

```go
query := customers.NewListRequestBuilder().
    // Pagination
    PerPage(100).  // Max 100
    Page(1).
    
    // Optional filters
    // Note: Specific filters depend on API capabilities
    
    Build()
```

### Update Customer Builder

```go
request := customers.NewUpdateRequestBuilder(customerID).
    // Update any field
    FirstName("UpdatedFirstName").
    LastName("UpdatedLastName").
    Phone("+2348087654321").
    
    // Update metadata
    Metadata(map[string]any{
        "updated_at": time.Now().Format(time.RFC3339),
        "source": "mobile_app",
    }).
    
    Build()
```

### Validation Builder

```go
request := customers.NewValidateRequestBuilder().
    CustomerCode("CUS_abc123xyz").
    
    // Identity verification
    FirstName("John").
    LastName("Doe").
    Type("bvn").           // bvn, bank_account, etc.
    Value("12345678901").  // BVN, account number, etc.
    Country("NG").
    
    // Additional fields for bank account validation
    BankCode("011").       // For bank account validation
    AccountNumber("0123456789").
    
    Build()
```

## Response Data Types

### Customer Object

```go
type Customer struct {
    ID           data.Uint       `json:"id"`
    FirstName    data.NullString `json:"first_name"`
    LastName     data.NullString `json:"last_name"`
    Email        data.String     `json:"email"`
    CustomerCode data.String     `json:"customer_code"`
    Phone        data.NullString `json:"phone"`
    Metadata     Metadata        `json:"metadata"`
    RiskAction   data.String     `json:"risk_action"`
    
    // Timestamps
    CreatedAt    data.Time       `json:"createdAt"`
    UpdatedAt    data.Time       `json:"updatedAt"`
    
    // Integration info
    Integration  data.Int        `json:"integration"`
    Domain       data.String     `json:"domain"`
    
    // Identification (if validated)
    Identified              data.Bool       `json:"identified"`
    Identifications         []Identification `json:"identifications"`
    
    // Additional fields may be present
}
```

### Authorization Object

```go
type Authorization struct {
    AuthorizationCode data.String     `json:"authorization_code"`
    Bin              data.String     `json:"bin"`
    Last4            data.String     `json:"last4"`
    ExpMonth         data.String     `json:"exp_month"`
    ExpYear          data.String     `json:"exp_year"`
    Channel          enums.Channel   `json:"channel"`
    CardType         data.String     `json:"card_type"`
    Bank             data.String     `json:"bank"`
    CountryCode      data.String     `json:"country_code"`
    Brand            data.String     `json:"brand"`
    Reusable         data.Bool       `json:"reusable"`
    Signature        data.String     `json:"signature"`
    AccountName      data.NullString `json:"account_name"`
}
```

## Error Handling

### Common Error Scenarios

```go
result, err := client.Customers.Create(ctx, request)

// Network errors
if err != nil {
    return fmt.Errorf("network error: %w", err)
}

// API errors
if err := result.Err(); err != nil {
    switch {
    case strings.Contains(err.Error(), "email"):
        return fmt.Errorf("invalid email address: %w", err)
    case strings.Contains(err.Error(), "duplicate"):
        return fmt.Errorf("customer already exists: %w", err)
    case strings.Contains(err.Error(), "phone"):
        return fmt.Errorf("invalid phone number: %w", err)
    default:
        return fmt.Errorf("customer creation failed: %w", err)
    }
}
```

### Validation Errors

```go
result, err := client.Customers.Validate(ctx, request)
if err != nil {
    return err
}

if err := result.Err(); err != nil {
    switch {
    case strings.Contains(err.Error(), "bvn"):
        return fmt.Errorf("BVN validation failed: %w", err)
    case strings.Contains(err.Error(), "mismatch"):
        return fmt.Errorf("customer details don't match records: %w", err)
    default:
        return fmt.Errorf("validation failed: %w", err)
    }
}

// Check validation status
if result.Data.Status.String() != "success" {
    return fmt.Errorf("validation unsuccessful: %s", result.Data.Status.String())
}
```

## Best Practices

### 1. Email Validation

Always validate email addresses before creating customers:

```go
import "net/mail"

func isValidEmail(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}

if !isValidEmail(email) {
    return fmt.Errorf("invalid email format: %s", email)
}
```

### 2. Phone Number Format

Use international format for phone numbers:

```go
// Good: International format
phone := "+2348012345678"

// Avoid: Local format
// phone := "08012345678"
```

### 3. Metadata Usage

Use metadata for internal tracking:

```go
metadata := map[string]any{
    "internal_customer_id": userID,
    "source": "mobile_app",
    "signup_date": time.Now().Format(time.RFC3339),
    "plan": "premium",
}

request := customers.NewCreateRequestBuilder().
    Email(email).
    Metadata(metadata).
    Build()
```

### 4. Customer Lookup

Implement efficient customer lookup:

```go
type CustomerService struct {
    client *paystack.Client
    cache  map[string]*types.Customer // Simple cache
}

func (s *CustomerService) GetOrCreateCustomer(ctx context.Context, email string) (*types.Customer, error) {
    // Check cache first
    if customer, exists := s.cache[email]; exists {
        return customer, nil
    }
    
    // Try to find existing customer
    // Note: You may need to implement search by email
    // or maintain your own customer mapping
    
    // Create new customer if not found
    request := customers.NewCreateRequestBuilder().
        Email(email).
        Build()
        
    result, err := s.client.Customers.Create(ctx, *request)
    if err != nil {
        return nil, err
    }
    
    if err := result.Err(); err != nil {
        // If customer already exists, you might want to handle this
        if strings.Contains(err.Error(), "duplicate") {
            // Handle duplicate customer scenario
        }
        return nil, err
    }
    
    // Cache the result
    s.cache[email] = &result.Data
    return &result.Data, nil
}
```

### 5. Risk Management

Implement automated risk management:

```go
func (s *CustomerService) EvaluateCustomerRisk(ctx context.Context, customerID uint64, transaction *types.Transaction) error {
    // Example risk evaluation logic
    if transaction.Amount.Int64() > 1000000 { // Large transaction
        request := customers.NewRiskActionRequestBuilder().
            Customer(customerID).
            Action(customers.RiskActionDeny).
            Build()
            
        result, err := s.client.Customers.RiskAction(ctx, *request)
        if err != nil {
            return err
        }
        
        if err := result.Err(); err != nil {
            return err
        }
    }
    
    return nil
}
```

## Testing

See `api/customers/*_test.go` for comprehensive examples using real JSON fixtures from `resources/examples/responses/customers/`.

Tests cover:
- Customer creation and updates
- List operations with pagination
- Validation workflows
- Risk management operations
- Error handling scenarios
