package transferrecipients

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// TransferRecipientUpdateRequest represents the request to update a transfer recipient
type TransferRecipientUpdateRequest struct {
	Name  string  `json:"name"`            // Required: recipient name
	Email *string `json:"email,omitempty"` // Optional: email address
}

// TransferRecipientUpdateRequestBuilder provides a fluent interface for building TransferRecipientUpdateRequest
type TransferRecipientUpdateRequestBuilder struct {
	req *TransferRecipientUpdateRequest
}

// NewTransferRecipientUpdateRequest creates a new builder for TransferRecipientUpdateRequest
func NewTransferRecipientUpdateRequest(name string) *TransferRecipientUpdateRequestBuilder {
	return &TransferRecipientUpdateRequestBuilder{
		req: &TransferRecipientUpdateRequest{
			Name: name,
		},
	}
}

// Email sets the recipient email
func (b *TransferRecipientUpdateRequestBuilder) Email(email string) *TransferRecipientUpdateRequestBuilder {
	b.req.Email = &email
	return b
}

// Build returns the constructed TransferRecipientUpdateRequest
func (b *TransferRecipientUpdateRequestBuilder) Build() *TransferRecipientUpdateRequest {
	return b.req
}

// TransferRecipientUpdateResponse represents the response from updating a transfer recipient
type TransferRecipientUpdateResponse = types.Response[types.TransferRecipient]

// Update updates a transfer recipient
func (c *Client) Update(ctx context.Context, idOrCode string, builder *TransferRecipientUpdateRequestBuilder) (*TransferRecipientUpdateResponse, error) {
	req := builder.Build()
	return net.Put[TransferRecipientUpdateRequest, types.TransferRecipient](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), req, c.BaseURL)
}
