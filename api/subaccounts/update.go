package subaccounts

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type updateRequest struct {
	BusinessName        *string        `json:"business_name,omitempty"`         // Optional: Business name
	BankCode            *string        `json:"settlement_bank,omitempty"`       // Optional: Bank code
	AccountNumber       *string        `json:"account_number,omitempty"`        // Optional: Account number
	PercentageCharge    *float64       `json:"percentage_charge,omitempty"`     // Optional: Percentage charge
	Description         *string        `json:"description,omitempty"`           // Optional: Description
	PrimaryContactEmail *string        `json:"primary_contact_email,omitempty"` // Optional: Primary contact email
	PrimaryContactName  *string        `json:"primary_contact_name,omitempty"`  // Optional: Primary contact name
	PrimaryContactPhone *string        `json:"primary_contact_phone,omitempty"` // Optional: Primary contact phone
	Active              *bool          `json:"active,omitempty"`                // Optional: Active status
	Metadata            map[string]any `json:"metadata,omitempty"`              // Optional: Metadata
}

type UpdateRequestBuilder struct {
	req *updateRequest
}

func NewUpdateRequestBuilder() *UpdateRequestBuilder {
	return &UpdateRequestBuilder{
		req: &updateRequest{},
	}
}

func (b *UpdateRequestBuilder) BusinessName(businessName string) *UpdateRequestBuilder {
	b.req.BusinessName = &businessName

	return b
}

func (b *UpdateRequestBuilder) BankCode(bankCode string) *UpdateRequestBuilder {
	b.req.BankCode = &bankCode

	return b
}

func (b *UpdateRequestBuilder) AccountNumber(accountNumber string) *UpdateRequestBuilder {
	b.req.AccountNumber = &accountNumber

	return b
}

func (b *UpdateRequestBuilder) PercentageCharge(percentageCharge float64) *UpdateRequestBuilder {
	b.req.PercentageCharge = &percentageCharge

	return b
}

func (b *UpdateRequestBuilder) Description(description string) *UpdateRequestBuilder {
	b.req.Description = &description

	return b
}

func (b *UpdateRequestBuilder) PrimaryContactEmail(email string) *UpdateRequestBuilder {
	b.req.PrimaryContactEmail = &email

	return b
}

func (b *UpdateRequestBuilder) PrimaryContactName(name string) *UpdateRequestBuilder {
	b.req.PrimaryContactName = &name

	return b
}

func (b *UpdateRequestBuilder) PrimaryContactPhone(phone string) *UpdateRequestBuilder {
	b.req.PrimaryContactPhone = &phone

	return b
}

func (b *UpdateRequestBuilder) Active(active bool) *UpdateRequestBuilder {
	b.req.Active = &active

	return b
}

func (b *UpdateRequestBuilder) Metadata(metadata map[string]any) *UpdateRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *UpdateRequestBuilder) Build() *updateRequest {
	return b.req
}

type UpdateResponseData = types.Subaccount
type UpdateResponse = types.Response[UpdateResponseData]

func (c *Client) Update(ctx context.Context, idOrCode string, builder UpdateRequestBuilder) (*UpdateResponse, error) {
	return net.Put[updateRequest, UpdateResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), builder.Build(), c.BaseURL)
}
