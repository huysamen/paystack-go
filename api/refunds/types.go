package refunds

import (
	"time"

	"github.com/huysamen/paystack-go/types"
)

// RefundStatus represents the status of a refund
type RefundStatus string

const (
	RefundStatusPending   RefundStatus = "pending"
	RefundStatusProcessed RefundStatus = "processed"
	RefundStatusFailed    RefundStatus = "failed"
)

// String returns the string representation of RefundStatus
func (s RefundStatus) String() string {
	return string(s)
}

// RefundChannel represents the payment channel for a refund
type RefundChannel string

const (
	RefundChannelCard         RefundChannel = "card"
	RefundChannelBank         RefundChannel = "bank"
	RefundChannelUSSD         RefundChannel = "ussd"
	RefundChannelQR           RefundChannel = "qr"
	RefundChannelMobileMoney  RefundChannel = "mobile_money"
	RefundChannelBankTransfer RefundChannel = "bank_transfer"
	RefundChannelApplePay     RefundChannel = "apple_pay"
	RefundChannelMigs         RefundChannel = "migs"
)

// String returns the string representation of RefundChannel
func (c RefundChannel) String() string {
	return string(c)
}

// RefundCreateRequest represents the request payload for creating a refund
type RefundCreateRequest struct {
	Transaction  string  `json:"transaction"`
	Amount       *int    `json:"amount,omitempty"`
	Currency     *string `json:"currency,omitempty"`
	CustomerNote *string `json:"customer_note,omitempty"`
	MerchantNote *string `json:"merchant_note,omitempty"`
}

// RefundCreateRequestBuilder provides a fluent interface for building RefundCreateRequest
type RefundCreateRequestBuilder struct {
	req *RefundCreateRequest
}

// NewRefundCreateRequest creates a new builder for RefundCreateRequest
func NewRefundCreateRequest(transaction string) *RefundCreateRequestBuilder {
	return &RefundCreateRequestBuilder{
		req: &RefundCreateRequest{
			Transaction: transaction,
		},
	}
}

// Amount sets the refund amount (optional - defaults to full transaction amount)
func (b *RefundCreateRequestBuilder) Amount(amount int) *RefundCreateRequestBuilder {
	b.req.Amount = &amount
	return b
}

// Currency sets the currency for the refund
func (b *RefundCreateRequestBuilder) Currency(currency string) *RefundCreateRequestBuilder {
	b.req.Currency = &currency
	return b
}

// CustomerNote sets a note for the customer
func (b *RefundCreateRequestBuilder) CustomerNote(note string) *RefundCreateRequestBuilder {
	b.req.CustomerNote = &note
	return b
}

// MerchantNote sets a note for the merchant
func (b *RefundCreateRequestBuilder) MerchantNote(note string) *RefundCreateRequestBuilder {
	b.req.MerchantNote = &note
	return b
}

// Build returns the constructed RefundCreateRequest
func (b *RefundCreateRequestBuilder) Build() *RefundCreateRequest {
	return b.req
}

// RefundListRequest represents the request payload for listing refunds
type RefundListRequest struct {
	Transaction *string    `json:"transaction,omitempty"`
	Currency    *string    `json:"currency,omitempty"`
	From        *time.Time `json:"from,omitempty"`
	To          *time.Time `json:"to,omitempty"`
	PerPage     *int       `json:"perPage,omitempty"`
	Page        *int       `json:"page,omitempty"`
}

// RefundListRequestBuilder provides a fluent interface for building RefundListRequest
type RefundListRequestBuilder struct {
	req *RefundListRequest
}

// NewRefundListRequest creates a new builder for RefundListRequest
func NewRefundListRequest() *RefundListRequestBuilder {
	return &RefundListRequestBuilder{
		req: &RefundListRequest{},
	}
}

// Transaction filters by transaction reference
func (b *RefundListRequestBuilder) Transaction(transaction string) *RefundListRequestBuilder {
	b.req.Transaction = &transaction
	return b
}

// Currency filters by currency
func (b *RefundListRequestBuilder) Currency(currency string) *RefundListRequestBuilder {
	b.req.Currency = &currency
	return b
}

