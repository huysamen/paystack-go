package subscriptions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type GenerateUpdateLinkResponse struct {
	Link string `json:"link"`
}

type SubscriptionGenerateUpdateLinkResponse = types.Response[GenerateUpdateLinkResponse]

func (c *Client) GenerateUpdateLink(ctx context.Context, code string) (*SubscriptionGenerateUpdateLinkResponse, error) {
	return net.Get[GenerateUpdateLinkResponse](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/manage/link", basePath, code), c.BaseURL)
}
