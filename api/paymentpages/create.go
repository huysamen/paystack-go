package paymentpages

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type createRequest struct {
	Name              string              `json:"name"`
	Description       string              `json:"description,omitempty"`
	Amount            *int                `json:"amount,omitempty"`
	Currency          string              `json:"currency,omitempty"`
	Slug              string              `json:"slug,omitempty"`
	Type              string              `json:"type,omitempty"`
	Plan              string              `json:"plan,omitempty"`
	FixedAmount       *bool               `json:"fixed_amount,omitempty"`
	SplitCode         string              `json:"split_code,omitempty"`
	Metadata          *types.Metadata     `json:"metadata,omitempty"`
	RedirectURL       string              `json:"redirect_url,omitempty"`
	SuccessMessage    string              `json:"success_message,omitempty"`
	NotificationEmail string              `json:"notification_email,omitempty"`
	CollectPhone      *bool               `json:"collect_phone,omitempty"`
	CustomFields      []types.CustomField `json:"custom_fields,omitempty"`
}

type CreateRequestBuilder struct {
	req *createRequest
}

func NewCreateRequestBuilder(name string) *CreateRequestBuilder {
	return &CreateRequestBuilder{
		req: &createRequest{
			Name: name,
		},
	}
}

func (b *CreateRequestBuilder) Description(description string) *CreateRequestBuilder {
	b.req.Description = description

	return b
}

func (b *CreateRequestBuilder) Amount(amount int) *CreateRequestBuilder {
	b.req.Amount = &amount

	return b
}

func (b *CreateRequestBuilder) Currency(currency string) *CreateRequestBuilder {
	b.req.Currency = currency

	return b
}

func (b *CreateRequestBuilder) Slug(slug string) *CreateRequestBuilder {
	b.req.Slug = slug

	return b
}

func (b *CreateRequestBuilder) Type(pageType string) *CreateRequestBuilder {
	b.req.Type = pageType

	return b
}

func (b *CreateRequestBuilder) Plan(plan string) *CreateRequestBuilder {
	b.req.Plan = plan

	return b
}

func (b *CreateRequestBuilder) FixedAmount(fixed bool) *CreateRequestBuilder {
	b.req.FixedAmount = &fixed

	return b
}

func (b *CreateRequestBuilder) SplitCode(splitCode string) *CreateRequestBuilder {
	b.req.SplitCode = splitCode

	return b
}

func (b *CreateRequestBuilder) Metadata(metadata *types.Metadata) *CreateRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *CreateRequestBuilder) RedirectURL(url string) *CreateRequestBuilder {
	b.req.RedirectURL = url

	return b
}

func (b *CreateRequestBuilder) SuccessMessage(message string) *CreateRequestBuilder {
	b.req.SuccessMessage = message

	return b
}

func (b *CreateRequestBuilder) NotificationEmail(email string) *CreateRequestBuilder {
	b.req.NotificationEmail = email

	return b
}

func (b *CreateRequestBuilder) CollectPhone(collect bool) *CreateRequestBuilder {
	b.req.CollectPhone = &collect

	return b
}

func (b *CreateRequestBuilder) CustomFields(fields []types.CustomField) *CreateRequestBuilder {
	b.req.CustomFields = fields

	return b
}

func (b *CreateRequestBuilder) AddCustomField(displayName, variableName string, required bool) *CreateRequestBuilder {
	if b.req.CustomFields == nil {
		b.req.CustomFields = []types.CustomField{}
	}

	b.req.CustomFields = append(b.req.CustomFields, types.CustomField{
		DisplayName:  displayName,
		VariableName: variableName,
		Required:     required,
	})

	return b
}

func (b *CreateRequestBuilder) Build() *createRequest {
	return b.req
}

type CreateResponseData = types.PaymentPage
type CreateResponse = types.Response[CreateResponseData]

func (c *Client) Create(ctx context.Context, builder CreateRequestBuilder) (*CreateResponse, error) {
	return net.Post[createRequest, CreateResponseData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
