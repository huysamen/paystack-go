package transfers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type VerifyResponse = types.Response[types.Transfer]

func (c *Client) Verify(ctx context.Context, reference string) (*VerifyResponse, error) {

	return net.Get[types.Transfer](ctx, c.Client, c.Secret, fmt.Sprintf("%s/verify/%s", basePath, reference), "", c.BaseURL)
}
