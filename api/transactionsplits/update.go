package transactionsplits

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionSplitUpdateRequest struct {
	Name             *string                           `json:"name,omitempty"`              // Name of the transaction split (optional)
	Active           *bool                             `json:"active,omitempty"`            // Active status (optional)
	BearerType       *types.TransactionSplitBearerType `json:"bearer_type,omitempty"`       // Who bears the charges (optional)
	BearerSubaccount *string                           `json:"bearer_subaccount,omitempty"` // Subaccount code if bearer_type is subaccount (optional)
}

type TransactionSplitUpdateRequestBuilder struct {
	name             *string
	active           *bool
	bearerType       *types.TransactionSplitBearerType
	bearerSubaccount *string
}

func NewTransactionSplitUpdateRequest() *TransactionSplitUpdateRequestBuilder {
	return &TransactionSplitUpdateRequestBuilder{}
}

func (b *TransactionSplitUpdateRequestBuilder) Name(name string) *TransactionSplitUpdateRequestBuilder {
	b.name = &name
	return b
}

func (b *TransactionSplitUpdateRequestBuilder) Active(active bool) *TransactionSplitUpdateRequestBuilder {
	b.active = &active
	return b
}

func (b *TransactionSplitUpdateRequestBuilder) BearerType(bearerType types.TransactionSplitBearerType) *TransactionSplitUpdateRequestBuilder {
	b.bearerType = &bearerType
	return b
}

func (b *TransactionSplitUpdateRequestBuilder) BearerSubaccount(subaccount string) *TransactionSplitUpdateRequestBuilder {
	b.bearerSubaccount = &subaccount
	return b
}

func (b *TransactionSplitUpdateRequestBuilder) Build() *TransactionSplitUpdateRequest {
	return &TransactionSplitUpdateRequest{
		Name:             b.name,
		Active:           b.active,
		BearerType:       b.bearerType,
		BearerSubaccount: b.bearerSubaccount,
	}
}

type TransactionSplitUpdateResponse = types.Response[types.TransactionSplit]

func (c *Client) Update(ctx context.Context, id string, builder *TransactionSplitUpdateRequestBuilder) (*types.Response[types.TransactionSplit], error) {
	return net.Put[TransactionSplitUpdateRequest, types.TransactionSplit](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, id), builder.Build(), c.BaseURL)
}
