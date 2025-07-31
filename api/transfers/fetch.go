package transfers

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

func (c *Client) Fetch(ctx context.Context, idOrCode string) (*types.Response[Transfer], error) {
	if idOrCode == "" {
		return nil, errors.New("transfer ID or code is required")
	}

	path := fmt.Sprintf("%s/%s", transferBasePath, idOrCode)

	return net.Get[Transfer](
		ctx,
		c.client,
		c.secret,
		path,
		c.baseURL,
	)
}
