package transactionsplits

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionSplitSubaccountRemoveRequest struct {
	Subaccount string `json:"subaccount"` // Subaccount code
}

type TransactionSplitSubaccountRemoveRequestBuilder struct {
	subaccount string
}

func NewTransactionSplitSubaccountRemoveRequest(subaccount string) *TransactionSplitSubaccountRemoveRequestBuilder {
	return &TransactionSplitSubaccountRemoveRequestBuilder{
		subaccount: subaccount,
	}
}

func (b *TransactionSplitSubaccountRemoveRequestBuilder) Subaccount(subaccount string) *TransactionSplitSubaccountRemoveRequestBuilder {
	b.subaccount = subaccount
	return b
}

func (b *TransactionSplitSubaccountRemoveRequestBuilder) Build() *TransactionSplitSubaccountRemoveRequest {
	return &TransactionSplitSubaccountRemoveRequest{
		Subaccount: b.subaccount,
	}
}

type TransactionSplitSubaccountRemoveResponse = types.Response[any]

func (c *Client) RemoveSubaccount(ctx context.Context, id string, builder *TransactionSplitSubaccountRemoveRequestBuilder) (*TransactionSplitSubaccountRemoveResponse, error) {
	return net.Post[TransactionSplitSubaccountRemoveRequest, any](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/subaccount/remove", basePath, id), builder.Build(), c.BaseURL)
}
