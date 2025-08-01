package subaccounts

import (
	"context"
	"fmt"

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

// CreateWithBuilder creates a new subaccount using the builder pattern
func (c *Client) CreateWithBuilder(ctx context.Context, builder *SubaccountCreateRequestBuilder) (*SubaccountCreateResponse, error) {
	if builder == nil {
		return nil, fmt.Errorf("builder cannot be nil")
	}
	return c.Create(ctx, builder.Build())
}
