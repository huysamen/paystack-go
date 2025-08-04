package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TerminalEventStatusResponse = types.Response[types.TerminalEventStatus]

func (c *Client) FetchEventStatus(ctx context.Context, terminalID, eventID string) (*TerminalEventStatusResponse, error) {
	endpoint := fmt.Sprintf("%s/%s/events/%s", basePath, terminalID, eventID)

	return net.Get[types.TerminalEventStatus](ctx, c.Client, c.Secret, endpoint, "", c.BaseURL)
}
