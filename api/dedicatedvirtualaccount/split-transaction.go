package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SplitDedicatedAccountTransactionRequest struct {
	Customer      string `json:"customer"`
	Subaccount    string `json:"subaccount,omitempty"`
	SplitCode     string `json:"split_code,omitempty"`
	PreferredBank string `json:"preferred_bank,omitempty"`
}

type SplitDedicatedAccountTransactionBuilder struct {
	request *SplitDedicatedAccountTransactionRequest
}

func NewSplitDedicatedAccountTransactionBuilder() *SplitDedicatedAccountTransactionBuilder {
	return &SplitDedicatedAccountTransactionBuilder{
		request: &SplitDedicatedAccountTransactionRequest{},
	}
}

func (b *SplitDedicatedAccountTransactionBuilder) Customer(customer string) *SplitDedicatedAccountTransactionBuilder {
	b.request.Customer = customer

	return b
}

func (b *SplitDedicatedAccountTransactionBuilder) Subaccount(subaccount string) *SplitDedicatedAccountTransactionBuilder {
	b.request.Subaccount = subaccount

	return b
}

func (b *SplitDedicatedAccountTransactionBuilder) SplitCode(splitCode string) *SplitDedicatedAccountTransactionBuilder {
	b.request.SplitCode = splitCode

	return b
}

func (b *SplitDedicatedAccountTransactionBuilder) PreferredBank(preferredBank string) *SplitDedicatedAccountTransactionBuilder {
	b.request.PreferredBank = preferredBank

	return b
}

func (b *SplitDedicatedAccountTransactionBuilder) Build() *SplitDedicatedAccountTransactionRequest {
	return b.request
}

type SplitDedicatedAccountTransactionResponse = types.Response[types.DedicatedVirtualAccount]

func (c *Client) SplitTransaction(ctx context.Context, builder *SplitDedicatedAccountTransactionBuilder) (*SplitDedicatedAccountTransactionResponse, error) {
	return net.Post[SplitDedicatedAccountTransactionRequest, types.DedicatedVirtualAccount](ctx, c.Client, c.Secret, basePath+"/split", builder.Build(), c.BaseURL)
}
