package types

import "github.com/huysamen/paystack-go/types/data"

// WebhookEndpoint is a webhook endpoint as Paystack lists it: everything but
// its events (Fetch, Create and Update answer those beside it).
type WebhookEndpoint struct {
	WebhookURLCode data.String `json:"webhook_url_code"`
	Name           data.String `json:"name"`
	WebhookURL     data.String `json:"webhook_url"`
	Enabled        data.Bool   `json:"enabled"`
	EventCount     data.Int    `json:"event_count"`
}

// WebhookCatalogEvent is one event a webhook endpoint can subscribe to.
type WebhookCatalogEvent struct {
	EventName   data.String     `json:"event_name"`
	Description data.NullString `json:"description"`
}
