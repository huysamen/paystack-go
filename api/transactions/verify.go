package transactions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Response type alias
type VerifyResponse = types.Response[types.Transaction]

func (c *Client) Verify(ctx context.Context, reference string) (*VerifyResponse, error) {
	return net.Get[types.Transaction](ctx, c.Client, c.Secret, fmt.Sprintf("%s%s/%s", basePath, transactionVerifyPath, reference), "", c.BaseURL)
}
