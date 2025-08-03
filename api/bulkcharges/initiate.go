package bulkcharges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// BulkChargeItem represents a single charge in a bulk charge request
type BulkChargeItem struct {
	Authorization string `json:"authorization"`
	Amount        int64  `json:"amount"`
	Reference     string `json:"reference"`
}

// InitiateBulkChargeRequest represents the request to initiate a bulk charge
type InitiateBulkChargeRequest []BulkChargeItem

// InitiateBulkChargeRequestBuilder provides a fluent interface for building InitiateBulkChargeRequest
type InitiateBulkChargeRequestBuilder struct {
	req *InitiateBulkChargeRequest
}

// NewInitiateBulkChargeRequest creates a new builder for InitiateBulkChargeRequest
func NewInitiateBulkChargeRequest() *InitiateBulkChargeRequestBuilder {
	return &InitiateBulkChargeRequestBuilder{
		req: &InitiateBulkChargeRequest{},
	}
}

// AddItem adds a bulk charge item to the request
func (b *InitiateBulkChargeRequestBuilder) AddItem(authorization string, amount int64, reference string) *InitiateBulkChargeRequestBuilder {
	*b.req = append(*b.req, BulkChargeItem{
		Authorization: authorization,
		Amount:        amount,
		Reference:     reference,
	})
	return b
}

// AddItems adds multiple bulk charge items to the request
func (b *InitiateBulkChargeRequestBuilder) AddItems(items []BulkChargeItem) *InitiateBulkChargeRequestBuilder {
	*b.req = append(*b.req, items...)
	return b
}

// Build returns the constructed InitiateBulkChargeRequest
func (b *InitiateBulkChargeRequestBuilder) Build() *InitiateBulkChargeRequest {
	return b.req
}

type InitiateBulkChargeResponse = types.Response[types.BulkChargeBatch]

// Initiate sends an array of objects with authorization codes and amounts for batch processing
func (c *Client) Initiate(ctx context.Context, builder *InitiateBulkChargeRequestBuilder) (*InitiateBulkChargeResponse, error) {
	return net.Post[InitiateBulkChargeRequest, types.BulkChargeBatch](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
