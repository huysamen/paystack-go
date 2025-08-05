package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type assignDestinationRequest struct {
	Destinations []types.VirtualTerminalDestination `json:"destinations"`
}

type AssignDestinationRequestBuilder struct {
	destinations []types.VirtualTerminalDestination
}

func NewAssignDestinationRequestBuilder() *AssignDestinationRequestBuilder {
	return &AssignDestinationRequestBuilder{
		destinations: make([]types.VirtualTerminalDestination, 0),
	}
}

func (b *AssignDestinationRequestBuilder) AddDestination(destination types.VirtualTerminalDestination) *AssignDestinationRequestBuilder {
	b.destinations = append(b.destinations, destination)

	return b
}

func (b *AssignDestinationRequestBuilder) Destinations(destinations []types.VirtualTerminalDestination) *AssignDestinationRequestBuilder {
	b.destinations = destinations

	return b
}

func (b *AssignDestinationRequestBuilder) Build() *assignDestinationRequest {
	return &assignDestinationRequest{
		Destinations: b.destinations,
	}
}

type AssignDestinationResponseData = []types.VirtualTerminalDestination
type AssignDestinationResponse = types.Response[AssignDestinationResponseData]

func (c *Client) AssignDestination(ctx context.Context, code string, builder AssignDestinationRequestBuilder) (*AssignDestinationResponse, error) {
	return net.Post[assignDestinationRequest, AssignDestinationResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/destination/assign", basePath, code), builder.Build(), c.BaseURL)
}
