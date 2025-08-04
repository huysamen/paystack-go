package disputes

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ExportDisputesRequest struct {
	From        *time.Time           `json:"from,omitempty"`
	To          *time.Time           `json:"to,omitempty"`
	PerPage     *int                 `json:"per_page,omitempty"`
	Page        *int                 `json:"page,omitempty"`
	Transaction *string              `json:"transaction,omitempty"`
	Status      *types.DisputeStatus `json:"status,omitempty"`
}

type ExportDisputesBuilder struct {
	request *ExportDisputesRequest
}

func NewExportDisputesBuilder() *ExportDisputesBuilder {
	return &ExportDisputesBuilder{
		request: &ExportDisputesRequest{},
	}
}

func (b *ExportDisputesBuilder) From(from time.Time) *ExportDisputesBuilder {
	b.request.From = &from

	return b
}

func (b *ExportDisputesBuilder) To(to time.Time) *ExportDisputesBuilder {
	b.request.To = &to

	return b
}

func (b *ExportDisputesBuilder) DateRange(from, to time.Time) *ExportDisputesBuilder {
	b.request.From = &from
	b.request.To = &to

	return b
}

func (b *ExportDisputesBuilder) PerPage(perPage int) *ExportDisputesBuilder {
	b.request.PerPage = &perPage

	return b
}

func (b *ExportDisputesBuilder) Page(page int) *ExportDisputesBuilder {
	b.request.Page = &page

	return b
}

func (b *ExportDisputesBuilder) Transaction(transaction string) *ExportDisputesBuilder {
	b.request.Transaction = &transaction

	return b
}

func (b *ExportDisputesBuilder) Status(status types.DisputeStatus) *ExportDisputesBuilder {
	b.request.Status = &status

	return b
}

func (b *ExportDisputesBuilder) Build() *ExportDisputesRequest {
	return b.request
}

type ExportDisputesResponse = types.Response[ExportData]

func (c *Client) Export(ctx context.Context, builder *ExportDisputesBuilder) (*ExportDisputesResponse, error) {
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

	return net.Get[ExportData](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
