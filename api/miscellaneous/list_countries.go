package miscellaneous

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ListCountriesResponseData = []types.Country
type ListCountriesResponse = types.Response[ListCountriesResponseData]

func (c *Client) ListCountries(ctx context.Context) (*ListCountriesResponse, error) {
	return net.Get[ListCountriesResponseData](ctx, c.Client, c.Secret, countryPath, c.BaseURL)
}
