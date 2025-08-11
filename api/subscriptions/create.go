package subscriptions

import (
	"context"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
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

// When creating, API returns numeric IDs for customer and plan; use a narrow shape
type CreateResponseData struct {
	Customer         data.Int            `json:"customer"`
	Plan             data.Int            `json:"plan"`
	Integration      data.Int            `json:"integration"`
	Domain           data.String         `json:"domain"`
	Start            data.Int            `json:"start"`
	Status           data.String         `json:"status"`
	Quantity         data.Int            `json:"quantity"`
	Amount           data.Int            `json:"amount"`
	Authorization    types.Authorization `json:"authorization"`
	SubscriptionCode data.String         `json:"subscription_code"`
	EmailToken       data.String         `json:"email_token"`
	ID               data.Int            `json:"id"`
	CreatedAt        data.Time           `json:"createdAt"`
	UpdatedAt        data.Time           `json:"updatedAt"`
}
type CreateResponse = types.Response[CreateResponseData]

func (c *Client) Create(ctx context.Context, builder CreateRequestBuilder) (*CreateResponse, error) {
	return net.Post[createRequest, CreateResponseData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
