package customers

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CustomerListRequest struct {
	// Optional
	PerPage *int
	Page    *int
	From    *time.Time
	To      *time.Time
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
	Data []Customer `json:"data"`
	Meta types.Meta `json:"meta"`
}

func (c *Client) List(ctx context.Context, req *CustomerListRequest) (*types.Response[CustomerListResponse], error) {
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