// DateRange sets both start and end date filters
func (b *RefundListRequestBuilder) DateRange(from, to time.Time) *RefundListRequestBuilder {
	b.req.From = &from
	b.req.To = &to
	return b
}

// From sets the start date filter
func (b *RefundListRequestBuilder) From(from time.Time) *RefundListRequestBuilder {
	b.req.From = &from
	return b
}

// To sets the end date filter
func (b *RefundListRequestBuilder) To(to time.Time) *RefundListRequestBuilder {
	b.req.To = &to
	return b
}

// PerPage sets the number of records per page
func (b *RefundListRequestBuilder) PerPage(perPage int) *RefundListRequestBuilder {
	b.req.PerPage = &perPage
	return b
}

// Page sets the page number
func (b *RefundListRequestBuilder) Page(page int) *RefundListRequestBuilder {
	b.req.Page = &page
	return b
}

// Build returns the constructed RefundListRequest
func (b *RefundListRequestBuilder) Build() *RefundListRequest {
	return b.req
}

// RefundTransaction represents transaction details in a refund
type RefundTransaction struct {
	ID              int                  `json:"id"`
	Domain          string               `json:"domain"`
	Reference       string               `json:"reference"`
	Amount          int                  `json:"amount"`
	PaidAt          *types.DateTime      `json:"paid_at"`
	Channel         string               `json:"channel"`
	Currency        string               `json:"currency"`
	Authorization   *RefundAuthorization `json:"authorization"`
	Customer        *RefundCustomer      `json:"customer"`
	Plan            any                  `json:"plan"`
	Split           any                  `json:"split"`
	OrderID         any                  `json:"order_id"`
	CreatedAt       *types.DateTime      `json:"created_at"`
	RequestedAmount int                  `json:"requested_amount"`
	Source          any                  `json:"source"`
	SourceDetails   any                  `json:"source_details"`
}

// RefundAuthorization represents authorization details in a refund
type RefundAuthorization struct {
	ExpMonth    any `json:"exp_month"`
	ExpYear     any `json:"exp_year"`
	AccountName any `json:"account_name"`
}

// RefundCustomer represents customer details in a refund
type RefundCustomer struct {
	ID                       int     `json:"id"`
	FirstName                *string `json:"first_name"`
	LastName                 *string `json:"last_name"`
	Email                    string  `json:"email"`
	CustomerCode             string  `json:"customer_code"`
	Phone                    *string `json:"phone"`
	Metadata                 any     `json:"metadata"`
	RiskAction               string  `json:"risk_action"`
	InternationalFormatPhone any     `json:"international_format_phone"`
}

// Refund represents a refund object
type Refund struct {
	ID             int             `json:"id"`
	Integration    int             `json:"integration"`
	Domain         string          `json:"domain"`
	Transaction    int             `json:"transaction"`
	Dispute        *int            `json:"dispute"`
	Settlement     *int            `json:"settlement"`
	Amount         int             `json:"amount"`
	DeductedAmount int             `json:"deducted_amount"`
	Currency       string          `json:"currency"`
	Channel        RefundChannel   `json:"channel"`
	FullyDeducted  bool            `json:"fully_deducted"`
	Status         RefundStatus    `json:"status"`
	RefundedBy     string          `json:"refunded_by"`
	RefundedAt     *types.DateTime `json:"refunded_at"`
	ExpectedAt     *types.DateTime `json:"expected_at"`
	CreatedAt      *types.DateTime `json:"created_at"`
	UpdatedAt      *types.DateTime `json:"updated_at"`
	CustomerNote   *string         `json:"customer_note"`
	MerchantNote   *string         `json:"merchant_note"`
}

// RefundCreateData represents the data returned when creating a refund
type RefundCreateData struct {
	Transaction *RefundTransaction `json:"transaction"`
	Amount      int                `json:"amount"`
	Currency    string             `json:"currency"`
	RefundedBy  string             `json:"refunded_by"`
	RefundedAt  *types.DateTime    `json:"refunded_at"`
	CreatedAt   *types.DateTime    `json:"created_at"`
}

// Response type aliases using generic types
type RefundCreateResponse = types.Response[RefundCreateData]
type RefundListResponse = types.Response[[]Refund]
type RefundFetchResponse = types.Response[Refund]
