package miscellaneous

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type BankListRequest struct {
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

type BankListRequestBuilder struct {
	req *BankListRequest
}

func NewBankListRequest() *BankListRequestBuilder {
	return &BankListRequestBuilder{
		req: &BankListRequest{},
	}
}

func (b *BankListRequestBuilder) Country(country string) *BankListRequestBuilder {
	b.req.Country = &country

	return b
}

func (b *BankListRequestBuilder) UseCursor(useCursor bool) *BankListRequestBuilder {
	b.req.UseCursor = &useCursor

	return b
}

func (b *BankListRequestBuilder) PerPage(perPage int) *BankListRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *BankListRequestBuilder) PayWithBankTransfer(payWithBankTransfer bool) *BankListRequestBuilder {
	b.req.PayWithBankTransfer = &payWithBankTransfer

	return b
}

func (b *BankListRequestBuilder) PayWithBank(payWithBank bool) *BankListRequestBuilder {
	b.req.PayWithBank = &payWithBank

	return b
}

func (b *BankListRequestBuilder) EnabledForVerification(enabled bool) *BankListRequestBuilder {
	b.req.EnabledForVerification = &enabled

	return b
}

func (b *BankListRequestBuilder) Next(next string) *BankListRequestBuilder {
	b.req.Next = &next

	return b
}

func (b *BankListRequestBuilder) Previous(previous string) *BankListRequestBuilder {
	b.req.Previous = &previous

	return b
}

func (b *BankListRequestBuilder) Gateway(gateway string) *BankListRequestBuilder {
	b.req.Gateway = &gateway

	return b
}

func (b *BankListRequestBuilder) Type(channelType string) *BankListRequestBuilder {
	b.req.Type = &channelType

	return b
}

func (b *BankListRequestBuilder) Currency(currency string) *BankListRequestBuilder {
	b.req.Currency = &currency

	return b
}

func (b *BankListRequestBuilder) IncludeNIPSortCode(include bool) *BankListRequestBuilder {
	b.req.IncludeNIPSortCode = &include

	return b
}

func (b *BankListRequestBuilder) Build() *BankListRequest {
	return b.req
}

type BankListResponse = types.Response[[]types.Bank]

func (c *Client) ListBanks(ctx context.Context, builder *BankListRequestBuilder) (*BankListResponse, error) {
	params := url.Values{}

	if builder != nil {
		req := builder.Build()
		if req.Country != nil {
			params.Set("country", *req.Country)
		}
		if req.UseCursor != nil {
			params.Set("use_cursor", strconv.FormatBool(*req.UseCursor))
		}
		if req.PerPage != nil {
			params.Set("perPage", strconv.Itoa(*req.PerPage))
		}
		if req.PayWithBankTransfer != nil {
			params.Set("pay_with_bank_transfer", strconv.FormatBool(*req.PayWithBankTransfer))
		}
		if req.PayWithBank != nil {
			params.Set("pay_with_bank", strconv.FormatBool(*req.PayWithBank))
		}
		if req.EnabledForVerification != nil {
			params.Set("enabled_for_verification", strconv.FormatBool(*req.EnabledForVerification))
		}
		if req.Next != nil {
			params.Set("next", *req.Next)
		}
		if req.Previous != nil {
			params.Set("previous", *req.Previous)
		}
		if req.Gateway != nil {
			params.Set("gateway", *req.Gateway)
		}
		if req.Type != nil {
			params.Set("type", *req.Type)
		}
		if req.Currency != nil {
			params.Set("currency", *req.Currency)
		}
		if req.IncludeNIPSortCode != nil {
			params.Set("include_nip_sort_code", strconv.FormatBool(*req.IncludeNIPSortCode))
		}
	}

	endpoint := bankPath
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	return net.Get[[]types.Bank](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
