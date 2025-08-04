package dedicatedvirtualaccount

import (
	"context"
	"net/url"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type RequeryDedicatedAccountRequest struct {
	AccountNumber string `json:"account_number"`
	ProviderSlug  string `json:"provider_slug"`
	Date          string `json:"date,omitempty"`
}

type RequeryDedicatedAccountBuilder struct {
	request *RequeryDedicatedAccountRequest
}

func NewRequeryDedicatedAccountBuilder() *RequeryDedicatedAccountBuilder {
	return &RequeryDedicatedAccountBuilder{
		request: &RequeryDedicatedAccountRequest{},
	}
}

func (b *RequeryDedicatedAccountBuilder) AccountNumber(accountNumber string) *RequeryDedicatedAccountBuilder {
	b.request.AccountNumber = accountNumber

	return b
}

func (b *RequeryDedicatedAccountBuilder) ProviderSlug(providerSlug string) *RequeryDedicatedAccountBuilder {
	b.request.ProviderSlug = providerSlug

	return b
}

func (b *RequeryDedicatedAccountBuilder) Date(date string) *RequeryDedicatedAccountBuilder {
	b.request.Date = date

	return b
}

func (b *RequeryDedicatedAccountBuilder) Build() *RequeryDedicatedAccountRequest {
	return b.request
}

type RequeryResponse = types.Response[any]

func (c *Client) Requery(ctx context.Context, builder *RequeryDedicatedAccountBuilder) (*RequeryResponse, error) {
	req := builder.Build()
	params := url.Values{}
	params.Set("account_number", req.AccountNumber)
	params.Set("provider_slug", req.ProviderSlug)
	if req.Date != "" {
		params.Set("date", req.Date)
	}

	endpoint := basePath + "/requery?" + params.Encode()

	return net.Get[any](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
