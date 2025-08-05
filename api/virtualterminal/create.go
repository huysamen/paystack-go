package virtualterminal

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type createRequest struct {
	Name         string                             `json:"name"`
	Destinations []types.VirtualTerminalDestination `json:"destinations,omitempty"`
	Metadata     *types.Metadata                    `json:"metadata,omitempty"`
	Currency     string                             `json:"currency,omitempty"`
	CustomFields []types.VirtualTerminalCustomField `json:"custom_fields,omitempty"`
}

type CreateRequestBuilder struct {
	req *createRequest
}

func NewRequestBuilder(name string) *CreateRequestBuilder {
	return &CreateRequestBuilder{
		req: &createRequest{
			Name: name,
		},
	}
}

func (b *CreateRequestBuilder) Build() *createRequest {
	return b.req
}

type CreateResponseData = types.VirtualTerminal
type CreateResponse = types.Response[CreateResponseData]

func (c *Client) Create(ctx context.Context, builder CreateRequestBuilder) (*CreateResponse, error) {
	return net.Post[createRequest, CreateResponseData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
