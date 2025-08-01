package settlements

import (
	"time"

	"github.com/huysamen/paystack-go/types"
)

// SettlementStatus represents the status of a settlement
type SettlementStatus string

const (
	SettlementStatusSuccess    SettlementStatus = "success"    // Successfully settled
	SettlementStatusProcessing SettlementStatus = "processing" // Currently being processed
	SettlementStatusPending    SettlementStatus = "pending"    // Pending settlement
	SettlementStatusFailed     SettlementStatus = "failed"     // Failed settlement
)

// String returns the string representation of SettlementStatus
func (s SettlementStatus) String() string {
	return string(s)
}

// Settlement represents a settlement record
type Settlement struct {
	ID              uint64           `json:"id"`
	Domain          string           `json:"domain"`
	Status          SettlementStatus `json:"status"`
	Currency        string           `json:"currency"`
	Integration     uint64           `json:"integration"`
	TotalAmount     int64            `json:"total_amount"`     // Amount after fees in kobo
	EffectiveAmount int64            `json:"effective_amount"` // Amount actually settled in kobo
	TotalFees       int64            `json:"total_fees"`       // Total fees charged in kobo
	TotalProcessed  int64            `json:"total_processed"`  // Total amount processed in kobo
	Deductions      map[string]any   `json:"deductions"`       // Any deductions applied
	SettlementDate  time.Time        `json:"settlement_date"`  // Date settlement was made
	SettledBy       *string          `json:"settled_by"`       // Who processed the settlement
	CreatedAt       time.Time        `json:"createdAt"`        // When settlement record was created
	UpdatedAt       time.Time        `json:"updatedAt"`        // When settlement was last updated
}

// SettlementTransaction represents a transaction within a settlement
type SettlementTransaction struct {
	ID              uint64         `json:"id"`
	Domain          string         `json:"domain"`
	Status          string         `json:"status"`
	Reference       string         `json:"reference"`
	Amount          int64          `json:"amount"` // Transaction amount in kobo
	Message         string         `json:"message"`
	GatewayResponse string         `json:"gateway_response"`
	PaidAt          *time.Time     `json:"paid_at"`
	CreatedAt       time.Time      `json:"created_at"`
	Channel         string         `json:"channel"`
	Currency        string         `json:"currency"`
	IPAddress       string         `json:"ip_address"`
	Metadata        map[string]any `json:"metadata"`
	Log             map[string]any `json:"log"`
	Fees            int64          `json:"fees"`          // Fees charged for this transaction
	FeesSplit       map[string]any `json:"fees_split"`    // Breakdown of fees
	Customer        map[string]any `json:"customer"`      // Customer information
	Authorization   map[string]any `json:"authorization"` // Authorization details
	Plan            map[string]any `json:"plan"`          // Plan details if subscription
	Subaccount      map[string]any `json:"subaccount"`    // Subaccount details if applicable
}

// Settlement List

// SettlementListRequest represents the request to list settlements
type SettlementListRequest struct {
	PerPage    *int              `json:"perPage,omitempty"`    // Optional: records per page (default: 50)
	Page       *int              `json:"page,omitempty"`       // Optional: page number (default: 1)
	Status     *SettlementStatus `json:"status,omitempty"`     // Optional: filter by status
	Subaccount *string           `json:"subaccount,omitempty"` // Optional: filter by subaccount ID (use "none" for main account only)
	From       *time.Time        `json:"from,omitempty"`       // Optional: start date filter
	To         *time.Time        `json:"to,omitempty"`         // Optional: end date filter
}

// SettlementListRequestBuilder provides a fluent interface for building SettlementListRequest
type SettlementListRequestBuilder struct {
	req *SettlementListRequest
}

// NewSettlementListRequest creates a new builder for SettlementListRequest
func NewSettlementListRequest() *SettlementListRequestBuilder {
	return &SettlementListRequestBuilder{
		req: &SettlementListRequest{},
	}
}

// PerPage sets the number of records per page
func (b *SettlementListRequestBuilder) PerPage(perPage int) *SettlementListRequestBuilder {
	b.req.PerPage = &perPage
	return b
}

