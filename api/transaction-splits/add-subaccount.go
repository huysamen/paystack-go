package transaction_splits

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// AddSubaccount adds or updates a subaccount in a transaction split
func (c *Client) AddSubaccount(ctx context.Context, id string, req *TransactionSplitSubaccountAddRequest) (*TransactionSplitSubaccountAddResponse, error) {
	if err := validateTransactionSplitID(id); err != nil {
		return nil, err
	}
	if err := validateSubaccountAddRequest(req); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s/subaccount/add", transactionSplitBasePath, id)
	resp, err := net.Post[TransactionSplitSubaccountAddRequest, TransactionSplitSubaccountAddResponse](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
