package transfer_recipients

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Delete deletes a transfer recipient (sets it to inactive)
func (c *Client) Delete(ctx context.Context, idOrCode string) (*TransferRecipientDeleteResponse, error) {
	if idOrCode == "" {
		return nil, fmt.Errorf("id_or_code is required")
	}

	endpoint := fmt.Sprintf("%s/%s", transferRecipientBasePath, idOrCode)

	resp, err := net.Delete[TransferRecipientDeleteResponse](ctx, c.client, c.secret, endpoint, c.baseURL)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
