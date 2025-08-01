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
	Transaction  string  `json:"transaction" validate:"required"`
	Amount       *int    `json:"amount,omitempty"`
	Currency     *string `json:"currency,omitempty"`
	CustomerNote *string `json:"customer_note,omitempty"`
	MerchantNote *string `json:"merchant_note,omitempty"`
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
