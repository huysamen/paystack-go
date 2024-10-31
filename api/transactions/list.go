package transactions

import (
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionListResponse []types.Transaction

func (c *Client) List() (*types.Response[TransactionListResponse], error) {
	return net.Get[TransactionListResponse](
		c.client,
		c.secret,
		fmt.Sprintf("%s", transactionBasePath),
	)
}

func (c *Client) ListPage(perPage, page int) (*types.Response[TransactionListResponse], error) {
	return net.Get[TransactionListResponse](
		c.client,
		c.secret,
		fmt.Sprintf("%s?perPage=%d&page=%d", transactionBasePath, perPage, page),
	)
}
