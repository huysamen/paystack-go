package types

import "time"

// RecipientType represents the type of transfer recipient
type RecipientType string

const (
	RecipientTypeNuban       RecipientType = "nuban"        // Nigerian bank account
	RecipientTypeGhipss      RecipientType = "ghipss"       // Ghana bank account
	RecipientTypeMobileMoney RecipientType = "mobile_money" // Mobile money account
	RecipientTypeBasa        RecipientType = "basa"         // South African bank account
)

// String returns the string representation of RecipientType
func (r RecipientType) String() string {
	return string(r)
}

// TransferRecipient represents a transfer recipient
type TransferRecipient struct {
	ID            uint64           `json:"id"`
	Domain        string           `json:"domain"`
	Type          RecipientType    `json:"type"`
	Currency      string           `json:"currency"`
	Name          string           `json:"name"`
	Description   string           `json:"description"`
	Details       RecipientDetails `json:"details"`
	Metadata      map[string]any   `json:"metadata"`
	RecipientCode string           `json:"recipient_code"`
	Active        bool             `json:"active"`
	IsDeleted     bool             `json:"is_deleted"`
	CreatedAt     time.Time        `json:"createdAt"`
	UpdatedAt     time.Time        `json:"updatedAt"`
	Integration   uint64           `json:"integration"`
	Email         *string          `json:"email,omitempty"`
}

// BulkRecipientItem represents a single recipient in a bulk create request
type BulkRecipientItem struct {
	Type              RecipientType  `json:"type"`                         // Required: nuban, ghipss, mobile_money, basa
	Name              string         `json:"name"`                         // Required: recipient's name
	AccountNumber     string         `json:"account_number"`               // Required for all types except authorization
	BankCode          string         `json:"bank_code"`                    // Required for all types except authorization
	Description       *string        `json:"description,omitempty"`        // Optional: description
	Currency          *string        `json:"currency,omitempty"`           // Optional: currency
	AuthorizationCode *string        `json:"authorization_code,omitempty"` // Optional: authorization code
	Metadata          map[string]any `json:"metadata,omitempty"`           // Optional: additional data
}

// BulkCreateResult represents the result of a bulk create operation
type BulkCreateResult struct {
	Success []TransferRecipient `json:"success"` // Successfully created recipients
	Errors  []struct {
		Type          string `json:"type"`
		Name          string `json:"name"`
		AccountNumber string `json:"account_number"`
		BankCode      string `json:"bank_code"`
		Message       string `json:"message"`
	} `json:"errors"` // Failed recipient creations
}
