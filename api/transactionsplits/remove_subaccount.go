package transactionsplits

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type removeSubaccountRequest struct {
	Subaccount string `json:"subaccount"` // Subaccount code
}

type RemoveSubaccountRequestBuilder struct {
	subaccount string
}

func NewRemoveSubaccountRequestBuilder(subaccount string) *RemoveSubaccountRequestBuilder {
	return &RemoveSubaccountRequestBuilder{
		subaccount: subaccount,
	}
}

func (b *RemoveSubaccountRequestBuilder) Subaccount(subaccount string) *RemoveSubaccountRequestBuilder {
	b.subaccount = subaccount

	return b
}

func (b *RemoveSubaccountRequestBuilder) Build() *removeSubaccountRequest {
	return &removeSubaccountRequest{
		Subaccount: b.subaccount,
	}
}

type RemoveSubaccountResponseData = any
type RemoveSubaccountResponse = types.Response[RemoveSubaccountResponseData]

func (c *Client) RemoveSubaccount(ctx context.Context, id string, builder RemoveSubaccountRequestBuilder) (*RemoveSubaccountResponse, error) {
	return net.Post[removeSubaccountRequest, RemoveSubaccountResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/subaccount/remove", basePath, id), builder.Build(), c.BaseURL)
}
