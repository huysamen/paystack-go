package miscellaneous

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type listBanksRequest struct {
	Country                *string `json:"country,omitempty"`                  // Optional: country filter (ghana, kenya, nigeria, south africa)
	UseCursor              *bool   `json:"use_cursor,omitempty"`               // Optional: enable cursor pagination
	PerPage                *int    `json:"perPage,omitempty"`                  // Optional: records per page (default: 50, max: 100)
	PayWithBankTransfer    *bool   `json:"pay_with_bank_transfer,omitempty"`   // Optional: filter for transfer payment banks
	PayWithBank            *bool   `json:"pay_with_bank,omitempty"`            // Optional: filter for direct payment banks
	EnabledForVerification *bool   `json:"enabled_for_verification,omitempty"` // Optional: filter for verification-supported banks
	Next                   *string `json:"next,omitempty"`                     // Optional: cursor for next page
	Previous               *string `json:"previous,omitempty"`                 // Optional: cursor for previous page
	Gateway                *string `json:"gateway,omitempty"`                  // Optional: gateway type filter
	Type                   *string `json:"type,omitempty"`                     // Optional: financial channel type
	Currency               *string `json:"currency,omitempty"`                 // Optional: currency filter
	IncludeNIPSortCode     *bool   `json:"include_nip_sort_code,omitempty"`    // Optional: include NIP institution codes
}

type ListBanksRequestBuilder struct {
	req *listBanksRequest
}

func NewListBanksRequestBuilder() *ListBanksRequestBuilder {
	return &ListBanksRequestBuilder{
		req: &listBanksRequest{},
	}
}

func (b *ListBanksRequestBuilder) Country(country string) *ListBanksRequestBuilder {
	b.req.Country = &country

	return b
}

func (b *ListBanksRequestBuilder) UseCursor(useCursor bool) *ListBanksRequestBuilder {
	b.req.UseCursor = &useCursor

	return b
}

func (b *ListBanksRequestBuilder) PerPage(perPage int) *ListBanksRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *ListBanksRequestBuilder) PayWithBankTransfer(payWithBankTransfer bool) *ListBanksRequestBuilder {
	b.req.PayWithBankTransfer = &payWithBankTransfer

	return b
}

func (b *ListBanksRequestBuilder) PayWithBank(payWithBank bool) *ListBanksRequestBuilder {
	b.req.PayWithBank = &payWithBank

	return b
}

func (b *ListBanksRequestBuilder) EnabledForVerification(enabled bool) *ListBanksRequestBuilder {
	b.req.EnabledForVerification = &enabled

	return b
}

func (b *ListBanksRequestBuilder) Next(next string) *ListBanksRequestBuilder {
	b.req.Next = &next

	return b
}

func (b *ListBanksRequestBuilder) Previous(previous string) *ListBanksRequestBuilder {
	b.req.Previous = &previous

	return b
}

func (b *ListBanksRequestBuilder) Gateway(gateway string) *ListBanksRequestBuilder {
	b.req.Gateway = &gateway

	return b
}

func (b *ListBanksRequestBuilder) Type(channelType string) *ListBanksRequestBuilder {
	b.req.Type = &channelType

	return b
}

func (b *ListBanksRequestBuilder) Currency(currency string) *ListBanksRequestBuilder {
	b.req.Currency = &currency

	return b
}

func (b *ListBanksRequestBuilder) IncludeNIPSortCode(include bool) *ListBanksRequestBuilder {
	b.req.IncludeNIPSortCode = &include

	return b
}

func (b *ListBanksRequestBuilder) Build() *listBanksRequest {
	return b.req
}

func (r *listBanksRequest) toQuery() string {
	params := url.Values{}
	if r.Country != nil {
		params.Set("country", *r.Country)
	}
	if r.UseCursor != nil {
		params.Set("use_cursor", strconv.FormatBool(*r.UseCursor))
	}
	if r.PerPage != nil {
		params.Set("perPage", strconv.Itoa(*r.PerPage))
	}
	if r.PayWithBankTransfer != nil {
		params.Set("pay_with_bank_transfer", strconv.FormatBool(*r.PayWithBankTransfer))
	}
	if r.PayWithBank != nil {
		params.Set("pay_with_bank", strconv.FormatBool(*r.PayWithBank))
	}
	if r.EnabledForVerification != nil {
		params.Set("enabled_for_verification", strconv.FormatBool(*r.EnabledForVerification))
	}
	if r.Next != nil {
		params.Set("next", *r.Next)
	}
	if r.Previous != nil {
		params.Set("previous", *r.Previous)
	}
	if r.Gateway != nil {
		params.Set("gateway", *r.Gateway)
	}
	if r.Type != nil {
		params.Set("type", *r.Type)
	}
	if r.Currency != nil {
		params.Set("currency", *r.Currency)
	}
	if r.IncludeNIPSortCode != nil {
		params.Set("include_nip_sort_code", strconv.FormatBool(*r.IncludeNIPSortCode))
	}

	return params.Encode()
}

type ListBanksResponseData = []types.Bank
type ListBanksResponse = types.Response[ListBanksResponseData]

func (c *Client) ListBanks(ctx context.Context, builder ListBanksRequestBuilder) (*ListBanksResponse, error) {
	path := bankPath

	req := builder.Build()
	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[ListBanksResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
