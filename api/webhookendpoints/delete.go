package webhookendpoints

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
)

type DeleteResponseData struct {
	WebhookURLCode data.String `json:"webhook_url_code"`
}

type DeleteResponse = types.Response[DeleteResponseData]

// Delete removes an endpoint for good. To stop events for a while, Update it
// with Enabled(false) instead.
func (c *Client) Delete(ctx context.Context, code string) (*DeleteResponse, error) {
	return net.Delete[DeleteResponseData](ctx, c.Client, c.Secret, endpointPath(code), c.BaseURL)
}
