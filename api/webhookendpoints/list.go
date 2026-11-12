package webhookendpoints

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type listRequest struct {
	Limit    *int    // Optional: endpoints per page, 1 to 50 (default 50)
	Next     *string // Optional: cursor from a previous meta.next
	Previous *string // Optional: cursor from a previous meta.previous
}

type ListRequestBuilder struct {
	req *listRequest
}

func NewListRequestBuilder() *ListRequestBuilder {
	return &ListRequestBuilder{
		req: &listRequest{},
	}
}

func (b *ListRequestBuilder) Limit(limit int) *ListRequestBuilder {
	b.req.Limit = &limit

	return b
}

// Next asks for the page after a previous answer's meta.next. Don't combine
// with Previous.
func (b *ListRequestBuilder) Next(cursor string) *ListRequestBuilder {
	b.req.Next = &cursor

	return b
}

// Previous asks for the page before a previous answer's meta.previous. Don't
// combine with Next.
func (b *ListRequestBuilder) Previous(cursor string) *ListRequestBuilder {
	b.req.Previous = &cursor

	return b
}

func (b *ListRequestBuilder) Build() *listRequest {
	return b.req
}

func (r *listRequest) toQuery() string {
	params := url.Values{}
	if r.Limit != nil {
		params.Set("limit", strconv.Itoa(*r.Limit))
	}
	if r.Next != nil {
		params.Set("next", *r.Next)
	}
	if r.Previous != nil {
		params.Set("previous", *r.Previous)
	}

	return params.Encode()
}

type ListResponseData = []types.WebhookEndpoint
type ListResponse = types.Response[ListResponseData]

// List lists the integration's webhook endpoints, without their events. The
// cursors for other pages are on the answer's Meta (Next, Previous), null when
// there is no page in that direction.
func (c *Client) List(ctx context.Context, builder ListRequestBuilder) (*ListResponse, error) {
	path := basePath + "/endpoints"

	if req := builder.Build(); req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[ListResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
