package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// SplitDedicatedAccountTransactionRequest represents the request to add split to dedicated account
type SplitDedicatedAccountTransactionRequest struct {
	Customer      string `json:"customer"`
	Subaccount    string `json:"subaccount,omitempty"`
	SplitCode     string `json:"split_code,omitempty"`
	PreferredBank string `json:"preferred_bank,omitempty"`
}

// SplitDedicatedAccountTransactionBuilder builds requests for splitting dedicated account transactions
type SplitDedicatedAccountTransactionBuilder struct {
	request *SplitDedicatedAccountTransactionRequest
}

// NewSplitDedicatedAccountTransactionBuilder creates a new builder for splitting dedicated account transactions
func NewSplitDedicatedAccountTransactionBuilder() *SplitDedicatedAccountTransactionBuilder {
	return &SplitDedicatedAccountTransactionBuilder{
		request: &SplitDedicatedAccountTransactionRequest{},
	}
}

// Customer sets the customer for splitting the dedicated account transaction
func (b *SplitDedicatedAccountTransactionBuilder) Customer(customer string) *SplitDedicatedAccountTransactionBuilder {
	b.request.Customer = customer

	return b
}

// Subaccount sets the subaccount for splitting the dedicated account transaction
func (b *SplitDedicatedAccountTransactionBuilder) Subaccount(subaccount string) *SplitDedicatedAccountTransactionBuilder {
	b.request.Subaccount = subaccount

	return b
}

// SplitCode sets the split code for splitting the dedicated account transaction
func (b *SplitDedicatedAccountTransactionBuilder) SplitCode(splitCode string) *SplitDedicatedAccountTransactionBuilder {
	b.request.SplitCode = splitCode

	return b
}

// PreferredBank sets the preferred bank for splitting the dedicated account transaction
func (b *SplitDedicatedAccountTransactionBuilder) PreferredBank(preferredBank string) *SplitDedicatedAccountTransactionBuilder {
	b.request.PreferredBank = preferredBank

	return b
}

// Build returns the built request
func (b *SplitDedicatedAccountTransactionBuilder) Build() *SplitDedicatedAccountTransactionRequest {
	return b.request
}

// SplitDedicatedAccountTransactionResponse represents the response type for splitting a dedicated account transaction
type SplitDedicatedAccountTransactionResponse = types.Response[types.DedicatedVirtualAccount]

// SplitTransaction splits a dedicated virtual account transaction with one or more accounts
func (c *Client) SplitTransaction(ctx context.Context, builder *SplitDedicatedAccountTransactionBuilder) (*SplitDedicatedAccountTransactionResponse, error) {
	return net.Post[SplitDedicatedAccountTransactionRequest, types.DedicatedVirtualAccount](ctx, c.Client, c.Secret, basePath+"/split", builder.Build(), c.BaseURL)
}
