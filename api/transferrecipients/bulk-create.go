package transferrecipients

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// BulkCreateTransferRecipientRequest represents the request to create multiple transfer recipients
type BulkCreateTransferRecipientRequest struct {
	Batch []types.BulkRecipientItem `json:"batch"` // Required: list of recipients
}

// BulkCreateTransferRecipientRequestBuilder builds a BulkCreateTransferRecipientRequest
type BulkCreateTransferRecipientRequestBuilder struct {
	request BulkCreateTransferRecipientRequest
}

// NewBulkCreateTransferRecipientRequestBuilder creates a new builder
func NewBulkCreateTransferRecipientRequestBuilder() *BulkCreateTransferRecipientRequestBuilder {
	return &BulkCreateTransferRecipientRequestBuilder{}
}

// Batch sets the batch of recipients
func (b *BulkCreateTransferRecipientRequestBuilder) Batch(batch []types.BulkRecipientItem) *BulkCreateTransferRecipientRequestBuilder {
	b.request.Batch = batch
	return b
}

// AddRecipient adds a single recipient to the batch
func (b *BulkCreateTransferRecipientRequestBuilder) AddRecipient(item types.BulkRecipientItem) *BulkCreateTransferRecipientRequestBuilder {
	if b.request.Batch == nil {
		b.request.Batch = make([]types.BulkRecipientItem, 0)
	}
	b.request.Batch = append(b.request.Batch, item)
	return b
}

// Build returns the built BulkCreateTransferRecipientRequest
func (b *BulkCreateTransferRecipientRequestBuilder) Build() *BulkCreateTransferRecipientRequest {
	return &b.request
}

// BulkCreateTransferRecipientResponse represents the response from bulk creating transfer recipients
type BulkCreateTransferRecipientResponse = types.Response[types.BulkCreateResult]

// BulkCreate creates multiple transfer recipients in a single request
func (c *Client) BulkCreate(ctx context.Context, builder *BulkCreateTransferRecipientRequestBuilder) (*BulkCreateTransferRecipientResponse, error) {
	req := builder.Build()
	return net.Post[BulkCreateTransferRecipientRequest, types.BulkCreateResult](ctx, c.Client, c.Secret, basePath+"/bulk", req, c.BaseURL)
}
