package disputes

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Update updates the details of a dispute on your integration
func (c *Client) Update(ctx context.Context, disputeID string, req *DisputeUpdateRequest) (*DisputeUpdateResponse, error) {
	if disputeID == "" {
		return nil, fmt.Errorf("dispute ID is required")
	}

	if req == nil {
		return nil, fmt.Errorf("update dispute request cannot be nil")
	}

	if err := validateDisputeUpdateRequest(req); err != nil {
		return nil, err
	}

	url := c.baseURL + disputesBasePath + "/" + disputeID
	return net.Put[DisputeUpdateRequest, []Dispute](ctx, c.client, c.secret, url, req)
}

// validateDisputeUpdateRequest validates the dispute update request
func validateDisputeUpdateRequest(req *DisputeUpdateRequest) error {
	if req.RefundAmount != nil && *req.RefundAmount < 0 {
		return fmt.Errorf("refund amount must be greater than or equal to 0")
	}

	return nil
}
