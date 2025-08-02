package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// ArchivePaymentRequestResponse represents the response from archiving a payment request
type ArchivePaymentRequestResponse = types.Response[any]

// Archive archives a payment request. A payment request will no longer be fetched on list or returned on verify
func (c *Client) Archive(ctx context.Context, code string) (*ArchivePaymentRequestResponse, error) {
	return net.Post[any, any](ctx, c.Client, c.Secret, basePath+"/archive/"+code, nil, c.BaseURL)
}
