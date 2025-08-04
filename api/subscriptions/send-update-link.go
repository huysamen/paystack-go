package subscriptions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// SendUpdateSubscriptionLinkResponse represents the response from sending update link
type SendUpdateSubscriptionLinkResponse = types.Response[any]

func (c *Client) SendUpdateLink(ctx context.Context, code string) (*SendUpdateSubscriptionLinkResponse, error) {
	return net.Post[any, any](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/manage/email", basePath, code), nil, c.BaseURL)
}
