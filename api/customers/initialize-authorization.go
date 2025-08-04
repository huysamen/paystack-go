package customers

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type initializeAuthorizationRequest struct {
	Email       string   `json:"email"`
	Channel     string   `json:"channel"` // "direct_debit" is the only supported option
	CallbackURL *string  `json:"callback_url,omitempty"`
	Account     *Account `json:"account,omitempty"`
	Address     *Address `json:"address,omitempty"`
}

type InitializeAuthorizationRequestBuilder struct {
	req *initializeAuthorizationRequest
}

func NewInitializeAuthorizationRequestBuilder(email, channel string) *InitializeAuthorizationRequestBuilder {
	return &InitializeAuthorizationRequestBuilder{
		req: &initializeAuthorizationRequest{
			Email:   email,
			Channel: channel,
		},
	}
}

func (b *InitializeAuthorizationRequestBuilder) CallbackURL(callbackURL string) *InitializeAuthorizationRequestBuilder {
	b.req.CallbackURL = &callbackURL

	return b
}

func (b *InitializeAuthorizationRequestBuilder) Account(number, bankCode string) *InitializeAuthorizationRequestBuilder {
	b.req.Account = &Account{
		Number:   number,
		BankCode: bankCode,
	}

	return b
}

func (b *InitializeAuthorizationRequestBuilder) Address(street, city, state string) *InitializeAuthorizationRequestBuilder {
	b.req.Address = &Address{
		Street: street,
		City:   city,
		State:  state,
	}

	return b
}

func (b *InitializeAuthorizationRequestBuilder) Build() *initializeAuthorizationRequest {
	return b.req
}

type InitializeAuthorizationResponseData struct {
	RedirectURL string `json:"redirect_url"`
	AccessCode  string `json:"access_code"`
	Reference   string `json:"reference"`
}

type InitializeAuthorizationResponse = types.Response[InitializeAuthorizationResponseData]

func (c *Client) InitializeAuthorization(ctx context.Context, builder InitializeAuthorizationRequestBuilder) (*InitializeAuthorizationResponse, error) {
	return net.Post[initializeAuthorizationRequest, InitializeAuthorizationResponseData](ctx, c.Client, c.Secret, basePath+"/authorization/initialize", builder.Build(), c.BaseURL)
}
