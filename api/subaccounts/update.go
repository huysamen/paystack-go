package subaccounts

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// SubaccountUpdateRequestBuilder provides a fluent interface for building SubaccountUpdateRequest
type SubaccountUpdateRequestBuilder struct {
	req *SubaccountUpdateRequest
}

// NewSubaccountUpdateRequest creates a new builder for SubaccountUpdateRequest
func NewSubaccountUpdateRequest() *SubaccountUpdateRequestBuilder {
	return &SubaccountUpdateRequestBuilder{
		req: &SubaccountUpdateRequest{},
	}
}

// BusinessName sets the business name
func (b *SubaccountUpdateRequestBuilder) BusinessName(businessName string) *SubaccountUpdateRequestBuilder {
	b.req.BusinessName = &businessName
	return b
}

// BankCode sets the bank code
func (b *SubaccountUpdateRequestBuilder) BankCode(bankCode string) *SubaccountUpdateRequestBuilder {
	b.req.BankCode = &bankCode
	return b
}

// AccountNumber sets the account number
func (b *SubaccountUpdateRequestBuilder) AccountNumber(accountNumber string) *SubaccountUpdateRequestBuilder {
	b.req.AccountNumber = &accountNumber
	return b
}

// PercentageCharge sets the percentage charge
func (b *SubaccountUpdateRequestBuilder) PercentageCharge(percentageCharge float64) *SubaccountUpdateRequestBuilder {
	b.req.PercentageCharge = &percentageCharge
	return b
}

// Description sets the description
func (b *SubaccountUpdateRequestBuilder) Description(description string) *SubaccountUpdateRequestBuilder {
	b.req.Description = &description
	return b
}

// PrimaryContactEmail sets the primary contact email
func (b *SubaccountUpdateRequestBuilder) PrimaryContactEmail(email string) *SubaccountUpdateRequestBuilder {
	b.req.PrimaryContactEmail = &email
	return b
}

// PrimaryContactName sets the primary contact name
func (b *SubaccountUpdateRequestBuilder) PrimaryContactName(name string) *SubaccountUpdateRequestBuilder {
	b.req.PrimaryContactName = &name
	return b
}

// PrimaryContactPhone sets the primary contact phone
func (b *SubaccountUpdateRequestBuilder) PrimaryContactPhone(phone string) *SubaccountUpdateRequestBuilder {
	b.req.PrimaryContactPhone = &phone
	return b
}

// Active sets the active status
func (b *SubaccountUpdateRequestBuilder) Active(active bool) *SubaccountUpdateRequestBuilder {
	b.req.Active = &active
	return b
}

// Metadata sets the metadata
func (b *SubaccountUpdateRequestBuilder) Metadata(metadata map[string]any) *SubaccountUpdateRequestBuilder {
	b.req.Metadata = metadata
	return b
}

// Build returns the constructed SubaccountUpdateRequest
func (b *SubaccountUpdateRequestBuilder) Build() *SubaccountUpdateRequest {
	return b.req
}

// Update updates an existing subaccount using the builder pattern
func (c *Client) Update(ctx context.Context, idOrCode string, builder *SubaccountUpdateRequestBuilder) (*SubaccountUpdateResponse, error) {
	return net.Put[SubaccountUpdateRequest, Subaccount](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, idOrCode), builder.Build(), c.BaseURL)
}
