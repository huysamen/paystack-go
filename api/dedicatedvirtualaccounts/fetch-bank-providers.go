package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchBankProvidersResponseData = []types.BankProvider
type FetchBankProvidersResponse = types.Response[FetchBankProvidersResponseData]

func (c *Client) FetchBankProviders(ctx context.Context) (*FetchBankProvidersResponse, error) {
	return net.Get[FetchBankProvidersResponseData](ctx, c.Client, c.Secret, basePath+"/available_providers", c.BaseURL)
}
