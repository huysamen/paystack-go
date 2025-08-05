package disputes

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type exportRequest struct {
	From        *time.Time           `json:"from,omitempty"`
	To          *time.Time           `json:"to,omitempty"`
	PerPage     *int                 `json:"per_page,omitempty"`
	Page        *int                 `json:"page,omitempty"`
	Transaction *string              `json:"transaction,omitempty"`
	Status      *types.DisputeStatus `json:"status,omitempty"`
}

type ExportRequestBuilder struct {
	request *exportRequest
}

func NewExportRequestBuilder() *ExportRequestBuilder {
	return &ExportRequestBuilder{
		request: &exportRequest{},
	}
}

func (b *ExportRequestBuilder) From(from time.Time) *ExportRequestBuilder {
	b.request.From = &from

	return b
}

func (b *ExportRequestBuilder) To(to time.Time) *ExportRequestBuilder {
	b.request.To = &to

	return b
}

func (b *ExportRequestBuilder) DateRange(from, to time.Time) *ExportRequestBuilder {
	b.request.From = &from
	b.request.To = &to

	return b
}

func (b *ExportRequestBuilder) PerPage(perPage int) *ExportRequestBuilder {
	b.request.PerPage = &perPage

	return b
}

func (b *ExportRequestBuilder) Page(page int) *ExportRequestBuilder {
	b.request.Page = &page

	return b
}

func (b *ExportRequestBuilder) Transaction(transaction string) *ExportRequestBuilder {
	b.request.Transaction = &transaction

	return b
}

func (b *ExportRequestBuilder) Status(status types.DisputeStatus) *ExportRequestBuilder {
	b.request.Status = &status

	return b
}

func (b *ExportRequestBuilder) Build() *exportRequest {
	return b.request
}

func (r *exportRequest) toQuery() string {
	params := url.Values{}
	if r.From != nil {
		params.Set("from", r.From.Format("2006-01-02"))
	}
	if r.To != nil {
		params.Set("to", r.To.Format("2006-01-02"))
	}
	if r.PerPage != nil {
		params.Set("per_page", strconv.Itoa(*r.PerPage))
	}
	if r.Page != nil {
		params.Set("page", strconv.Itoa(*r.Page))
	}
	if r.Transaction != nil {
		params.Set("transaction", *r.Transaction)
	}
	if r.Status != nil {
		params.Set("status", string(*r.Status))
	}

	return params.Encode()
}

type ExportResponseData struct {
	Path      string          `json:"path"`
	ExpiresAt *types.DateTime `json:"expires_at,omitempty"`
}

type ExportResponse = types.Response[ExportResponseData]

func (c *Client) Export(ctx context.Context, builder ExportRequestBuilder) (*ExportResponse, error) {
	path := basePath + "/export"

	req := builder.Build()
	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[ExportResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
