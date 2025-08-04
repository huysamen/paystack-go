package directdebit

import (
	"context"
	"fmt"
	"net/url"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ListMandateAuthorizationsRequest struct {
	Cursor  string                           `json:"cursor,omitempty"`
	Status  types.MandateAuthorizationStatus `json:"status,omitempty"`
	PerPage int                              `json:"per_page,omitempty"`
}

type ListMandateAuthorizationsRequestBuilder struct {
	req *ListMandateAuthorizationsRequest
}

func NewListMandateAuthorizationsRequestBuilder() *ListMandateAuthorizationsRequestBuilder {
	return &ListMandateAuthorizationsRequestBuilder{
		req: &ListMandateAuthorizationsRequest{},
	}
}

func (b *ListMandateAuthorizationsRequestBuilder) Cursor(cursor string) *ListMandateAuthorizationsRequestBuilder {
	b.req.Cursor = cursor

	return b
}

func (b *ListMandateAuthorizationsRequestBuilder) Status(status types.MandateAuthorizationStatus) *ListMandateAuthorizationsRequestBuilder {
	b.req.Status = status

	return b
}

func (b *ListMandateAuthorizationsRequestBuilder) PerPage(perPage int) *ListMandateAuthorizationsRequestBuilder {
	b.req.PerPage = perPage

	return b
}

func (b *ListMandateAuthorizationsRequestBuilder) Build() *ListMandateAuthorizationsRequest {
	return b.req
}

func (r *ListMandateAuthorizationsRequest) toQuery() string {
	params := url.Values{}

	if r.Cursor != "" {
		params.Set("cursor", r.Cursor)
	}

	if r.Status != "" {
		params.Set("status", string(r.Status))
	}

	if r.PerPage > 0 {
		params.Set("per_page", fmt.Sprintf("%d", r.PerPage))
	}

	return params.Encode()
}

type ListMandateAuthorizationsResponseData = []types.MandateAuthorization
type ListMandateAuthorizationsResponse = types.Response[ListMandateAuthorizationsResponseData]

func (c *Client) ListMandateAuthorizations(ctx context.Context, builder ListMandateAuthorizationsRequestBuilder) (*ListMandateAuthorizationsResponse, error) {
	req := builder.Build()
	path := basePath + "/mandate-authorizations"

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[ListMandateAuthorizationsResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
