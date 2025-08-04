package miscellaneous

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CountryListResponse = types.Response[[]types.Country]

func (c *Client) ListCountries(ctx context.Context) (*CountryListResponse, error) {
	return net.Get[[]types.Country](ctx, c.Client, c.Secret, countryPath, c.BaseURL)
}
