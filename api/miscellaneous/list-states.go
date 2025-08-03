package miscellaneous

import (
	"context"
	"net/url"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// State represents a state for address verification
type State struct {
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Abbreviation string `json:"abbreviation"`
}

// StateListRequest represents the request to list states
type StateListRequest struct {
	Country string `json:"country"` // Required: country code
}

// StateListRequestBuilder provides a fluent interface for building StateListRequest
type StateListRequestBuilder struct {
	req *StateListRequest
}

// NewStateListRequest creates a new builder for StateListRequest
func NewStateListRequest(country string) *StateListRequestBuilder {
	return &StateListRequestBuilder{
		req: &StateListRequest{
			Country: country,
		},
	}
}

// Build returns the constructed StateListRequest
func (b *StateListRequestBuilder) Build() *StateListRequest {
	return b.req
}

// StateListResponse represents the response from listing states
type StateListResponse = types.Response[[]State]

// ListStates retrieves a list of states for a country (for address verification)
func (c *Client) ListStates(ctx context.Context, builder *StateListRequestBuilder) (*StateListResponse, error) {
	req := builder.Build()
	params := url.Values{}
	params.Set("country", req.Country)

	endpoint := statesPath + "?" + params.Encode()
	return net.Get[[]State](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
