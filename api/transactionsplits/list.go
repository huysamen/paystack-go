package transactionsplits

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type listRequest struct {
	Name    *string    `json:"name,omitempty"`    // Filter by name (optional)
	Active  *bool      `json:"active,omitempty"`  // Filter by active status (optional)
	SortBy  *string    `json:"sort_by,omitempty"` // Sort by field, defaults to createdAt (optional)
	PerPage *int       `json:"perPage,omitempty"` // Number of splits per page (default: 50)
	Page    *int       `json:"page,omitempty"`    // Page number (default: 1)
	From    *time.Time `json:"from,omitempty"`    // Start date filter (optional)
	To      *time.Time `json:"to,omitempty"`      // End date filter (optional)
}

type ListRequestBuilder struct {
	req *listRequest
}

func NewListRequestBuilder() *ListRequestBuilder {
	return &ListRequestBuilder{
		req: &listRequest{},
	}
}

func (b *ListRequestBuilder) Name(name string) *ListRequestBuilder {
	b.req.Name = &name

	return b
}

func (b *ListRequestBuilder) Active(active bool) *ListRequestBuilder {
	b.req.Active = &active

	return b
}

func (b *ListRequestBuilder) SortBy(sortBy string) *ListRequestBuilder {
	b.req.SortBy = &sortBy

	return b
}

func (b *ListRequestBuilder) PerPage(perPage int) *ListRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *ListRequestBuilder) Page(page int) *ListRequestBuilder {
	b.req.Page = &page

	return b
}

func (b *ListRequestBuilder) DateRange(from, to time.Time) *ListRequestBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

func (b *ListRequestBuilder) From(from time.Time) *ListRequestBuilder {
	b.req.From = &from

	return b
}

func (b *ListRequestBuilder) To(to time.Time) *ListRequestBuilder {
	b.req.To = &to

	return b
}

func (b *ListRequestBuilder) Build() *listRequest {
	return b.req
}

func (r *listRequest) toQuery() string {
	params := url.Values{}

	if r.Name != nil {
		params.Set("name", *r.Name)
	}

	if r.Active != nil {
		params.Set("active", strconv.FormatBool(*r.Active))
	}

	if r.SortBy != nil {
		params.Set("sort_by", *r.SortBy)
	}

	if r.PerPage != nil {
		params.Set("perPage", strconv.Itoa(*r.PerPage))
	}

	if r.Page != nil {
		params.Set("page", strconv.Itoa(*r.Page))
	}

	if r.From != nil {
		params.Set("from", r.From.Format(time.RFC3339))
	}

	if r.To != nil {
		params.Set("to", r.To.Format(time.RFC3339))
	}

	return params.Encode()
}

type ListResponseData = []types.TransactionSplit
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
