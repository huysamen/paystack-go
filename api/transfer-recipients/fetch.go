package transfer_recipients

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Fetch retrieves a specific transfer recipient by ID or code
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*TransferRecipientFetchResponse, error) {
	if idOrCode == "" {
		return nil, fmt.Errorf("id_or_code is required")
	}

	endpoint := fmt.Sprintf("%s/%s", transferRecipientBasePath, idOrCode)

	resp, err := net.Get[TransferRecipientFetchResponse](ctx, c.client, c.secret, endpoint, c.baseURL)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
