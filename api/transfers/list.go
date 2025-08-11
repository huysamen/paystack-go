package transfers

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/optional"
	"github.com/huysamen/paystack-go/types"
)

type listRequest struct {
	PerPage   *int
	Page      *int
	Recipient *int       // Filter by recipient ID
	From      *time.Time // Start date filter
	To        *time.Time // End date filter
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

func (b *ListRequestBuilder) Recipient(recipient int) *ListRequestBuilder {
	b.req.Recipient = optional.Int(recipient)

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

func (b *ListRequestBuilder) Build() *listRequest {
	return b.req
}

func (r *listRequest) toQuery() string {
	params := url.Values{}

	if r.PerPage != nil {
		params.Add("perPage", fmt.Sprintf("%d", *r.PerPage))
	}
	if r.Page != nil {
		params.Add("page", fmt.Sprintf("%d", *r.Page))
	}
	if r.Recipient != nil {
		params.Add("recipient", fmt.Sprintf("%d", *r.Recipient))
	}
	if r.From != nil {
		params.Add("from", r.From.Format("2006-01-02T15:04:05.999Z"))
	}
	if r.To != nil {
		params.Add("to", r.To.Format("2006-01-02T15:04:05.999Z"))
	}

	return params.Encode()
}

type ListResponseData = []types.Transfer
type ListResponse = types.Response[ListResponseData]

func (c *Client) List(ctx context.Context, builder ListRequestBuilder) (*ListResponse, error) {
	req := builder.Build()
	query := ""

	if req != nil {
		query = req.toQuery()
	}

	return net.Get[ListResponseData](ctx, c.Client, c.Secret, basePath, query, c.BaseURL)
}
