package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// SendNotificationResponse represents the response from sending a notification
type SendNotificationResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// SendNotification sends notification of a payment request to your customers
func (c *Client) SendNotification(ctx context.Context, code string) (*types.Response[any], error) {
	return net.Post[any, any](
		ctx, c.client, c.secret, paymentRequestsBasePath+"/notify/"+code, nil, c.baseURL,
	)
}
