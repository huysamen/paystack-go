package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type UnassignDestinationRequest struct {
	Targets []string `json:"targets"`
}

type UnassignDestinationRequestBuilder struct {
	targets []string
}

func NewUnassignDestinationRequest() *UnassignDestinationRequestBuilder {
	return &UnassignDestinationRequestBuilder{
		targets: make([]string, 0),
	}
}

func (b *UnassignDestinationRequestBuilder) AddTarget(target string) *UnassignDestinationRequestBuilder {
	b.targets = append(b.targets, target)

	return b
}

func (b *UnassignDestinationRequestBuilder) Targets(targets []string) *UnassignDestinationRequestBuilder {
	b.targets = targets

	return b
}

func (b *UnassignDestinationRequestBuilder) Build() *UnassignDestinationRequest {
	return &UnassignDestinationRequest{
		Targets: b.targets,
	}
}

type UnassignDestinationResponse = types.Response[any]

func (c *Client) UnassignDestination(ctx context.Context, code string, builder *UnassignDestinationRequestBuilder) (*UnassignDestinationResponse, error) {
	return net.Post[UnassignDestinationRequest, any](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/destination/unassign", basePath, code), builder.Build(), c.BaseURL)
}
