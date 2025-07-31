package plans

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
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

func (c *Client) List(ctx context.Context, req *PlanListRequest) (*types.Response[PlanListResponse], error) {
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
