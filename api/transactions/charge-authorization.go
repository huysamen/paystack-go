package transactions

import (
	"context"
	"fmt"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionChargeAuthorizationRequest struct {
	// Required
	Amount            int    `json:"amount"`
	Email             string `json:"email"`
	AuthorizationCode string `json:"authorization_code"`

	// Optional
	Reference         string          `json:"reference,omitempty"`
	Currency          types.Currency  `json:"currency,omitempty"`
	Metadata          types.Metadata  `json:"metadata,omitempty"`
	Channels          []types.Channel `json:"channels,omitempty"`
	Subaccount        string          `json:"subaccount,omitempty"`
	TransactionCharge int             `json:"transaction_charge,omitempty"`
	Bearer            types.Bearer    `json:"bearer,omitempty"`
	Queue             bool            `json:"queue,omitempty"`
}

type TransactionChargeAuthorizationResponse struct {
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
	Plan            types.Plan          `json:"plan"`
}

// TransactionChargeAuthorizationRequestBuilder builds a TransactionChargeAuthorizationRequest
type TransactionChargeAuthorizationRequestBuilder struct {
	request TransactionChargeAuthorizationRequest
}

// NewTransactionChargeAuthorizationRequestBuilder creates a new builder
func NewTransactionChargeAuthorizationRequestBuilder() *TransactionChargeAuthorizationRequestBuilder {
	return &TransactionChargeAuthorizationRequestBuilder{}
}

// Amount sets the amount for the transaction
func (b *TransactionChargeAuthorizationRequestBuilder) Amount(amount int) *TransactionChargeAuthorizationRequestBuilder {
	b.request.Amount = amount
	return b
}

// Email sets the email for the transaction
func (b *TransactionChargeAuthorizationRequestBuilder) Email(email string) *TransactionChargeAuthorizationRequestBuilder {
	b.request.Email = email
	return b
}

// AuthorizationCode sets the authorization code for the transaction
func (b *TransactionChargeAuthorizationRequestBuilder) AuthorizationCode(authorizationCode string) *TransactionChargeAuthorizationRequestBuilder {
	b.request.AuthorizationCode = authorizationCode
	return b
}

// Reference sets the reference for the transaction
func (b *TransactionChargeAuthorizationRequestBuilder) Reference(reference string) *TransactionChargeAuthorizationRequestBuilder {
	b.request.Reference = reference
	return b
}

// Currency sets the currency for the transaction
func (b *TransactionChargeAuthorizationRequestBuilder) Currency(currency types.Currency) *TransactionChargeAuthorizationRequestBuilder {
	b.request.Currency = currency
	return b
}

// Metadata sets the metadata for the transaction
func (b *TransactionChargeAuthorizationRequestBuilder) Metadata(metadata types.Metadata) *TransactionChargeAuthorizationRequestBuilder {
	b.request.Metadata = metadata
	return b
}

// Channels sets the channels for the transaction
func (b *TransactionChargeAuthorizationRequestBuilder) Channels(channels []types.Channel) *TransactionChargeAuthorizationRequestBuilder {
	b.request.Channels = channels
	return b
}

// Subaccount sets the subaccount for the transaction
func (b *TransactionChargeAuthorizationRequestBuilder) Subaccount(subaccount string) *TransactionChargeAuthorizationRequestBuilder {
	b.request.Subaccount = subaccount
	return b
}

// TransactionCharge sets the transaction charge
func (b *TransactionChargeAuthorizationRequestBuilder) TransactionCharge(transactionCharge int) *TransactionChargeAuthorizationRequestBuilder {
	b.request.TransactionCharge = transactionCharge
	return b
}

// Bearer sets the bearer for the transaction
func (b *TransactionChargeAuthorizationRequestBuilder) Bearer(bearer types.Bearer) *TransactionChargeAuthorizationRequestBuilder {
	b.request.Bearer = bearer
	return b
}

// Queue sets whether to queue the transaction
func (b *TransactionChargeAuthorizationRequestBuilder) Queue(queue bool) *TransactionChargeAuthorizationRequestBuilder {
	b.request.Queue = queue
	return b
}

// Build returns the built TransactionChargeAuthorizationRequest
func (b *TransactionChargeAuthorizationRequestBuilder) Build() *TransactionChargeAuthorizationRequest {
	return &b.request
}

func (c *Client) ChargeAuthorization(ctx context.Context, builder *TransactionChargeAuthorizationRequestBuilder) (*types.Response[TransactionChargeAuthorizationResponse], error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	return net.Post[TransactionChargeAuthorizationRequest, TransactionChargeAuthorizationResponse](
		ctx,
		c.client,
		c.secret,
		fmt.Sprintf("%s%s", transactionBasePath, transactionChargeAuthorizationPath),
		req,
		c.baseURL,
	)
}
