package transactionsplits

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionSplitSubaccountAddRequest struct {
	Subaccount string `json:"subaccount"` // Subaccount code
	Share      int    `json:"share"`      // Share amount (percentage or flat amount)
}

type TransactionSplitSubaccountAddRequestBuilder struct {
	subaccount string
	share      int
}

func NewTransactionSplitSubaccountAddRequest(subaccount string, share int) *TransactionSplitSubaccountAddRequestBuilder {
	return &TransactionSplitSubaccountAddRequestBuilder{
		subaccount: subaccount,
		share:      share,
	}
}

func (b *TransactionSplitSubaccountAddRequestBuilder) Subaccount(subaccount string) *TransactionSplitSubaccountAddRequestBuilder {
	b.subaccount = subaccount
	return b
}

func (b *TransactionSplitSubaccountAddRequestBuilder) Share(share int) *TransactionSplitSubaccountAddRequestBuilder {
	b.share = share
	return b
}

func (b *TransactionSplitSubaccountAddRequestBuilder) Build() *TransactionSplitSubaccountAddRequest {
	return &TransactionSplitSubaccountAddRequest{
		Subaccount: b.subaccount,
		Share:      b.share,
	}
}

type AddSubaccountResponse = types.Response[types.TransactionSplit]

func (c *Client) AddSubaccount(ctx context.Context, id string, builder *TransactionSplitSubaccountAddRequestBuilder) (*AddSubaccountResponse, error) {
	req := builder.Build()
	return net.Post[TransactionSplitSubaccountAddRequest, types.TransactionSplit](
		ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/subaccount/add", basePath, id), req, c.BaseURL,
	)
}
