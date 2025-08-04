package transactionsplits

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionSplitCreateRequest struct {
	Name             string                             `json:"name"`                        // Name of the transaction split
	Type             types.TransactionSplitType         `json:"type"`                        // Type of split (percentage or flat)
	Currency         types.Currency                     `json:"currency"`                    // Currency for the split
	Subaccounts      []types.TransactionSplitSubaccount `json:"subaccounts"`                 // List of subaccounts and their shares
	BearerType       *types.TransactionSplitBearerType  `json:"bearer_type,omitempty"`       // Who bears the charges (optional)
	BearerSubaccount *string                            `json:"bearer_subaccount,omitempty"` // Subaccount code if bearer_type is subaccount (optional)
}

type TransactionSplitCreateRequestBuilder struct {
	name             string
	splitType        types.TransactionSplitType
	currency         types.Currency
	subaccounts      []types.TransactionSplitSubaccount
	bearerType       *types.TransactionSplitBearerType
	bearerSubaccount *string
}

func NewTransactionSplitCreateRequest(name string, splitType types.TransactionSplitType, currency types.Currency) *TransactionSplitCreateRequestBuilder {
	return &TransactionSplitCreateRequestBuilder{
		name:        name,
		splitType:   splitType,
		currency:    currency,
		subaccounts: make([]types.TransactionSplitSubaccount, 0),
	}
}

func (b *TransactionSplitCreateRequestBuilder) AddSubaccount(subaccount string, share int) *TransactionSplitCreateRequestBuilder {
	b.subaccounts = append(b.subaccounts, types.TransactionSplitSubaccount{
		Subaccount: subaccount,
		Share:      share,
	})
	return b
}

func (b *TransactionSplitCreateRequestBuilder) Subaccounts(subaccounts []types.TransactionSplitSubaccount) *TransactionSplitCreateRequestBuilder {
	b.subaccounts = subaccounts
	return b
}

func (b *TransactionSplitCreateRequestBuilder) BearerType(bearerType types.TransactionSplitBearerType) *TransactionSplitCreateRequestBuilder {
	b.bearerType = &bearerType
	return b
}

func (b *TransactionSplitCreateRequestBuilder) BearerSubaccount(subaccount string) *TransactionSplitCreateRequestBuilder {
	b.bearerSubaccount = &subaccount
	return b
}

func (b *TransactionSplitCreateRequestBuilder) Build() *TransactionSplitCreateRequest {
	return &TransactionSplitCreateRequest{
		Name:             b.name,
		Type:             b.splitType,
		Currency:         b.currency,
		Subaccounts:      b.subaccounts,
		BearerType:       b.bearerType,
		BearerSubaccount: b.bearerSubaccount,
	}
}

type TransactionSplitCreateResponse = types.Response[types.TransactionSplit]

func (c *Client) Create(ctx context.Context, builder *TransactionSplitCreateRequestBuilder) (*types.Response[types.TransactionSplit], error) {
	req := builder.Build()
	return net.Post[TransactionSplitCreateRequest, types.TransactionSplit](
		ctx, c.Client, c.Secret, basePath, req, c.BaseURL,
	)
}
