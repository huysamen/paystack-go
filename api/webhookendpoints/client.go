// Package webhookendpoints covers Paystack's Webhook Endpoints API: webhook
// URLs, each with its own events, managed with the integration's secret key.
//
// The API is in pilot: an integration outside it gets 403 on every call.
package webhookendpoints

import (
	"github.com/huysamen/paystack-go/api"
)

const basePath = "/integration/webhooks"

type Client api.API
