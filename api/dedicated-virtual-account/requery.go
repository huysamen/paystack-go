package dedicatedvirtualaccount

import (
	"context"
	"net/url"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// RequeryDedicatedAccountRequest represents the request to requery a dedicated account
type RequeryDedicatedAccountRequest struct {
	AccountNumber string `json:"account_number"`
	ProviderSlug  string `json:"provider_slug"`
	Date          string `json:"date,omitempty"`
}

// RequeryDedicatedAccountResponse represents the response from requerying a dedicated account
type RequeryDedicatedAccountResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// RequeryDedicatedAccountBuilder builds requests for requerying dedicated accounts
type RequeryDedicatedAccountBuilder struct {
	request *RequeryDedicatedAccountRequest
}

// NewRequeryDedicatedAccountBuilder creates a new builder for requerying dedicated accounts
func NewRequeryDedicatedAccountBuilder() *RequeryDedicatedAccountBuilder {
	return &RequeryDedicatedAccountBuilder{
		request: &RequeryDedicatedAccountRequest{},
	}
}

// AccountNumber sets the account number for requerying the dedicated account
func (b *RequeryDedicatedAccountBuilder) AccountNumber(accountNumber string) *RequeryDedicatedAccountBuilder {
	b.request.AccountNumber = accountNumber
	return b
}

// ProviderSlug sets the provider slug for requerying the dedicated account
func (b *RequeryDedicatedAccountBuilder) ProviderSlug(providerSlug string) *RequeryDedicatedAccountBuilder {
	b.request.ProviderSlug = providerSlug
	return b
}

// Date sets the date for requerying the dedicated account
func (b *RequeryDedicatedAccountBuilder) Date(date string) *RequeryDedicatedAccountBuilder {
	b.request.Date = date
	return b
}

// Build returns the built request
func (b *RequeryDedicatedAccountBuilder) Build() *RequeryDedicatedAccountRequest {
	return b.request
}

// Requery requerying dedicated virtual account for new transactions
func (c *Client) Requery(ctx context.Context, builder *RequeryDedicatedAccountBuilder) (*types.Response[any], error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	params := url.Values{}
	params.Set("account_number", req.AccountNumber)
	params.Set("provider_slug", req.ProviderSlug)
	if req.Date != "" {
		params.Set("date", req.Date)
	}

	endpoint := dedicatedVirtualAccountBasePath + "/requery?" + params.Encode()
	resp, err := net.Get[any](
		ctx, c.client, c.secret, endpoint, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
