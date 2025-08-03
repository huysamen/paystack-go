package types

// DedicatedVirtualAccountBank represents a bank provider for dedicated virtual accounts
type DedicatedVirtualAccountBank struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// DedicatedVirtualAccount represents a dedicated virtual account
type DedicatedVirtualAccount struct {
	ID            int                         `json:"id"`
	AccountName   string                      `json:"account_name"`
	AccountNumber string                      `json:"account_number"`
	Assigned      bool                        `json:"assigned"`
	Currency      string                      `json:"currency"`
	Metadata      *Metadata                   `json:"metadata,omitempty"`
	Active        bool                        `json:"active"`
	Bank          DedicatedVirtualAccountBank `json:"bank"`
	Customer      *Customer                   `json:"customer,omitempty"`
	CreatedAt     string                      `json:"created_at,omitempty"`
	UpdatedAt     string                      `json:"updated_at,omitempty"`
	SplitConfig   any                         `json:"split_config,omitempty"`
}
