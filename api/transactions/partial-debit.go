package transactions

import (
	"context"
	"fmt"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionPartialDebitRequest struct {
	// Required
	AuthorizationCode string         `json:"authorization_code"`
	Currency          types.Currency `json:"currency"`
	Amount            int            `json:"amount"`
	Email             string         `json:"email"`

	// Optional
	Reference string `json:"reference,omitempty"`
	AtLeast   string `json:"at_least,omitempty"`
}

type TransactionPartialDebitResponse struct {
	ID              uint64              `json:"id"`
	Amount          int                 `json:"amount"`
	Currency        types.Currency      `json:"currency"`
	TransactionDate time.Time           `json:"transaction_date"`
	Status          string              `json:"status"`
	Reference       string              `json:"reference"`
	Domain          string              `json:"domain"`
	Metadata        types.Metadata      `json:"metadata"`
	GatewayResponse string              `json:"gateway_response"`
	Message         string              `json:"message"`
	Channel         types.Channel       `json:"channel"`
	IPAddress       string              `json:"ip_address"`
	Log             types.Log           `json:"log"`
	Fees            int                 `json:"fees"`
	Authorization   types.Authorization `json:"authorization"`
	Customer        types.Customer      `json:"customer"`
	Plan            uint64              `json:"plan"`
	RequestedAmount int                 `json:"requested_amount"`
}

// TransactionPartialDebitRequestBuilder builds a TransactionPartialDebitRequest
type TransactionPartialDebitRequestBuilder struct {
	request TransactionPartialDebitRequest
}

// NewTransactionPartialDebitRequestBuilder creates a new builder
func NewTransactionPartialDebitRequestBuilder() *TransactionPartialDebitRequestBuilder {
	return &TransactionPartialDebitRequestBuilder{}
}

// AuthorizationCode sets the authorization code for the transaction
func (b *TransactionPartialDebitRequestBuilder) AuthorizationCode(authorizationCode string) *TransactionPartialDebitRequestBuilder {
	b.request.AuthorizationCode = authorizationCode
	return b
}

// Currency sets the currency for the transaction
func (b *TransactionPartialDebitRequestBuilder) Currency(currency types.Currency) *TransactionPartialDebitRequestBuilder {
	b.request.Currency = currency
	return b
}

// Amount sets the amount for the transaction
func (b *TransactionPartialDebitRequestBuilder) Amount(amount int) *TransactionPartialDebitRequestBuilder {
	b.request.Amount = amount
	return b
}

// Email sets the email for the transaction
func (b *TransactionPartialDebitRequestBuilder) Email(email string) *TransactionPartialDebitRequestBuilder {
	b.request.Email = email
	return b
}

// Reference sets the reference for the transaction
func (b *TransactionPartialDebitRequestBuilder) Reference(reference string) *TransactionPartialDebitRequestBuilder {
	b.request.Reference = reference
	return b
}

// AtLeast sets the at least amount for the transaction
func (b *TransactionPartialDebitRequestBuilder) AtLeast(atLeast string) *TransactionPartialDebitRequestBuilder {
	b.request.AtLeast = atLeast
	return b
}

// Build returns the built TransactionPartialDebitRequest
func (b *TransactionPartialDebitRequestBuilder) Build() *TransactionPartialDebitRequest {
	return &b.request
}

func (c *Client) PartialDebit(ctx context.Context, builder *TransactionPartialDebitRequestBuilder) (*types.Response[TransactionPartialDebitResponse], error) {
	req := builder.Build()
	return net.Post[TransactionPartialDebitRequest, TransactionPartialDebitResponse](
		ctx, c.Client, c.Secret, fmt.Sprintf("%s%s", basePath, transactionPartialDebitPath), req, c.BaseURL,
	)
}
