package transferrecipients

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransferRecipientFetchResponse = types.Response[types.TransferRecipient]

func (c *Client) Fetch(ctx context.Context, idOrCode string) (*TransferRecipientFetchResponse, error) {
	return net.Get[types.TransferRecipient](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), "", c.BaseURL)
}
