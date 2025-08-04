package miscellaneous

import (
	"context"
	"net/url"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type listStatesRequest struct {
	Country string `json:"country"` // Required: country code
}

type ListStatesRequestBuilder struct {
	req *listStatesRequest
}

func NewListStatesRequestBuilder(country string) *ListStatesRequestBuilder {
	return &ListStatesRequestBuilder{
		req: &listStatesRequest{
			Country: country,
		},
	}
}

func (b *ListStatesRequestBuilder) Build() *listStatesRequest {
	return b.req
}

func (r *listStatesRequest) toQuery() string {
	params := url.Values{}
	params.Set("country", r.Country)

	return params.Encode()
}

type ListStatesResponseData = []types.State
type ListStatesResponse = types.Response[ListStatesResponseData]

func (c *Client) ListStates(ctx context.Context, builder *ListStatesRequestBuilder) (*ListStatesResponse, error) {
	req := builder.Build()
	path := statesPath

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[ListStatesResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
