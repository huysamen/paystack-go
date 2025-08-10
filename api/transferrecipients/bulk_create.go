package transferrecipients

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
)

type bulkCreateRequest struct {
	Batch []types.BulkRecipientItem `json:"batch"` // Required: list of recipients
}

type BulkCreateRequestBuilder struct {
	req bulkCreateRequest
}

func NewBulkCreateRequestBuilder() *BulkCreateRequestBuilder {
	return &BulkCreateRequestBuilder{}
}

func (b *BulkCreateRequestBuilder) Batch(batch []types.BulkRecipientItem) *BulkCreateRequestBuilder {
	b.req.Batch = batch

	return b
}

func (b *BulkCreateRequestBuilder) AddRecipient(item types.BulkRecipientItem) *BulkCreateRequestBuilder {
	if b.req.Batch == nil {
		b.req.Batch = make([]types.BulkRecipientItem, 0)
	}

	b.req.Batch = append(b.req.Batch, item)

	return b
}

func (b *BulkCreateRequestBuilder) Build() *bulkCreateRequest {
	return &b.req
}

type BulkCreateResponseData struct {
	Success []types.Recipient `json:"success"`
	Errors  []struct {
		Error   data.String             `json:"error"`
		Payload types.BulkRecipientItem `json:"payload"`
	} `json:"errors"`
}

type BulkCreateResponse = types.Response[BulkCreateResponseData]

func (c *Client) BulkCreate(ctx context.Context, builder BulkCreateRequestBuilder) (*BulkCreateResponse, error) {
	return net.Post[bulkCreateRequest, BulkCreateResponseData](ctx, c.Client, c.Secret, basePath+"/bulk", builder.Build(), c.BaseURL)
}
