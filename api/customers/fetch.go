package customers

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

func (c *Client) Fetch(ctx context.Context, emailOrCode string) (*types.Response[CustomerWithRelations], error) {
	if emailOrCode == "" {
		return nil, errors.New("email or customer code is required")
	}

	path := fmt.Sprintf("%s/%s", customerBasePath, emailOrCode)

	return net.Get[CustomerWithRelations](
		ctx,
		c.client,
		c.secret,
		path,
		c.baseURL,
	)
}
