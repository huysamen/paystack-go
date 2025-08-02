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

// TransferRecipientCreateRequestBuilder provides a fluent interface for building TransferRecipientCreateRequest
type TransferRecipientCreateRequestBuilder struct {
	req *TransferRecipientCreateRequest
}

// NewTransferRecipientCreateRequest creates a new builder for TransferRecipientCreateRequest
func NewTransferRecipientCreateRequest(recipientType RecipientType, name, accountNumber, bankCode string) *TransferRecipientCreateRequestBuilder {
	return &TransferRecipientCreateRequestBuilder{
		req: &TransferRecipientCreateRequest{
			Type:          recipientType,
			Name:          name,
			AccountNumber: accountNumber,
			BankCode:      bankCode,
		},
	}
}

// Description sets the recipient description
func (b *TransferRecipientCreateRequestBuilder) Description(description string) *TransferRecipientCreateRequestBuilder {
	b.req.Description = &description
	return b
}

// Currency sets the currency
func (b *TransferRecipientCreateRequestBuilder) Currency(currency string) *TransferRecipientCreateRequestBuilder {
	b.req.Currency = &currency
	return b
}

// AuthorizationCode sets the authorization code
func (b *TransferRecipientCreateRequestBuilder) AuthorizationCode(authCode string) *TransferRecipientCreateRequestBuilder {
	b.req.AuthorizationCode = &authCode
	return b
}

// Metadata sets the recipient metadata
func (b *TransferRecipientCreateRequestBuilder) Metadata(metadata map[string]any) *TransferRecipientCreateRequestBuilder {
	b.req.Metadata = metadata
	return b
}

// Build returns the constructed TransferRecipientCreateRequest
func (b *TransferRecipientCreateRequestBuilder) Build() *TransferRecipientCreateRequest {
	return b.req
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

// TransferRecipientListRequestBuilder provides a fluent interface for building TransferRecipientListRequest
type TransferRecipientListRequestBuilder struct {
	req *TransferRecipientListRequest
}

// NewTransferRecipientListRequest creates a new builder for TransferRecipientListRequest
func NewTransferRecipientListRequest() *TransferRecipientListRequestBuilder {
	return &TransferRecipientListRequestBuilder{
		req: &TransferRecipientListRequest{},
	}
}

// PerPage sets the number of recipients per page
func (b *TransferRecipientListRequestBuilder) PerPage(perPage int) *TransferRecipientListRequestBuilder {
	b.req.PerPage = &perPage
	return b
}

// Page sets the page number
func (b *TransferRecipientListRequestBuilder) Page(page int) *TransferRecipientListRequestBuilder {
	b.req.Page = &page
	return b
}

// From sets the start date filter
func (b *TransferRecipientListRequestBuilder) From(from time.Time) *TransferRecipientListRequestBuilder {
	b.req.From = &from
	return b
}

// To sets the end date filter
func (b *TransferRecipientListRequestBuilder) To(to time.Time) *TransferRecipientListRequestBuilder {
	b.req.To = &to
	return b
}

// DateRange sets both from and to dates for convenience
func (b *TransferRecipientListRequestBuilder) DateRange(from, to time.Time) *TransferRecipientListRequestBuilder {
	b.req.From = &from
	b.req.To = &to
	return b
}

// Build returns the constructed TransferRecipientListRequest
func (b *TransferRecipientListRequestBuilder) Build() *TransferRecipientListRequest {
	return b.req
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

// TransferRecipientUpdateRequestBuilder provides a fluent interface for building TransferRecipientUpdateRequest
type TransferRecipientUpdateRequestBuilder struct {
	req *TransferRecipientUpdateRequest
}

// NewTransferRecipientUpdateRequest creates a new builder for TransferRecipientUpdateRequest
func NewTransferRecipientUpdateRequest(name string) *TransferRecipientUpdateRequestBuilder {
	return &TransferRecipientUpdateRequestBuilder{
		req: &TransferRecipientUpdateRequest{
			Name: name,
		},
	}
}

// Email sets the recipient email
func (b *TransferRecipientUpdateRequestBuilder) Email(email string) *TransferRecipientUpdateRequestBuilder {
	b.req.Email = &email
	return b
}

// Build returns the constructed TransferRecipientUpdateRequest
func (b *TransferRecipientUpdateRequestBuilder) Build() *TransferRecipientUpdateRequest {
	return b.req
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

// BulkCreateTransferRecipientRequestBuilder builds a BulkCreateTransferRecipientRequest
type BulkCreateTransferRecipientRequestBuilder struct {
	request BulkCreateTransferRecipientRequest
}

// NewBulkCreateTransferRecipientRequestBuilder creates a new builder
func NewBulkCreateTransferRecipientRequestBuilder() *BulkCreateTransferRecipientRequestBuilder {
	return &BulkCreateTransferRecipientRequestBuilder{}
}

// Batch sets the batch of recipients
func (b *BulkCreateTransferRecipientRequestBuilder) Batch(batch []BulkRecipientItem) *BulkCreateTransferRecipientRequestBuilder {
	b.request.Batch = batch
	return b
}

// AddRecipient adds a single recipient to the batch
func (b *BulkCreateTransferRecipientRequestBuilder) AddRecipient(item BulkRecipientItem) *BulkCreateTransferRecipientRequestBuilder {
	if b.request.Batch == nil {
		b.request.Batch = make([]BulkRecipientItem, 0)
	}
	b.request.Batch = append(b.request.Batch, item)
	return b
}

// Build returns the built BulkCreateTransferRecipientRequest
func (b *BulkCreateTransferRecipientRequestBuilder) Build() *BulkCreateTransferRecipientRequest {
	return &b.request
}
