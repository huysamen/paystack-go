package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Assign creates a customer, validates the customer, and assigns a dedicated virtual account
func (c *Client) Assign(ctx context.Context, req *AssignDedicatedVirtualAccountRequest) (*types.Response[interface{}], error) {
	if err := validateAssignRequest(req); err != nil {
		return nil, err
	}

	endpoint := dedicatedVirtualAccountBasePath + "/assign"
	resp, err := net.Post[AssignDedicatedVirtualAccountRequest, interface{}](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
