package verification

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CardBINResolveResponse = types.Response[types.CardBINResolution]

func (c *Client) ResolveCardBIN(ctx context.Context, bin string) (*CardBINResolveResponse, error) {
	if len(bin) > 6 {
		bin = bin[:6]
	}
	return net.Get[types.CardBINResolution](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", cardBINResolveBasePath, bin), "", c.BaseURL)
}
