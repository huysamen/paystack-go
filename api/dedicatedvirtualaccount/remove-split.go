package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type RemoveSplitFromDedicatedAccountRequest struct {
	AccountNumber string `json:"account_number"`
}

type RemoveSplitFromDedicatedAccountBuilder struct {
	request *RemoveSplitFromDedicatedAccountRequest
}

func NewRemoveSplitFromDedicatedAccountBuilder() *RemoveSplitFromDedicatedAccountBuilder {
	return &RemoveSplitFromDedicatedAccountBuilder{
		request: &RemoveSplitFromDedicatedAccountRequest{},
	}
}

func (b *RemoveSplitFromDedicatedAccountBuilder) AccountNumber(accountNumber string) *RemoveSplitFromDedicatedAccountBuilder {
	b.request.AccountNumber = accountNumber

	return b
}

func (b *RemoveSplitFromDedicatedAccountBuilder) Build() *RemoveSplitFromDedicatedAccountRequest {
	return b.request
}

type RemoveSplitFromDedicatedAccountResponse = types.Response[types.DedicatedVirtualAccount]

func (c *Client) RemoveSplit(ctx context.Context, builder *RemoveSplitFromDedicatedAccountBuilder) (*RemoveSplitFromDedicatedAccountResponse, error) {
	return net.DeleteWithBody[RemoveSplitFromDedicatedAccountRequest, types.DedicatedVirtualAccount](ctx, c.Client, c.Secret, basePath+"/split", builder.Build(), c.BaseURL)
}
