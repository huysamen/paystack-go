package customers

import (
	"context"
	"errors"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CustomerCreateRequest struct {
	Email     string         `json:"email"`
	FirstName *string        `json:"first_name,omitempty"`
	LastName  *string        `json:"last_name,omitempty"`
	Phone     *string        `json:"phone,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

func (r *CustomerCreateRequest) Validate() error {
	if r.Email == "" {
		return errors.New("email is required")
	}
	return nil
}

func (c *Client) Create(ctx context.Context, req *CustomerCreateRequest) (*types.Response[Customer], error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	return net.Post[CustomerCreateRequest, Customer](
		ctx,
		c.client,
		c.secret,
		customerBasePath,
		req,
		c.baseURL,
	)
}
