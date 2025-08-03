package disputes

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// ExportDisputesRequest represents the request to export disputes
type ExportDisputesRequest struct {
	From        *time.Time     `json:"from,omitempty"`
	To          *time.Time     `json:"to,omitempty"`
	PerPage     *int           `json:"per_page,omitempty"`
	Page        *int           `json:"page,omitempty"`
	Transaction *string        `json:"transaction,omitempty"`
	Status      *DisputeStatus `json:"status,omitempty"`
}

// ExportDisputesResponse represents the response from exporting disputes
type ExportDisputesResponse = types.Response[ExportData]

// ExportDisputesBuilder builds requests for exporting disputes
type ExportDisputesBuilder struct {
	request *ExportDisputesRequest
}

// NewExportDisputesBuilder creates a new builder for exporting disputes
func NewExportDisputesBuilder() *ExportDisputesBuilder {
	return &ExportDisputesBuilder{
		request: &ExportDisputesRequest{},
	}
}

// From sets the start date filter
func (b *ExportDisputesBuilder) From(from time.Time) *ExportDisputesBuilder {
	b.request.From = &from
	return b
}

// To sets the end date filter
func (b *ExportDisputesBuilder) To(to time.Time) *ExportDisputesBuilder {
	b.request.To = &to
	return b
}

// DateRange sets both from and to dates for convenience
func (b *ExportDisputesBuilder) DateRange(from, to time.Time) *ExportDisputesBuilder {
	b.request.From = &from
	b.request.To = &to
	return b
}

// PerPage sets the number of disputes per page
func (b *ExportDisputesBuilder) PerPage(perPage int) *ExportDisputesBuilder {
	b.request.PerPage = &perPage
	return b
}

// Page sets the page number
func (b *ExportDisputesBuilder) Page(page int) *ExportDisputesBuilder {
	b.request.Page = &page
	return b
}

// Transaction filters by transaction ID
func (b *ExportDisputesBuilder) Transaction(transaction string) *ExportDisputesBuilder {
	b.request.Transaction = &transaction
	return b
}

// Status filters by dispute status
func (b *ExportDisputesBuilder) Status(status DisputeStatus) *ExportDisputesBuilder {
	b.request.Status = &status
	return b
}

// Build returns the built request
func (b *ExportDisputesBuilder) Build() *ExportDisputesRequest {
	return b.request
}

// Export exports disputes available on your integration
func (c *Client) Export(ctx context.Context, builder *ExportDisputesBuilder) (*ExportDisputesResponse, error) {
	endpoint := basePath + "/export"
	req := builder.Build()

	// Build query parameters
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
