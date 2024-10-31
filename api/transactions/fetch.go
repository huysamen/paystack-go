package transactions

import (
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

func (c *Client) Fetch(id uint64) (*types.Response[types.Transaction], error) {
	return net.Get[types.Transaction](
		c.client,
		c.secret,
		fmt.Sprintf("%s/%d", transactionBasePath, id),
	)
}
