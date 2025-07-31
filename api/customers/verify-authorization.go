package customers

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type AuthorizationVerifyResponse struct {
	AuthorizationCode string            `json:"authorization_code"`
	Channel           types.Channel     `json:"channel"`
	Bank              string            `json:"bank"`
	Active            bool              `json:"active"`
	Customer          CustomerReference `json:"customer"`
}

type CustomerReference struct {
	Code  string `json:"code"`
	Email string `json:"email"`
}

func (c *Client) VerifyAuthorization(ctx context.Context, reference string) (*types.Response[AuthorizationVerifyResponse], error) {
	if reference == "" {
		return nil, errors.New("reference is required")
	}

	path := fmt.Sprintf("%s/authorization/verify/%s", customerBasePath, reference)

	return net.Get[AuthorizationVerifyResponse](
		ctx,
		c.client,
		c.secret,
		path,
		c.baseURL,
	)
}
