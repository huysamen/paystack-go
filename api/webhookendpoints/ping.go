package webhookendpoints

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
)

// PingResponseData is what the endpoint's URL answered a signed "ping" event
// with. OK is true only for a 2xx; HTTPStatus is null when Paystack could not
// reach the URL.
type PingResponseData struct {
	OK         data.Bool    `json:"ok"`
	HTTPStatus data.NullInt `json:"http_status"`
	LatencyMS  data.Int     `json:"latency_ms"`
}

type PingResponse = types.Response[PingResponseData]

// Ping sends a test event to the endpoint. Paystack allows 10 pings every 60
// seconds in each mode and answers 429 beyond that.
func (c *Client) Ping(ctx context.Context, code string) (*PingResponse, error) {
	return net.Post[any, PingResponseData](ctx, c.Client, c.Secret, endpointPath(code)+"/ping", nil, c.BaseURL)
}
