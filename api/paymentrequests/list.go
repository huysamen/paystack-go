package paymentrequests

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type listRequest struct {
	PerPage        int    `json:"perPage,omitempty"`
	Page           int    `json:"page,omitempty"`
	Customer       string `json:"customer,omitempty"`
	Status         string `json:"status,omitempty"`
	Currency       string `json:"currency,omitempty"`
	IncludeArchive string `json:"include_archive,omitempty"`
	From           string `json:"from,omitempty"`
	To             string `json:"to,omitempty"`
}

type ListRequestBuilder struct {
	req *listRequest
}

func NewListRequestBuilder() *ListRequestBuilder {
	return &ListRequestBuilder{
		req: &listRequest{},
	}
}

func (b *ListRequestBuilder) PerPage(perPage int) *ListRequestBuilder {
	b.req.PerPage = perPage

	return b
}

func (b *ListRequestBuilder) Page(page int) *ListRequestBuilder {
	b.req.Page = page

	return b
}

func (b *ListRequestBuilder) Customer(customer string) *ListRequestBuilder {
	b.req.Customer = customer

	return b
}

func (b *ListRequestBuilder) Status(status string) *ListRequestBuilder {
	b.req.Status = status

	return b
}

func (b *ListRequestBuilder) Currency(currency string) *ListRequestBuilder {
	b.req.Currency = currency

	return b
}

func (b *ListRequestBuilder) IncludeArchive(includeArchive string) *ListRequestBuilder {
	b.req.IncludeArchive = includeArchive

	return b
}

func (b *ListRequestBuilder) From(from string) *ListRequestBuilder {
	b.req.From = from

	return b
}

func (b *ListRequestBuilder) To(to string) *ListRequestBuilder {
	b.req.To = to

	return b
}

func (b *ListRequestBuilder) DateRange(from, to string) *ListRequestBuilder {
	b.req.From = from
	b.req.To = to

	return b
}

func (b *ListRequestBuilder) Build() *listRequest {
	return b.req
}

func (r *listRequest) toQuery() string {
	params := url.Values{}
	if r.PerPage > 0 {
		params.Set("perPage", strconv.Itoa(r.PerPage))
	}
	if r.Page > 0 {
		params.Set("page", strconv.Itoa(r.Page))
	}
	if r.Customer != "" {
		params.Set("customer", r.Customer)
	}
	if r.Status != "" {
		params.Set("status", r.Status)
	}
	if r.Currency != "" {
		params.Set("currency", r.Currency)
	}
	if r.IncludeArchive != "" {
		params.Set("include_archive", r.IncludeArchive)
	}
	if r.From != "" {
		params.Set("from", r.From)
	}
	if r.To != "" {
		params.Set("to", r.To)
	}

	return params.Encode()
}

type ListResponseData = []types.PaymentRequest
type ListResponse = types.Response[ListResponseData]

func (c *Client) List(ctx context.Context, builder ListRequestBuilder) (*ListResponse, error) {
	req := builder.Build()
	path := basePath

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[ListResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
