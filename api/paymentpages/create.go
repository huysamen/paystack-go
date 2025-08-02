package paymentpages

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// CustomField represents a custom field for payment pages
type CustomField struct {
	DisplayName  string `json:"display_name"`
	VariableName string `json:"variable_name"`
	Required     bool   `json:"required"`
}

// CreatePaymentPageRequest represents the request to create a payment page
type CreatePaymentPageRequest struct {
	Name              string          `json:"name"`
	Description       string          `json:"description,omitempty"`
	Amount            *int            `json:"amount,omitempty"`
	Currency          string          `json:"currency,omitempty"`
	Slug              string          `json:"slug,omitempty"`
	Type              string          `json:"type,omitempty"`
	Plan              string          `json:"plan,omitempty"`
	FixedAmount       *bool           `json:"fixed_amount,omitempty"`
	SplitCode         string          `json:"split_code,omitempty"`
	Metadata          *types.Metadata `json:"metadata,omitempty"`
	RedirectURL       string          `json:"redirect_url,omitempty"`
	SuccessMessage    string          `json:"success_message,omitempty"`
	NotificationEmail string          `json:"notification_email,omitempty"`
	CollectPhone      *bool           `json:"collect_phone,omitempty"`
	CustomFields      []CustomField   `json:"custom_fields,omitempty"`
}

// CreatePaymentPageRequestBuilder provides a fluent interface for building CreatePaymentPageRequest
type CreatePaymentPageRequestBuilder struct {
	req *CreatePaymentPageRequest
}

// NewCreatePaymentPageRequest creates a new builder for CreatePaymentPageRequest
func NewCreatePaymentPageRequest(name string) *CreatePaymentPageRequestBuilder {
	return &CreatePaymentPageRequestBuilder{
		req: &CreatePaymentPageRequest{
			Name: name,
		},
	}
}

// Description sets the description for the payment page
func (b *CreatePaymentPageRequestBuilder) Description(description string) *CreatePaymentPageRequestBuilder {
	b.req.Description = description
	return b
}

// Amount sets the amount for the payment page (in kobo/cents)
func (b *CreatePaymentPageRequestBuilder) Amount(amount int) *CreatePaymentPageRequestBuilder {
	b.req.Amount = &amount
	return b
}

// Currency sets the currency for the payment page
func (b *CreatePaymentPageRequestBuilder) Currency(currency string) *CreatePaymentPageRequestBuilder {
	b.req.Currency = currency
	return b
}

// Slug sets the custom slug for the payment page
func (b *CreatePaymentPageRequestBuilder) Slug(slug string) *CreatePaymentPageRequestBuilder {
	b.req.Slug = slug
	return b
}

// Type sets the type of payment page
func (b *CreatePaymentPageRequestBuilder) Type(pageType string) *CreatePaymentPageRequestBuilder {
	b.req.Type = pageType
	return b
}

// Plan sets the plan for subscription-based payment pages
func (b *CreatePaymentPageRequestBuilder) Plan(plan string) *CreatePaymentPageRequestBuilder {
	b.req.Plan = plan
	return b
}

// FixedAmount sets whether the payment page has a fixed amount
func (b *CreatePaymentPageRequestBuilder) FixedAmount(fixed bool) *CreatePaymentPageRequestBuilder {
	b.req.FixedAmount = &fixed
	return b
}

// SplitCode sets the split code for automatic payment splitting
func (b *CreatePaymentPageRequestBuilder) SplitCode(splitCode string) *CreatePaymentPageRequestBuilder {
	b.req.SplitCode = splitCode
	return b
}

// Metadata sets metadata for the payment page
func (b *CreatePaymentPageRequestBuilder) Metadata(metadata *types.Metadata) *CreatePaymentPageRequestBuilder {
	b.req.Metadata = metadata
	return b
}

// RedirectURL sets the redirect URL after successful payment
func (b *CreatePaymentPageRequestBuilder) RedirectURL(url string) *CreatePaymentPageRequestBuilder {
	b.req.RedirectURL = url
	return b
}

// SuccessMessage sets the success message after payment
func (b *CreatePaymentPageRequestBuilder) SuccessMessage(message string) *CreatePaymentPageRequestBuilder {
	b.req.SuccessMessage = message
	return b
}

// NotificationEmail sets the notification email for payments
func (b *CreatePaymentPageRequestBuilder) NotificationEmail(email string) *CreatePaymentPageRequestBuilder {
	b.req.NotificationEmail = email
	return b
}

// CollectPhone sets whether to collect phone numbers
func (b *CreatePaymentPageRequestBuilder) CollectPhone(collect bool) *CreatePaymentPageRequestBuilder {
	b.req.CollectPhone = &collect
	return b
}

// CustomFields sets custom fields for the payment page
func (b *CreatePaymentPageRequestBuilder) CustomFields(fields []CustomField) *CreatePaymentPageRequestBuilder {
	b.req.CustomFields = fields
	return b
}

// AddCustomField adds a custom field to the payment page
func (b *CreatePaymentPageRequestBuilder) AddCustomField(displayName, variableName string, required bool) *CreatePaymentPageRequestBuilder {
	if b.req.CustomFields == nil {
		b.req.CustomFields = []CustomField{}
	}
	b.req.CustomFields = append(b.req.CustomFields, CustomField{
		DisplayName:  displayName,
		VariableName: variableName,
		Required:     required,
	})
	return b
}

// Build returns the constructed CreatePaymentPageRequest
func (b *CreatePaymentPageRequestBuilder) Build() *CreatePaymentPageRequest {
	return b.req
}

// CreatePaymentPageResponse represents the response from creating a payment page
type CreatePaymentPageResponse = types.Response[PaymentPage]

// Create creates a new payment page using the builder pattern
func (c *Client) Create(ctx context.Context, builder *CreatePaymentPageRequestBuilder) (*CreatePaymentPageResponse, error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()

	resp, err := net.Post[CreatePaymentPageRequest, PaymentPage](
		ctx,
		c.client,
		c.secret,
		paymentPagesBasePath,
		req,
		c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
