package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TerminalCommissionRequest struct {
	SerialNumber string `json:"serial_number"` // Device serial number
}

type TerminalCommissionResponse = types.Response[types.Terminal]

func (c *Client) CommissionDevice(ctx context.Context, req *TerminalCommissionRequest) (*TerminalCommissionResponse, error) {
	endpoint := fmt.Sprintf("%s/commission_device", basePath)
	return net.Post[TerminalCommissionRequest, types.Terminal](
		ctx, c.Client, c.Secret, endpoint, req, c.BaseURL,
	)
}
