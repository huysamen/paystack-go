package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type splitTransactionRequest struct {
	Customer      string `json:"customer"`
	Subaccount    string `json:"subaccount,omitempty"`
	SplitCode     string `json:"split_code,omitempty"`
	PreferredBank string `json:"preferred_bank,omitempty"`
}

type SplitTransactioRequestnBuilder struct {
	request *splitTransactionRequest
}

func NewSplitTransactionRequestBuilder() *SplitTransactioRequestnBuilder {
	return &SplitTransactioRequestnBuilder{
		request: &splitTransactionRequest{},
	}
}

func (b *SplitTransactioRequestnBuilder) Customer(customer string) *SplitTransactioRequestnBuilder {
	b.request.Customer = customer

	return b
}

func (b *SplitTransactioRequestnBuilder) Subaccount(subaccount string) *SplitTransactioRequestnBuilder {
	b.request.Subaccount = subaccount

	return b
}

func (b *SplitTransactioRequestnBuilder) SplitCode(splitCode string) *SplitTransactioRequestnBuilder {
	b.request.SplitCode = splitCode

	return b
}

func (b *SplitTransactioRequestnBuilder) PreferredBank(preferredBank string) *SplitTransactioRequestnBuilder {
	b.request.PreferredBank = preferredBank

	return b
}

func (b *SplitTransactioRequestnBuilder) Build() *splitTransactionRequest {
	return b.request
}

type SplitTransactionResponseData = types.DedicatedVirtualAccount
type SplitTransactionResponse = types.Response[SplitTransactionResponseData]

func (c *Client) SplitTransaction(ctx context.Context, builder SplitTransactioRequestnBuilder) (*SplitTransactionResponse, error) {
	return net.Post[splitTransactionRequest, SplitTransactionResponseData](ctx, c.Client, c.Secret, basePath+"/split", builder.Build(), c.BaseURL)
}
