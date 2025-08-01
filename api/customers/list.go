package customers

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/options"
	"github.com/huysamen/paystack-go/types"
)

type CustomerListRequest struct {
	// Optional
	PerPage *int
	Page    *int
	From    *time.Time
	To      *time.Time
}

// CustomerListRequestBuilder provides a fluent interface for building CustomerListRequest
type CustomerListRequestBuilder struct {
	req *CustomerListRequest
}

// NewCustomerListRequest creates a new builder for CustomerListRequest
func NewCustomerListRequest() *CustomerListRequestBuilder {
	return &CustomerListRequestBuilder{
		req: &CustomerListRequest{},
	}
}

// PerPage sets the number of records per page
func (b *CustomerListRequestBuilder) PerPage(perPage int) *CustomerListRequestBuilder {
	b.req.PerPage = options.Int(perPage)
	return b
}

// Page sets the page number
func (b *CustomerListRequestBuilder) Page(page int) *CustomerListRequestBuilder {
	b.req.Page = options.Int(page)
	return b
}

// DateRange sets both start and end date filters
func (b *CustomerListRequestBuilder) DateRange(from, to time.Time) *CustomerListRequestBuilder {
	b.req.From = options.Time(from)
	b.req.To = options.Time(to)
	return b
}

// From sets the start date filter
func (b *CustomerListRequestBuilder) From(from time.Time) *CustomerListRequestBuilder {
	b.req.From = options.Time(from)
	return b
}

// To sets the end date filter
func (b *CustomerListRequestBuilder) To(to time.Time) *CustomerListRequestBuilder {
	b.req.To = options.Time(to)
	return b
}

// Build returns the constructed CustomerListRequest
func (b *CustomerListRequestBuilder) Build() *CustomerListRequest {
	return b.req
}

func (r *CustomerListRequest) toQuery() string {
	params := url.Values{}

	if r.PerPage != nil {
		params.Add("perPage", fmt.Sprintf("%d", *r.PerPage))
	}
	if r.Page != nil {
		params.Add("page", fmt.Sprintf("%d", *r.Page))
	}
	if r.From != nil {
		params.Add("from", r.From.Format("2006-01-02T15:04:05.999Z"))
	}
	if r.To != nil {
		params.Add("to", r.To.Format("2006-01-02T15:04:05.999Z"))
	}

	return params.Encode()
}

type CustomerListResponse struct {
	Data []types.Customer `json:"data"`
	Meta types.Meta       `json:"meta"`
}

// List lists customers using a builder (fluent interface)
func (c *Client) List(ctx context.Context, builder *CustomerListRequestBuilder) (*types.Response[CustomerListResponse], error) {
	req := builder.Build()
	path := customerBasePath

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[CustomerListResponse](
		ctx,
		c.client,
		c.secret,
		path,
		c.baseURL,
	)
}
