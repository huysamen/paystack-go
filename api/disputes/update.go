package disputes

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type updateRequest struct {
	RefundAmount     *int    `json:"refund_amount,omitempty"`
	UploadedFileName *string `json:"uploaded_filename,omitempty"`
}

type UpdateRequestBuilder struct {
	request *updateRequest
}

func NewUpdateRequestBuilder() *UpdateRequestBuilder {
	return &UpdateRequestBuilder{
		request: &updateRequest{},
	}
}

func (b *UpdateRequestBuilder) RefundAmount(amount int) *UpdateRequestBuilder {
	b.request.RefundAmount = &amount

	return b
}

func (b *UpdateRequestBuilder) UploadedFileName(fileName string) *UpdateRequestBuilder {
	b.request.UploadedFileName = &fileName

	return b
}

func (b *UpdateRequestBuilder) Build() *updateRequest {
	return b.request
}

type UpdateResponseData = types.Dispute
type UpdateResponse = types.Response[UpdateResponseData]

func (c *Client) Update(ctx context.Context, disputeID string, builder *UpdateRequestBuilder) (*UpdateResponse, error) {
	return net.Put[updateRequest, UpdateResponseData](ctx, c.Client, c.Secret, basePath+"/"+disputeID, builder.Build(), c.BaseURL)
}
