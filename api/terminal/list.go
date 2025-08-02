package terminal

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// List retrieves a list of terminals
func (c *Client) List(ctx context.Context, builder *TerminalListRequestBuilder) (*types.Response[[]Terminal], error) {
	params := url.Values{}

	if builder != nil {
		req := builder.Build()
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

	return net.Get[[]Terminal](ctx, c.client, c.secret, endpoint, c.baseURL)
}
