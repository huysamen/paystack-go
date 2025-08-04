package disputes

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type resolveRequest struct {
	Resolution       types.DisputeResolution `json:"resolution"`
	Message          string                  `json:"message"`
	RefundAmount     int                     `json:"refund_amount"`
	UploadedFileName string                  `json:"uploaded_filename"`
	Evidence         *int                    `json:"evidence,omitempty"`
}

type ResolveRequestBuilder struct {
	req *resolveRequest
}

func NewResolveRequestBuilder(resolution types.DisputeResolution, message string, refundAmount int, uploadedFileName string) *ResolveRequestBuilder {
	return &ResolveRequestBuilder{
		req: &resolveRequest{
			Resolution:       resolution,
			Message:          message,
			RefundAmount:     refundAmount,
			UploadedFileName: uploadedFileName,
		},
	}
}

func (b *ResolveRequestBuilder) Evidence(evidence int) *ResolveRequestBuilder {
	b.req.Evidence = &evidence

	return b
}

func (b *ResolveRequestBuilder) Build() *resolveRequest {
	return b.req
}

type ResolveResponseData = types.Dispute
type ResolveResponse = types.Response[ResolveResponseData]

func (c *Client) Resolve(ctx context.Context, disputeID string, builder *ResolveRequestBuilder) (*ResolveResponse, error) {
	return net.Put[resolveRequest, ResolveResponseData](ctx, c.Client, c.Secret, basePath+"/"+disputeID+"/resolve", builder.Build(), c.BaseURL)
}
