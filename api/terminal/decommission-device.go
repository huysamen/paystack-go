package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TerminalDecommissionRequest struct {
	SerialNumber string `json:"serial_number"` // Device serial number
}

type TerminalDecommissionResponse = types.Response[any]

func (c *Client) DecommissionDevice(ctx context.Context, req *TerminalDecommissionRequest) (*TerminalDecommissionResponse, error) {
	endpoint := fmt.Sprintf("%s/decommission_device", basePath)
	return net.Post[TerminalDecommissionRequest, any](
		ctx, c.Client, c.Secret, endpoint, req, c.BaseURL,
	)
}
