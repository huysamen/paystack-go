package refunds

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

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

// fixme: this is not correct
// RefundCreateResponseData represents the data returned when creating a refund
type RefundCreateResponseData struct {
	Transaction *types.Transaction `json:"transaction"`
	Amount      int                `json:"amount"`
	Currency    string             `json:"currency"`
	RefundedBy  string             `json:"refunded_by"`
	RefundedAt  *types.DateTime    `json:"refunded_at"`
	CreatedAt   *types.DateTime    `json:"created_at"`
}

// RefundCreateRequest represents the request to create a refund
type RefundCreateResponse = types.Response[RefundCreateResponseData]

// Create initiates a refund on a transaction using a builder
func (c *Client) Create(ctx context.Context, builder *RefundCreateRequestBuilder) (*RefundCreateResponse, error) {
	return net.Post[RefundCreateRequest, RefundCreateResponseData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
