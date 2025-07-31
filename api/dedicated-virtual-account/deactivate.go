package dedicatedvirtualaccount

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Deactivate deactivates a dedicated virtual account on your integration
func (c *Client) Deactivate(ctx context.Context, dedicatedAccountID string) (*DedicatedVirtualAccount, error) {
	if err := validateDedicatedAccountID(dedicatedAccountID); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s", dedicatedVirtualAccountBasePath, dedicatedAccountID)
	resp, err := net.Delete[DedicatedVirtualAccount](
		ctx, c.client, c.secret, endpoint, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
