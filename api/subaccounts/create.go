package subaccounts

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type createRequest struct {
	BusinessName        string         `json:"business_name"`                   // Required: Name of business
	BankCode            string         `json:"settlement_bank"`                 // Required: Bank Code (use settlement_bank as per API docs)
	AccountNumber       string         `json:"account_number"`                  // Required: Bank Account Number
	PercentageCharge    float64        `json:"percentage_charge"`               // Required: Percentage the main account receives
	Description         *string        `json:"description,omitempty"`           // Optional: Description
	PrimaryContactEmail *string        `json:"primary_contact_email,omitempty"` // Optional: Contact email
	PrimaryContactName  *string        `json:"primary_contact_name,omitempty"`  // Optional: Contact name
	PrimaryContactPhone *string        `json:"primary_contact_phone,omitempty"` // Optional: Contact phone
	Metadata            map[string]any `json:"metadata,omitempty"`              // Optional: Additional data
}

type CreateRequestBuilder struct {
	req *createRequest
}

func NewCreateRequestBuilder(businessName, bankCode, accountNumber string, percentageCharge float64) *CreateRequestBuilder {
	return &CreateRequestBuilder{
		req: &createRequest{
			BusinessName:     businessName,
			BankCode:         bankCode,
			AccountNumber:    accountNumber,
			PercentageCharge: percentageCharge,
		},
	}
}

func (b *CreateRequestBuilder) Description(description string) *CreateRequestBuilder {
	b.req.Description = &description

	return b
}

func (b *CreateRequestBuilder) PrimaryContactEmail(email string) *CreateRequestBuilder {
	b.req.PrimaryContactEmail = &email

	return b
}

func (b *CreateRequestBuilder) PrimaryContactName(name string) *CreateRequestBuilder {
	b.req.PrimaryContactName = &name

	return b
}

func (b *CreateRequestBuilder) PrimaryContactPhone(phone string) *CreateRequestBuilder {
	b.req.PrimaryContactPhone = &phone

	return b
}

func (b *CreateRequestBuilder) Metadata(metadata map[string]any) *CreateRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *CreateRequestBuilder) Build() *createRequest {
	return b.req
}

type CreateResponseData = types.Subaccount
type CreateResponse = types.Response[CreateResponseData]

func (c *Client) Create(ctx context.Context, builder CreateRequestBuilder) (*CreateResponse, error) {
	return net.Post[createRequest, CreateResponseData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
