package miscellaneous

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// ListCountries retrieves a list of countries supported by Paystack
func (c *Client) ListCountries(ctx context.Context) (*CountryListResponse, error) {
	resp, err := net.Get[CountryListResponse](ctx, c.client, c.secret, countryBasePath, c.baseURL)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
