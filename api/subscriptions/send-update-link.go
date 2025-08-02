package subscriptions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SendUpdateLinkResponse struct {
	Message string `json:"message"`
}

// SubscriptionSendUpdateLinkResponse represents the response from sending update link
type SubscriptionSendUpdateLinkResponse = types.Response[SendUpdateLinkResponse]

func (c *Client) SendUpdateLink(ctx context.Context, code string) (*SubscriptionSendUpdateLinkResponse, error) {
	return net.Post[struct{}, SendUpdateLinkResponse](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s/manage/email", basePath, code), &struct{}{}, c.BaseURL)
}
