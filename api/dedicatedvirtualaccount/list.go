package dedicatedvirtualaccount

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// ListDedicatedVirtualAccountsRequest represents the request to list dedicated virtual accounts
type ListDedicatedVirtualAccountsRequest struct {
	Active       *bool  `json:"active,omitempty"`
	Currency     string `json:"currency,omitempty"`
	ProviderSlug string `json:"provider_slug,omitempty"`
	BankID       string `json:"bank_id,omitempty"`
	Customer     string `json:"customer,omitempty"`
}

type ListDedicatedVirtualAccountsResponse = types.Response[[]types.DedicatedVirtualAccount]

// ListDedicatedVirtualAccountsBuilder builds requests for listing dedicated virtual accounts
type ListDedicatedVirtualAccountsBuilder struct {
	request *ListDedicatedVirtualAccountsRequest
}

// NewListDedicatedVirtualAccountsBuilder creates a new builder for listing dedicated virtual accounts
func NewListDedicatedVirtualAccountsBuilder() *ListDedicatedVirtualAccountsBuilder {
	return &ListDedicatedVirtualAccountsBuilder{
		request: &ListDedicatedVirtualAccountsRequest{},
	}
}

// Active sets the active filter for listing dedicated virtual accounts
func (b *ListDedicatedVirtualAccountsBuilder) Active(active bool) *ListDedicatedVirtualAccountsBuilder {
	b.request.Active = &active
	return b
}

// Currency sets the currency filter for listing dedicated virtual accounts
func (b *ListDedicatedVirtualAccountsBuilder) Currency(currency string) *ListDedicatedVirtualAccountsBuilder {
	b.request.Currency = currency
	return b
}

// ProviderSlug sets the provider slug filter for listing dedicated virtual accounts
func (b *ListDedicatedVirtualAccountsBuilder) ProviderSlug(providerSlug string) *ListDedicatedVirtualAccountsBuilder {
	b.request.ProviderSlug = providerSlug
	return b
}

// BankID sets the bank ID filter for listing dedicated virtual accounts
func (b *ListDedicatedVirtualAccountsBuilder) BankID(bankID string) *ListDedicatedVirtualAccountsBuilder {
	b.request.BankID = bankID
	return b
}

// Customer sets the customer filter for listing dedicated virtual accounts
func (b *ListDedicatedVirtualAccountsBuilder) Customer(customer string) *ListDedicatedVirtualAccountsBuilder {
	b.request.Customer = customer
	return b
}

// Build returns the built request
func (b *ListDedicatedVirtualAccountsBuilder) Build() *ListDedicatedVirtualAccountsRequest {
	return b.request
}

// List retrieves dedicated virtual accounts available on your integration
func (c *Client) List(ctx context.Context, builder *ListDedicatedVirtualAccountsBuilder) (*ListDedicatedVirtualAccountsResponse, error) {
	endpoint := basePath

	if builder != nil {
		req := builder.Build()
		params := url.Values{}
		if req.Active != nil {
			params.Set("active", strconv.FormatBool(*req.Active))
		}
		if req.Currency != "" {
			params.Set("currency", req.Currency)
		}
		if req.ProviderSlug != "" {
			params.Set("provider_slug", req.ProviderSlug)
		}
		if req.BankID != "" {
			params.Set("bank_id", req.BankID)
		}
		if req.Customer != "" {
			params.Set("customer", req.Customer)
		}

		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	return net.Get[[]types.DedicatedVirtualAccount](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
