package transactionsplits

import (
	"context"

	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
)

type createRequest struct {
	Name             string                             `json:"name"`                        // Name of the transaction split
	Type             enums.TransactionSplitType         `json:"type"`                        // Type of split (percentage or flat)
	Currency         enums.Currency                     `json:"currency"`                    // Currency for the split
	Subaccounts      []types.TransactionSplitSubaccount `json:"subaccounts"`                 // List of subaccounts and their shares
	BearerType       *enums.TransactionSplitBearerType  `json:"bearer_type,omitempty"`       // Who bears the charges (optional)
	BearerSubaccount *string                            `json:"bearer_subaccount,omitempty"` // Subaccount code if bearer_type is subaccount (optional)
}

type CreateRequestBuilder struct {
	name             string
	splitType        enums.TransactionSplitType
	currency         enums.Currency
	subaccounts      []types.TransactionSplitSubaccount
	bearerType       *enums.TransactionSplitBearerType
	bearerSubaccount *string
}

func NewCreateRequest(name string, splitType enums.TransactionSplitType, currency enums.Currency) *CreateRequestBuilder {
	return &CreateRequestBuilder{
		name:        name,
		splitType:   splitType,
		currency:    currency,
		subaccounts: make([]types.TransactionSplitSubaccount, 0),
	}
}

func (b *CreateRequestBuilder) AddSubaccount(subaccount string, share int) *CreateRequestBuilder {
	b.subaccounts = append(b.subaccounts, types.TransactionSplitSubaccount{
		Subaccount: data.NewString(subaccount),
		Share:      data.NewInt(int64(share)),
	})
	return b
}

func (b *CreateRequestBuilder) Subaccounts(subaccounts []types.TransactionSplitSubaccount) *CreateRequestBuilder {
	b.subaccounts = subaccounts
	return b
}

func (b *CreateRequestBuilder) BearerType(bearerType enums.TransactionSplitBearerType) *CreateRequestBuilder {
	b.bearerType = &bearerType
	return b
}

func (b *CreateRequestBuilder) BearerSubaccount(subaccount string) *CreateRequestBuilder {
	b.bearerSubaccount = &subaccount
	return b
}

func (b *CreateRequestBuilder) Build() *createRequest {
	return &createRequest{
		Name:             b.name,
		Type:             b.splitType,
		Currency:         b.currency,
		Subaccounts:      b.subaccounts,
		BearerType:       b.bearerType,
		BearerSubaccount: b.bearerSubaccount,
	}
}

type CreateResponseData = types.TransactionSplit
type CreateResponse = types.Response[CreateResponseData]

func (c *Client) Create(ctx context.Context, builder CreateRequestBuilder) (*CreateResponse, error) {
	req := builder.Build()
	return net.Post[createRequest, CreateResponseData](ctx, c.Client, c.Secret, basePath, req, c.BaseURL)
}
