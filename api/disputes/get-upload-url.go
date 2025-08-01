package disputes

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// GetUploadURL gets a signed URL for uploading dispute evidence files
func (c *Client) GetUploadURL(ctx context.Context, disputeID string, req *DisputeUploadURLRequest) (*DisputeUploadURLResponse, error) {
	if disputeID == "" {
		return nil, fmt.Errorf("dispute ID is required")
	}

	if req == nil {
		return nil, fmt.Errorf("upload URL request cannot be nil")
	}

	if err := validateDisputeUploadURLRequest(req); err != nil {
		return nil, err
	}

	url := c.baseURL + disputesBasePath + "/" + disputeID + "/upload_url"
	return net.Get[UploadURLData](ctx, c.client, c.secret, url)
}

// validateDisputeUploadURLRequest validates the dispute upload URL request
func validateDisputeUploadURLRequest(req *DisputeUploadURLRequest) error {
	if req.UploadFileName == "" {
		return fmt.Errorf("upload filename is required")
	}

	return nil
}
