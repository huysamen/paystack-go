package dedicatedvirtualaccount

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ListDedicatedVirtualAccountsRequest struct {
	Active       *bool  `json:"active,omitempty"`
	Currency     string `json:"currency,omitempty"`
	ProviderSlug string `json:"provider_slug,omitempty"`
	BankID       string `json:"bank_id,omitempty"`
	Customer     string `json:"customer,omitempty"`
}

type ListDedicatedVirtualAccountsBuilder struct {
	request *ListDedicatedVirtualAccountsRequest
}

func NewListDedicatedVirtualAccountsBuilder() *ListDedicatedVirtualAccountsBuilder {
	return &ListDedicatedVirtualAccountsBuilder{
		request: &ListDedicatedVirtualAccountsRequest{},
	}
}

func (b *ListDedicatedVirtualAccountsBuilder) Active(active bool) *ListDedicatedVirtualAccountsBuilder {
	b.request.Active = &active

	return b
}

func (b *ListDedicatedVirtualAccountsBuilder) Currency(currency string) *ListDedicatedVirtualAccountsBuilder {
	b.request.Currency = currency

	return b
}

func (b *ListDedicatedVirtualAccountsBuilder) ProviderSlug(providerSlug string) *ListDedicatedVirtualAccountsBuilder {
	b.request.ProviderSlug = providerSlug

	return b
}

func (b *ListDedicatedVirtualAccountsBuilder) BankID(bankID string) *ListDedicatedVirtualAccountsBuilder {
	b.request.BankID = bankID

	return b
}

func (b *ListDedicatedVirtualAccountsBuilder) Customer(customer string) *ListDedicatedVirtualAccountsBuilder {
	b.request.Customer = customer

	return b
}

func (b *ListDedicatedVirtualAccountsBuilder) Build() *ListDedicatedVirtualAccountsRequest {
	return b.request
}

type ListDedicatedVirtualAccountsResponse = types.Response[[]types.DedicatedVirtualAccount]

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
