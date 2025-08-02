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
	From        *time.Time     `json:"from,omitempty"`
	To          *time.Time     `json:"to,omitempty"`
	PerPage     *int           `json:"per_page,omitempty"`
	Page        *int           `json:"page,omitempty"`
	Transaction *string        `json:"transaction,omitempty"`
	Status      *DisputeStatus `json:"status,omitempty"`
}

// ListDisputesResponse represents the response from listing disputes
type ListDisputesResponse = types.Response[[]Dispute]

// ListDisputesBuilder builds requests for listing disputes
type ListDisputesBuilder struct {
	request *ListDisputesRequest
}

// NewListDisputesBuilder creates a new builder for listing disputes
func NewListDisputesBuilder() *ListDisputesBuilder {
	return &ListDisputesBuilder{
		request: &ListDisputesRequest{},
	}
}

// From sets the start date filter
func (b *ListDisputesBuilder) From(from time.Time) *ListDisputesBuilder {
	b.request.From = &from
	return b
}

// To sets the end date filter
func (b *ListDisputesBuilder) To(to time.Time) *ListDisputesBuilder {
	b.request.To = &to
	return b
}

// DateRange sets both from and to dates for convenience
func (b *ListDisputesBuilder) DateRange(from, to time.Time) *ListDisputesBuilder {
	b.request.From = &from
	b.request.To = &to
	return b
}

// PerPage sets the number of disputes per page
func (b *ListDisputesBuilder) PerPage(perPage int) *ListDisputesBuilder {
	b.request.PerPage = &perPage
	return b
}

// Page sets the page number
func (b *ListDisputesBuilder) Page(page int) *ListDisputesBuilder {
	b.request.Page = &page
	return b
}

// Transaction filters by transaction ID
func (b *ListDisputesBuilder) Transaction(transaction string) *ListDisputesBuilder {
	b.request.Transaction = &transaction
	return b
}

// Status filters by dispute status
func (b *ListDisputesBuilder) Status(status DisputeStatus) *ListDisputesBuilder {
	b.request.Status = &status
	return b
}

// Build returns the built request
func (b *ListDisputesBuilder) Build() *ListDisputesRequest {
	return b.request
}

// List retrieves disputes filed against your integration
func (c *Client) List(ctx context.Context, builder *ListDisputesBuilder) (*types.Response[[]Dispute], error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	endpoint := c.baseURL + disputesBasePath
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

	resp, err := net.Get[[]Dispute](ctx, c.client, c.secret, endpoint, c.baseURL)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
