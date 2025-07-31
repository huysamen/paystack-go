package verification

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// ResolveAccount resolves a bank account number to get account details
func (c *Client) ResolveAccount(ctx context.Context, req *AccountResolveRequest) (*AccountResolveResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AccountNumber == "" {
		return nil, fmt.Errorf("account_number is required")
	}
	if req.BankCode == "" {
		return nil, fmt.Errorf("bank_code is required")
	}

	endpoint := fmt.Sprintf("%s?account_number=%s&bank_code=%s", accountResolveBasePath, req.AccountNumber, req.BankCode)

	resp, err := net.Get[AccountResolveResponse](ctx, c.client, c.secret, endpoint, c.baseURL)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
