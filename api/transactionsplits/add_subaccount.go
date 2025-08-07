package transactionsplits

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type addSubaccountRequest struct {
	Subaccount string `json:"subaccount"` // Subaccount code
	Share      int    `json:"share"`      // Share amount (percentage or flat amount)
}

type AddSubaccountRequestBuilder struct {
	subaccount string
	share      int
}

func NewAddSubaccountRequestBuilder(subaccount string, share int) *AddSubaccountRequestBuilder {
	return &AddSubaccountRequestBuilder{
		subaccount: subaccount,
		share:      share,
	}
}

func (b *AddSubaccountRequestBuilder) Subaccount(subaccount string) *AddSubaccountRequestBuilder {
	b.subaccount = subaccount
	return b
}

func (b *AddSubaccountRequestBuilder) Share(share int) *AddSubaccountRequestBuilder {
	b.share = share
	return b
}

func (b *AddSubaccountRequestBuilder) Build() *addSubaccountRequest {
	return &addSubaccountRequest{
		Subaccount: b.subaccount,
		Share:      b.share,
	}
}

type AddSubaccountResponseData = types.TransactionSplit
type AddSubaccountResponse = types.Response[AddSubaccountResponseData]

func (c *Client) AddSubaccount(ctx context.Context, id string, builder AddSubaccountRequestBuilder) (*AddSubaccountResponse, error) {
	req := builder.Build()
	return net.Post[addSubaccountRequest, AddSubaccountResponseData](
		ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/subaccount/add", basePath, id), req, c.BaseURL,
	)
}
