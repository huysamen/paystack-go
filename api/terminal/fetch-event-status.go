package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// FetchEventStatus checks the status of an event sent to terminal
func (c *Client) FetchEventStatus(ctx context.Context, terminalID, eventID string) (*TerminalEventStatusResponse, error) {
	if err := validateTerminalID(terminalID); err != nil {
		return nil, err
	}
	if err := validateEventID(eventID); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s/event/%s", terminalBasePath, terminalID, eventID)
	resp, err := net.Get[TerminalEventStatusResponse](ctx, c.client, c.secret, endpoint, c.baseURL)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
