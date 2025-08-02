package bulkcharges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

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

// Initiate sends an array of objects with authorization codes and amounts for batch processing
func (c *Client) Initiate(ctx context.Context, builder *InitiateBulkChargeRequestBuilder) (*types.Response[BulkChargeBatch], error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()

	return net.Post[InitiateBulkChargeRequest, BulkChargeBatch](
		ctx, c.client, c.secret, bulkChargesBasePath, req, c.baseURL,
	)
}
