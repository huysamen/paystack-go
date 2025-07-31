package transaction_splits

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// RemoveSubaccount removes a subaccount from a transaction split
func (c *Client) RemoveSubaccount(ctx context.Context, id string, req *TransactionSplitSubaccountRemoveRequest) (*TransactionSplitSubaccountRemoveResponse, error) {
	if err := validateTransactionSplitID(id); err != nil {
		return nil, err
	}
	if err := validateSubaccountRemoveRequest(req); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s/subaccount/remove", transactionSplitBasePath, id)
	resp, err := net.Post[TransactionSplitSubaccountRemoveRequest, TransactionSplitSubaccountRemoveResponse](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
