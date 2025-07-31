package miscellaneous

import (
	"context"
	"fmt"
	"net/url"

	"github.com/huysamen/paystack-go/net"
)

// ListStates retrieves a list of states for a country (for address verification)
func (c *Client) ListStates(ctx context.Context, req *StateListRequest) (*StateListResponse, error) {
	if req == nil || req.Country == "" {
		return nil, fmt.Errorf("country is required")
	}

	params := url.Values{}
	params.Set("country", req.Country)

	endpoint := statesBasePath + "?" + params.Encode()

	resp, err := net.Get[StateListResponse](ctx, c.client, c.secret, endpoint, c.baseURL)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
