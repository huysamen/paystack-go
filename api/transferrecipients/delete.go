package transferrecipients

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransferRecipientDeleteResponse = types.Response[any]

func (c *Client) Delete(ctx context.Context, idOrCode string) (*TransferRecipientDeleteResponse, error) {
	return net.Delete[any](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), c.BaseURL)
}
