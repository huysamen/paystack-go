package dedicatedvirtualaccount

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// FetchDedicatedVirtualAccountResponse represents the response from fetching a dedicated virtual account
type FetchDedicatedVirtualAccountResponse struct {
	Status  bool                     `json:"status"`
	Message string                   `json:"message"`
	Data    *DedicatedVirtualAccount `json:"data"`
}

// Fetch gets details of a dedicated virtual account on your integration
func (c *Client) Fetch(ctx context.Context, dedicatedAccountID string) (*DedicatedVirtualAccount, error) {
	endpoint := fmt.Sprintf("%s/%s", dedicatedVirtualAccountBasePath, dedicatedAccountID)
	resp, err := net.Get[FetchDedicatedVirtualAccountResponse](
		ctx, c.client, c.secret, endpoint, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return resp.Data.Data, nil
}
