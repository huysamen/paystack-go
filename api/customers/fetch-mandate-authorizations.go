package customers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchMandateAuthorizationsResponseData = []types.MandateAuthorization
type FetchMandateAuthorizationsResponse = types.Response[FetchMandateAuthorizationsResponseData]

func (c *Client) FetchMandateAuthorizations(ctx context.Context, customerID string) (*FetchMandateAuthorizationsResponse, error) {
	path := fmt.Sprintf("%s/%s/directdebit-mandate-authorizations", basePath, customerID)

	return net.Get[FetchMandateAuthorizationsResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
