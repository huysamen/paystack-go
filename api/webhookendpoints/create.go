package webhookendpoints

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type createRequest struct {
	Name   string   `json:"name"`   // Required: label, up to 120 characters
	URL    string   `json:"url"`    // Required: https URL, up to 2,048 characters
	Events []string `json:"events"` // Required: event names, or ["*"] for the whole catalog
}

type CreateRequestBuilder struct {
	req *createRequest
}

func NewCreateRequestBuilder(name, url string, events []string) *CreateRequestBuilder {
	return &CreateRequestBuilder{
		req: &createRequest{Name: name, URL: url, Events: events},
	}
}

func (b *CreateRequestBuilder) Build() *createRequest {
	return b.req
}

type CreateResponse = types.Response[WriteResponseData]

// Create creates an endpoint. Paystack answers 409 when another endpoint in
// the same environment has the URL (scheme and host compared ignoring case).
func (c *Client) Create(ctx context.Context, builder CreateRequestBuilder) (*CreateResponse, error) {
	return net.Post[createRequest, WriteResponseData](ctx, c.Client, c.Secret, basePath+"/endpoints", builder.Build(), c.BaseURL)
}
