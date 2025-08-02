package dedicatedvirtualaccount

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FetchDedicatedVirtualAccountResponse represents the response from fetching a dedicated virtual account
type FetchDedicatedVirtualAccountResponse struct {
	Status  bool                     `json:"status"`
	Message string                   `json:"message"`
	Data    *DedicatedVirtualAccount `json:"data"`
}

// Fetch gets details of a dedicated virtual account on your integration
func (c *Client) Fetch(ctx context.Context, dedicatedAccountID string) (*types.Response[FetchDedicatedVirtualAccountResponse], error) {
	endpoint := fmt.Sprintf("%s/%s", basePath, dedicatedAccountID)

	return net.Get[FetchDedicatedVirtualAccountResponse](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
