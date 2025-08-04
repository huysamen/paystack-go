package transactions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionInitializeRequest struct {
	Amount int    `json:"amount"`
	Email  string `json:"email"`

	Currency          types.Currency  `json:"currency,omitempty"`
	Reference         string          `json:"reference,omitempty"`
	CallbackURL       string          `json:"callback_url,omitempty"`
	Plan              string          `json:"plan,omitempty"`
	InvoiceLimit      int             `json:"invoice_limit,omitempty"`
	Metadata          types.Metadata  `json:"metadata,omitempty"`
	Channels          []types.Channel `json:"channels,omitempty"`
	SplitCode         []string        `json:"split_code,omitempty"`
	Subaccount        string          `json:"subaccount,omitempty"`
	TransactionCharge int             `json:"transaction_charge,omitempty"`
	Bearer            types.Bearer    `json:"bearer,omitempty"`
}

type TransactionInitializeRequestBuilder struct {
	request TransactionInitializeRequest
}

func NewTransactionInitializeRequestBuilder() *TransactionInitializeRequestBuilder {
	return &TransactionInitializeRequestBuilder{}
}

func (b *TransactionInitializeRequestBuilder) Amount(amount int) *TransactionInitializeRequestBuilder {
	b.request.Amount = amount

	return b
}

func (b *TransactionInitializeRequestBuilder) Email(email string) *TransactionInitializeRequestBuilder {
	b.request.Email = email

	return b
}

func (b *TransactionInitializeRequestBuilder) Currency(currency types.Currency) *TransactionInitializeRequestBuilder {
	b.request.Currency = currency

	return b
}

func (b *TransactionInitializeRequestBuilder) Reference(reference string) *TransactionInitializeRequestBuilder {
	b.request.Reference = reference

	return b
}

func (b *TransactionInitializeRequestBuilder) CallbackURL(callbackURL string) *TransactionInitializeRequestBuilder {
	b.request.CallbackURL = callbackURL

	return b
}

func (b *TransactionInitializeRequestBuilder) Plan(plan string) *TransactionInitializeRequestBuilder {
	b.request.Plan = plan

	return b
}

func (b *TransactionInitializeRequestBuilder) InvoiceLimit(invoiceLimit int) *TransactionInitializeRequestBuilder {
	b.request.InvoiceLimit = invoiceLimit

	return b
}

func (b *TransactionInitializeRequestBuilder) Metadata(metadata types.Metadata) *TransactionInitializeRequestBuilder {
	b.request.Metadata = metadata

	return b
}

func (b *TransactionInitializeRequestBuilder) Channels(channels []types.Channel) *TransactionInitializeRequestBuilder {
	b.request.Channels = channels

	return b
}

func (b *TransactionInitializeRequestBuilder) SplitCode(splitCode []string) *TransactionInitializeRequestBuilder {
	b.request.SplitCode = splitCode

	return b
}

func (b *TransactionInitializeRequestBuilder) Subaccount(subaccount string) *TransactionInitializeRequestBuilder {
	b.request.Subaccount = subaccount

	return b
}

func (b *TransactionInitializeRequestBuilder) TransactionCharge(transactionCharge int) *TransactionInitializeRequestBuilder {
	b.request.TransactionCharge = transactionCharge

	return b
}

func (b *TransactionInitializeRequestBuilder) Bearer(bearer types.Bearer) *TransactionInitializeRequestBuilder {
	b.request.Bearer = bearer

	return b
}

func (b *TransactionInitializeRequestBuilder) Build() *TransactionInitializeRequest {
	return &b.request
}

type InitializeResponseData struct {
	AuthorizationURL string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

type InitializeResponse = types.Response[InitializeResponseData]

func (c *Client) Initialize(ctx context.Context, builder *TransactionInitializeRequestBuilder) (*InitializeResponse, error) {
	req := builder.Build()
	return net.Post[TransactionInitializeRequest, InitializeResponseData](
		ctx, c.Client, c.Secret, fmt.Sprintf("%s%s", basePath, transactionInitializePath), req, c.BaseURL,
	)
}
