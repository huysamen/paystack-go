package transfers

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type BulkTransferRequest struct {
	Source    string             `json:"source"`    // Only "balance" supported for now
	Currency  *string            `json:"currency"`  // Optional, defaults to NGN
	Transfers []BulkTransferItem `json:"transfers"` // List of transfers
}

type BulkTransferItem struct {
	Amount    int    `json:"amount"`
	Reference string `json:"reference"`
	Reason    string `json:"reason"`
	Recipient string `json:"recipient"`
}

type BulkTransferRequestBuilder struct {
	req *BulkTransferRequest
}

func NewBulkTransferRequest(source string) *BulkTransferRequestBuilder {
	return &BulkTransferRequestBuilder{
		req: &BulkTransferRequest{
			Source:    source,
			Transfers: make([]BulkTransferItem, 0),
		},
	}
}

func (b *BulkTransferRequestBuilder) Currency(currency string) *BulkTransferRequestBuilder {
	b.req.Currency = &currency

	return b
}

func (b *BulkTransferRequestBuilder) AddTransfer(item BulkTransferItem) *BulkTransferRequestBuilder {
	b.req.Transfers = append(b.req.Transfers, item)

	return b
}

func (b *BulkTransferRequestBuilder) Transfers(transfers []BulkTransferItem) *BulkTransferRequestBuilder {
	b.req.Transfers = transfers

	return b
}

func (b *BulkTransferRequestBuilder) Build() *BulkTransferRequest {
	return b.req
}

type BulkTransferResponseData struct {
	Reference    string         `json:"reference"`
	Recipient    string         `json:"recipient"`
	Amount       int            `json:"amount"`
	TransferCode string         `json:"transfer_code"`
	Currency     types.Currency `json:"currency"`
	Status       string         `json:"status"`
}

type BulkTransferResponse = types.Response[[]BulkTransferResponseData]

func (c *Client) Bulk(ctx context.Context, builder *BulkTransferRequestBuilder) (*BulkTransferResponse, error) {
	return net.Post[BulkTransferRequest, []BulkTransferResponseData](ctx, c.Client, c.Secret, basePath+"/bulk", builder.Build(), c.BaseURL)
}
