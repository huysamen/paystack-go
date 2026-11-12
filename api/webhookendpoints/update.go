package webhookendpoints

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type updateRequest struct {
	Name    *string  `json:"name,omitempty"`    // Optional: label
	URL     *string  `json:"url,omitempty"`     // Optional: https URL
	Events  []string `json:"events,omitempty"`  // Optional: replaces the whole list; can't be empty
	Enabled *bool    `json:"enabled,omitempty"` // Optional: false pauses, true resumes
}

type UpdateRequestBuilder struct {
	req *updateRequest
}

func NewUpdateRequestBuilder() *UpdateRequestBuilder {
	return &UpdateRequestBuilder{
		req: &updateRequest{},
	}
}

func (b *UpdateRequestBuilder) Name(name string) *UpdateRequestBuilder {
	b.req.Name = &name

	return b
}

func (b *UpdateRequestBuilder) URL(url string) *UpdateRequestBuilder {
	b.req.URL = &url

	return b
}

// Events replaces the endpoint's events.
func (b *UpdateRequestBuilder) Events(events []string) *UpdateRequestBuilder {
	b.req.Events = events

	return b
}

func (b *UpdateRequestBuilder) Enabled(enabled bool) *UpdateRequestBuilder {
	b.req.Enabled = &enabled

	return b
}

func (b *UpdateRequestBuilder) Build() *updateRequest {
	return b.req
}

type UpdateResponse = types.Response[WriteResponseData]

// Update changes an endpoint's name, URL, events or status. Send at least one.
func (c *Client) Update(ctx context.Context, code string, builder UpdateRequestBuilder) (*UpdateResponse, error) {
	return net.Patch[updateRequest, WriteResponseData](ctx, c.Client, c.Secret, endpointPath(code), builder.Build(), c.BaseURL)
}
