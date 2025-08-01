package plans

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/options"
	"github.com/huysamen/paystack-go/types"
)

type PlanListRequest struct {
	// Optional
	PerPage  *int
	Page     *int
	Status   *string
	Interval *types.Interval
	Amount   *int
}

// PlanListRequestBuilder provides a fluent interface for building PlanListRequest
type PlanListRequestBuilder struct {
	req *PlanListRequest
}

// NewPlanListRequest creates a new builder for PlanListRequest
func NewPlanListRequest() *PlanListRequestBuilder {
	return &PlanListRequestBuilder{
		req: &PlanListRequest{},
	}
}

// PerPage sets the number of records per page
func (b *PlanListRequestBuilder) PerPage(perPage int) *PlanListRequestBuilder {
	b.req.PerPage = options.Int(perPage)
	return b
}

// Page sets the page number
func (b *PlanListRequestBuilder) Page(page int) *PlanListRequestBuilder {
	b.req.Page = options.Int(page)
	return b
}

// Status filters by plan status
func (b *PlanListRequestBuilder) Status(status string) *PlanListRequestBuilder {
	b.req.Status = options.String(status)
	return b
}

// Interval filters by plan interval
func (b *PlanListRequestBuilder) Interval(interval types.Interval) *PlanListRequestBuilder {
	b.req.Interval = &interval
	return b
}

// Amount filters by plan amount
func (b *PlanListRequestBuilder) Amount(amount int) *PlanListRequestBuilder {
	b.req.Amount = options.Int(amount)
	return b
}

// Build returns the constructed PlanListRequest
func (b *PlanListRequestBuilder) Build() *PlanListRequest {
	return b.req
}

func (r *PlanListRequest) toQuery() string {
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

type PlanListResponse []types.Plan

// List lists plans using a builder (fluent interface)
func (c *Client) List(ctx context.Context, builder *PlanListRequestBuilder) (*types.Response[PlanListResponse], error) {
	req := builder.Build()
	path := planBasePath

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[PlanListResponse](
		ctx,
		c.client,
		c.secret,
		path,
		c.baseURL,
	)
}
