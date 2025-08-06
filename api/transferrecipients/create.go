package transferrecipients

import (
	"context"

	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type createRequest struct {
	Type              enums.TransferRecipientType `json:"type"`                         // Required: nuban, ghipss, mobile_money, basa
	Name              string                      `json:"name"`                         // Required: recipient's name
	AccountNumber     string                      `json:"account_number"`               // Required for all types except authorization
	BankCode          string                      `json:"bank_code"`                    // Required for all types except authorization
	Description       *string                     `json:"description,omitempty"`        // Optional: description
	Currency          *string                     `json:"currency,omitempty"`           // Optional: currency
	AuthorizationCode *string                     `json:"authorization_code,omitempty"` // Optional: authorization code
	Metadata          map[string]any              `json:"metadata,omitempty"`           // Optional: additional data
}

type CreateRequestBuilder struct {
	req *createRequest
}

func NewCreateRequestBuilder(recipientType enums.TransferRecipientType, name, accountNumber, bankCode string) *CreateRequestBuilder {
	return &CreateRequestBuilder{
		req: &createRequest{
			Type:          recipientType,
			Name:          name,
			AccountNumber: accountNumber,
			BankCode:      bankCode,
		},
	}
}

func (b *CreateRequestBuilder) Description(description string) *CreateRequestBuilder {
	b.req.Description = &description

	return b
}

func (b *CreateRequestBuilder) Currency(currency string) *CreateRequestBuilder {
	b.req.Currency = &currency

	return b
}

func (b *CreateRequestBuilder) AuthorizationCode(authCode string) *CreateRequestBuilder {
	b.req.AuthorizationCode = &authCode

	return b
}

func (b *CreateRequestBuilder) Metadata(metadata map[string]any) *CreateRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *CreateRequestBuilder) Build() *createRequest {
	return b.req
}

type CreateResponseData = types.TransferRecipient
type CreateResponse = types.Response[CreateResponseData]

func (c *Client) Create(ctx context.Context, builder CreateRequestBuilder) (*CreateResponse, error) {
	return net.Post[createRequest, CreateResponseData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
