package terminal

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// DecommissionDevice unlinks a debug device from your integration
func (c *Client) DecommissionDevice(ctx context.Context, req *TerminalDecommissionRequest) (*types.Response[any], error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	endpoint := fmt.Sprintf("%s/decommission_device", terminalBasePath)
	return net.Post[TerminalDecommissionRequest, any](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
}
