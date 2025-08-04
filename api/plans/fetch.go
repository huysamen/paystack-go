package plans

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchPlanResponse = types.Response[types.Plan]

func (c *Client) FetchByID(ctx context.Context, id uint64) (*FetchPlanResponse, error) {
	return net.Get[types.Plan](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%d", basePath, id), c.BaseURL)
}

func (c *Client) FetchByCode(ctx context.Context, code string) (*FetchPlanResponse, error) {
	return net.Get[types.Plan](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, code), c.BaseURL)
}

func (c *Client) Fetch(ctx context.Context, idOrCode string) (*FetchPlanResponse, error) {
	return net.Get[types.Plan](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), c.BaseURL)
}
