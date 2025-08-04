package disputes

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
) // ResolveDisputeRequest represents the request to resolve a dispute
type ResolveDisputeRequest struct {
	Resolution       types.DisputeResolution `json:"resolution"`
	Message          string                  `json:"message"`
	RefundAmount     int                     `json:"refund_amount"`
	UploadedFileName string                  `json:"uploaded_filename"`
	Evidence         *int                    `json:"evidence,omitempty"`
}

type ResolveDisputeBuilder struct {
	req *ResolveDisputeRequest
}

func NewResolveDisputeBuilder(resolution types.DisputeResolution, message string, refundAmount int, uploadedFileName string) *ResolveDisputeBuilder {
	return &ResolveDisputeBuilder{
		req: &ResolveDisputeRequest{
			Resolution:       resolution,
			Message:          message,
			RefundAmount:     refundAmount,
			UploadedFileName: uploadedFileName,
		},
	}
}

func (b *ResolveDisputeBuilder) Evidence(evidence int) *ResolveDisputeBuilder {
	b.req.Evidence = &evidence

	return b
}

func (b *ResolveDisputeBuilder) Build() *ResolveDisputeRequest {
	return b.req
}

type ResolveDisputeResponse = types.Response[types.Dispute]

func (c *Client) Resolve(ctx context.Context, disputeID string, builder *ResolveDisputeBuilder) (*ResolveDisputeResponse, error) {
	return net.Put[ResolveDisputeRequest, types.Dispute](ctx, c.Client, c.Secret, basePath+"/"+disputeID+"/resolve", builder.Build(), c.BaseURL)
}
