package webhookendpoints

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// EventsCatalogResponseData maps a category ("transactions") to its events.
type EventsCatalogResponseData = map[string][]types.WebhookCatalogEvent
type EventsCatalogResponse = types.Response[EventsCatalogResponseData]

// EventsCatalog lists the events an endpoint can subscribe to.
func (c *Client) EventsCatalog(ctx context.Context) (*EventsCatalogResponse, error) {
	return net.Get[EventsCatalogResponseData](ctx, c.Client, c.Secret, basePath+"/events-catalog", c.BaseURL)
}
