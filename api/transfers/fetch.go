package transfers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchResponse = types.Response[types.Transfer]

func (c *Client) Fetch(ctx context.Context, idOrCode string) (*FetchResponse, error) {
	return net.Get[types.Transfer](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), "", c.BaseURL)
}
