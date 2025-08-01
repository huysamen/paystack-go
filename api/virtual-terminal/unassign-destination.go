package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// UnassignDestination unassigns destinations from a virtual terminal
func (c *Client) UnassignDestination(ctx context.Context, code string, req *UnassignDestinationRequest) (*types.Response[any], error) {
	if err := validateCode(code); err != nil {
		return nil, err
	}
	if err := validateUnassignDestinationRequest(req); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s/destination/unassign", virtualTerminalBasePath, code)
	resp, err := net.Post[UnassignDestinationRequest, any](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
