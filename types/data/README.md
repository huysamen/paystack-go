# Data Types Package

This package contains flexible data types that handle inconsistencies in JSON responses from APIs.

## MultiBool

Handles boolean fields that may come as strings or booleans from the API.

**Accepts as `true`:**
- `true` (boolean)
- `"true"` (string)
- `"success"` (string)

**Treats everything else as `false`:**
- `false` (boolean)
- `"false"` (string)
- `"failed"` (string)
- `"error"` (string)
- `""` (empty string)
- `null`
- Numbers
- Any other value

**Usage:**
```go
type Response struct {
    Status MultiBool `json:"status"`
}

// Use .Bool() to get the boolean value
if response.Status.Bool() {
    // success
}
```

## MultiString

Handles string fields that may come as strings or numbers from the API.

**Handles:**
- String values: `"hello"` → `"hello"`
- Number values: `42` → `"42"`
- Float values: `12.34` → `"12"` (truncated to integer)
- Null values: `null` → `""`

**Usage:**
```go
type Authorization struct {
    ExpMonth MultiString `json:"exp_month"`
    ExpYear  MultiString `json:"exp_year"`
}

// Use .String() to get the string value
month := auth.ExpMonth.String()
```

## Design Philosophy

These types are designed to be:
- **Resilient**: Handle API inconsistencies gracefully
- **Safe**: Default to sensible values when parsing fails
- **Transparent**: Work seamlessly with existing code
- **Well-tested**: Comprehensive test coverage for edge cases
