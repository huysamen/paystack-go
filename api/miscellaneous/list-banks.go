package miscellaneous

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Bank represents a bank in the system
type Bank struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Code        string    `json:"code"`
	LongCode    string    `json:"longcode"`
	Gateway     *string   `json:"gateway"`
	PayWithBank bool      `json:"pay_with_bank"`
	Active      bool      `json:"active"`
	IsDeleted   bool      `json:"is_deleted"`
	Country     string    `json:"country"`
	Currency    string    `json:"currency"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// BankListRequest represents the request to list banks
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

// BankListRequestBuilder provides a fluent interface for building BankListRequest
type BankListRequestBuilder struct {
	req *BankListRequest
}

// NewBankListRequest creates a new builder for BankListRequest
func NewBankListRequest() *BankListRequestBuilder {
	return &BankListRequestBuilder{
		req: &BankListRequest{},
	}
}

// Country sets the country filter
func (b *BankListRequestBuilder) Country(country string) *BankListRequestBuilder {
	b.req.Country = &country
	return b
}

// UseCursor enables cursor-based pagination
func (b *BankListRequestBuilder) UseCursor(useCursor bool) *BankListRequestBuilder {
	b.req.UseCursor = &useCursor
	return b
}

// PerPage sets the number of records per page
func (b *BankListRequestBuilder) PerPage(perPage int) *BankListRequestBuilder {
	b.req.PerPage = &perPage
	return b
}

// PayWithBankTransfer filters for transfer payment banks
func (b *BankListRequestBuilder) PayWithBankTransfer(payWithBankTransfer bool) *BankListRequestBuilder {
	b.req.PayWithBankTransfer = &payWithBankTransfer
	return b
}

// PayWithBank filters for direct payment banks
func (b *BankListRequestBuilder) PayWithBank(payWithBank bool) *BankListRequestBuilder {
	b.req.PayWithBank = &payWithBank
	return b
}

// EnabledForVerification filters for verification-supported banks
func (b *BankListRequestBuilder) EnabledForVerification(enabled bool) *BankListRequestBuilder {
	b.req.EnabledForVerification = &enabled
	return b
}

// Next sets the cursor for next page
func (b *BankListRequestBuilder) Next(next string) *BankListRequestBuilder {
	b.req.Next = &next
	return b
}

// Previous sets the cursor for previous page
func (b *BankListRequestBuilder) Previous(previous string) *BankListRequestBuilder {
	b.req.Previous = &previous
	return b
}

// Gateway sets the gateway type filter
func (b *BankListRequestBuilder) Gateway(gateway string) *BankListRequestBuilder {
	b.req.Gateway = &gateway
	return b
}

// Type sets the financial channel type
func (b *BankListRequestBuilder) Type(channelType string) *BankListRequestBuilder {
	b.req.Type = &channelType
	return b
}

// Currency sets the currency filter
func (b *BankListRequestBuilder) Currency(currency string) *BankListRequestBuilder {
	b.req.Currency = &currency
	return b
}

// IncludeNIPSortCode includes NIP institution codes
func (b *BankListRequestBuilder) IncludeNIPSortCode(include bool) *BankListRequestBuilder {
	b.req.IncludeNIPSortCode = &include
	return b
}

// Build returns the constructed BankListRequest
func (b *BankListRequestBuilder) Build() *BankListRequest {
	return b.req
}

// BankListResponse represents the response from listing banks
type BankListResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    []Bank      `json:"data"`
	Meta    *types.Meta `json:"meta,omitempty"`
}

// ListBanks retrieves a list of banks using the builder pattern
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

	endpoint := bankBasePath
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	resp, err := net.Get[BankListResponse](
		ctx,
		c.client,
		c.secret,
		endpoint,
		c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
