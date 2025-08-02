package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// UnassignDestination unassigns destinations from a virtual terminal
func (c *Client) UnassignDestination(ctx context.Context, code string, builder *UnassignDestinationRequestBuilder) (*types.Response[any], error) {
	return net.Post[UnassignDestinationRequest, any](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/destination/unassign", basePath, code), builder.Build(), c.BaseURL)
}
