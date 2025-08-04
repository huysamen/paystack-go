package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// AssignDestinationRequest represents the request to assign destinations to a virtual terminal
type AssignDestinationRequest struct {
	Destinations []types.VirtualTerminalDestination `json:"destinations"`
}

// AssignDestinationRequestBuilder provides a fluent interface for building AssignDestinationRequest
type AssignDestinationRequestBuilder struct {
	destinations []types.VirtualTerminalDestination
}

// NewAssignDestinationRequest creates a new builder for assigning destinations
func NewAssignDestinationRequest() *AssignDestinationRequestBuilder {
	return &AssignDestinationRequestBuilder{
		destinations: make([]types.VirtualTerminalDestination, 0),
	}
}

// AddDestination adds a destination to the assignment request
func (b *AssignDestinationRequestBuilder) AddDestination(destination types.VirtualTerminalDestination) *AssignDestinationRequestBuilder {
	b.destinations = append(b.destinations, destination)

	return b
}

// Destinations sets all destinations at once
func (b *AssignDestinationRequestBuilder) Destinations(destinations []types.VirtualTerminalDestination) *AssignDestinationRequestBuilder {
	b.destinations = destinations

	return b
}

// Build creates the AssignDestinationRequest
func (b *AssignDestinationRequestBuilder) Build() *AssignDestinationRequest {
	return &AssignDestinationRequest{
		Destinations: b.destinations,
	}
}

// AssignDestinationResponse represents the response from assigning destinations
type AssignDestinationResponse = types.Response[[]types.VirtualTerminalDestination]

// AssignDestination assigns destinations to a virtual terminal
func (c *Client) AssignDestination(ctx context.Context, code string, builder *AssignDestinationRequestBuilder) (*AssignDestinationResponse, error) {
	return net.Post[AssignDestinationRequest, []types.VirtualTerminalDestination](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/destination/assign", basePath, code), builder.Build(), c.BaseURL)
}
