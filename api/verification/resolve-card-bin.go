package verification

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// CardBINResolveResponse represents the response from resolving a card BIN
type CardBINResolveResponse = types.Response[types.CardBINResolution]

// ResolveCardBIN resolves card BIN information
func (c *Client) ResolveCardBIN(ctx context.Context, bin string) (*types.Response[types.CardBINResolution], error) {
	// Use only first 6 characters for BIN resolution
	if len(bin) > 6 {
		bin = bin[:6]
	}
	return net.Get[types.CardBINResolution](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", cardBINResolveBasePath, bin), "", c.BaseURL)
}
