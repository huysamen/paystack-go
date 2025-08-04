package subaccounts

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SubaccountUpdateRequest struct {
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

type SubaccountUpdateRequestBuilder struct {
	req *SubaccountUpdateRequest
}

func NewSubaccountUpdateRequest() *SubaccountUpdateRequestBuilder {
	return &SubaccountUpdateRequestBuilder{
		req: &SubaccountUpdateRequest{},
	}
}

func (b *SubaccountUpdateRequestBuilder) BusinessName(businessName string) *SubaccountUpdateRequestBuilder {
	b.req.BusinessName = &businessName

	return b
}

func (b *SubaccountUpdateRequestBuilder) BankCode(bankCode string) *SubaccountUpdateRequestBuilder {
	b.req.BankCode = &bankCode

	return b
}

func (b *SubaccountUpdateRequestBuilder) AccountNumber(accountNumber string) *SubaccountUpdateRequestBuilder {
	b.req.AccountNumber = &accountNumber

	return b
}

func (b *SubaccountUpdateRequestBuilder) PercentageCharge(percentageCharge float64) *SubaccountUpdateRequestBuilder {
	b.req.PercentageCharge = &percentageCharge

	return b
}

func (b *SubaccountUpdateRequestBuilder) Description(description string) *SubaccountUpdateRequestBuilder {
	b.req.Description = &description

	return b
}

func (b *SubaccountUpdateRequestBuilder) PrimaryContactEmail(email string) *SubaccountUpdateRequestBuilder {
	b.req.PrimaryContactEmail = &email

	return b
}

func (b *SubaccountUpdateRequestBuilder) PrimaryContactName(name string) *SubaccountUpdateRequestBuilder {
	b.req.PrimaryContactName = &name

	return b
}

func (b *SubaccountUpdateRequestBuilder) PrimaryContactPhone(phone string) *SubaccountUpdateRequestBuilder {
	b.req.PrimaryContactPhone = &phone

	return b
}

func (b *SubaccountUpdateRequestBuilder) Active(active bool) *SubaccountUpdateRequestBuilder {
	b.req.Active = &active

	return b
}

func (b *SubaccountUpdateRequestBuilder) Metadata(metadata map[string]any) *SubaccountUpdateRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *SubaccountUpdateRequestBuilder) Build() *SubaccountUpdateRequest {
	return b.req
}

type SubaccountUpdateResponse = types.Response[types.Subaccount]

func (c *Client) Update(ctx context.Context, idOrCode string, builder *SubaccountUpdateRequestBuilder) (*SubaccountUpdateResponse, error) {
	return net.Put[SubaccountUpdateRequest, types.Subaccount](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), builder.Build(), c.BaseURL)
}
