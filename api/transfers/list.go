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

type TransferListRequest struct {
	PerPage   *int
	Page      *int
	Recipient *int       // Filter by recipient ID
	From      *time.Time // Start date filter
	To        *time.Time // End date filter
}

type TransferListRequestBuilder struct {
	req *TransferListRequest
}

func NewTransferListRequest() *TransferListRequestBuilder {
	return &TransferListRequestBuilder{
		req: &TransferListRequest{},
	}
}

func (b *TransferListRequestBuilder) PerPage(perPage int) *TransferListRequestBuilder {
	b.req.PerPage = optional.Int(perPage)

	return b
}

func (b *TransferListRequestBuilder) Page(page int) *TransferListRequestBuilder {
	b.req.Page = optional.Int(page)

	return b
}

func (b *TransferListRequestBuilder) Recipient(recipient int) *TransferListRequestBuilder {
	b.req.Recipient = optional.Int(recipient)

	return b
}

func (b *TransferListRequestBuilder) DateRange(from, to time.Time) *TransferListRequestBuilder {
	b.req.From = optional.Time(from)
	b.req.To = optional.Time(to)

	return b
}

func (b *TransferListRequestBuilder) From(from time.Time) *TransferListRequestBuilder {
	b.req.From = optional.Time(from)

	return b
}

func (b *TransferListRequestBuilder) To(to time.Time) *TransferListRequestBuilder {
	b.req.To = optional.Time(to)

	return b
}

func (b *TransferListRequestBuilder) Build() *TransferListRequest {
	return b.req
}

func (r *TransferListRequest) toQuery() string {
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

type ListResponse = types.Response[[]types.Transfer]

func (c *Client) List(ctx context.Context, builder *TransferListRequestBuilder) (*ListResponse, error) {
	req := builder.Build()
	query := ""
	if req != nil {
		query = req.toQuery()
	}

	return net.Get[[]types.Transfer](ctx, c.Client, c.Secret, basePath, query, c.BaseURL)
}
