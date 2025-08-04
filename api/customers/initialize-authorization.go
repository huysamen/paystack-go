package customers

import (
	"context"

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

type AuthorizationInitializeRequestBuilder struct {
	req *AuthorizationInitializeRequest
}

func NewInitializeAuthorizationRequest(email, channel string) *AuthorizationInitializeRequestBuilder {
	return &AuthorizationInitializeRequestBuilder{
		req: &AuthorizationInitializeRequest{
			Email:   email,
			Channel: channel,
		},
	}
}

func (b *AuthorizationInitializeRequestBuilder) CallbackURL(callbackURL string) *AuthorizationInitializeRequestBuilder {
	b.req.CallbackURL = &callbackURL

	return b
}

func (b *AuthorizationInitializeRequestBuilder) Account(number, bankCode string) *AuthorizationInitializeRequestBuilder {
	b.req.Account = &Account{
		Number:   number,
		BankCode: bankCode,
	}

	return b
}

func (b *AuthorizationInitializeRequestBuilder) Address(street, city, state string) *AuthorizationInitializeRequestBuilder {
	b.req.Address = &Address{
		Street: street,
		City:   city,
		State:  state,
	}

	return b
}

func (b *AuthorizationInitializeRequestBuilder) Build() *AuthorizationInitializeRequest {
	return b.req
}

type InitializeAuthorizationResponseData struct {
	RedirectURL string `json:"redirect_url"`
	AccessCode  string `json:"access_code"`
	Reference   string `json:"reference"`
}

type InitializeAuthorizationResponse = types.Response[InitializeAuthorizationResponseData]

func (c *Client) InitializeAuthorization(ctx context.Context, builder *AuthorizationInitializeRequestBuilder) (*InitializeAuthorizationResponse, error) {
	return net.Post[AuthorizationInitializeRequest, InitializeAuthorizationResponseData](ctx, c.Client, c.Secret, basePath+"/authorization/initialize", builder.Build(), c.BaseURL)
}
