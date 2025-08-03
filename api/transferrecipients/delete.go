package transferrecipients

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// TransferRecipientDeleteResponse represents the response from deleting a transfer recipient
type TransferRecipientDeleteResponse = types.Response[any]

// Delete deletes a transfer recipient (sets it to inactive)
func (c *Client) Delete(ctx context.Context, idOrCode string) (*TransferRecipientDeleteResponse, error) {
	return net.Delete[any](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), c.BaseURL)
}
