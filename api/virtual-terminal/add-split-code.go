package virtualterminal

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// AddSplitCode adds a split code to a virtual terminal
func (c *Client) AddSplitCode(ctx context.Context, code string, builder *AddSplitCodeRequestBuilder) (*types.Response[any], error) {
	if code == "" {
		return nil, errors.New("virtual terminal code is required")
	}
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	endpoint := fmt.Sprintf("%s/%s/split_code", virtualTerminalBasePath, code)
	return net.Put[AddSplitCodeRequest, any](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
}
