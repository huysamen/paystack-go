package subaccounts

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SubaccountCreateRequest struct {
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

type SubaccountCreateRequestBuilder struct {
	req *SubaccountCreateRequest
}

func NewSubaccountCreateRequest(businessName, bankCode, accountNumber string, percentageCharge float64) *SubaccountCreateRequestBuilder {
	return &SubaccountCreateRequestBuilder{
		req: &SubaccountCreateRequest{
			BusinessName:     businessName,
			BankCode:         bankCode,
			AccountNumber:    accountNumber,
			PercentageCharge: percentageCharge,
		},
	}
}

func (b *SubaccountCreateRequestBuilder) Description(description string) *SubaccountCreateRequestBuilder {
	b.req.Description = &description

	return b
}

func (b *SubaccountCreateRequestBuilder) PrimaryContactEmail(email string) *SubaccountCreateRequestBuilder {
	b.req.PrimaryContactEmail = &email

	return b
}

func (b *SubaccountCreateRequestBuilder) PrimaryContactName(name string) *SubaccountCreateRequestBuilder {
	b.req.PrimaryContactName = &name

	return b
}

func (b *SubaccountCreateRequestBuilder) PrimaryContactPhone(phone string) *SubaccountCreateRequestBuilder {
	b.req.PrimaryContactPhone = &phone

	return b
}

func (b *SubaccountCreateRequestBuilder) Metadata(metadata map[string]any) *SubaccountCreateRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *SubaccountCreateRequestBuilder) Build() *SubaccountCreateRequest {
	return b.req
}

type SubaccountCreateResponse = types.Response[types.Subaccount]

func (c *Client) Create(ctx context.Context, builder *SubaccountCreateRequestBuilder) (*SubaccountCreateResponse, error) {
	return net.Post[SubaccountCreateRequest, types.Subaccount](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
