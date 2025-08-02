package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// AddSplitCode adds a split code to a virtual terminal
func (c *Client) AddSplitCode(ctx context.Context, code string, builder *AddSplitCodeRequestBuilder) (*types.Response[any], error) {
	return net.Put[AddSplitCodeRequest, any](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/split_code", basePath, code), builder.Build(), c.BaseURL)
}
