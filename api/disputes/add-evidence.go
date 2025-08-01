package disputes

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// AddEvidence provides evidence for a dispute
func (c *Client) AddEvidence(ctx context.Context, disputeID string, req *DisputeEvidenceRequest) (*DisputeEvidenceResponse, error) {
	if disputeID == "" {
		return nil, fmt.Errorf("dispute ID is required")
	}

	if req == nil {
		return nil, fmt.Errorf("add evidence request cannot be nil")
	}

	url := c.baseURL + disputesBasePath + "/" + disputeID + "/evidence"
	return net.Post[DisputeEvidenceRequest, Evidence](ctx, c.client, c.secret, url, req)
}
