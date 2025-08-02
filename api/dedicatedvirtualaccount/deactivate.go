package dedicatedvirtualaccount

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// DeactivateDedicatedAccountResponse represents the response from deactivating a dedicated account
type DeactivateDedicatedAccountResponse struct {
	Status  bool                     `json:"status"`
	Message string                   `json:"message"`
	Data    *DedicatedVirtualAccount `json:"data"`
}

// Deactivate deactivates a dedicated virtual account on your integration
func (c *Client) Deactivate(ctx context.Context, dedicatedAccountID string) (*types.Response[DeactivateDedicatedAccountResponse], error) {
	endpoint := fmt.Sprintf("%s/%s", basePath, dedicatedAccountID)

	return net.Delete[DeactivateDedicatedAccountResponse](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
