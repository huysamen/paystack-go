package terminal

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TerminalListRequest struct {
	PerPage  *int    `json:"perPage,omitempty"`  // Number of terminals per page (default: 50)
	Next     *string `json:"next,omitempty"`     // Cursor for next page (optional)
	Previous *string `json:"previous,omitempty"` // Cursor for previous page (optional)
}

type TerminalListRequestBuilder struct {
	req *TerminalListRequest
}

func NewTerminalListRequest() *TerminalListRequestBuilder {
	return &TerminalListRequestBuilder{
		req: &TerminalListRequest{},
	}
}

func (b *TerminalListRequestBuilder) PerPage(perPage int) *TerminalListRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *TerminalListRequestBuilder) Next(next string) *TerminalListRequestBuilder {
	b.req.Next = &next

	return b
}

func (b *TerminalListRequestBuilder) Previous(previous string) *TerminalListRequestBuilder {
	b.req.Previous = &previous

	return b
}

func (b *TerminalListRequestBuilder) Build() *TerminalListRequest {
	return b.req
}

type TerminalListResponse = types.Response[[]types.Terminal]

func (c *Client) List(ctx context.Context, builder *TerminalListRequestBuilder) (*TerminalListResponse, error) {
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

	endpoint := basePath
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	return net.Get[[]types.Terminal](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
