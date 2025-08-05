package transferrecipients

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type updateRequest struct {
	Name  string  `json:"name"`            // Required: recipient name
	Email *string `json:"email,omitempty"` // Optional: email address
}

type UpdateRequestBuilder struct {
	req *updateRequest
}

func NewUpdateRequestBuilder(name string) *UpdateRequestBuilder {
	return &UpdateRequestBuilder{
		req: &updateRequest{
			Name: name,
		},
	}
}

func (b *UpdateRequestBuilder) Email(email string) *UpdateRequestBuilder {
	b.req.Email = &email

	return b
}

func (b *UpdateRequestBuilder) Build() *updateRequest {
	return b.req
}

type UpdateResponseData = types.TransferRecipient
type UpdateResponse = types.Response[UpdateResponseData]

func (c *Client) Update(ctx context.Context, idOrCode string, builder UpdateRequestBuilder) (*UpdateResponse, error) {
	return net.Put[updateRequest, UpdateResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), builder.Build(), c.BaseURL)
}
