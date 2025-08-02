package virtualterminal

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// UnassignDestination unassigns destinations from a virtual terminal
func (c *Client) UnassignDestination(ctx context.Context, code string, builder *UnassignDestinationRequestBuilder) (*types.Response[any], error) {
	if code == "" {
		return nil, errors.New("virtual terminal code is required")
	}
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	endpoint := fmt.Sprintf("%s/%s/destination/unassign", virtualTerminalBasePath, code)
	return net.Post[UnassignDestinationRequest, any](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
}
