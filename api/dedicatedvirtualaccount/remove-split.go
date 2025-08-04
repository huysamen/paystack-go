package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type RemoveSplitRequest struct {
	AccountNumber string `json:"account_number"`
}

type RemoveSplitRequestBuilder struct {
	request *RemoveSplitRequest
}

func NewRemoveSplitRequestBuilder() *RemoveSplitRequestBuilder {
	return &RemoveSplitRequestBuilder{
		request: &RemoveSplitRequest{},
	}
}

func (b *RemoveSplitRequestBuilder) AccountNumber(accountNumber string) *RemoveSplitRequestBuilder {
	b.request.AccountNumber = accountNumber

	return b
}

func (b *RemoveSplitRequestBuilder) Build() *RemoveSplitRequest {
	return b.request
}

type RemoveSplitResponseData = types.DedicatedVirtualAccount
type RemoveSplitResponse = types.Response[RemoveSplitResponseData]

func (c *Client) RemoveSplit(ctx context.Context, builder RemoveSplitRequestBuilder) (*RemoveSplitResponse, error) {
	return net.DeleteWithBody[RemoveSplitRequest, RemoveSplitResponseData](ctx, c.Client, c.Secret, basePath+"/split", builder.Build(), c.BaseURL)
}
