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

// Builder for AuthorizationInitializeRequest
type AuthorizationInitializeRequestBuilder struct {
	req *AuthorizationInitializeRequest
}

// NewInitializeAuthorizationRequest creates a new builder for authorization initialization
func NewInitializeAuthorizationRequest(email, channel string) *AuthorizationInitializeRequestBuilder {
	return &AuthorizationInitializeRequestBuilder{
		req: &AuthorizationInitializeRequest{
			Email:   email,
			Channel: channel,
		},
	}
}

// CallbackURL sets the callback URL
func (b *AuthorizationInitializeRequestBuilder) CallbackURL(callbackURL string) *AuthorizationInitializeRequestBuilder {
	b.req.CallbackURL = &callbackURL

	return b
}

// Account sets the account details
func (b *AuthorizationInitializeRequestBuilder) Account(number, bankCode string) *AuthorizationInitializeRequestBuilder {
	b.req.Account = &Account{
		Number:   number,
		BankCode: bankCode,
	}

	return b
}

// Address sets the address details
func (b *AuthorizationInitializeRequestBuilder) Address(street, city, state string) *AuthorizationInitializeRequestBuilder {
	b.req.Address = &Address{
		Street: street,
		City:   city,
		State:  state,
	}

	return b
}

// Build creates the AuthorizationInitializeRequest
func (b *AuthorizationInitializeRequestBuilder) Build() *AuthorizationInitializeRequest {
	return b.req
}

type InitializeAuthorizationResponseData struct {
	RedirectURL string `json:"redirect_url"`
	AccessCode  string `json:"access_code"`
	Reference   string `json:"reference"`
}

// InitializeAuthorizationResponse represents the response for initializing authorization
type InitializeAuthorizationResponse = types.Response[InitializeAuthorizationResponseData]

// InitializeAuthorization initializes authorization for a customer
func (c *Client) InitializeAuthorization(ctx context.Context, builder *AuthorizationInitializeRequestBuilder) (*InitializeAuthorizationResponse, error) {
	return net.Post[AuthorizationInitializeRequest, InitializeAuthorizationResponseData](ctx, c.Client, c.Secret, basePath+"/authorization/initialize", builder.Build(), c.BaseURL)
}
