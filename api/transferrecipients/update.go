package transferrecipients

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransferRecipientUpdateRequest struct {
	Name  string  `json:"name"`            // Required: recipient name
	Email *string `json:"email,omitempty"` // Optional: email address
}

type TransferRecipientUpdateRequestBuilder struct {
	req *TransferRecipientUpdateRequest
}

func NewTransferRecipientUpdateRequest(name string) *TransferRecipientUpdateRequestBuilder {
	return &TransferRecipientUpdateRequestBuilder{
		req: &TransferRecipientUpdateRequest{
			Name: name,
		},
	}
}

func (b *TransferRecipientUpdateRequestBuilder) Email(email string) *TransferRecipientUpdateRequestBuilder {
	b.req.Email = &email

	return b
}

func (b *TransferRecipientUpdateRequestBuilder) Build() *TransferRecipientUpdateRequest {
	return b.req
}

type TransferRecipientUpdateResponse = types.Response[types.TransferRecipient]

func (c *Client) Update(ctx context.Context, idOrCode string, builder *TransferRecipientUpdateRequestBuilder) (*TransferRecipientUpdateResponse, error) {
	req := builder.Build()
	return net.Put[TransferRecipientUpdateRequest, types.TransferRecipient](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), req, c.BaseURL)
}
