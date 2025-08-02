package terminal

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// CommissionDevice activates a debug device by linking it to your integration
func (c *Client) CommissionDevice(ctx context.Context, req *TerminalCommissionRequest) (*types.Response[Terminal], error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	endpoint := fmt.Sprintf("%s/commission_device", terminalBasePath)
	return net.Post[TerminalCommissionRequest, Terminal](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
}
