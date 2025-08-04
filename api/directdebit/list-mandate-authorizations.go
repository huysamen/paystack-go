package directdebit

import (
	"context"
	"fmt"
	"net/url"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// ListMandateAuthorizationsRequest represents the request to list mandate authorizations
type ListMandateAuthorizationsRequest struct {
	Cursor  string                     `json:"cursor,omitempty"`
	Status  MandateAuthorizationStatus `json:"status,omitempty"`
	PerPage int                        `json:"per_page,omitempty"`
}

// ListMandateAuthorizationsBuilder builds requests for listing mandate authorizations
type ListMandateAuthorizationsBuilder struct {
	req *ListMandateAuthorizationsRequest
}

// NewListMandateAuthorizationsBuilder creates a new builder for listing mandate authorizations
func NewListMandateAuthorizationsBuilder() *ListMandateAuthorizationsBuilder {
	return &ListMandateAuthorizationsBuilder{
		req: &ListMandateAuthorizationsRequest{},
	}
}

// Cursor sets the cursor for pagination
func (b *ListMandateAuthorizationsBuilder) Cursor(cursor string) *ListMandateAuthorizationsBuilder {
	b.req.Cursor = cursor

	return b
}

// Status sets the status filter for mandate authorizations
func (b *ListMandateAuthorizationsBuilder) Status(status MandateAuthorizationStatus) *ListMandateAuthorizationsBuilder {
	b.req.Status = status

	return b
}

// PerPage sets the number of items per page
func (b *ListMandateAuthorizationsBuilder) PerPage(perPage int) *ListMandateAuthorizationsBuilder {
	b.req.PerPage = perPage

	return b
}

// Build returns the built request
func (b *ListMandateAuthorizationsBuilder) Build() *ListMandateAuthorizationsRequest {
	return b.req
}

// MandateAuthorization represents a direct debit mandate authorization
type ListMandateAuthorizationsResponse = types.Response[[]MandateAuthorization]

// ListMandateAuthorizations retrieves a list of direct debit mandate authorizations
func (c *Client) ListMandateAuthorizations(ctx context.Context, builder *ListMandateAuthorizationsBuilder) (*ListMandateAuthorizationsResponse, error) {
	endpoint := basePath + "/mandate-authorizations"

	if builder != nil {
		req := builder.Build()
		params := url.Values{}
		if req.Cursor != "" {
			params.Set("cursor", req.Cursor)
		}
		if req.Status != "" {
			params.Set("status", string(req.Status))
		}
		if req.PerPage > 0 {
			params.Set("per_page", fmt.Sprintf("%d", req.PerPage))
		}

		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	return net.Get[[]MandateAuthorization](ctx, c.Client, c.Secret, endpoint, c.BaseURL)
}
