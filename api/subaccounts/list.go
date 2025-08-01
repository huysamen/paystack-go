package subaccounts

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
)

// SubaccountListRequestBuilder provides a fluent interface for building SubaccountListRequest
type SubaccountListRequestBuilder struct {
	req *SubaccountListRequest
}

// NewSubaccountListRequest creates a new builder for SubaccountListRequest
func NewSubaccountListRequest() *SubaccountListRequestBuilder {
	return &SubaccountListRequestBuilder{
		req: &SubaccountListRequest{},
	}
}

// PerPage sets the number of subaccounts per page
func (b *SubaccountListRequestBuilder) PerPage(perPage int) *SubaccountListRequestBuilder {
	b.req.PerPage = &perPage
	return b
}

// Page sets the page number
func (b *SubaccountListRequestBuilder) Page(page int) *SubaccountListRequestBuilder {
	b.req.Page = &page
	return b
}

// From sets the start date filter
func (b *SubaccountListRequestBuilder) From(from time.Time) *SubaccountListRequestBuilder {
	b.req.From = &from
	return b
}

// To sets the end date filter
func (b *SubaccountListRequestBuilder) To(to time.Time) *SubaccountListRequestBuilder {
	b.req.To = &to
	return b
}

// DateRange sets both from and to dates for convenience
func (b *SubaccountListRequestBuilder) DateRange(from, to time.Time) *SubaccountListRequestBuilder {
	b.req.From = &from
	b.req.To = &to
	return b
}

// Build returns the constructed SubaccountListRequest
func (b *SubaccountListRequestBuilder) Build() *SubaccountListRequest {
	return b.req
}

// List retrieves a list of subaccounts using the builder pattern
func (c *Client) List(ctx context.Context, builder *SubaccountListRequestBuilder) (*SubaccountListResponse, error) {
	var req *SubaccountListRequest
	if builder != nil {
		req = builder.Build()
	}

	params := url.Values{}
	if req != nil {
		if req.PerPage != nil {
			params.Set("perPage", strconv.Itoa(*req.PerPage))
		}
		if req.Page != nil {
			params.Set("page", strconv.Itoa(*req.Page))
		}
		if req.From != nil {
			params.Set("from", req.From.Format(time.RFC3339))
		}
		if req.To != nil {
			params.Set("to", req.To.Format(time.RFC3339))
		}
	}

	endpoint := subaccountBasePath
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	return net.Get[[]Subaccount](ctx, c.client, c.secret, endpoint, c.baseURL)
}
