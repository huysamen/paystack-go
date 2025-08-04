package disputes

import (
	"context"
	"net/url"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
) // GetUploadURLRequest represents the request to get upload URL for a dispute
type GetUploadURLRequest struct {
	UploadFileName string `json:"upload_filename"`
}

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

// GetUploadURLResponse represents the response from getting upload URL
type GetUploadURLResponse = types.Response[UploadURLData]

// GetUploadURL gets a signed URL for uploading dispute evidence files
func (c *Client) GetUploadURL(ctx context.Context, disputeID string, builder *GetUploadURLBuilder) (*GetUploadURLResponse, error) {
	req := builder.Build()

	// Build query parameters
	params := url.Values{}
	params.Set("upload_filename", req.UploadFileName)
	endpoint := basePath + "/" + disputeID + "/upload_url?" + params.Encode()

	return net.Get[UploadURLData](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
