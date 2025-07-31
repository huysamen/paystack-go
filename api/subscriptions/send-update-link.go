package subscriptions

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SendUpdateLinkResponse struct {
	Message string `json:"message"`
}

func (c *Client) SendUpdateLink(ctx context.Context, code string) (*types.Response[SendUpdateLinkResponse], error) {
	if code == "" {
		return nil, errors.New("subscription code is required")
	}

	path := fmt.Sprintf("%s/%s/manage/email", subscriptionBasePath, code)

	// This endpoint doesn't require a request body, but we need to pass something to Post
	var emptyRequest struct{}

	return net.Post[struct{}, SendUpdateLinkResponse](
		ctx,
		c.client,
		c.secret,
		path,
		&emptyRequest,
		c.baseURL,
	)
}
