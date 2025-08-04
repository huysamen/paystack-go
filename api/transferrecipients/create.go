package transferrecipients

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// TransferRecipientCreateRequest represents the request to create a transfer recipient
type TransferRecipientCreateRequest struct {
	Type              types.TransferRecipientType `json:"type"`                         // Required: nuban, ghipss, mobile_money, basa
	Name              string                      `json:"name"`                         // Required: recipient's name
	AccountNumber     string                      `json:"account_number"`               // Required for all types except authorization
	BankCode          string                      `json:"bank_code"`                    // Required for all types except authorization
	Description       *string                     `json:"description,omitempty"`        // Optional: description
	Currency          *string                     `json:"currency,omitempty"`           // Optional: currency
	AuthorizationCode *string                     `json:"authorization_code,omitempty"` // Optional: authorization code
	Metadata          map[string]any              `json:"metadata,omitempty"`           // Optional: additional data
}

// TransferRecipientCreateRequestBuilder provides a fluent interface for building TransferRecipientCreateRequest
type TransferRecipientCreateRequestBuilder struct {
	req *TransferRecipientCreateRequest
}

// NewTransferRecipientCreateRequest creates a new builder for TransferRecipientCreateRequest
func NewTransferRecipientCreateRequest(recipientType types.TransferRecipientType, name, accountNumber, bankCode string) *TransferRecipientCreateRequestBuilder {
	return &TransferRecipientCreateRequestBuilder{
		req: &TransferRecipientCreateRequest{
			Type:          recipientType,
			Name:          name,
			AccountNumber: accountNumber,
			BankCode:      bankCode,
		},
	}
}

// Description sets the recipient description
func (b *TransferRecipientCreateRequestBuilder) Description(description string) *TransferRecipientCreateRequestBuilder {
	b.req.Description = &description

	return b
}

// Currency sets the currency
func (b *TransferRecipientCreateRequestBuilder) Currency(currency string) *TransferRecipientCreateRequestBuilder {
	b.req.Currency = &currency

	return b
}

// AuthorizationCode sets the authorization code
func (b *TransferRecipientCreateRequestBuilder) AuthorizationCode(authCode string) *TransferRecipientCreateRequestBuilder {
	b.req.AuthorizationCode = &authCode

	return b
}

// Metadata sets the recipient metadata
func (b *TransferRecipientCreateRequestBuilder) Metadata(metadata map[string]any) *TransferRecipientCreateRequestBuilder {
	b.req.Metadata = metadata

	return b
}

// Build returns the constructed TransferRecipientCreateRequest
func (b *TransferRecipientCreateRequestBuilder) Build() *TransferRecipientCreateRequest {
	return b.req
}

// TransferRecipientCreateResponse represents the response from creating a transfer recipient
type TransferRecipientCreateResponse = types.Response[types.TransferRecipient]

// Create creates a new transfer recipient
func (c *Client) Create(ctx context.Context, builder *TransferRecipientCreateRequestBuilder) (*TransferRecipientCreateResponse, error) {
	req := builder.Build()
	return net.Post[TransferRecipientCreateRequest, types.TransferRecipient](ctx, c.Client, c.Secret, basePath, req, c.BaseURL)
}
