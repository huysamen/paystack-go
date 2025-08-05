package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type DeactivateResponseData = any
type DeactivateResponse = types.Response[DeactivateResponseData]

func (c *Client) Deactivate(ctx context.Context, code string) (*DeactivateResponse, error) {
	return net.Put[any, DeactivateResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/deactivate", basePath, code), nil, c.BaseURL)
}
