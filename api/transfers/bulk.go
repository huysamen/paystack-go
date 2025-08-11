package transfers

import (
	"context"

	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
)

type bulkRequest struct {
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

type BulkRequestBuilder struct {
	req *bulkRequest
}

func NewBulkRequestBuilder(source string) *BulkRequestBuilder {
	return &BulkRequestBuilder{
		req: &bulkRequest{
			Source:    source,
			Transfers: make([]BulkTransferItem, 0),
		},
	}
}

func (b *BulkRequestBuilder) Currency(currency string) *BulkRequestBuilder {
	b.req.Currency = &currency

	return b
}

func (b *BulkRequestBuilder) AddTransfer(item BulkTransferItem) *BulkRequestBuilder {
	b.req.Transfers = append(b.req.Transfers, item)

	return b
}

func (b *BulkRequestBuilder) Transfers(transfers []BulkTransferItem) *BulkRequestBuilder {
	b.req.Transfers = transfers

	return b
}

func (b *BulkRequestBuilder) Build() *bulkRequest {
	return b.req
}

type BulkResponseData struct {
	Reference    data.String    `json:"reference"`
	Recipient    data.String    `json:"recipient"` // In bulk responses recipient is code string
	Amount       data.Int       `json:"amount"`
	TransferCode data.String    `json:"transfer_code"`
	Currency     enums.Currency `json:"currency"`
	Status       data.String    `json:"status"`
}

type BulkResponse = types.Response[[]BulkResponseData]

func (c *Client) Bulk(ctx context.Context, builder BulkRequestBuilder) (*BulkResponse, error) {
	return net.Post[bulkRequest, []BulkResponseData](ctx, c.Client, c.Secret, basePath+"/bulk", builder.Build(), c.BaseURL)
}
