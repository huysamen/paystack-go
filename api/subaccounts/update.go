package subaccounts

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Update updates an existing subaccount
func (c *Client) Update(ctx context.Context, idOrCode string, req *SubaccountUpdateRequest) (*SubaccountUpdateResponse, error) {
	if idOrCode == "" {
		return nil, fmt.Errorf("id_or_code is required")
	}

	if err := validateUpdateRequest(req); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s", subaccountBasePath, idOrCode)

	resp, err := net.Put[SubaccountUpdateRequest, SubaccountUpdateResponse](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// UpdateWithBuilder updates an existing subaccount using the builder pattern
func (c *Client) UpdateWithBuilder(ctx context.Context, idOrCode string, builder *SubaccountUpdateRequestBuilder) (*SubaccountUpdateResponse, error) {
	if builder == nil {
		return nil, fmt.Errorf("builder cannot be nil")
	}
	return c.Update(ctx, idOrCode, builder.Build())
}
