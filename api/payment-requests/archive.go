package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Archive archives a payment request. A payment request will no longer be fetched on list or returned on verify
func (c *Client) Archive(ctx context.Context, code string) (*types.Response[any], error) {

	return net.Post[any, any](
		ctx, c.client, c.secret, paymentRequestsBasePath+"/archive/"+code, nil, c.baseURL,
	)
}
