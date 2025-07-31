package transfers

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/huysamen/paystack-go/net"
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

func (c *Client) List(ctx context.Context, req *TransferListRequest) (*types.Response[TransferListResponse], error) {
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
