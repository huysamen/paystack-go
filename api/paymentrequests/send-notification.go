package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SendNotificationResponse = types.Response[any]

func (c *Client) SendNotification(ctx context.Context, code string) (*SendNotificationResponse, error) {
	return net.Post[any, any](ctx, c.Client, c.Secret, basePath+"/notify/"+code, nil, c.BaseURL)
}
