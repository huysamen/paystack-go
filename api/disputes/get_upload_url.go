package disputes

import (
	"context"
	"net/url"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
)

type getUploadURLRequest struct {
	UploadFileName string `json:"upload_filename"`
}

type GetUploadURLRequestBuilder struct {
	request *getUploadURLRequest
}

func NewGetUploadURLRequestBuilder(uploadFileName string) *GetUploadURLRequestBuilder {
	return &GetUploadURLRequestBuilder{
		request: &getUploadURLRequest{
			UploadFileName: uploadFileName,
		},
	}
}

func (b *GetUploadURLRequestBuilder) Build() *getUploadURLRequest {
	return b.request
}

type GetUploadURLResponseData struct {
	SignedURL data.String `json:"signedUrl"`
	FileName  data.String `json:"fileName"`
	ExpiresIn data.Int    `json:"expiresIn"`
}

type GetUploadURLResponse = types.Response[GetUploadURLResponseData]

func (c *Client) GetUploadURL(ctx context.Context, disputeID string, builder *GetUploadURLRequestBuilder) (*GetUploadURLResponse, error) {
	req := builder.Build()

	params := url.Values{}
	params.Set("upload_filename", req.UploadFileName)
	endpoint := basePath + "/" + disputeID + "/upload_url?" + params.Encode()

	return net.Get[GetUploadURLResponseData](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
