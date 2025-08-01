package plans

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// ListPlansRequest represents the request to list plans
type ListPlansRequest struct {
	// Optional
	PerPage  *int            `json:"perPage,omitempty"`
	Page     *int            `json:"page,omitempty"`
	Status   *string         `json:"status,omitempty"`
	Interval *types.Interval `json:"interval,omitempty"`
	Amount   *int            `json:"amount,omitempty"`
}

// ListPlansRequestBuilder provides a fluent interface for building ListPlansRequest
type ListPlansRequestBuilder struct {
	req *ListPlansRequest
}

// NewListPlansRequest creates a new builder for ListPlansRequest
func NewListPlansRequest() *ListPlansRequestBuilder {
	return &ListPlansRequestBuilder{
		req: &ListPlansRequest{},
	}
}

// PerPage sets the number of records per page
func (b *ListPlansRequestBuilder) PerPage(perPage int) *ListPlansRequestBuilder {
	b.req.PerPage = &perPage
	return b
}

// Page sets the page number
func (b *ListPlansRequestBuilder) Page(page int) *ListPlansRequestBuilder {
	b.req.Page = &page
	return b
}

// Status filters by plan status
func (b *ListPlansRequestBuilder) Status(status string) *ListPlansRequestBuilder {
	b.req.Status = &status
	return b
}

// Interval filters by plan interval
func (b *ListPlansRequestBuilder) Interval(interval types.Interval) *ListPlansRequestBuilder {
	b.req.Interval = &interval
	return b
}

// Amount filters by plan amount
func (b *ListPlansRequestBuilder) Amount(amount int) *ListPlansRequestBuilder {
	b.req.Amount = &amount
	return b
}

// Build returns the constructed ListPlansRequest
func (b *ListPlansRequestBuilder) Build() *ListPlansRequest {
	return b.req
}

// toQuery converts the request to URL query parameters
func (r *ListPlansRequest) toQuery() string {
	params := url.Values{}

	if r.PerPage != nil {
		params.Add("perPage", strconv.Itoa(*r.PerPage))
	}
	if r.Page != nil {
		params.Add("page", strconv.Itoa(*r.Page))
	}
	if r.Status != nil {
		params.Add("status", *r.Status)
	}
	if r.Interval != nil {
		params.Add("interval", r.Interval.String())
	}
	if r.Amount != nil {
		params.Add("amount", strconv.Itoa(*r.Amount))
	}

	return params.Encode()
}

// ListPlansResponse represents the response from listing plans
type ListPlansResponse struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Data    []types.Plan `json:"data"`
	Meta    *types.Meta  `json:"meta,omitempty"`
}

// List lists plans using a builder (fluent interface)
func (c *Client) List(ctx context.Context, builder *ListPlansRequestBuilder) (*types.Response[[]types.Plan], error) {
	var req *ListPlansRequest
	if builder != nil {
		req = builder.Build()
	}

	path := planBasePath

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[[]types.Plan](
		ctx,
		c.client,
		c.secret,
		path,
		c.baseURL,
	)
}
