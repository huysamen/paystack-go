package transferrecipients

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

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

type TransferRecipientCreateRequestBuilder struct {
	req *TransferRecipientCreateRequest
}

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

func (b *TransferRecipientCreateRequestBuilder) Description(description string) *TransferRecipientCreateRequestBuilder {
	b.req.Description = &description

	return b
}

func (b *TransferRecipientCreateRequestBuilder) Currency(currency string) *TransferRecipientCreateRequestBuilder {
	b.req.Currency = &currency

	return b
}

func (b *TransferRecipientCreateRequestBuilder) AuthorizationCode(authCode string) *TransferRecipientCreateRequestBuilder {
	b.req.AuthorizationCode = &authCode

	return b
}

func (b *TransferRecipientCreateRequestBuilder) Metadata(metadata map[string]any) *TransferRecipientCreateRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *TransferRecipientCreateRequestBuilder) Build() *TransferRecipientCreateRequest {
	return b.req
}

type TransferRecipientCreateResponse = types.Response[types.TransferRecipient]

func (c *Client) Create(ctx context.Context, builder *TransferRecipientCreateRequestBuilder) (*TransferRecipientCreateResponse, error) {
	req := builder.Build()
	return net.Post[TransferRecipientCreateRequest, types.TransferRecipient](ctx, c.Client, c.Secret, basePath, req, c.BaseURL)
}
