package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// SplitTransaction splits a dedicated virtual account transaction with one or more accounts
func (c *Client) SplitTransaction(ctx context.Context, req *SplitDedicatedAccountTransactionRequest) (*DedicatedVirtualAccount, error) {
	if err := validateSplitRequest(req); err != nil {
		return nil, err
	}

	endpoint := dedicatedVirtualAccountBasePath + "/split"
	resp, err := net.Post[SplitDedicatedAccountTransactionRequest, DedicatedVirtualAccount](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
