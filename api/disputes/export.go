package disputes

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Export exports disputes available on your integration
func (c *Client) Export(ctx context.Context, req *DisputeExportRequest) (*DisputeExportResponse, error) {
	if req == nil {
		req = &DisputeExportRequest{}
	}

	if err := validateDisputeExportRequest(req); err != nil {
		return nil, err
	}

	url := c.baseURL + disputesBasePath + "/export"
	return net.Get[ExportData](ctx, c.client, c.secret, url)
}

// validateDisputeExportRequest validates the dispute export request
func validateDisputeExportRequest(req *DisputeExportRequest) error {
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
