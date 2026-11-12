package webhookendpoints

import (
	"context"
	"net/url"
	"slices"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FetchResponseData is an endpoint with its events grouped by category.
type FetchResponseData struct {
	types.WebhookEndpoint
	Events map[string][]string `json:"events"`
}

// EventNames returns the endpoint's events across all categories, sorted and
// without duplicates.
func (d FetchResponseData) EventNames() []string {
	var names []string
	for _, events := range d.Events {
		names = append(names, events...)
	}

	slices.Sort(names)

	return slices.Compact(names)
}

type FetchResponse = types.Response[FetchResponseData]

// Fetch gets one endpoint, by its webhook_url_code, with its events.
func (c *Client) Fetch(ctx context.Context, code string) (*FetchResponse, error) {
	return net.Get[FetchResponseData](ctx, c.Client, c.Secret, endpointPath(code), c.BaseURL)
}

func endpointPath(code string) string {
	return basePath + "/endpoints/" + url.PathEscape(code)
}
