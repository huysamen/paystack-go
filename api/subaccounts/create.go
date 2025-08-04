package subaccounts

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// SubaccountCreateRequestBuilder provides a fluent interface for building SubaccountCreateRequest
type SubaccountCreateRequestBuilder struct {
	req *SubaccountCreateRequest
}

// NewSubaccountCreateRequest creates a new builder for SubaccountCreateRequest
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

// Description sets the subaccount description
func (b *SubaccountCreateRequestBuilder) Description(description string) *SubaccountCreateRequestBuilder {
	b.req.Description = &description
	return b
}

// PrimaryContactEmail sets the primary contact email
func (b *SubaccountCreateRequestBuilder) PrimaryContactEmail(email string) *SubaccountCreateRequestBuilder {
	b.req.PrimaryContactEmail = &email
	return b
}

// PrimaryContactName sets the primary contact name
func (b *SubaccountCreateRequestBuilder) PrimaryContactName(name string) *SubaccountCreateRequestBuilder {
	b.req.PrimaryContactName = &name
	return b
}

// PrimaryContactPhone sets the primary contact phone
func (b *SubaccountCreateRequestBuilder) PrimaryContactPhone(phone string) *SubaccountCreateRequestBuilder {
	b.req.PrimaryContactPhone = &phone
	return b
}

// Metadata sets the subaccount metadata
func (b *SubaccountCreateRequestBuilder) Metadata(metadata map[string]any) *SubaccountCreateRequestBuilder {
	b.req.Metadata = metadata
	return b
}

// Build returns the constructed SubaccountCreateRequest
func (b *SubaccountCreateRequestBuilder) Build() *SubaccountCreateRequest {
	return b.req
}

// Create creates a new subaccount using the builder pattern
func (c *Client) Create(ctx context.Context, builder *SubaccountCreateRequestBuilder) (*SubaccountCreateResponse, error) {
	return net.Post[SubaccountCreateRequest, types.Subaccount](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
