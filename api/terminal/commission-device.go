package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// CommissionDevice activates a debug device by linking it to your integration
func (c *Client) CommissionDevice(ctx context.Context, req *TerminalCommissionRequest) (*TerminalCommissionResponse, error) {
	endpoint := fmt.Sprintf("%s/commission_device", basePath)
	return net.Post[TerminalCommissionRequest, Terminal](
		ctx, c.Client, c.Secret, endpoint, req, c.BaseURL,
	)
}
