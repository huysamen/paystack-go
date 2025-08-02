package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// EventStatus represents the data returned from fetching event status
type EventStatus struct {
	Delivered bool `json:"delivered"` // Whether event was delivered to terminal
}

// FetchEventStatus fetches the status of a terminal event
func (c *Client) FetchEventStatus(ctx context.Context, terminalID, eventID string) (*types.Response[EventStatus], error) {
	endpoint := fmt.Sprintf("%s/%s/events/%s", basePath, terminalID, eventID)
	return net.Get[EventStatus](ctx, c.Client, c.Secret, endpoint, "", c.BaseURL)
}
