package dedicatedvirtualaccount

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// List retrieves dedicated virtual accounts available on your integration
func (c *Client) List(ctx context.Context, req *ListDedicatedVirtualAccountsRequest) (*types.Response[[]DedicatedVirtualAccount], error) {
	endpoint := dedicatedVirtualAccountBasePath

	if req != nil {
		params := url.Values{}
		if req.Active != nil {
			params.Set("active", strconv.FormatBool(*req.Active))
		}
		if req.Currency != "" {
			params.Set("currency", req.Currency)
		}
		if req.ProviderSlug != "" {
			params.Set("provider_slug", req.ProviderSlug)
		}
		if req.BankID != "" {
			params.Set("bank_id", req.BankID)
		}
		if req.Customer != "" {
			params.Set("customer", req.Customer)
		}

		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	resp, err := net.Get[[]DedicatedVirtualAccount](
		ctx, c.client, c.secret, endpoint, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
