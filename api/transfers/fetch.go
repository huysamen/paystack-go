package transfers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FetchResponse represents the response for fetching a transfer
type FetchResponse = types.Response[types.Transfer]

// Fetch retrieves a transfer by ID or code
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*FetchResponse, error) {
	return net.Get[types.Transfer](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), "", c.BaseURL)
}
