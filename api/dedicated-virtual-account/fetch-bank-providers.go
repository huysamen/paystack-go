package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FetchBankProviders gets available bank providers for a dedicated virtual account
func (c *Client) FetchBankProviders(ctx context.Context) (*types.Response[[]BankProvider], error) {
	endpoint := dedicatedVirtualAccountBasePath + "/available_providers"
	resp, err := net.Get[[]BankProvider](
		ctx, c.client, c.secret, endpoint, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
