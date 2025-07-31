package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// DecommissionDevice unlinks a debug device from your integration
func (c *Client) DecommissionDevice(ctx context.Context, req *TerminalDecommissionRequest) (*TerminalDecommissionResponse, error) {
	if err := validateDecommissionRequest(req); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/decommission_device", terminalBasePath)
	resp, err := net.Post[TerminalDecommissionRequest, TerminalDecommissionResponse](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
