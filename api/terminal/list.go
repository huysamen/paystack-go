package terminal

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
)

// List retrieves a list of terminals
func (c *Client) List(ctx context.Context, req *TerminalListRequest) (*TerminalListResponse, error) {
	params := url.Values{}

	if req != nil {
		if req.PerPage != nil {
			params.Set("perPage", strconv.Itoa(*req.PerPage))
		}
		if req.Next != nil {
			params.Set("next", *req.Next)
		}
		if req.Previous != nil {
			params.Set("previous", *req.Previous)
		}
	}

	endpoint := terminalBasePath
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	resp, err := net.Get[TerminalListResponse](ctx, c.client, c.secret, endpoint, c.baseURL)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// ListWithBuilder retrieves a list of terminals using the builder pattern
func (c *Client) ListWithBuilder(ctx context.Context, builder *TerminalListRequestBuilder) (*TerminalListResponse, error) {
	if builder == nil {
		return c.List(ctx, nil)
	}
	return c.List(ctx, builder.Build())
}
