package dedicatedvirtualaccounts

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

type SplitTransactionRequestBuilder struct {
	request *splitTransactionRequest
}

func NewSplitTransactionRequestBuilder() *SplitTransactionRequestBuilder {
	return &SplitTransactionRequestBuilder{
		request: &splitTransactionRequest{},
	}
}

func (b *SplitTransactionRequestBuilder) Customer(customer string) *SplitTransactionRequestBuilder {
	b.request.Customer = customer

	return b
}

func (b *SplitTransactionRequestBuilder) Subaccount(subaccount string) *SplitTransactionRequestBuilder {
	b.request.Subaccount = subaccount

	return b
}

func (b *SplitTransactionRequestBuilder) SplitCode(splitCode string) *SplitTransactionRequestBuilder {
	b.request.SplitCode = splitCode

	return b
}

func (b *SplitTransactionRequestBuilder) PreferredBank(preferredBank string) *SplitTransactionRequestBuilder {
	b.request.PreferredBank = preferredBank

	return b
}

func (b *SplitTransactionRequestBuilder) Build() *splitTransactionRequest {
	return b.request
}

type SplitTransactionResponseData = types.DedicatedVirtualAccount
type SplitTransactionResponse = types.Response[SplitTransactionResponseData]

func (c *Client) SplitTransaction(ctx context.Context, builder SplitTransactionRequestBuilder) (*SplitTransactionResponse, error) {
	return net.Post[splitTransactionRequest, SplitTransactionResponseData](ctx, c.Client, c.Secret, basePath+"/split", builder.Build(), c.BaseURL)
}
