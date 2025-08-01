package disputes

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Resolve resolves a dispute on your integration
func (c *Client) Resolve(ctx context.Context, disputeID string, req *DisputeResolveRequest) (*DisputeResolveResponse, error) {
	if disputeID == "" {
		return nil, fmt.Errorf("dispute ID is required")
	}

	if req == nil {
		return nil, fmt.Errorf("resolve dispute request cannot be nil")
	}

	if err := validateDisputeResolveRequest(req); err != nil {
		return nil, err
	}

	url := c.baseURL + disputesBasePath + "/" + disputeID + "/resolve"
	return net.Put[DisputeResolveRequest, Dispute](ctx, c.client, c.secret, url, req)
}

// validateDisputeResolveRequest validates the dispute resolve request
func validateDisputeResolveRequest(req *DisputeResolveRequest) error {
	if req.Resolution == "" {
		return fmt.Errorf("resolution is required")
	}

	validResolutions := []string{
		string(DisputeResolutionMerchantAccepted),
		string(DisputeResolutionDeclined),
	}

	valid := false
	for _, resolution := range validResolutions {
		if string(req.Resolution) == resolution {
			valid = true
			break
		}
	}

	if !valid {
		return fmt.Errorf("invalid dispute resolution: %s", req.Resolution)
	}

	if req.Message == "" {
		return fmt.Errorf("message is required")
	}

	if req.RefundAmount < 0 {
		return fmt.Errorf("refund amount must be greater than or equal to 0")
	}

	if req.UploadedFileName == "" {
		return fmt.Errorf("uploaded filename is required")
	}

	return nil
}
