# Miscellaneous API

Access supporting data like banks, countries, and states.

## Available Operations

- **List Banks** - Get supported banks
- **List Countries** - Get supported countries  
- **List States** - Get states/provinces

## Quick Examples

### List Banks

```go
import "github.com/huysamen/paystack-go/api/miscellaneous"

query := miscellaneous.NewListBanksRequestBuilder().
    Country("nigeria").
    UseCursor(false).
    PerPage(100).
    Build()

result, err := client.Miscellaneous.ListBanks(ctx, *query)
if err != nil {
    return err
}

for _, bank := range result.Data {
    fmt.Printf("Bank: %s (Code: %s)\n", 
        bank.Name.String(), 
        bank.Code.String())
}
```

### List Countries

```go
result, err := client.Miscellaneous.ListCountries(ctx)
if err != nil {
    return err
}

for _, country := range result.Data {
    fmt.Printf("Country: %s (Code: %s)\n",
        country.Name.String(),
        country.IsoCode.String())
}
```

### List States

```go
result, err := client.Miscellaneous.ListStates(ctx, "NG")  // Nigeria
if err != nil {
    return err
}

for _, state := range result.Data {
    fmt.Printf("State: %s\n", state.Name.String())
}
```

## Use Cases

### Bank Selection UI

```go
func getBankOptions() ([]BankOption, error) {
    query := miscellaneous.NewListBanksRequestBuilder().
        Country("nigeria").
        PerPage(100).
        Build()
        
    result, err := client.Miscellaneous.ListBanks(ctx, *query)
    if err != nil {
        return nil, err
    }
    
    var options []BankOption
    for _, bank := range result.Data {
        options = append(options, BankOption{
            Code: bank.Code.String(),
            Name: bank.Name.String(),
        })
    }
    
    return options, nil
}
```

### Address Forms

```go
func getLocationData(countryCode string) (*LocationData, error) {
    states, err := client.Miscellaneous.ListStates(ctx, countryCode)
    if err != nil {
        return nil, err
    }
    
    return &LocationData{
        CountryCode: countryCode,
        States:      states.Data,
    }, nil
}
```

Essential for building user interfaces that need bank selection, country/state dropdowns, and address forms.
