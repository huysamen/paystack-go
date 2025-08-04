package paymentrequests

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ListPaymentRequestsRequest struct {
	PerPage        int    `json:"perPage,omitempty"`
	Page           int    `json:"page,omitempty"`
	Customer       string `json:"customer,omitempty"`
	Status         string `json:"status,omitempty"`
	Currency       string `json:"currency,omitempty"`
	IncludeArchive string `json:"include_archive,omitempty"`
	From           string `json:"from,omitempty"`
	To             string `json:"to,omitempty"`
}

type ListPaymentRequestsRequestBuilder struct {
	req *ListPaymentRequestsRequest
}

func NewListPaymentRequestsRequest() *ListPaymentRequestsRequestBuilder {
	return &ListPaymentRequestsRequestBuilder{
		req: &ListPaymentRequestsRequest{},
	}
}

func (b *ListPaymentRequestsRequestBuilder) PerPage(perPage int) *ListPaymentRequestsRequestBuilder {
	b.req.PerPage = perPage

	return b
}

func (b *ListPaymentRequestsRequestBuilder) Page(page int) *ListPaymentRequestsRequestBuilder {
	b.req.Page = page

	return b
}

func (b *ListPaymentRequestsRequestBuilder) Customer(customer string) *ListPaymentRequestsRequestBuilder {
	b.req.Customer = customer

	return b
}

func (b *ListPaymentRequestsRequestBuilder) Status(status string) *ListPaymentRequestsRequestBuilder {
	b.req.Status = status

	return b
}

func (b *ListPaymentRequestsRequestBuilder) Currency(currency string) *ListPaymentRequestsRequestBuilder {
	b.req.Currency = currency

	return b
}

func (b *ListPaymentRequestsRequestBuilder) IncludeArchive(includeArchive string) *ListPaymentRequestsRequestBuilder {
	b.req.IncludeArchive = includeArchive

	return b
}

func (b *ListPaymentRequestsRequestBuilder) From(from string) *ListPaymentRequestsRequestBuilder {
	b.req.From = from

	return b
}

func (b *ListPaymentRequestsRequestBuilder) To(to string) *ListPaymentRequestsRequestBuilder {
	b.req.To = to

	return b
}

func (b *ListPaymentRequestsRequestBuilder) DateRange(from, to string) *ListPaymentRequestsRequestBuilder {
	b.req.From = from
	b.req.To = to

	return b
}

func (b *ListPaymentRequestsRequestBuilder) Build() *ListPaymentRequestsRequest {
	return b.req
}

type ListPaymentRequestsResponse = types.Response[[]types.PaymentRequest]

func (c *Client) List(ctx context.Context, builder *ListPaymentRequestsRequestBuilder) (*ListPaymentRequestsResponse, error) {
	params := url.Values{}

	if builder != nil {
		req := builder.Build()
		if req.PerPage > 0 {
			params.Set("perPage", strconv.Itoa(req.PerPage))
		}
		if req.Page > 0 {
			params.Set("page", strconv.Itoa(req.Page))
		}
		if req.Customer != "" {
			params.Set("customer", req.Customer)
		}
		if req.Status != "" {
			params.Set("status", req.Status)
		}
		if req.Currency != "" {
			params.Set("currency", req.Currency)
		}
		if req.IncludeArchive != "" {
			params.Set("include_archive", req.IncludeArchive)
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

	return net.Get[[]types.PaymentRequest](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
