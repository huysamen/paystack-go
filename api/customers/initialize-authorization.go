package customers

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Request and Response types
type AuthorizationInitializeRequest struct {
	Email       string   `json:"email"`
	Channel     string   `json:"channel"` // "direct_debit" is the only supported option
	CallbackURL *string  `json:"callback_url,omitempty"`
	Account     *Account `json:"account,omitempty"`
	Address     *Address `json:"address,omitempty"`
}

type AuthorizationInitializeResponse struct {
	RedirectURL string `json:"redirect_url"`
	AccessCode  string `json:"access_code"`
	Reference   string `json:"reference"`
}

type InitializeAuthorizationResponse = types.Response[AuthorizationInitializeResponse]

// Builder for AuthorizationInitializeRequest
type AuthorizationInitializeRequestBuilder struct {
	email       string
	channel     string
	callbackURL *string
	account     *Account
	address     *Address
}

// NewInitializeAuthorizationRequest creates a new builder for authorization initialization
func NewInitializeAuthorizationRequest(email, channel string) *AuthorizationInitializeRequestBuilder {
	return &AuthorizationInitializeRequestBuilder{
		email:   email,
		channel: channel,
	}
}

// CallbackURL sets the callback URL
func (b *AuthorizationInitializeRequestBuilder) CallbackURL(callbackURL string) *AuthorizationInitializeRequestBuilder {
	b.callbackURL = &callbackURL
	return b
}

// Account sets the account details
func (b *AuthorizationInitializeRequestBuilder) Account(number, bankCode string) *AuthorizationInitializeRequestBuilder {
	b.account = &Account{
		Number:   number,
		BankCode: bankCode,
	}
	return b
}

// Address sets the address details
func (b *AuthorizationInitializeRequestBuilder) Address(street, city, state string) *AuthorizationInitializeRequestBuilder {
	b.address = &Address{
		Street: street,
		City:   city,
		State:  state,
	}
	return b
}

// Build creates the AuthorizationInitializeRequest
func (b *AuthorizationInitializeRequestBuilder) Build() *AuthorizationInitializeRequest {
	return &AuthorizationInitializeRequest{
		Email:       b.email,
		Channel:     b.channel,
		CallbackURL: b.callbackURL,
		Account:     b.account,
		Address:     b.address,
	}
}

// InitializeAuthorization initializes authorization for a customer
func (c *Client) InitializeAuthorization(ctx context.Context, builder *AuthorizationInitializeRequestBuilder) (*InitializeAuthorizationResponse, error) {
	return net.Post[AuthorizationInitializeRequest, AuthorizationInitializeResponse](ctx, c.Client, c.Secret, basePath+"/authorization/initialize", builder.Build(), c.BaseURL)
}
