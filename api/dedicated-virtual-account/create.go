package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// Create creates a dedicated virtual account for an existing customer
func (c *Client) Create(ctx context.Context, req *CreateDedicatedVirtualAccountRequest) (*DedicatedVirtualAccount, error) {
	if err := validateCreateRequest(req); err != nil {
		return nil, err
	}

	resp, err := net.Post[CreateDedicatedVirtualAccountRequest, DedicatedVirtualAccount](
		ctx, c.client, c.secret, dedicatedVirtualAccountBasePath, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
