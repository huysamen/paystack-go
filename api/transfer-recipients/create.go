package transfer_recipients

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// Create creates a new transfer recipient
func (c *Client) Create(ctx context.Context, req *TransferRecipientCreateRequest) (*TransferRecipientCreateResponse, error) {
	if err := validateCreateRequest(req); err != nil {
		return nil, err
	}

	resp, err := net.Post[TransferRecipientCreateRequest, TransferRecipientCreateResponse](
		ctx, c.client, c.secret, transferRecipientBasePath, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
