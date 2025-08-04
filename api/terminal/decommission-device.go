package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// TerminalDecommissionRequest represents the request to decommission a terminal
type TerminalDecommissionRequest struct {
	SerialNumber string `json:"serial_number"` // Device serial number
}

// TerminalDecommissionResponse represents the response from decommissioning a terminal
type TerminalDecommissionResponse = types.Response[any]

// DecommissionDevice unlinks a debug device from your integration
func (c *Client) DecommissionDevice(ctx context.Context, req *TerminalDecommissionRequest) (*TerminalDecommissionResponse, error) {
	endpoint := fmt.Sprintf("%s/decommission_device", basePath)
	return net.Post[TerminalDecommissionRequest, any](
		ctx, c.Client, c.Secret, endpoint, req, c.BaseURL,
	)
}
