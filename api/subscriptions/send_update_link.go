package subscriptions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SendUpdateLinkResponseData = any
type SendUpdateLinkResponse = types.Response[SendUpdateLinkResponseData]

func (c *Client) SendUpdateLink(ctx context.Context, code string) (*SendUpdateLinkResponse, error) {
	return net.Post[any, SendUpdateLinkResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/manage/email", basePath, code), nil, c.BaseURL)
}
