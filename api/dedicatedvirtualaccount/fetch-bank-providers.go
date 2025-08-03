package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchBankProvidersResponse = types.Response[[]BankProvider]

// FetchBankProviders gets available bank providers for a dedicated virtual account
func (c *Client) FetchBankProviders(ctx context.Context) (*FetchBankProvidersResponse, error) {
	return net.Get[[]BankProvider](ctx, c.Client, c.Secret, basePath+"/available_providers", c.BaseURL)
}
