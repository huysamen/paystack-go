package verification

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// ResolveCardBIN resolves card BIN information
func (c *Client) ResolveCardBIN(ctx context.Context, bin string) (*types.Response[CardBINResolution], error) {
	if bin == "" {
		return nil, errors.New("bin is required")
	}
	if len(bin) < 6 {
		return nil, errors.New("bin must be at least 6 characters")
	}

	// Use only first 6 characters for BIN resolution
	if len(bin) > 6 {
		bin = bin[:6]
	}

	endpoint := fmt.Sprintf("%s/%s", cardBINResolveBasePath, bin)
	return net.Get[CardBINResolution](ctx, c.client, c.secret, endpoint, "", c.baseURL)
}
