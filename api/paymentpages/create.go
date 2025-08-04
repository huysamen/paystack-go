package paymentpages

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CustomField struct {
	DisplayName  string `json:"display_name"`
	VariableName string `json:"variable_name"`
	Required     bool   `json:"required"`
}

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

type CreatePaymentPageRequestBuilder struct {
	req *CreatePaymentPageRequest
}

func NewCreatePaymentPageRequest(name string) *CreatePaymentPageRequestBuilder {
	return &CreatePaymentPageRequestBuilder{
		req: &CreatePaymentPageRequest{
			Name: name,
		},
	}
}

func (b *CreatePaymentPageRequestBuilder) Description(description string) *CreatePaymentPageRequestBuilder {
	b.req.Description = description

	return b
}

func (b *CreatePaymentPageRequestBuilder) Amount(amount int) *CreatePaymentPageRequestBuilder {
	b.req.Amount = &amount

	return b
}

func (b *CreatePaymentPageRequestBuilder) Currency(currency string) *CreatePaymentPageRequestBuilder {
	b.req.Currency = currency

	return b
}

func (b *CreatePaymentPageRequestBuilder) Slug(slug string) *CreatePaymentPageRequestBuilder {
	b.req.Slug = slug

	return b
}

func (b *CreatePaymentPageRequestBuilder) Type(pageType string) *CreatePaymentPageRequestBuilder {
	b.req.Type = pageType

	return b
}

func (b *CreatePaymentPageRequestBuilder) Plan(plan string) *CreatePaymentPageRequestBuilder {
	b.req.Plan = plan

	return b
}

func (b *CreatePaymentPageRequestBuilder) FixedAmount(fixed bool) *CreatePaymentPageRequestBuilder {
	b.req.FixedAmount = &fixed

	return b
}

func (b *CreatePaymentPageRequestBuilder) SplitCode(splitCode string) *CreatePaymentPageRequestBuilder {
	b.req.SplitCode = splitCode

	return b
}

func (b *CreatePaymentPageRequestBuilder) Metadata(metadata *types.Metadata) *CreatePaymentPageRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *CreatePaymentPageRequestBuilder) RedirectURL(url string) *CreatePaymentPageRequestBuilder {
	b.req.RedirectURL = url

	return b
}

func (b *CreatePaymentPageRequestBuilder) SuccessMessage(message string) *CreatePaymentPageRequestBuilder {
	b.req.SuccessMessage = message

	return b
}

func (b *CreatePaymentPageRequestBuilder) NotificationEmail(email string) *CreatePaymentPageRequestBuilder {
	b.req.NotificationEmail = email

	return b
}

func (b *CreatePaymentPageRequestBuilder) CollectPhone(collect bool) *CreatePaymentPageRequestBuilder {
	b.req.CollectPhone = &collect

	return b
}

func (b *CreatePaymentPageRequestBuilder) CustomFields(fields []CustomField) *CreatePaymentPageRequestBuilder {
	b.req.CustomFields = fields

	return b
}

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

func (b *CreatePaymentPageRequestBuilder) Build() *CreatePaymentPageRequest {
	return b.req
}

type CreatePaymentPageResponse = types.Response[types.PaymentPage]

func (c *Client) Create(ctx context.Context, builder *CreatePaymentPageRequestBuilder) (*CreatePaymentPageResponse, error) {
	return net.Post[CreatePaymentPageRequest, types.PaymentPage](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
