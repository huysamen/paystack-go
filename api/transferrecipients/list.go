package transferrecipients

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransferRecipientListRequest struct {
	PerPage *int       `json:"perPage,omitempty"` // Optional: records per page (default: 50)
	Page    *int       `json:"page,omitempty"`    // Optional: page number (default: 1)
	From    *time.Time `json:"from,omitempty"`    // Optional: start date filter
	To      *time.Time `json:"to,omitempty"`      // Optional: end date filter
}

type TransferRecipientListRequestBuilder struct {
	req *TransferRecipientListRequest
}

func NewTransferRecipientListRequest() *TransferRecipientListRequestBuilder {
	return &TransferRecipientListRequestBuilder{
		req: &TransferRecipientListRequest{},
	}
}

func (b *TransferRecipientListRequestBuilder) PerPage(perPage int) *TransferRecipientListRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *TransferRecipientListRequestBuilder) Page(page int) *TransferRecipientListRequestBuilder {
	b.req.Page = &page

	return b
}

func (b *TransferRecipientListRequestBuilder) From(from time.Time) *TransferRecipientListRequestBuilder {
	b.req.From = &from

	return b
}

func (b *TransferRecipientListRequestBuilder) To(to time.Time) *TransferRecipientListRequestBuilder {
	b.req.To = &to

	return b
}

func (b *TransferRecipientListRequestBuilder) DateRange(from, to time.Time) *TransferRecipientListRequestBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

func (b *TransferRecipientListRequestBuilder) Build() *TransferRecipientListRequest {
	return b.req
}

type TransferRecipientListResponse = types.Response[[]types.TransferRecipient]

func (c *Client) List(ctx context.Context, builder *TransferRecipientListRequestBuilder) (*TransferRecipientListResponse, error) {
	req := builder.Build()
	params := url.Values{}

	if req != nil {
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

	queryParams := ""
	if len(params) > 0 {
		queryParams = params.Encode()
	}

	return net.Get[[]types.TransferRecipient](ctx, c.Client, c.Secret, basePath, queryParams, c.BaseURL)
}
