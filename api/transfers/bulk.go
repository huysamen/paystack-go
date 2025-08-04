package transfers

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Request type
type BulkTransferRequest struct {
	Source    string             `json:"source"`    // Only "balance" supported for now
	Currency  *string            `json:"currency"`  // Optional, defaults to NGN
	Transfers []BulkTransferItem `json:"transfers"` // List of transfers
}

// BulkTransferItem represents a single transfer in a bulk transfer request
type BulkTransferItem struct {
	Amount    int    `json:"amount"`
	Reference string `json:"reference"`
	Reason    string `json:"reason"`
	Recipient string `json:"recipient"`
}

// Builder for creating BulkTransferRequest
type BulkTransferRequestBuilder struct {
	req *BulkTransferRequest
}

// NewBulkTransferRequest creates a new builder for bulk transfer
func NewBulkTransferRequest(source string) *BulkTransferRequestBuilder {
	return &BulkTransferRequestBuilder{
		req: &BulkTransferRequest{
			Source:    source,
			Transfers: make([]BulkTransferItem, 0),
		},
	}
}

// Currency sets the currency
func (b *BulkTransferRequestBuilder) Currency(currency string) *BulkTransferRequestBuilder {
	b.req.Currency = &currency

	return b
}

// AddTransfer adds a transfer to the bulk request
func (b *BulkTransferRequestBuilder) AddTransfer(item BulkTransferItem) *BulkTransferRequestBuilder {
	b.req.Transfers = append(b.req.Transfers, item)

	return b
}

// Transfers sets all transfers at once
func (b *BulkTransferRequestBuilder) Transfers(transfers []BulkTransferItem) *BulkTransferRequestBuilder {
	b.req.Transfers = transfers

	return b
}

// Build creates the BulkTransferRequest
func (b *BulkTransferRequestBuilder) Build() *BulkTransferRequest {
	return b.req
}

// BulkTransferResponseData represents a single transfer result in a bulk transfer response
type BulkTransferResponseData struct {
	Reference    string         `json:"reference"`
	Recipient    string         `json:"recipient"`
	Amount       int            `json:"amount"`
	TransferCode string         `json:"transfer_code"`
	Currency     types.Currency `json:"currency"`
	Status       string         `json:"status"`
}

// BulkTransferResponse represents the response for bulk transfer
type BulkTransferResponse = types.Response[[]BulkTransferResponseData]

// Bulk creates a bulk transfer with the provided builder
func (c *Client) Bulk(ctx context.Context, builder *BulkTransferRequestBuilder) (*BulkTransferResponse, error) {
	return net.Post[BulkTransferRequest, []BulkTransferResponseData](ctx, c.Client, c.Secret, basePath+"/bulk", builder.Build(), c.BaseURL)
}
