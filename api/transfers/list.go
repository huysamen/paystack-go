package transfers

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/options"
	"github.com/huysamen/paystack-go/types"
)

type TransferListRequest struct {
	// Optional
	PerPage   *int
	Page      *int
	Recipient *int       // Filter by recipient ID
	From      *time.Time // Start date filter
	To        *time.Time // End date filter
}

// TransferListRequestBuilder provides a fluent interface for building TransferListRequest
type TransferListRequestBuilder struct {
	req *TransferListRequest
}

// NewTransferListRequest creates a new builder for TransferListRequest
func NewTransferListRequest() *TransferListRequestBuilder {
	return &TransferListRequestBuilder{
		req: &TransferListRequest{},
	}
}

// PerPage sets the number of records per page
func (b *TransferListRequestBuilder) PerPage(perPage int) *TransferListRequestBuilder {
	b.req.PerPage = options.Int(perPage)
	return b
}

// Page sets the page number
func (b *TransferListRequestBuilder) Page(page int) *TransferListRequestBuilder {
	b.req.Page = options.Int(page)
	return b
}

// Recipient filters by recipient ID
func (b *TransferListRequestBuilder) Recipient(recipient int) *TransferListRequestBuilder {
	b.req.Recipient = options.Int(recipient)
	return b
}

// DateRange sets both start and end date filters
func (b *TransferListRequestBuilder) DateRange(from, to time.Time) *TransferListRequestBuilder {
	b.req.From = options.Time(from)
	b.req.To = options.Time(to)
	return b
}

// From sets the start date filter
func (b *TransferListRequestBuilder) From(from time.Time) *TransferListRequestBuilder {
	b.req.From = options.Time(from)
	return b
}

// To sets the end date filter
func (b *TransferListRequestBuilder) To(to time.Time) *TransferListRequestBuilder {
	b.req.To = options.Time(to)
	return b
}

// Build returns the constructed TransferListRequest
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

type TransferListResponse struct {
	Data []Transfer `json:"data"`
	Meta types.Meta `json:"meta"`
}

// List lists transfers using a builder (fluent interface)
func (c *Client) List(ctx context.Context, builder *TransferListRequestBuilder) (*types.Response[TransferListResponse], error) {
	req := builder.Build()
	path := transferBasePath

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[TransferListResponse](
		ctx,
		c.client,
		c.secret,
		path,
		c.baseURL,
	)
}
