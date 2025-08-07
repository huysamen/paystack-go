package types

import "github.com/huysamen/paystack-go/enums"

// Bank represents a bank in the system
type Bank struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Slug        string         `json:"slug"`
	Code        string         `json:"code"`
	LongCode    string         `json:"longcode"`
	Gateway     *string        `json:"gateway"`
	PayWithBank bool           `json:"pay_with_bank"`
	Active      bool           `json:"active"`
	IsDeleted   bool           `json:"is_deleted"`
	Country     string         `json:"country"`
	Currency    enums.Currency `json:"currency"`
	Type        string         `json:"type"`
	CreatedAt   DateTime       `json:"createdAt"`
	UpdatedAt   DateTime       `json:"updatedAt"`
}

// BankProvider represents a bank provider
type BankProvider struct {
	ID           int    `json:"id"`
	ProviderSlug string `json:"provider_slug"`
	BankID       int    `json:"bank_id"`
	BankName     string `json:"bank_name"`
}
