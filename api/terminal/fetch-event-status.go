package terminal

import (
	"context"
	"errors"
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
	if terminalID == "" {
		return nil, errors.New("terminal ID is required")
	}
	if eventID == "" {
		return nil, errors.New("event ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/events/%s", terminalBasePath, terminalID, eventID)
	return net.Get[EventStatus](ctx, c.client, c.secret, endpoint, "", c.baseURL)
}
