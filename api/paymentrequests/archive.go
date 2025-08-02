package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// ArchivePaymentRequestResponse represents the response from archiving a payment request
type ArchivePaymentRequestResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// Archive archives a payment request. A payment request will no longer be fetched on list or returned on verify
func (c *Client) Archive(ctx context.Context, code string) (*types.Response[any], error) {
	return net.Post[any, any](
		ctx, c.client, c.secret, paymentRequestsBasePath+"/archive/"+code, nil, c.baseURL,
	)
}
