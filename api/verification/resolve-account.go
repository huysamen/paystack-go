package verification

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// ResolveAccount resolves a bank account number to get account details
func (c *Client) ResolveAccount(ctx context.Context, builder *AccountResolveRequestBuilder) (*types.Response[AccountResolution], error) {
	req := builder.Build()
	endpoint := fmt.Sprintf("%s?account_number=%s&bank_code=%s", accountResolveBasePath, req.AccountNumber, req.BankCode)
	return net.Get[AccountResolution](ctx, c.Client, c.Secret, endpoint, "", c.BaseURL)
}
