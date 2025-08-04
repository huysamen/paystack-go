package dedicatedvirtualaccount

import (
	"context"
	"net/url"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type RequeryRequest struct {
	AccountNumber string `json:"account_number"`
	ProviderSlug  string `json:"provider_slug"`
	Date          string `json:"date,omitempty"`
}

type RequeryRequestBuilder struct {
	request *RequeryRequest
}

func NewRequeryRequestBuilder() *RequeryRequestBuilder {
	return &RequeryRequestBuilder{
		request: &RequeryRequest{},
	}
}

func (b *RequeryRequestBuilder) AccountNumber(accountNumber string) *RequeryRequestBuilder {
	b.request.AccountNumber = accountNumber

	return b
}

func (b *RequeryRequestBuilder) ProviderSlug(providerSlug string) *RequeryRequestBuilder {
	b.request.ProviderSlug = providerSlug

	return b
}

func (b *RequeryRequestBuilder) Date(date string) *RequeryRequestBuilder {
	b.request.Date = date

	return b
}

func (b *RequeryRequestBuilder) Build() *RequeryRequest {
	return b.request
}

func (r *RequeryRequest) toQuery() string {
	params := url.Values{}

	params.Set("account_number", r.AccountNumber)
	params.Set("provider_slug", r.ProviderSlug)

	if r.Date != "" {
		params.Set("date", r.Date)
	}

	return params.Encode()
}

type RequeryResponseData = any
type RequeryResponse = types.Response[RequeryResponseData]

func (c *Client) Requery(ctx context.Context, builder RequeryRequestBuilder) (*RequeryResponse, error) {
	req := builder.Build()
	path := basePath + "/requery"

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[RequeryResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
