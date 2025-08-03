package types

// BankProvider represents a bank provider
type BankProvider struct {
	ID           int    `json:"id"`
	ProviderSlug string `json:"provider_slug"`
	BankID       int    `json:"bank_id"`
	BankName     string `json:"bank_name"`
}
