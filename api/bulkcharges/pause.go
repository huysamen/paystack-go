package bulkcharges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type PauseResponseData = any
type PauseResponse = types.Response[PauseResponseData]

func (c *Client) Pause(ctx context.Context, batchCode string) (*PauseResponse, error) {
	return net.Get[PauseResponseData](ctx, c.Client, c.Secret, pausePath+"/"+batchCode, c.BaseURL)
}
