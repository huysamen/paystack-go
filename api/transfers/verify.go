package transfers

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

func (c *Client) Verify(ctx context.Context, reference string) (*types.Response[Transfer], error) {
	if reference == "" {
		return nil, errors.New("transfer reference is required")
	}

	path := fmt.Sprintf("%s/verify/%s", transferBasePath, reference)

	return net.Get[Transfer](
		ctx,
		c.client,
		c.secret,
		path,
		c.baseURL,
	)
}
