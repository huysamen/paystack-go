package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FetchBankProvidersResponse represents the response from fetching bank providers
type FetchBankProvidersResponse struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    []BankProvider `json:"data"`
}

// FetchBankProviders gets available bank providers for a dedicated virtual account
func (c *Client) FetchBankProviders(ctx context.Context) (*types.Response[[]BankProvider], error) {
	return net.Get[[]BankProvider](ctx, c.Client, c.Secret, basePath+"/available_providers", c.BaseURL)
}
