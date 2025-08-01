package disputes

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// List retrieves disputes filed against your integration
func (c *Client) List(ctx context.Context, req *DisputeListRequest) (*DisputeListResponse, error) {
	if req == nil {
		req = &DisputeListRequest{}
	}

	if err := validateDisputeListRequest(req); err != nil {
		return nil, err
	}

	url := c.baseURL + disputesBasePath
	return net.Get[[]Dispute](ctx, c.client, c.secret, url)
}

// validateDisputeListRequest validates the dispute list request
func validateDisputeListRequest(req *DisputeListRequest) error {
	if req.PerPage != nil && (*req.PerPage < 1 || *req.PerPage > 100) {
		return fmt.Errorf("per_page must be between 1 and 100")
	}

	if req.Page != nil && *req.Page < 1 {
		return fmt.Errorf("page must be greater than 0")
	}

	if req.From != nil && req.To != nil && req.From.After(*req.To) {
		return fmt.Errorf("from date must be before to date")
	}

	if req.Status != nil {
		validStatuses := []string{
			string(DisputeStatusAwaitingMerchantFeedback),
			string(DisputeStatusAwaitingBankFeedback),
			string(DisputeStatusPending),
			string(DisputeStatusResolved),
			string(DisputeStatusArchived),
		}

		valid := false
		for _, status := range validStatuses {
			if string(*req.Status) == status {
				valid = true
				break
			}
		}

		if !valid {
			return fmt.Errorf("invalid dispute status: %s", *req.Status)
		}
	}

	return nil
}
