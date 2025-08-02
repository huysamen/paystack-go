package transfers

import "github.com/huysamen/paystack-go/api"

const basePath = "/transfer"

type Client api.API

func NewClient(c api.API) *Client {
	client := Client(c)
	return &client
}
