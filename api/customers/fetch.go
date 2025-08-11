package customers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CustomerWithRelations struct {
	types.Customer
	Subscriptions  []types.Subscription  `json:"subscriptions"`
	Authorizations []types.Authorization `json:"authorizations"`
	Transactions   []types.Transaction   `json:"transactions"`
}

type FetchCustomerResponseData = CustomerWithRelations
type FetchCustomerResponse = types.Response[FetchCustomerResponseData]

func (c *Client) Fetch(ctx context.Context, emailOrCode string) (*FetchCustomerResponse, error) {
	path := fmt.Sprintf("%s/%s", basePath, emailOrCode)

	return net.Get[FetchCustomerResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
