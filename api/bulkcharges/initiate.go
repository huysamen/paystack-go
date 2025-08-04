package bulkcharges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type BulkChargeItem struct {
	Authorization string `json:"authorization"`
	Amount        int64  `json:"amount"`
	Reference     string `json:"reference"`
}

type InitiateRequest []BulkChargeItem

type InitiateRequestBuilder struct {
	req *InitiateRequest
}

func NewInitiateRequest() *InitiateRequestBuilder {
	return &InitiateRequestBuilder{
		req: &InitiateRequest{},
	}
}

func (b *InitiateRequestBuilder) AddItem(authorization string, amount int64, reference string) *InitiateRequestBuilder {
	*b.req = append(*b.req, BulkChargeItem{
		Authorization: authorization,
		Amount:        amount,
		Reference:     reference,
	})
	return b
}

func (b *InitiateRequestBuilder) AddItems(items []BulkChargeItem) *InitiateRequestBuilder {
	*b.req = append(*b.req, items...)

	return b
}

func (b *InitiateRequestBuilder) Build() *InitiateRequest {
	return b.req
}

type InitiateResponseData = types.BulkChargeBatch
type InitiateResponse = types.Response[InitiateResponseData]

func (c *Client) Initiate(ctx context.Context, builder InitiateRequestBuilder) (*InitiateResponse, error) {
	return net.Post[InitiateRequest, InitiateResponseData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
