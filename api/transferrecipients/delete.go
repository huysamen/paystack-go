package transferrecipients

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type DeleteResponseData = any
type DeleteResponse = types.Response[DeleteResponseData]

func (c *Client) Delete(ctx context.Context, idOrCode string) (*DeleteResponse, error) {
	return net.Delete[DeleteResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), c.BaseURL)
}
