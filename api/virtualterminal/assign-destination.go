package virtualterminal

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// AssignDestination assigns destinations to a virtual terminal
func (c *Client) AssignDestination(ctx context.Context, code string, builder *AssignDestinationRequestBuilder) (*types.Response[[]VirtualTerminalDestination], error) {
	if code == "" {
		return nil, errors.New("virtual terminal code is required")
	}
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	endpoint := fmt.Sprintf("%s/%s/destination/assign", virtualTerminalBasePath, code)
	return net.Post[AssignDestinationRequest, []VirtualTerminalDestination](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
}
