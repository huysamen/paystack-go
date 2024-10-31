package transactions

import (
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

func (c *Client) ViewTimelineByID(id uint64) (*types.Response[types.Log], error) {
	return net.Get[types.Log](
		c.client,
		c.secret,
		fmt.Sprintf("%s%s/%d", transactionBasePath, transactionViewTimelinePath, id),
	)
}

func (c *Client) ViewTimelineByReference(reference string) (*types.Response[types.Log], error) {
	return net.Get[types.Log](
		c.client,
		c.secret,
		fmt.Sprintf("%s%s/%s", transactionBasePath, transactionViewTimelinePath, reference),
	)
}
