package transferrecipients

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type BulkCreateTransferRecipientRequest struct {
	Batch []types.BulkRecipientItem `json:"batch"` // Required: list of recipients
}

type BulkCreateTransferRecipientRequestBuilder struct {
	request BulkCreateTransferRecipientRequest
}

func NewBulkCreateTransferRecipientRequestBuilder() *BulkCreateTransferRecipientRequestBuilder {
	return &BulkCreateTransferRecipientRequestBuilder{}
}

func (b *BulkCreateTransferRecipientRequestBuilder) Batch(batch []types.BulkRecipientItem) *BulkCreateTransferRecipientRequestBuilder {
	b.request.Batch = batch
	return b
}

func (b *BulkCreateTransferRecipientRequestBuilder) AddRecipient(item types.BulkRecipientItem) *BulkCreateTransferRecipientRequestBuilder {
	if b.request.Batch == nil {
		b.request.Batch = make([]types.BulkRecipientItem, 0)
	}
	b.request.Batch = append(b.request.Batch, item)
	return b
}

func (b *BulkCreateTransferRecipientRequestBuilder) Build() *BulkCreateTransferRecipientRequest {
	return &b.request
}

type BulkCreateTransferRecipientResponse = types.Response[types.BulkCreateResult]

func (c *Client) BulkCreate(ctx context.Context, builder *BulkCreateTransferRecipientRequestBuilder) (*BulkCreateTransferRecipientResponse, error) {
	req := builder.Build()
	return net.Post[BulkCreateTransferRecipientRequest, types.BulkCreateResult](ctx, c.Client, c.Secret, basePath+"/bulk", req, c.BaseURL)
}
