package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// TerminalCommissionRequest represents the request to commission a terminal
type TerminalCommissionRequest struct {
	SerialNumber string `json:"serial_number"` // Device serial number
}

// TerminalCommissionResponse represents the response from commissioning a terminal
type TerminalCommissionResponse = types.Response[types.Terminal]

// CommissionDevice activates a debug device by linking it to your integration
func (c *Client) CommissionDevice(ctx context.Context, req *TerminalCommissionRequest) (*TerminalCommissionResponse, error) {
	endpoint := fmt.Sprintf("%s/commission_device", basePath)
	return net.Post[TerminalCommissionRequest, types.Terminal](
		ctx, c.Client, c.Secret, endpoint, req, c.BaseURL,
	)
}
