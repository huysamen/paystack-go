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

type GetUploadURLBuilder struct {
	request *GetUploadURLRequest
}

func NewGetUploadURLBuilder(uploadFileName string) *GetUploadURLBuilder {
	return &GetUploadURLBuilder{
		request: &GetUploadURLRequest{
			UploadFileName: uploadFileName,
		},
	}
}

func (b *GetUploadURLBuilder) Build() *GetUploadURLRequest {
	return b.request
}

type GetUploadURLResponse = types.Response[UploadURLData]

func (c *Client) GetUploadURL(ctx context.Context, disputeID string, builder *GetUploadURLBuilder) (*GetUploadURLResponse, error) {
	req := builder.Build()

	params := url.Values{}
	params.Set("upload_filename", req.UploadFileName)
	endpoint := basePath + "/" + disputeID + "/upload_url?" + params.Encode()

	return net.Get[UploadURLData](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
