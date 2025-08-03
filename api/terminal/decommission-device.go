package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// DecommissionDevice unlinks a debug device from your integration
func (c *Client) DecommissionDevice(ctx context.Context, req *TerminalDecommissionRequest) (*TerminalDecommissionResponse, error) {
	endpoint := fmt.Sprintf("%s/decommission_device", basePath)
	return net.Post[TerminalDecommissionRequest, any](
		ctx, c.Client, c.Secret, endpoint, req, c.BaseURL,
	)
}
