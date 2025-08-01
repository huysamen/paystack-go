package plans

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FetchPlanResponse represents the response from fetching a plan
type FetchPlanResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    types.Plan `json:"data"`
}

// FetchByID fetches a plan by its ID
func (c *Client) FetchByID(ctx context.Context, id uint64) (*types.Plan, error) {
	resp, err := net.Get[types.Plan](
		ctx,
		c.client,
		c.secret,
		fmt.Sprintf("%s/%d", planBasePath, id),
		c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

// FetchByCode fetches a plan by its code
func (c *Client) FetchByCode(ctx context.Context, code string) (*types.Plan, error) {
	resp, err := net.Get[types.Plan](
		ctx,
		c.client,
		c.secret,
		fmt.Sprintf("%s/%s", planBasePath, code),
		c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

// Fetch fetches a plan by ID or code (convenience method)
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*types.Plan, error) {
	resp, err := net.Get[types.Plan](
		ctx,
		c.client,
		c.secret,
		fmt.Sprintf("%s/%s", planBasePath, idOrCode),
		c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
