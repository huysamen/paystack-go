package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// UnassignDestinationRequest represents the request to unassign destinations from a virtual terminal
type UnassignDestinationRequest struct {
	Targets []string `json:"targets"`
}

// UnassignDestinationRequestBuilder provides a fluent interface for building UnassignDestinationRequest
type UnassignDestinationRequestBuilder struct {
	targets []string
}

// NewUnassignDestinationRequest creates a new builder for unassigning destinations
func NewUnassignDestinationRequest() *UnassignDestinationRequestBuilder {
	return &UnassignDestinationRequestBuilder{
		targets: make([]string, 0),
	}
}

// AddTarget adds a target to the unassignment request
func (b *UnassignDestinationRequestBuilder) AddTarget(target string) *UnassignDestinationRequestBuilder {
	b.targets = append(b.targets, target)

	return b
}

// Targets sets all targets at once
func (b *UnassignDestinationRequestBuilder) Targets(targets []string) *UnassignDestinationRequestBuilder {
	b.targets = targets

	return b
}

// Build creates the UnassignDestinationRequest
func (b *UnassignDestinationRequestBuilder) Build() *UnassignDestinationRequest {
	return &UnassignDestinationRequest{
		Targets: b.targets,
	}
}

// UnassignDestinationResponse represents the response from unassigning destinations
type UnassignDestinationResponse = types.Response[any]

// UnassignDestination unassigns destinations from a virtual terminal
func (c *Client) UnassignDestination(ctx context.Context, code string, builder *UnassignDestinationRequestBuilder) (*UnassignDestinationResponse, error) {
	return net.Post[UnassignDestinationRequest, any](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/destination/unassign", basePath, code), builder.Build(), c.BaseURL)
}
