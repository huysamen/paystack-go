package transfer_recipients

import (
	"time"

	"github.com/huysamen/paystack-go/types"
)

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

// RecipientDetails contains account details for a transfer recipient
type RecipientDetails struct {
	AccountNumber *string `json:"account_number,omitempty"`
	AccountName   *string `json:"account_name,omitempty"`
	BankCode      *string `json:"bank_code,omitempty"`
	BankName      *string `json:"bank_name,omitempty"`
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

// Transfer Recipient Create

// TransferRecipientCreateRequest represents the request to create a transfer recipient
type TransferRecipientCreateRequest struct {
	Type              RecipientType  `json:"type"`                         // Required: nuban, ghipss, mobile_money, basa
	Name              string         `json:"name"`                         // Required: recipient's name
	AccountNumber     string         `json:"account_number"`               // Required for all types except authorization
	BankCode          string         `json:"bank_code"`                    // Required for all types except authorization
	Description       *string        `json:"description,omitempty"`        // Optional: description
	Currency          *string        `json:"currency,omitempty"`           // Optional: currency
	AuthorizationCode *string        `json:"authorization_code,omitempty"` // Optional: authorization code
	Metadata          map[string]any `json:"metadata,omitempty"`           // Optional: additional data
}

// TransferRecipientCreateResponse represents the response from creating a transfer recipient
type TransferRecipientCreateResponse struct {
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Data    TransferRecipient `json:"data"`
}

// Transfer Recipient Bulk Create

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

// BulkCreateTransferRecipientRequest represents the request to create multiple transfer recipients
type BulkCreateTransferRecipientRequest struct {
	Batch []BulkRecipientItem `json:"batch"` // Required: list of recipients
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

// BulkCreateTransferRecipientResponse represents the response from bulk creating transfer recipients
type BulkCreateTransferRecipientResponse struct {
	Status  bool             `json:"status"`
	Message string           `json:"message"`
	Data    BulkCreateResult `json:"data"`
}

// Transfer Recipient List

// TransferRecipientListRequest represents the request to list transfer recipients
type TransferRecipientListRequest struct {
	PerPage *int       `json:"perPage,omitempty"` // Optional: records per page (default: 50)
	Page    *int       `json:"page,omitempty"`    // Optional: page number (default: 1)
	From    *time.Time `json:"from,omitempty"`    // Optional: start date filter
	To      *time.Time `json:"to,omitempty"`      // Optional: end date filter
}

// TransferRecipientListResponse represents the response from listing transfer recipients
type TransferRecipientListResponse struct {
	Status  bool                `json:"status"`
	Message string              `json:"message"`
	Data    []TransferRecipient `json:"data"`
	Meta    types.Meta          `json:"meta"`
}

// Transfer Recipient Fetch

// TransferRecipientFetchResponse represents the response from fetching a transfer recipient
type TransferRecipientFetchResponse struct {
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Data    TransferRecipient `json:"data"`
}

// Transfer Recipient Update

// TransferRecipientUpdateRequest represents the request to update a transfer recipient
type TransferRecipientUpdateRequest struct {
	Name  string  `json:"name"`            // Required: recipient name
	Email *string `json:"email,omitempty"` // Optional: email address
}

// TransferRecipientUpdateResponse represents the response from updating a transfer recipient
type TransferRecipientUpdateResponse struct {
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Data    TransferRecipient `json:"data"`
}

// Transfer Recipient Delete

// TransferRecipientDeleteResponse represents the response from deleting a transfer recipient
type TransferRecipientDeleteResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
