package subscriptions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type GenerateUpdateLinkResponseData struct {
	Link string `json:"link"`
}

type GenerateUpdateLinkResponse = types.Response[GenerateUpdateLinkResponseData]

func (c *Client) GenerateUpdateLink(ctx context.Context, code string) (*GenerateUpdateLinkResponse, error) {
	return net.Get[GenerateUpdateLinkResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/manage/link", basePath, code), c.BaseURL)
}
