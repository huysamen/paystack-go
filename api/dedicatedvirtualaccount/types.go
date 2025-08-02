package dedicatedvirtualaccount

import (
	"github.com/huysamen/paystack-go/types"
)

// Bank represents a bank provider for dedicated virtual accounts
type Bank struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// BankProvider represents a bank provider
type BankProvider struct {
	ID           int    `json:"id"`
	ProviderSlug string `json:"provider_slug"`
	BankID       int    `json:"bank_id"`
	BankName     string `json:"bank_name"`
}

// DedicatedVirtualAccount represents a dedicated virtual account
type DedicatedVirtualAccount struct {
	ID            int             `json:"id"`
	AccountName   string          `json:"account_name"`
	AccountNumber string          `json:"account_number"`
	Assigned      bool            `json:"assigned"`
	Currency      string          `json:"currency"`
	Metadata      *types.Metadata `json:"metadata,omitempty"`
	Active        bool            `json:"active"`
	Bank          Bank            `json:"bank"`
	Customer      *types.Customer `json:"customer,omitempty"`
	CreatedAt     string          `json:"created_at,omitempty"`
	UpdatedAt     string          `json:"updated_at,omitempty"`
	SplitConfig   any             `json:"split_config,omitempty"`
}
