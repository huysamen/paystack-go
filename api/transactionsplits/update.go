package transactionsplits

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type updateRequest struct {
	Name             *string                           `json:"name,omitempty"`
	Active           *bool                             `json:"active,omitempty"`
	BearerType       *types.TransactionSplitBearerType `json:"bearer_type,omitempty"`
	BearerSubaccount *string                           `json:"bearer_subaccount,omitempty"`
}

type UpdateRequestBuilder struct {
	req *updateRequest
}

func NewUpdateRequestBuilder() *UpdateRequestBuilder {
	return &UpdateRequestBuilder{
		req: &updateRequest{},
	}
}

func (b *UpdateRequestBuilder) Name(name string) *UpdateRequestBuilder {
	b.req.Name = &name

	return b
}

func (b *UpdateRequestBuilder) Active(active bool) *UpdateRequestBuilder {
	b.req.Active = &active

	return b
}

func (b *UpdateRequestBuilder) BearerType(bearerType types.TransactionSplitBearerType) *UpdateRequestBuilder {
	b.req.BearerType = &bearerType

	return b
}

func (b *UpdateRequestBuilder) BearerSubaccount(subaccount string) *UpdateRequestBuilder {
	b.req.BearerSubaccount = &subaccount

	return b
}

func (b *UpdateRequestBuilder) Build() *updateRequest {
	return b.req
}

type UpdateResponseData = types.TransactionSplit
type UpdateResponse = types.Response[UpdateResponseData]

func (c *Client) Update(ctx context.Context, id string, builder UpdateRequestBuilder) (*UpdateResponse, error) {
	return net.Put[updateRequest, UpdateResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, id), builder.Build(), c.BaseURL)
}
