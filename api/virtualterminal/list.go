package virtualterminal

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// ListVirtualTerminalsRequest represents the request to list virtual terminals
type ListVirtualTerminalsRequest struct {
	Status   string `json:"status,omitempty"`
	PerPage  int    `json:"perPage,omitempty"`
	Search   string `json:"search,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
}

// ListVirtualTerminalsRequestBuilder provides a fluent interface for building ListVirtualTerminalsRequest
type ListVirtualTerminalsRequestBuilder struct {
	req *ListVirtualTerminalsRequest
}

// NewListVirtualTerminalsRequest creates a new builder for ListVirtualTerminalsRequest
func NewListVirtualTerminalsRequest() *ListVirtualTerminalsRequestBuilder {
	return &ListVirtualTerminalsRequestBuilder{
		req: &ListVirtualTerminalsRequest{},
	}
}

// Build returns the constructed ListVirtualTerminalsRequest
func (b *ListVirtualTerminalsRequestBuilder) Build() *ListVirtualTerminalsRequest {
	return b.req
}

// ListVirtualTerminalsResponse represents the response from listing virtual terminals
type ListVirtualTerminalsResponse = types.Response[[]types.VirtualTerminal]

// List lists virtual terminals using the builder pattern
func (c *Client) List(ctx context.Context, builder *ListVirtualTerminalsRequestBuilder) (*types.Response[[]types.VirtualTerminal], error) {
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

	endpoint := basePath
	if params.Encode() != "" {
		endpoint += "?" + params.Encode()
	}

	return net.Get[[]types.VirtualTerminal](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
