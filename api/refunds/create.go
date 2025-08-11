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
	Transaction    *types.Transaction `json:"transaction"`
	Integration    data.Int           `json:"integration"`
	DeductedAmount data.Int           `json:"deducted_amount"`
	Channel        *string            `json:"channel"`
	MerchantNote   data.NullString    `json:"merchant_note"`
	CustomerNote   data.NullString    `json:"customer_note"`
	Status         data.String        `json:"status"`
	RefundedBy     data.String        `json:"refunded_by"`
	ExpectedAt     data.NullTime      `json:"expected_at"`
	Currency       data.String        `json:"currency"`
	Domain         data.String        `json:"domain"`
	Amount         data.Int           `json:"amount"`
	FullyDeducted  data.Bool          `json:"fully_deducted"`
	ID             data.Int           `json:"id"`
	CreatedAt      data.NullTime      `json:"createdAt"`
	UpdatedAt      data.NullTime      `json:"updatedAt"`
}

type CreateResponse = types.Response[CreateResponseData]

func (c *Client) Create(ctx context.Context, builder CreateRequestBuilder) (*CreateResponse, error) {
	return net.Post[createRequest, CreateResponseData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
