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

type listRequest struct {
	PerPage    *int
	Page       *int
	Customer   *uint64
	TerminalID *string
	Status     *string
	From       *time.Time
	To         *time.Time
	Amount     *int
}

type ListRequestBuilder struct {
	req *listRequest
}

func NewListRequestBuilder() *ListRequestBuilder {
	return &ListRequestBuilder{
		req: &listRequest{},
	}
}

func (b *ListRequestBuilder) PerPage(perPage int) *ListRequestBuilder {
	b.req.PerPage = optional.Int(perPage)

	return b
}

func (b *ListRequestBuilder) Page(page int) *ListRequestBuilder {
	b.req.Page = optional.Int(page)

	return b
}

func (b *ListRequestBuilder) Customer(customer uint64) *ListRequestBuilder {
	b.req.Customer = optional.Uint64(customer)

	return b
}

func (b *ListRequestBuilder) TerminalID(terminalID string) *ListRequestBuilder {
	b.req.TerminalID = optional.String(terminalID)

	return b
}

func (b *ListRequestBuilder) Status(status string) *ListRequestBuilder {
	b.req.Status = optional.String(status)

	return b
}

func (b *ListRequestBuilder) DateRange(from, to time.Time) *ListRequestBuilder {
	b.req.From = optional.Time(from)
	b.req.To = optional.Time(to)

	return b
}

func (b *ListRequestBuilder) From(from time.Time) *ListRequestBuilder {
	b.req.From = optional.Time(from)

	return b
}

func (b *ListRequestBuilder) To(to time.Time) *ListRequestBuilder {
	b.req.To = optional.Time(to)

	return b
}

func (b *ListRequestBuilder) Amount(amount int) *ListRequestBuilder {
	b.req.Amount = optional.Int(amount)

	return b
}

func (b *ListRequestBuilder) Build() *listRequest {
	return b.req
}

func (r *listRequest) toQuery() string {
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

type ListResponseData = []types.Transaction
type ListResponse = types.Response[ListResponseData]

func (c *Client) List(ctx context.Context, builder ListRequestBuilder) (*ListResponse, error) {
	path := basePath

	req := builder.Build()
	if query := req.toQuery(); query != "" {
		path += "?" + query
	}

	return net.Get[ListResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
