package dedicatedvirtualaccount

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// DeactivateDedicatedAccountResponse represents the response from deactivating a dedicated account
type DeactivateDedicatedAccountResponse struct {
	Status  bool                     `json:"status"`
	Message string                   `json:"message"`
	Data    *DedicatedVirtualAccount `json:"data"`
}

// Deactivate deactivates a dedicated virtual account on your integration
func (c *Client) Deactivate(ctx context.Context, dedicatedAccountID string) (*DedicatedVirtualAccount, error) {
	endpoint := fmt.Sprintf("%s/%s", dedicatedVirtualAccountBasePath, dedicatedAccountID)
	resp, err := net.Delete[DeactivateDedicatedAccountResponse](
		ctx, c.client, c.secret, endpoint, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return resp.Data.Data, nil
}
