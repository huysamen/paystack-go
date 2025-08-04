package miscellaneous

import (
	"context"
	"net/url"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type StateListRequest struct {
	Country string `json:"country"` // Required: country code
}

type StateListRequestBuilder struct {
	req *StateListRequest
}

func NewStateListRequest(country string) *StateListRequestBuilder {
	return &StateListRequestBuilder{
		req: &StateListRequest{
			Country: country,
		},
	}
}

func (b *StateListRequestBuilder) Build() *StateListRequest {
	return b.req
}

type StateListResponse = types.Response[[]types.State]

func (c *Client) ListStates(ctx context.Context, builder *StateListRequestBuilder) (*StateListResponse, error) {
	req := builder.Build()
	params := url.Values{}
	params.Set("country", req.Country)

	endpoint := statesPath + "?" + params.Encode()
	return net.Get[[]types.State](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
