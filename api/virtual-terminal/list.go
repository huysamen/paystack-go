package virtualterminal

import (
	"context"
	"fmt"
	"net/url"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// List retrieves a list of virtual terminals
func (c *Client) List(ctx context.Context, req *ListVirtualTerminalsRequest) (*types.Response[[]VirtualTerminal], error) {
	endpoint := virtualTerminalBasePath

	if req != nil {
		params := url.Values{}
		if req.Status != "" {
			params.Set("status", req.Status)
		}
		if req.PerPage > 0 {
			params.Set("perPage", fmt.Sprintf("%d", req.PerPage))
		}
		if req.Search != "" {
			params.Set("search", req.Search)
		}
		if req.Next != "" {
			params.Set("next", req.Next)
		}
		if req.Previous != "" {
			params.Set("previous", req.Previous)
		}

		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	resp, err := net.Get[[]VirtualTerminal](
		ctx, c.client, c.secret, endpoint, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
