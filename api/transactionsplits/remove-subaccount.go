package transactionsplits

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// TransactionSplitSubaccountRemoveRequest represents the request to remove a subaccount from a split
type TransactionSplitSubaccountRemoveRequest struct {
	Subaccount string `json:"subaccount"` // Subaccount code
}

// TransactionSplitSubaccountRemoveRequestBuilder provides a fluent interface for building TransactionSplitSubaccountRemoveRequest
type TransactionSplitSubaccountRemoveRequestBuilder struct {
	subaccount string
}

// NewTransactionSplitSubaccountRemoveRequest creates a new builder for removing a subaccount from a split
func NewTransactionSplitSubaccountRemoveRequest(subaccount string) *TransactionSplitSubaccountRemoveRequestBuilder {
	return &TransactionSplitSubaccountRemoveRequestBuilder{
		subaccount: subaccount,
	}
}

// Subaccount sets the subaccount code
func (b *TransactionSplitSubaccountRemoveRequestBuilder) Subaccount(subaccount string) *TransactionSplitSubaccountRemoveRequestBuilder {
	b.subaccount = subaccount
	return b
}

// Build creates the TransactionSplitSubaccountRemoveRequest
func (b *TransactionSplitSubaccountRemoveRequestBuilder) Build() *TransactionSplitSubaccountRemoveRequest {
	return &TransactionSplitSubaccountRemoveRequest{
		Subaccount: b.subaccount,
	}
}

// TransactionSplitSubaccountRemoveResponse represents the response from removing a subaccount from a split
type TransactionSplitSubaccountRemoveResponse = types.Response[any]

// RemoveSubaccount removes a subaccount from a transaction split
func (c *Client) RemoveSubaccount(ctx context.Context, id string, builder *TransactionSplitSubaccountRemoveRequestBuilder) (*TransactionSplitSubaccountRemoveResponse, error) {
	return net.Post[TransactionSplitSubaccountRemoveRequest, any](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/subaccount/remove", basePath, id), builder.Build(), c.BaseURL)
}
