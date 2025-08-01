package refunds

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

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

// Create initiates a refund on a transaction using a builder
func (c *Client) Create(ctx context.Context, builder *RefundCreateRequestBuilder) (*RefundCreateResponse, error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()

	url := c.baseURL + refundsBasePath
	return net.Post[RefundCreateRequest, RefundCreateData](ctx, c.client, c.secret, url, req)
}
