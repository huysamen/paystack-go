package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// AssignDestination assigns destinations to a virtual terminal
func (c *Client) AssignDestination(ctx context.Context, code string, builder *AssignDestinationRequestBuilder) (*types.Response[[]VirtualTerminalDestination], error) {
	return net.Post[AssignDestinationRequest, []VirtualTerminalDestination](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/destination/assign", basePath, code), builder.Build(), c.BaseURL)
}
