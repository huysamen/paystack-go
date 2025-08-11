package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type updateRequest struct {
	Name    *string `json:"name,omitempty"`    // Name of the terminal (optional)
	Address *string `json:"address,omitempty"` // Address of the terminal (optional)
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

func (b *UpdateRequestBuilder) Address(address string) *UpdateRequestBuilder {
	b.req.Address = &address

	return b
}

func (b *UpdateRequestBuilder) Build() *updateRequest {
	return b.req
}

type UpdateResponseData = types.Terminal
type UpdateResponse = types.Response[UpdateResponseData]

func (c *Client) Update(ctx context.Context, terminalID string, builder UpdateRequestBuilder) (*UpdateResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", basePath, terminalID)

	return net.Put[updateRequest, UpdateResponseData](ctx, c.Client, c.Secret, endpoint, builder.Build(), c.BaseURL)
}
