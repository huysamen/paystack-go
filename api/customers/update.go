package customers

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CustomerUpdateRequest struct {
	FirstName *string        `json:"first_name,omitempty"`
	LastName  *string        `json:"last_name,omitempty"`
	Phone     *string        `json:"phone,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

func (c *Client) Update(ctx context.Context, code string, req *CustomerUpdateRequest) (*types.Response[Customer], error) {
	if code == "" {
		return nil, errors.New("customer code is required")
	}

	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	path := fmt.Sprintf("%s/%s", customerBasePath, code)

	return net.Put[CustomerUpdateRequest, Customer](
		ctx,
		c.client,
		c.secret,
		path,
		req,
		c.baseURL,
	)
}
