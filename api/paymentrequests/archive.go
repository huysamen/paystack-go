package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ArchiveResponseData = any
type ArchiveResponse = types.Response[ArchiveResponseData]

func (c *Client) Archive(ctx context.Context, code string) (*ArchiveResponse, error) {
	return net.Post[any, ArchiveResponseData](ctx, c.Client, c.Secret, basePath+"/archive/"+code, nil, c.BaseURL)
}
