package plans

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FetchPlanResponse represents the response from fetching a plan
type FetchPlanResponse = types.Response[types.Plan]

// FetchByID fetches a plan by its ID
func (c *Client) FetchByID(ctx context.Context, id uint64) (*FetchPlanResponse, error) {
	return net.Get[types.Plan](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%d", basePath, id), c.BaseURL)
}

// FetchByCode fetches a plan by its code
func (c *Client) FetchByCode(ctx context.Context, code string) (*FetchPlanResponse, error) {
	return net.Get[types.Plan](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, code), c.BaseURL)
}

// Fetch fetches a plan by ID or code (convenience method)
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*FetchPlanResponse, error) {
	return net.Get[types.Plan](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), c.BaseURL)
}
