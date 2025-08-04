package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SendNotificationResponseData = any
type SendNotificationResponse = types.Response[SendNotificationResponseData]

func (c *Client) SendNotification(ctx context.Context, code string) (*SendNotificationResponse, error) {
	return net.Post[any, SendNotificationResponseData](ctx, c.Client, c.Secret, basePath+"/notify/"+code, nil, c.BaseURL)
}
