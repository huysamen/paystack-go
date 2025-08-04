package miscellaneous

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// CountryListResponse represents the response from listing countries
type CountryListResponse = types.Response[[]types.Country]

// ListCountries retrieves a list of countries supported by Paystack
func (c *Client) ListCountries(ctx context.Context) (*CountryListResponse, error) {
	return net.Get[[]types.Country](ctx, c.Client, c.Secret, countryPath, c.BaseURL)
}
