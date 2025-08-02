package virtualterminal

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// ListVirtualTerminalsResponse represents the response from listing virtual terminals.
type ListVirtualTerminalsResponse types.Response[[]VirtualTerminal]

// List lists virtual terminals using the builder pattern
func (c *Client) List(ctx context.Context, builder *ListVirtualTerminalsRequestBuilder) (*ListVirtualTerminalsResponse, error) {
	params := url.Values{}

	if builder != nil {
		req := builder.Build()
		if req.Status != "" {
			params.Set("status", req.Status)
		}
		if req.PerPage > 0 {
			params.Set("perPage", strconv.Itoa(req.PerPage))
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
	}

	endpoint := virtualTerminalBasePath
	if params.Encode() != "" {
		endpoint += "?" + params.Encode()
	}

	resp, err := net.Get[[]VirtualTerminal](ctx, c.client, c.secret, endpoint, c.baseURL)
	if err != nil {
		return nil, err
	}

	response := ListVirtualTerminalsResponse(*resp)
	return &response, nil
}
