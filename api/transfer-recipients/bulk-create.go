package transfer_recipients

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// BulkCreate creates multiple transfer recipients in a single request
func (c *Client) BulkCreate(ctx context.Context, req *BulkCreateTransferRecipientRequest) (*BulkCreateTransferRecipientResponse, error) {
	if err := validateBulkCreateRequest(req); err != nil {
		return nil, err
	}

	resp, err := net.Post[BulkCreateTransferRecipientRequest, BulkCreateTransferRecipientResponse](
		ctx, c.client, c.secret, transferRecipientBasePath+"/bulk", req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
