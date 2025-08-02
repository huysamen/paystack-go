package transactions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionInitializeRequest struct {
	// Required fields
	Amount int    `json:"amount"`
	Email  string `json:"email"`

	// Optional fields
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

type TransactionInitializeResponse struct {
	AuthorizationURL string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

// TransactionInitializeRequestBuilder builds a TransactionInitializeRequest
type TransactionInitializeRequestBuilder struct {
	request TransactionInitializeRequest
}

// NewTransactionInitializeRequestBuilder creates a new builder
func NewTransactionInitializeRequestBuilder() *TransactionInitializeRequestBuilder {
	return &TransactionInitializeRequestBuilder{}
}

// Amount sets the amount for the transaction
func (b *TransactionInitializeRequestBuilder) Amount(amount int) *TransactionInitializeRequestBuilder {
	b.request.Amount = amount
	return b
}

// Email sets the email for the transaction
func (b *TransactionInitializeRequestBuilder) Email(email string) *TransactionInitializeRequestBuilder {
	b.request.Email = email
	return b
}

// Currency sets the currency for the transaction
func (b *TransactionInitializeRequestBuilder) Currency(currency types.Currency) *TransactionInitializeRequestBuilder {
	b.request.Currency = currency
	return b
}

// Reference sets the reference for the transaction
func (b *TransactionInitializeRequestBuilder) Reference(reference string) *TransactionInitializeRequestBuilder {
	b.request.Reference = reference
	return b
}

// CallbackURL sets the callback URL for the transaction
func (b *TransactionInitializeRequestBuilder) CallbackURL(callbackURL string) *TransactionInitializeRequestBuilder {
	b.request.CallbackURL = callbackURL
	return b
}

// Plan sets the plan for the transaction
func (b *TransactionInitializeRequestBuilder) Plan(plan string) *TransactionInitializeRequestBuilder {
	b.request.Plan = plan
	return b
}

// InvoiceLimit sets the invoice limit for the transaction
func (b *TransactionInitializeRequestBuilder) InvoiceLimit(invoiceLimit int) *TransactionInitializeRequestBuilder {
	b.request.InvoiceLimit = invoiceLimit
	return b
}

// Metadata sets the metadata for the transaction
func (b *TransactionInitializeRequestBuilder) Metadata(metadata types.Metadata) *TransactionInitializeRequestBuilder {
	b.request.Metadata = metadata
	return b
}

// Channels sets the channels for the transaction
func (b *TransactionInitializeRequestBuilder) Channels(channels []types.Channel) *TransactionInitializeRequestBuilder {
	b.request.Channels = channels
	return b
}

// SplitCode sets the split codes for the transaction
func (b *TransactionInitializeRequestBuilder) SplitCode(splitCode []string) *TransactionInitializeRequestBuilder {
	b.request.SplitCode = splitCode
	return b
}

// Subaccount sets the subaccount for the transaction
func (b *TransactionInitializeRequestBuilder) Subaccount(subaccount string) *TransactionInitializeRequestBuilder {
	b.request.Subaccount = subaccount
	return b
}

// TransactionCharge sets the transaction charge
func (b *TransactionInitializeRequestBuilder) TransactionCharge(transactionCharge int) *TransactionInitializeRequestBuilder {
	b.request.TransactionCharge = transactionCharge
	return b
}

// Bearer sets the bearer for the transaction
func (b *TransactionInitializeRequestBuilder) Bearer(bearer types.Bearer) *TransactionInitializeRequestBuilder {
	b.request.Bearer = bearer
	return b
}

// Build returns the built TransactionInitializeRequest
func (b *TransactionInitializeRequestBuilder) Build() *TransactionInitializeRequest {
	return &b.request
}

func (c *Client) Initialize(ctx context.Context, builder *TransactionInitializeRequestBuilder) (*types.Response[TransactionInitializeResponse], error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	return net.Post[TransactionInitializeRequest, TransactionInitializeResponse](
		ctx,
		c.client,
		c.secret,
		fmt.Sprintf("%s%s", transactionBasePath, transactionInitializePath),
		req,
		c.baseURL,
	)
}
