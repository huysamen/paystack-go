package transferrecipients

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// TransferRecipientFetchResponse represents the response from fetching a transfer recipient
type TransferRecipientFetchResponse = types.Response[types.TransferRecipient]

// Fetch retrieves a specific transfer recipient by ID or code
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*TransferRecipientFetchResponse, error) {
	return net.Get[types.TransferRecipient](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), "", c.BaseURL)
}
