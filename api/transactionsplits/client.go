package transactionsplits

import "github.com/huysamen/paystack-go/api"

const basePath = "/split"

type Client api.API

func NewClient(c api.API) *Client {
	client := Client(c)
	return &client
}
