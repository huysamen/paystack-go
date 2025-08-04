package disputes

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// ListDisputesRequest represents the request to list disputes
type ListDisputesRequest struct {
	From        *time.Time           `json:"from,omitempty"`
	To          *time.Time           `json:"to,omitempty"`
	PerPage     *int                 `json:"per_page,omitempty"`
	Page        *int                 `json:"page,omitempty"`
	Transaction *string              `json:"transaction,omitempty"`
	Status      *types.DisputeStatus `json:"status,omitempty"`
}

// ListDisputesBuilder builds requests for listing disputes
type ListDisputesBuilder struct {
	req *ListDisputesRequest
}

// NewListDisputesBuilder creates a new builder for listing disputes
func NewListDisputesBuilder() *ListDisputesBuilder {
	return &ListDisputesBuilder{
		req: &ListDisputesRequest{},
	}
}

// From sets the start date filter
func (b *ListDisputesBuilder) From(from time.Time) *ListDisputesBuilder {
	b.req.From = &from

	return b
}

// To sets the end date filter
func (b *ListDisputesBuilder) To(to time.Time) *ListDisputesBuilder {
	b.req.To = &to

	return b
}

// DateRange sets both from and to dates for convenience
func (b *ListDisputesBuilder) DateRange(from, to time.Time) *ListDisputesBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

// PerPage sets the number of disputes per page
func (b *ListDisputesBuilder) PerPage(perPage int) *ListDisputesBuilder {
	b.req.PerPage = &perPage

	return b
}

// Page sets the page number
func (b *ListDisputesBuilder) Page(page int) *ListDisputesBuilder {
	b.req.Page = &page

	return b
}

// Transaction filters by transaction ID
func (b *ListDisputesBuilder) Transaction(transaction string) *ListDisputesBuilder {
	b.req.Transaction = &transaction

	return b
}

// Status filters by dispute status
func (b *ListDisputesBuilder) Status(status types.DisputeStatus) *ListDisputesBuilder {
	b.req.Status = &status

	return b
}

// Build returns the built request
func (b *ListDisputesBuilder) Build() *ListDisputesRequest {
	return b.req
}

// ListDisputesResponse represents the response from listing disputes
type ListDisputesResponse = types.Response[[]types.Dispute]

// List retrieves disputes filed against your integration
func (c *Client) List(ctx context.Context, builder *ListDisputesBuilder) (*ListDisputesResponse, error) {
	endpoint := basePath
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

	return net.Get[[]types.Dispute](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
