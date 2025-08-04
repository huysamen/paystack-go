package transactions

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/optional"
	"github.com/huysamen/paystack-go/types"
)

type TransactionListRequest struct {
	PerPage    *int
	Page       *int
	Customer   *uint64
	TerminalID *string
	Status     *string
	From       *time.Time
	To         *time.Time
	Amount     *int
}

type TransactionListRequestBuilder struct {
	req *TransactionListRequest
}

func NewTransactionListRequest() *TransactionListRequestBuilder {
	return &TransactionListRequestBuilder{
		req: &TransactionListRequest{},
	}
}

func (b *TransactionListRequestBuilder) PerPage(perPage int) *TransactionListRequestBuilder {
	b.req.PerPage = optional.Int(perPage)

	return b
}

func (b *TransactionListRequestBuilder) Page(page int) *TransactionListRequestBuilder {
	b.req.Page = optional.Int(page)

	return b
}

func (b *TransactionListRequestBuilder) Customer(customer uint64) *TransactionListRequestBuilder {
	b.req.Customer = optional.Uint64(customer)

	return b
}

func (b *TransactionListRequestBuilder) TerminalID(terminalID string) *TransactionListRequestBuilder {
	b.req.TerminalID = optional.String(terminalID)

	return b
}

func (b *TransactionListRequestBuilder) Status(status string) *TransactionListRequestBuilder {
	b.req.Status = optional.String(status)

	return b
}

func (b *TransactionListRequestBuilder) DateRange(from, to time.Time) *TransactionListRequestBuilder {
	b.req.From = optional.Time(from)
	b.req.To = optional.Time(to)

	return b
}

func (b *TransactionListRequestBuilder) From(from time.Time) *TransactionListRequestBuilder {
	b.req.From = optional.Time(from)

	return b
}

func (b *TransactionListRequestBuilder) To(to time.Time) *TransactionListRequestBuilder {
	b.req.To = optional.Time(to)

	return b
}

func (b *TransactionListRequestBuilder) Amount(amount int) *TransactionListRequestBuilder {
	b.req.Amount = optional.Int(amount)

	return b
}

func (b *TransactionListRequestBuilder) Build() *TransactionListRequest {
	return b.req
}

func (r *TransactionListRequest) toQuery() string {
	params := url.Values{}

	if r.PerPage != nil {
		params.Add("perPage", strconv.Itoa(*r.PerPage))
	}
	if r.Page != nil {
		params.Add("page", strconv.Itoa(*r.Page))
	}
	if r.Customer != nil {
		params.Add("customer", strconv.FormatUint(*r.Customer, 10))
	}
	if r.TerminalID != nil {
		params.Add("terminalid", *r.TerminalID)
	}
	if r.Status != nil {
		params.Add("status", *r.Status)
	}
	if r.From != nil {
		params.Add("from", r.From.Format("2006-01-02T15:04:05.999Z"))
	}
	if r.To != nil {
		params.Add("to", r.To.Format("2006-01-02T15:04:05.999Z"))
	}
	if r.Amount != nil {
		params.Add("amount", strconv.Itoa(*r.Amount))
	}

	return params.Encode()
}

type ListResponse = types.Response[[]types.Transaction]

func (c *Client) List(ctx context.Context, builder *TransactionListRequestBuilder) (*ListResponse, error) {
	req := builder.Build()
	query := ""
	if req != nil {
		query = req.toQuery()
	}
	return net.Get[[]types.Transaction](ctx, c.Client, c.Secret, basePath, query, c.BaseURL)
}
