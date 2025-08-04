package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchBankProvidersResponse = types.Response[[]types.BankProvider]

func (c *Client) FetchBankProviders(ctx context.Context) (*FetchBankProvidersResponse, error) {
	return net.Get[[]types.BankProvider](ctx, c.Client, c.Secret, basePath+"/available_providers", c.BaseURL)
}
