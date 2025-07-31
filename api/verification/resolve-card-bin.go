package verification

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// ResolveCardBIN resolves card BIN information
func (c *Client) ResolveCardBIN(ctx context.Context, bin string) (*CardBINResolveResponse, error) {
	if bin == "" {
		return nil, fmt.Errorf("bin is required")
	}
	if len(bin) < 6 {
		return nil, fmt.Errorf("bin must be at least 6 characters")
	}

	// Use only first 6 characters for BIN resolution
	if len(bin) > 6 {
		bin = bin[:6]
	}

	endpoint := fmt.Sprintf("%s/%s", cardBINResolveBasePath, bin)

	resp, err := net.Get[CardBINResolveResponse](ctx, c.client, c.secret, endpoint, c.baseURL)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
