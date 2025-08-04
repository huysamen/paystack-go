package subscriptions

import (
	"context"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type createRequest struct {
	Customer      string     `json:"customer"`                // Customer email or customer code
	Plan          string     `json:"plan"`                    // Plan code
	Authorization *string    `json:"authorization,omitempty"` // Authorization code if customer has multiple
	StartDate     *time.Time `json:"start_date,omitempty"`    // Date for first debit (ISO 8601)
}

type CreateRequestBuilder struct {
	req *createRequest
}

func NewCreateRequestBuilder(customer, plan string) *CreateRequestBuilder {
	return &CreateRequestBuilder{
		req: &createRequest{
			Customer: customer,
			Plan:     plan,
		},
	}
}

func (b *CreateRequestBuilder) Authorization(authorization string) *CreateRequestBuilder {
	b.req.Authorization = &authorization

	return b
}

func (b *CreateRequestBuilder) StartDate(startDate time.Time) *CreateRequestBuilder {
	b.req.StartDate = &startDate

	return b
}

func (b *CreateRequestBuilder) Build() *createRequest {
	return b.req
}

type CreateResponseData = types.Subscription
type CreateResponse = types.Response[CreateResponseData]

func (c *Client) Create(ctx context.Context, builder CreateRequestBuilder) (*CreateResponse, error) {
	return net.Post[createRequest, CreateResponseData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
