package transactions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type initializeRequest struct {
	Amount int    `json:"amount"`
	Email  string `json:"email"`

	Currency          enums.Currency  `json:"currency,omitempty"`
	Reference         string          `json:"reference,omitempty"`
	CallbackURL       string          `json:"callback_url,omitempty"`
	Plan              string          `json:"plan,omitempty"`
	InvoiceLimit      int             `json:"invoice_limit,omitempty"`
	Metadata          types.Metadata  `json:"metadata,omitempty"`
	Channels          []enums.Channel `json:"channels,omitempty"`
	SplitCode         []string        `json:"split_code,omitempty"`
	Subaccount        string          `json:"subaccount,omitempty"`
	TransactionCharge int             `json:"transaction_charge,omitempty"`
	Bearer            enums.Bearer    `json:"bearer,omitempty"`
}

type InitializeRequestBuilder struct {
	request initializeRequest
}

func NewInitializeRequestBuilder() *InitializeRequestBuilder {
	return &InitializeRequestBuilder{}
}

func (b *InitializeRequestBuilder) Amount(amount int) *InitializeRequestBuilder {
	b.request.Amount = amount

	return b
}

func (b *InitializeRequestBuilder) Email(email string) *InitializeRequestBuilder {
	b.request.Email = email

	return b
}

func (b *InitializeRequestBuilder) Currency(currency enums.Currency) *InitializeRequestBuilder {
	b.request.Currency = currency

	return b
}

func (b *InitializeRequestBuilder) Reference(reference string) *InitializeRequestBuilder {
	b.request.Reference = reference

	return b
}

func (b *InitializeRequestBuilder) CallbackURL(callbackURL string) *InitializeRequestBuilder {
	b.request.CallbackURL = callbackURL

	return b
}

func (b *InitializeRequestBuilder) Plan(plan string) *InitializeRequestBuilder {
	b.request.Plan = plan

	return b
}

func (b *InitializeRequestBuilder) InvoiceLimit(invoiceLimit int) *InitializeRequestBuilder {
	b.request.InvoiceLimit = invoiceLimit

	return b
}

func (b *InitializeRequestBuilder) Metadata(metadata types.Metadata) *InitializeRequestBuilder {
	b.request.Metadata = metadata

	return b
}

func (b *InitializeRequestBuilder) Channels(channels []enums.Channel) *InitializeRequestBuilder {
	b.request.Channels = channels

	return b
}

func (b *InitializeRequestBuilder) SplitCode(splitCode []string) *InitializeRequestBuilder {
	b.request.SplitCode = splitCode

	return b
}

func (b *InitializeRequestBuilder) Subaccount(subaccount string) *InitializeRequestBuilder {
	b.request.Subaccount = subaccount

	return b
}

func (b *InitializeRequestBuilder) TransactionCharge(transactionCharge int) *InitializeRequestBuilder {
	b.request.TransactionCharge = transactionCharge

	return b
}

func (b *InitializeRequestBuilder) Bearer(bearer enums.Bearer) *InitializeRequestBuilder {
	b.request.Bearer = bearer

	return b
}

func (b *InitializeRequestBuilder) Build() *initializeRequest {
	return &b.request
}

type InitializeResponseData struct {
	AuthorizationURL string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

type InitializeResponse = types.Response[InitializeResponseData]

func (c *Client) Initialize(ctx context.Context, builder InitializeRequestBuilder) (*InitializeResponse, error) {
	return net.Post[initializeRequest, InitializeResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s%s", basePath, transactionInitializePath), builder.Build(), c.BaseURL)
}
