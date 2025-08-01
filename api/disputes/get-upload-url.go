package disputes

import (
"context"
"errors"
"net/url"

"github.com/huysamen/paystack-go/net"
"github.com/huysamen/paystack-go/types"
)

// GetUploadURLRequest represents the request to get upload URL for a dispute
type GetUploadURLRequest struct {
	UploadFileName string `json:"upload_filename"`
}

// GetUploadURLResponse represents the response from getting upload URL
type GetUploadURLResponse = types.Response[UploadURLData]

// GetUploadURLBuilder builds requests for getting upload URLs
type GetUploadURLBuilder struct {
	request *GetUploadURLRequest
}

// NewGetUploadURLBuilder creates a new builder for getting upload URLs
func NewGetUploadURLBuilder(uploadFileName string) *GetUploadURLBuilder {
	return &GetUploadURLBuilder{
		request: &GetUploadURLRequest{
			UploadFileName: uploadFileName,
		},
	}
}

// Build returns the built request
func (b *GetUploadURLBuilder) Build() *GetUploadURLRequest {
	return b.request
}

// GetUploadURL gets a signed URL for uploading dispute evidence files
func (c *Client) GetUploadURL(ctx context.Context, disputeID string, builder *GetUploadURLBuilder) (*types.Response[UploadURLData], error) {
	if disputeID == "" {
		return nil, errors.New("dispute ID is required")
	}

	if builder == nil {
		return nil, errors.New(ErrBuilderRequired)
	}

	endpoint := c.baseURL + disputesBasePath + "/" + disputeID + "/upload_url"
	req := builder.Build()

	// Build query parameters
	params := url.Values{}
	params.Set("upload_filename", req.UploadFileName)
	endpoint += "?" + params.Encode()

	resp, err := net.Get[UploadURLData](ctx, c.client, c.secret, endpoint, c.baseURL)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
