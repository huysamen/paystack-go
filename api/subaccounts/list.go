package subaccounts

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SubaccountListRequest struct {
	PerPage *int       `json:"perPage,omitempty"` // Optional: records per page (default: 50)
	Page    *int       `json:"page,omitempty"`    // Optional: page number (default: 1)
	From    *time.Time `json:"from,omitempty"`    // Optional: start date filter
	To      *time.Time `json:"to,omitempty"`      // Optional: end date filter
}

type SubaccountListRequestBuilder struct {
	req *SubaccountListRequest
}

func NewSubaccountListRequest() *SubaccountListRequestBuilder {
	return &SubaccountListRequestBuilder{
		req: &SubaccountListRequest{},
	}
}

func (b *SubaccountListRequestBuilder) PerPage(perPage int) *SubaccountListRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *SubaccountListRequestBuilder) Page(page int) *SubaccountListRequestBuilder {
	b.req.Page = &page

	return b
}

func (b *SubaccountListRequestBuilder) From(from time.Time) *SubaccountListRequestBuilder {
	b.req.From = &from

	return b
}

func (b *SubaccountListRequestBuilder) To(to time.Time) *SubaccountListRequestBuilder {
	b.req.To = &to

	return b
}

func (b *SubaccountListRequestBuilder) DateRange(from, to time.Time) *SubaccountListRequestBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

func (b *SubaccountListRequestBuilder) Build() *SubaccountListRequest {
	return b.req
}

type SubaccountListResponse = types.Response[[]types.Subaccount]

func (c *Client) List(ctx context.Context, builder *SubaccountListRequestBuilder) (*SubaccountListResponse, error) {
	params := url.Values{}
	if builder != nil {
		req := builder.Build()
		if req.PerPage != nil {
			params.Set("perPage", strconv.Itoa(*req.PerPage))
		}
		if req.Page != nil {
			params.Set("page", strconv.Itoa(*req.Page))
		}
		if req.From != nil {
			params.Set("from", req.From.Format(time.RFC3339))
		}
		if req.To != nil {
			params.Set("to", req.To.Format(time.RFC3339))
		}
	}

	endpoint := basePath
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	return net.Get[[]types.Subaccount](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
