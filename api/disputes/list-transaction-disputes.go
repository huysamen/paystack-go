package disputes

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// ListTransactionDisputes retrieves disputes for a particular transaction
func (c *Client) ListTransactionDisputes(ctx context.Context, transactionID string) (*TransactionDisputeResponse, error) {
	if transactionID == "" {
		return nil, fmt.Errorf("transaction ID is required")
	}

	url := c.baseURL + disputesBasePath + "/transaction/" + transactionID
	return net.Get[TransactionDisputeData](ctx, c.client, c.secret, url)
}
