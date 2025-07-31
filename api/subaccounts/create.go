package subaccounts

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// Create creates a new subaccount
func (c *Client) Create(ctx context.Context, req *SubaccountCreateRequest) (*SubaccountCreateResponse, error) {
	if err := validateCreateRequest(req); err != nil {
		return nil, err
	}

	resp, err := net.Post[SubaccountCreateRequest, SubaccountCreateResponse](
		ctx, c.client, c.secret, subaccountBasePath, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
