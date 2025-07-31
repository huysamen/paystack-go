package customers

import (
	"context"
	"errors"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type AuthorizationInitializeRequest struct {
	Email       string   `json:"email"`
	Channel     string   `json:"channel"` // "direct_debit" is the only supported option
	CallbackURL *string  `json:"callback_url,omitempty"`
	Account     *Account `json:"account,omitempty"`
	Address     *Address `json:"address,omitempty"`
}

func (r *AuthorizationInitializeRequest) Validate() error {
	if r.Email == "" {
		return errors.New("email is required")
	}
	if r.Channel == "" {
		return errors.New("channel is required")
	}
	if r.Channel != "direct_debit" {
		return errors.New("only 'direct_debit' channel is supported")
	}
	return nil
}

type AuthorizationInitializeResponse struct {
	RedirectURL string `json:"redirect_url"`
	AccessCode  string `json:"access_code"`
	Reference   string `json:"reference"`
}

func (c *Client) InitializeAuthorization(ctx context.Context, req *AuthorizationInitializeRequest) (*types.Response[AuthorizationInitializeResponse], error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	path := customerBasePath + "/authorization/initialize"

	return net.Post[AuthorizationInitializeRequest, AuthorizationInitializeResponse](
		ctx,
		c.client,
		c.secret,
		path,
		req,
		c.baseURL,
	)
}
