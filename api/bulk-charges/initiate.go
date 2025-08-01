package bulkcharges

import (
	"context"
	"errors"

	"github.com/huysamen/paystack-go/net"
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

// InitiateBulkChargeResponse represents the response from initiating a bulk charge
type InitiateBulkChargeResponse struct {
	Status  bool            `json:"status"`
	Message string          `json:"message"`
	Data    BulkChargeBatch `json:"data"`
}

// Initiate sends an array of objects with authorization codes and amounts for batch processing
func (c *Client) Initiate(ctx context.Context, builder *InitiateBulkChargeRequestBuilder) (*InitiateBulkChargeResponse, error) {
	if builder == nil {
		return nil, errors.New("builder cannot be nil")
	}

	req := builder.Build()

	resp, err := net.Post[InitiateBulkChargeRequest, InitiateBulkChargeResponse](
		ctx, c.client, c.secret, bulkChargesBasePath, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
