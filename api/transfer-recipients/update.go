package transfer_recipients

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Update updates an existing transfer recipient
func (c *Client) Update(ctx context.Context, idOrCode string, req *TransferRecipientUpdateRequest) (*TransferRecipientUpdateResponse, error) {
	if idOrCode == "" {
		return nil, fmt.Errorf("id_or_code is required")
	}

	if err := validateUpdateRequest(req); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s", transferRecipientBasePath, idOrCode)

	resp, err := net.Put[TransferRecipientUpdateRequest, TransferRecipientUpdateResponse](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
