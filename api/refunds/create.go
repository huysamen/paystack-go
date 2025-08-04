package refunds

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type RefundCreateRequest struct {
	Transaction  string  `json:"transaction"`
	Amount       *int    `json:"amount,omitempty"`
	Currency     *string `json:"currency,omitempty"`
	CustomerNote *string `json:"customer_note,omitempty"`
	MerchantNote *string `json:"merchant_note,omitempty"`
}

type RefundCreateRequestBuilder struct {
	req *RefundCreateRequest
}

func NewRefundCreateRequest(transaction string) *RefundCreateRequestBuilder {
	return &RefundCreateRequestBuilder{
		req: &RefundCreateRequest{
			Transaction: transaction,
		},
	}
}

func (b *RefundCreateRequestBuilder) Amount(amount int) *RefundCreateRequestBuilder {
	b.req.Amount = &amount
	return b
}

func (b *RefundCreateRequestBuilder) Currency(currency string) *RefundCreateRequestBuilder {
	b.req.Currency = &currency
	return b
}

func (b *RefundCreateRequestBuilder) CustomerNote(note string) *RefundCreateRequestBuilder {
	b.req.CustomerNote = &note
	return b
}

func (b *RefundCreateRequestBuilder) MerchantNote(note string) *RefundCreateRequestBuilder {
	b.req.MerchantNote = &note
	return b
}

func (b *RefundCreateRequestBuilder) Build() *RefundCreateRequest {
	return b.req
}

type RefundCreateResponseData struct {
	Transaction *types.Transaction `json:"transaction"`
	Amount      int                `json:"amount"`
	Currency    string             `json:"currency"`
	RefundedBy  string             `json:"refunded_by"`
	RefundedAt  *types.DateTime    `json:"refunded_at"`
	CreatedAt   *types.DateTime    `json:"created_at"`
}

type RefundCreateResponse = types.Response[RefundCreateResponseData]

func (c *Client) Create(ctx context.Context, builder *RefundCreateRequestBuilder) (*RefundCreateResponse, error) {
	return net.Post[RefundCreateRequest, RefundCreateResponseData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
