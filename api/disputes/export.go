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

type ExportResponseData struct {
	Path      string          `json:"path"`
	ExpiresAt *types.DateTime `json:"expires_at,omitempty"`
}

type ExportResponse = types.Response[ExportResponseData]

func (c *Client) Export(ctx context.Context, builder *ExportRequestBuilder) (*ExportResponse, error) {
	endpoint := basePath + "/export"
	req := builder.Build()

	params := url.Values{}
	if req.From != nil {
		params.Set("from", req.From.Format("2006-01-02"))
	}
	if req.To != nil {
		params.Set("to", req.To.Format("2006-01-02"))
	}
	if req.PerPage != nil {
		params.Set("per_page", strconv.Itoa(*req.PerPage))
	}
	if req.Page != nil {
		params.Set("page", strconv.Itoa(*req.Page))
	}
	if req.Transaction != nil {
		params.Set("transaction", *req.Transaction)
	}
	if req.Status != nil {
		params.Set("status", string(*req.Status))
	}

	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	return net.Get[ExportResponseData](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
