package transactionsplits

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionSplitListRequest struct {
	Name    *string    `json:"name,omitempty"`    // Filter by name (optional)
	Active  *bool      `json:"active,omitempty"`  // Filter by active status (optional)
	SortBy  *string    `json:"sort_by,omitempty"` // Sort by field, defaults to createdAt (optional)
	PerPage *int       `json:"perPage,omitempty"` // Number of splits per page (default: 50)
	Page    *int       `json:"page,omitempty"`    // Page number (default: 1)
	From    *time.Time `json:"from,omitempty"`    // Start date filter (optional)
	To      *time.Time `json:"to,omitempty"`      // End date filter (optional)
}

type TransactionSplitListRequestBuilder struct {
	name    *string
	active  *bool
	sortBy  *string
	perPage *int
	page    *int
	from    *time.Time
	to      *time.Time
}

func NewTransactionSplitListRequest() *TransactionSplitListRequestBuilder {
	return &TransactionSplitListRequestBuilder{}
}

func (b *TransactionSplitListRequestBuilder) Name(name string) *TransactionSplitListRequestBuilder {
	b.name = &name
	return b
}

func (b *TransactionSplitListRequestBuilder) Active(active bool) *TransactionSplitListRequestBuilder {
	b.active = &active
	return b
}

func (b *TransactionSplitListRequestBuilder) SortBy(sortBy string) *TransactionSplitListRequestBuilder {
	b.sortBy = &sortBy
	return b
}

func (b *TransactionSplitListRequestBuilder) PerPage(perPage int) *TransactionSplitListRequestBuilder {
	b.perPage = &perPage
	return b
}

func (b *TransactionSplitListRequestBuilder) Page(page int) *TransactionSplitListRequestBuilder {
	b.page = &page
	return b
}

func (b *TransactionSplitListRequestBuilder) DateRange(from, to time.Time) *TransactionSplitListRequestBuilder {
	b.from = &from
	b.to = &to
	return b
}

func (b *TransactionSplitListRequestBuilder) From(from time.Time) *TransactionSplitListRequestBuilder {
	b.from = &from
	return b
}

func (b *TransactionSplitListRequestBuilder) To(to time.Time) *TransactionSplitListRequestBuilder {
	b.to = &to
	return b
}

func (b *TransactionSplitListRequestBuilder) Build() *TransactionSplitListRequest {
	return &TransactionSplitListRequest{
		Name:    b.name,
		Active:  b.active,
		SortBy:  b.sortBy,
		PerPage: b.perPage,
		Page:    b.page,
		From:    b.from,
		To:      b.to,
	}
}

type TransactionSplitListResponse = types.Response[[]types.TransactionSplit]

func (c *Client) List(ctx context.Context, builder *TransactionSplitListRequestBuilder) (*types.Response[[]types.TransactionSplit], error) {
	params := url.Values{}

	if builder != nil {
		req := builder.Build()
		if req.Name != nil {
			params.Set("name", *req.Name)
		}
		if req.Active != nil {
			params.Set("active", strconv.FormatBool(*req.Active))
		}
		if req.SortBy != nil {
			params.Set("sort_by", *req.SortBy)
		}
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

	query := ""
	if len(params) > 0 {
		query = params.Encode()
	}

	return net.Get[[]types.TransactionSplit](ctx, c.Client, c.Secret, basePath, query, c.BaseURL)
}
