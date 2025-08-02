package terminal

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// TerminalPresence represents the data returned from fetching terminal presence
type TerminalPresence struct {
	Online    bool `json:"online"`    // Whether terminal is online
	Available bool `json:"available"` // Whether terminal is available for events
}

// FetchTerminalStatus fetches the status of a terminal
func (c *Client) FetchTerminalStatus(ctx context.Context, terminalID string) (*types.Response[TerminalPresence], error) {
	if terminalID == "" {
		return nil, errors.New("terminal ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/presence", terminalBasePath, terminalID)
	return net.Get[TerminalPresence](ctx, c.client, c.secret, endpoint, "", c.baseURL)
}
