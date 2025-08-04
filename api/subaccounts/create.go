package subaccounts

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// SubaccountCreateRequest represents the request to create a subaccount
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

// SubaccountCreateResponse represents the response from creating a subaccount
type SubaccountCreateResponse = types.Response[types.Subaccount]

// Create creates a new subaccount using the builder pattern
func (c *Client) Create(ctx context.Context, builder *SubaccountCreateRequestBuilder) (*SubaccountCreateResponse, error) {
	return net.Post[SubaccountCreateRequest, types.Subaccount](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
