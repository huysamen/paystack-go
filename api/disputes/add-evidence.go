package disputes

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/utils"
)

// AddEvidence provides evidence for a dispute
func (c *Client) AddEvidence(ctx context.Context, disputeID string, req *DisputeEvidenceRequest) (*DisputeEvidenceResponse, error) {
	if disputeID == "" {
		return nil, fmt.Errorf("dispute ID is required")
	}

	if req == nil {
		return nil, fmt.Errorf("add evidence request cannot be nil")
	}

	if err := validateDisputeEvidenceRequest(req); err != nil {
		return nil, err
	}

	url := c.baseURL + disputesBasePath + "/" + disputeID + "/evidence"
	return net.Post[DisputeEvidenceRequest, Evidence](ctx, c.client, c.secret, url, req)
}

// validateDisputeEvidenceRequest validates the dispute evidence request
func validateDisputeEvidenceRequest(req *DisputeEvidenceRequest) error {
	if req.CustomerEmail == "" {
		return fmt.Errorf("customer email is required")
	}

	if err := utils.ValidateEmail(req.CustomerEmail); err != nil {
		return fmt.Errorf("invalid customer email: %w", err)
	}

	if req.CustomerName == "" {
		return fmt.Errorf("customer name is required")
	}

	if req.CustomerPhone == "" {
		return fmt.Errorf("customer phone is required")
	}

	if req.ServiceDetails == "" {
		return fmt.Errorf("service details are required")
	}

	return nil
}
