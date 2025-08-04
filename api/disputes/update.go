package disputes

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
) // UpdateDisputeRequest represents the request to update a dispute
type UpdateDisputeRequest struct {
	RefundAmount     *int    `json:"refund_amount,omitempty"`
	UploadedFileName *string `json:"uploaded_filename,omitempty"`
}

type UpdateDisputeBuilder struct {
	request *UpdateDisputeRequest
}

func NewUpdateDisputeBuilder() *UpdateDisputeBuilder {
	return &UpdateDisputeBuilder{
		request: &UpdateDisputeRequest{},
	}
}

func (b *UpdateDisputeBuilder) RefundAmount(amount int) *UpdateDisputeBuilder {
	b.request.RefundAmount = &amount

	return b
}

func (b *UpdateDisputeBuilder) UploadedFileName(fileName string) *UpdateDisputeBuilder {
	b.request.UploadedFileName = &fileName

	return b
}

func (b *UpdateDisputeBuilder) Build() *UpdateDisputeRequest {
	return b.request
}

type UpdateDisputeResponse = types.Response[types.Dispute]

func (c *Client) Update(ctx context.Context, disputeID string, builder *UpdateDisputeBuilder) (*UpdateDisputeResponse, error) {
	return net.Put[UpdateDisputeRequest, types.Dispute](ctx, c.Client, c.Secret, basePath+"/"+disputeID, builder.Build(), c.BaseURL)
}
