package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type updateRequest struct {
	Name string `json:"name"`
}

type UpdateRequestBuilder struct {
	name string
}

func NewUpdateRequestBuilder(name string) *UpdateRequestBuilder {
	return &UpdateRequestBuilder{
		name: name,
	}
}

func (b *UpdateRequestBuilder) Name(name string) *UpdateRequestBuilder {
	b.name = name
	return b
}

func (b *UpdateRequestBuilder) Build() *updateRequest {
	return &updateRequest{
		Name: b.name,
	}
}

type UpdateResponseData = types.VirtualTerminal
type UpdateResponse = types.Response[UpdateResponseData]

func (c *Client) Update(ctx context.Context, code string, builder UpdateRequestBuilder) (*UpdateResponse, error) {
	return net.Put[updateRequest, UpdateResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, code), builder.Build(), c.BaseURL)
}
