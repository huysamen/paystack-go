package dedicatedvirtualaccount

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ListRequest struct {
	Active       *bool  `json:"active,omitempty"`
	Currency     string `json:"currency,omitempty"`
	ProviderSlug string `json:"provider_slug,omitempty"`
	BankID       string `json:"bank_id,omitempty"`
	Customer     string `json:"customer,omitempty"`
}

type ListRequestBuilder struct {
	request *ListRequest
}

func NewListRequestBuilder() *ListRequestBuilder {
	return &ListRequestBuilder{
		request: &ListRequest{},
	}
}

func (b *ListRequestBuilder) Active(active bool) *ListRequestBuilder {
	b.request.Active = &active

	return b
}

func (b *ListRequestBuilder) Currency(currency string) *ListRequestBuilder {
	b.request.Currency = currency

	return b
}

func (b *ListRequestBuilder) ProviderSlug(providerSlug string) *ListRequestBuilder {
	b.request.ProviderSlug = providerSlug

	return b
}

func (b *ListRequestBuilder) BankID(bankID string) *ListRequestBuilder {
	b.request.BankID = bankID

	return b
}

func (b *ListRequestBuilder) Customer(customer string) *ListRequestBuilder {
	b.request.Customer = customer

	return b
}

func (b *ListRequestBuilder) Build() *ListRequest {
	return b.request
}

func (r *ListRequest) toQuery() string {
	params := url.Values{}
	if r.Active != nil {
		params.Set("active", strconv.FormatBool(*r.Active))
	}
	if r.Currency != "" {
		params.Set("currency", r.Currency)
	}
	if r.ProviderSlug != "" {
		params.Set("provider_slug", r.ProviderSlug)
	}
	if r.BankID != "" {
		params.Set("bank_id", r.BankID)
	}
	if r.Customer != "" {
		params.Set("customer", r.Customer)
	}

	return params.Encode()
}

type ListResponseData = []types.DedicatedVirtualAccount
type ListResponse = types.Response[ListResponseData]

func (c *Client) List(ctx context.Context, builder ListRequestBuilder) (*ListResponse, error) {
	req := builder.Build()
	path := basePath

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[ListResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
