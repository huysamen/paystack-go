package subscriptions

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type GenerateUpdateLinkResponse struct {
	Link string `json:"link"`
}

func (c *Client) GenerateUpdateLink(ctx context.Context, code string) (*types.Response[GenerateUpdateLinkResponse], error) {
	if code == "" {
		return nil, errors.New("subscription code is required")
	}

	path := fmt.Sprintf("%s/%s/manage/link", subscriptionBasePath, code)

	return net.Get[GenerateUpdateLinkResponse](
		ctx,
		c.client,
		c.secret,
		path,
		c.baseURL,
	)
}