// Page sets the page number
func (b *SettlementListRequestBuilder) Page(page int) *SettlementListRequestBuilder {
	b.req.Page = &page
	return b
}

// Status filters by settlement status
func (b *SettlementListRequestBuilder) Status(status SettlementStatus) *SettlementListRequestBuilder {
	b.req.Status = &status
	return b
}

// Subaccount filters by subaccount ID (use "none" for main account only)
func (b *SettlementListRequestBuilder) Subaccount(subaccount string) *SettlementListRequestBuilder {
	b.req.Subaccount = &subaccount
	return b
}

// MainAccountOnly filters for main account settlements only
func (b *SettlementListRequestBuilder) MainAccountOnly() *SettlementListRequestBuilder {
	none := "none"
	b.req.Subaccount = &none
	return b
}

// DateRange sets both start and end date filters
func (b *SettlementListRequestBuilder) DateRange(from, to time.Time) *SettlementListRequestBuilder {
	b.req.From = &from
	b.req.To = &to
	return b
}

// From sets the start date filter
func (b *SettlementListRequestBuilder) From(from time.Time) *SettlementListRequestBuilder {
	b.req.From = &from
	return b
}

// To sets the end date filter
func (b *SettlementListRequestBuilder) To(to time.Time) *SettlementListRequestBuilder {
	b.req.To = &to
	return b
}

// Build returns the constructed SettlementListRequest
func (b *SettlementListRequestBuilder) Build() *SettlementListRequest {
	return b.req
}

// SettlementListResponse represents the response from listing settlements
type SettlementListResponse struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Data    []Settlement `json:"data"`
	Meta    types.Meta   `json:"meta"`
}

// Settlement Transactions

// SettlementTransactionListRequest represents the request to list settlement transactions
type SettlementTransactionListRequest struct {
	PerPage *int       `json:"perPage,omitempty"` // Optional: records per page (default: 50)
	Page    *int       `json:"page,omitempty"`    // Optional: page number (default: 1)
	From    *time.Time `json:"from,omitempty"`    // Optional: start date filter
	To      *time.Time `json:"to,omitempty"`      // Optional: end date filter
}

// SettlementTransactionListRequestBuilder provides a fluent interface for building SettlementTransactionListRequest
type SettlementTransactionListRequestBuilder struct {
	req *SettlementTransactionListRequest
}

// NewSettlementTransactionListRequest creates a new builder for SettlementTransactionListRequest
func NewSettlementTransactionListRequest() *SettlementTransactionListRequestBuilder {
	return &SettlementTransactionListRequestBuilder{
		req: &SettlementTransactionListRequest{},
	}
}

// PerPage sets the number of records per page
func (b *SettlementTransactionListRequestBuilder) PerPage(perPage int) *SettlementTransactionListRequestBuilder {
	b.req.PerPage = &perPage
	return b
}

// Page sets the page number
func (b *SettlementTransactionListRequestBuilder) Page(page int) *SettlementTransactionListRequestBuilder {
	b.req.Page = &page
	return b
}

// DateRange sets both start and end date filters
func (b *SettlementTransactionListRequestBuilder) DateRange(from, to time.Time) *SettlementTransactionListRequestBuilder {
	b.req.From = &from
	b.req.To = &to
	return b
}

// From sets the start date filter
func (b *SettlementTransactionListRequestBuilder) From(from time.Time) *SettlementTransactionListRequestBuilder {
	b.req.From = &from
	return b
}

// To sets the end date filter
func (b *SettlementTransactionListRequestBuilder) To(to time.Time) *SettlementTransactionListRequestBuilder {
	b.req.To = &to
	return b
}

// Build returns the constructed SettlementTransactionListRequest
func (b *SettlementTransactionListRequestBuilder) Build() *SettlementTransactionListRequest {
	return b.req
}

// SettlementTransactionListResponse represents the response from listing settlement transactions
type SettlementTransactionListResponse struct {
	Status  bool                    `json:"status"`
	Message string                  `json:"message"`
	Data    []SettlementTransaction `json:"data"`
	Meta    types.Meta              `json:"meta"`
}
