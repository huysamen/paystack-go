package webhookendpoints

import "github.com/huysamen/paystack-go/types"

// WriteResponseData is an endpoint as Create and Update answer it: its events
// as a flat list.
type WriteResponseData struct {
	types.WebhookEndpoint
	Events []string `json:"events"`
}
