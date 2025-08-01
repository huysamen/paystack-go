package miscellaneous

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
)

// ListBanks retrieves a list of banks
func (c *Client) ListBanks(ctx context.Context, req *BankListRequest) (*BankListResponse, error) {
	params := url.Values{}

	if req != nil {
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

	resp, err := net.Get[BankListResponse](ctx, c.client, c.secret, endpoint, c.baseURL)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// ListBanksWithBuilder retrieves a list of banks using the builder pattern
func (c *Client) ListBanksWithBuilder(ctx context.Context, builder *BankListRequestBuilder) (*BankListResponse, error) {
	if builder == nil {
		return c.ListBanks(ctx, nil)
	}
	return c.ListBanks(ctx, builder.Build())
}
