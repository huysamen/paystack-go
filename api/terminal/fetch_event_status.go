package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchEventStatusResponseData = types.TerminalEventStatus
type FetchEventStatusResponse = types.Response[FetchEventStatusResponseData]

func (c *Client) FetchEventStatus(ctx context.Context, terminalID, eventID string) (*FetchEventStatusResponse, error) {
	endpoint := fmt.Sprintf("%s/%s/events/%s", basePath, terminalID, eventID)

	return net.Get[FetchEventStatusResponseData](ctx, c.Client, c.Secret, endpoint, "", c.BaseURL)
}
