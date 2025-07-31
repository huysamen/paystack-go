package dedicatedvirtualaccount

import (
	"context"
	"net/url"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Requery requerying dedicated virtual account for new transactions
func (c *Client) Requery(ctx context.Context, req *RequeryDedicatedAccountRequest) (*types.Response[interface{}], error) {
	if err := validateRequeryRequest(req); err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Set("account_number", req.AccountNumber)
	params.Set("provider_slug", req.ProviderSlug)
	if req.Date != "" {
		params.Set("date", req.Date)
	}

	endpoint := dedicatedVirtualAccountBasePath + "/requery?" + params.Encode()
	resp, err := net.Get[interface{}](
		ctx, c.client, c.secret, endpoint, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
