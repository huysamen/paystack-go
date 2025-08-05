package verification

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ResolveCardBINResponseData = types.CardBINResolution
type ResolveCardBINResponse = types.Response[ResolveCardBINResponseData]

func (c *Client) ResolveCardBIN(ctx context.Context, bin string) (*ResolveCardBINResponse, error) {
	return net.Get[ResolveCardBINResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", cardBINResolveBasePath, bin), "", c.BaseURL)
}
