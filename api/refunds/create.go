package refunds

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
)

type createRequest struct {
	Transaction  string  `json:"transaction"`
	Amount       *int    `json:"amount,omitempty"`
	Currency     *string `json:"currency,omitempty"`
	CustomerNote *string `json:"customer_note,omitempty"`
	MerchantNote *string `json:"merchant_note,omitempty"`
}

type CreateRequestBuilder struct {
	req *createRequest
}

func NewCreateRequestBuilder(transaction string) *CreateRequestBuilder {
	return &CreateRequestBuilder{
		req: &createRequest{
			Transaction: transaction,
		},
	}
}

func (b *CreateRequestBuilder) Amount(amount int) *CreateRequestBuilder {
	b.req.Amount = &amount
	return b
}

func (b *CreateRequestBuilder) Currency(currency string) *CreateRequestBuilder {
	b.req.Currency = &currency
	return b
}

func (b *CreateRequestBuilder) CustomerNote(note string) *CreateRequestBuilder {
	b.req.CustomerNote = &note
	return b
}

func (b *CreateRequestBuilder) MerchantNote(note string) *CreateRequestBuilder {
	b.req.MerchantNote = &note
	return b
}

func (b *CreateRequestBuilder) Build() *createRequest {
	return b.req
}

type CreateResponseData struct {
	Transaction *types.Transaction `json:"transaction"`
	Amount      int                `json:"amount"`
	Currency    string             `json:"currency"`
	RefundedBy  string             `json:"refunded_by"`
	RefundedAt  data.NullTime      `json:"refunded_at"`
	CreatedAt   data.NullTime      `json:"created_at"`
}

type CreateResponse = types.Response[CreateResponseData]

func (c *Client) Create(ctx context.Context, builder CreateRequestBuilder) (*CreateResponse, error) {
	return net.Post[createRequest, CreateResponseData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
