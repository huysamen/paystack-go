package types

// Bearer represents who bears the transaction charges
type Bearer string

const (
	BearerUnknown    Bearer = ""
	BearerAccount    Bearer = "account"
	BearerSubaccount Bearer = "subaccount"
)

// String returns the string representation of the bearer
func (b Bearer) String() string {
	return string(b)
}

// IsValid returns true if the bearer is a valid known value
func (b Bearer) IsValid() bool {
	switch b {
	case BearerAccount, BearerSubaccount:
		return true
	default:
		return false
	}
}
