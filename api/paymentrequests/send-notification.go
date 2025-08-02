package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// SendNotificationResponse represents the response from sending a notification
type SendNotificationResponse = types.Response[any]

// SendNotification sends notification of a payment request to your customers
func (c *Client) SendNotification(ctx context.Context, code string) (*SendNotificationResponse, error) {
	return net.Post[any, any](ctx, c.Client, c.Secret, basePath+"/notify/"+code, nil, c.BaseURL)
}
