package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchTerminalStatusResponseData = types.TerminalPresenceStatus
type FetchTerminalStatusResponse = types.Response[FetchTerminalStatusResponseData]

func (c *Client) FetchTerminalStatus(ctx context.Context, terminalID string) (*FetchTerminalStatusResponse, error) {
	endpoint := fmt.Sprintf("%s/%s/presence", basePath, terminalID)

	return net.Get[FetchTerminalStatusResponseData](ctx, c.Client, c.Secret, endpoint, "", c.BaseURL)
}
