package transactionsplits

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// TransactionSplitSubaccountAddRequest represents the request to add/update a subaccount in a split
type TransactionSplitSubaccountAddRequest struct {
	Subaccount string `json:"subaccount"` // Subaccount code
	Share      int    `json:"share"`      // Share amount (percentage or flat amount)
}

// TransactionSplitSubaccountAddRequestBuilder provides a fluent interface for building TransactionSplitSubaccountAddRequest
type TransactionSplitSubaccountAddRequestBuilder struct {
	subaccount string
	share      int
}

// NewTransactionSplitSubaccountAddRequest creates a new builder for adding a subaccount to a split
func NewTransactionSplitSubaccountAddRequest(subaccount string, share int) *TransactionSplitSubaccountAddRequestBuilder {
	return &TransactionSplitSubaccountAddRequestBuilder{
		subaccount: subaccount,
		share:      share,
	}
}

// Subaccount sets the subaccount code
func (b *TransactionSplitSubaccountAddRequestBuilder) Subaccount(subaccount string) *TransactionSplitSubaccountAddRequestBuilder {
	b.subaccount = subaccount
	return b
}

// Share sets the share amount
func (b *TransactionSplitSubaccountAddRequestBuilder) Share(share int) *TransactionSplitSubaccountAddRequestBuilder {
	b.share = share
	return b
}

// Build creates the TransactionSplitSubaccountAddRequest
func (b *TransactionSplitSubaccountAddRequestBuilder) Build() *TransactionSplitSubaccountAddRequest {
	return &TransactionSplitSubaccountAddRequest{
		Subaccount: b.subaccount,
		Share:      b.share,
	}
}

// AddSubaccountResponse represents the response for adding a subaccount
type AddSubaccountResponse = types.Response[types.TransactionSplit]

// AddSubaccount adds or updates a subaccount in a transaction split
func (c *Client) AddSubaccount(ctx context.Context, id string, builder *TransactionSplitSubaccountAddRequestBuilder) (*AddSubaccountResponse, error) {
	req := builder.Build()
	return net.Post[TransactionSplitSubaccountAddRequest, types.TransactionSplit](
		ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/subaccount/add", basePath, id), req, c.BaseURL,
	)
}
