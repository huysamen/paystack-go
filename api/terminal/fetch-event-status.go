package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// FetchEventStatus fetches the status of a terminal event
func (c *Client) FetchEventStatus(ctx context.Context, terminalID, eventID string) (*TerminalEventStatusResponse, error) {
	endpoint := fmt.Sprintf("%s/%s/events/%s", basePath, terminalID, eventID)
	return net.Get[TerminalEventStatus](ctx, c.Client, c.Secret, endpoint, "", c.BaseURL)
}
