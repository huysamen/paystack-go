package refunds

import (
	"context"
	"time"

	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
)

type listRequest struct {
	Transaction *string    `json:"transaction,omitempty"`
	Currency    *string    `json:"currency,omitempty"`
	From        *time.Time `json:"from,omitempty"`
	To          *time.Time `json:"to,omitempty"`
	PerPage     *int       `json:"perPage,omitempty"`
	Page        *int       `json:"page,omitempty"`
}

type ListRequestBuilder struct {
	req *listRequest
}

func NewListRequestBuilder() *ListRequestBuilder {
	return &ListRequestBuilder{
		req: &listRequest{},
	}
}

func (b *ListRequestBuilder) Transaction(transaction string) *ListRequestBuilder {
	b.req.Transaction = &transaction

	return b
}

func (b *ListRequestBuilder) Currency(currency string) *ListRequestBuilder {
	b.req.Currency = &currency

	return b
}

func (b *ListRequestBuilder) DateRange(from, to time.Time) *ListRequestBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

func (b *ListRequestBuilder) From(from time.Time) *ListRequestBuilder {
	b.req.From = &from

	return b
}

func (b *ListRequestBuilder) To(to time.Time) *ListRequestBuilder {
	b.req.To = &to

	return b
}

func (b *ListRequestBuilder) PerPage(perPage int) *ListRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *ListRequestBuilder) Page(page int) *ListRequestBuilder {
	b.req.Page = &page

	return b
}

func (b *ListRequestBuilder) Build() *listRequest {
	return b.req
}

// ListRefund represents a refund in list responses with different field types
type ListRefund struct {
	ID             data.Int             `json:"id"`
	Integration    data.Int             `json:"integration"`
	Domain         data.String          `json:"domain"`
	Transaction    data.Int             `json:"transaction"` // Transaction ID, not object in list response
	Dispute        data.NullInt         `json:"dispute"`
	Settlement     data.NullInt         `json:"settlement"`
	Amount         data.Int             `json:"amount"`
	DeductedAmount data.NullInt         `json:"deducted_amount"` // Can be null in list response
	Currency       enums.Currency       `json:"currency"`
	Channel        *enums.RefundChannel `json:"channel"`
	FullyDeducted  data.NullBool        `json:"fully_deducted"` // Can be null in list response
	Status         enums.RefundStatus   `json:"status"`
	RefundedBy     data.String          `json:"refunded_by"`
	RefundedAt     data.NullTime        `json:"refunded_at"`
	ExpectedAt     data.NullTime        `json:"expected_at"`
	CreatedAt      data.Time            `json:"created_at"` // Snake case in list response
	UpdatedAt      data.Time            `json:"updated_at"` // Snake case in list response
	CustomerNote   data.NullString      `json:"customer_note"`
	MerchantNote   data.NullString      `json:"merchant_note"`
}

type ListResponseData = []ListRefund
type ListResponse = types.Response[ListResponseData]

func (c *Client) List(ctx context.Context, builder ListRequestBuilder) (*ListResponse, error) {
	return net.Get[ListResponseData](ctx, c.Client, c.Secret, basePath, c.BaseURL)
}
