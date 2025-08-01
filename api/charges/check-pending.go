package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// CheckPendingChargeResponse represents the response from checking a pending charge
type CheckPendingChargeResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    ChargeData `json:"data"`
}

// CheckPending checks the status of a pending charge
func (c *Client) CheckPending(ctx context.Context, reference string) (*CheckPendingChargeResponse, error) {
	if reference == "" {
		return nil, ErrBuilderRequired // Reusing the error for consistency
	}

	url := c.baseURL + chargesBasePath + "/" + reference
	resp, err := net.Get[CheckPendingChargeResponse](ctx, c.client, c.secret, url)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
