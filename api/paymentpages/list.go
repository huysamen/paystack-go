package paymentpages

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ListPaymentPagesRequest struct {
	PerPage int    `json:"perPage,omitempty"`
	Page    int    `json:"page,omitempty"`
	From    string `json:"from,omitempty"`
	To      string `json:"to,omitempty"`
}

type ListPaymentPagesRequestBuilder struct {
	req *ListPaymentPagesRequest
}

func NewListPaymentPagesRequest() *ListPaymentPagesRequestBuilder {
	return &ListPaymentPagesRequestBuilder{
		req: &ListPaymentPagesRequest{},
	}
}

func (b *ListPaymentPagesRequestBuilder) PerPage(perPage int) *ListPaymentPagesRequestBuilder {
	b.req.PerPage = perPage

	return b
}

func (b *ListPaymentPagesRequestBuilder) Page(page int) *ListPaymentPagesRequestBuilder {
	b.req.Page = page

	return b
}

func (b *ListPaymentPagesRequestBuilder) From(from string) *ListPaymentPagesRequestBuilder {
	b.req.From = from

	return b
}

func (b *ListPaymentPagesRequestBuilder) To(to string) *ListPaymentPagesRequestBuilder {
	b.req.To = to

	return b
}

func (b *ListPaymentPagesRequestBuilder) Build() *ListPaymentPagesRequest {
	return b.req
}

type ListPaymentPagesResponse = types.Response[[]types.PaymentPage]

func (c *Client) List(ctx context.Context, builder *ListPaymentPagesRequestBuilder) (*ListPaymentPagesResponse, error) {
	params := url.Values{}

	if builder != nil {
		req := builder.Build()
		if req.PerPage > 0 {
			params.Set("perPage", strconv.Itoa(req.PerPage))
		}
		if req.Page > 0 {
			params.Set("page", strconv.Itoa(req.Page))
		}
		if req.From != "" {
			params.Set("from", req.From)
		}
		if req.To != "" {
			params.Set("to", req.To)
		}
	}

	endpoint := basePath
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	return net.Get[[]types.PaymentPage](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
