package virtualterminal

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ListVirtualTerminalsRequest struct {
	Status   string `json:"status,omitempty"`
	PerPage  int    `json:"perPage,omitempty"`
	Search   string `json:"search,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
}

type ListVirtualTerminalsRequestBuilder struct {
	req *ListVirtualTerminalsRequest
}

func NewListVirtualTerminalsRequest() *ListVirtualTerminalsRequestBuilder {
	return &ListVirtualTerminalsRequestBuilder{
		req: &ListVirtualTerminalsRequest{},
	}
}

func (b *ListVirtualTerminalsRequestBuilder) Build() *ListVirtualTerminalsRequest {
	return b.req
}

type ListVirtualTerminalsResponse = types.Response[[]types.VirtualTerminal]

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

	endpoint := basePath
	if params.Encode() != "" {
		endpoint += "?" + params.Encode()
	}

	return net.Get[[]types.VirtualTerminal](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
