package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ArchivePaymentRequestResponse = types.Response[any]

func (c *Client) Archive(ctx context.Context, code string) (*ArchivePaymentRequestResponse, error) {
	return net.Post[any, any](ctx, c.Client, c.Secret, basePath+"/archive/"+code, nil, c.BaseURL)
}
