package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// CommissionDevice activates a debug device by linking it to your integration
func (c *Client) CommissionDevice(ctx context.Context, req *TerminalCommissionRequest) (*TerminalCommissionResponse, error) {
	if err := validateCommissionRequest(req); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/commission_device", terminalBasePath)
	resp, err := net.Post[TerminalCommissionRequest, TerminalCommissionResponse](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
