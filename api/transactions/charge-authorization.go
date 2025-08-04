package transactions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionChargeAuthorizationRequest struct {
	Amount            int    `json:"amount"`
	Email             string `json:"email"`
	AuthorizationCode string `json:"authorization_code"`

	Reference         string          `json:"reference,omitempty"`
	Currency          types.Currency  `json:"currency,omitempty"`
	Metadata          types.Metadata  `json:"metadata,omitempty"`
	Channels          []types.Channel `json:"channels,omitempty"`
	Subaccount        string          `json:"subaccount,omitempty"`
	TransactionCharge int             `json:"transaction_charge,omitempty"`
	Bearer            types.Bearer    `json:"bearer,omitempty"`
	Queue             bool            `json:"queue,omitempty"`
}

type TransactionChargeAuthorizationRequestBuilder struct {
	request TransactionChargeAuthorizationRequest
}

func NewTransactionChargeAuthorizationRequestBuilder() *TransactionChargeAuthorizationRequestBuilder {
	return &TransactionChargeAuthorizationRequestBuilder{}
}

func (b *TransactionChargeAuthorizationRequestBuilder) Amount(amount int) *TransactionChargeAuthorizationRequestBuilder {
	b.request.Amount = amount

	return b
}

func (b *TransactionChargeAuthorizationRequestBuilder) Email(email string) *TransactionChargeAuthorizationRequestBuilder {
	b.request.Email = email

	return b
}

func (b *TransactionChargeAuthorizationRequestBuilder) AuthorizationCode(authorizationCode string) *TransactionChargeAuthorizationRequestBuilder {
	b.request.AuthorizationCode = authorizationCode

	return b
}

func (b *TransactionChargeAuthorizationRequestBuilder) Reference(reference string) *TransactionChargeAuthorizationRequestBuilder {
	b.request.Reference = reference

	return b
}

func (b *TransactionChargeAuthorizationRequestBuilder) Currency(currency types.Currency) *TransactionChargeAuthorizationRequestBuilder {
	b.request.Currency = currency

	return b
}

func (b *TransactionChargeAuthorizationRequestBuilder) Metadata(metadata types.Metadata) *TransactionChargeAuthorizationRequestBuilder {
	b.request.Metadata = metadata

	return b
}

func (b *TransactionChargeAuthorizationRequestBuilder) Channels(channels []types.Channel) *TransactionChargeAuthorizationRequestBuilder {
	b.request.Channels = channels

	return b
}

func (b *TransactionChargeAuthorizationRequestBuilder) Subaccount(subaccount string) *TransactionChargeAuthorizationRequestBuilder {
	b.request.Subaccount = subaccount

	return b
}

func (b *TransactionChargeAuthorizationRequestBuilder) TransactionCharge(transactionCharge int) *TransactionChargeAuthorizationRequestBuilder {
	b.request.TransactionCharge = transactionCharge

	return b
}

func (b *TransactionChargeAuthorizationRequestBuilder) Bearer(bearer types.Bearer) *TransactionChargeAuthorizationRequestBuilder {
	b.request.Bearer = bearer

	return b
}

func (b *TransactionChargeAuthorizationRequestBuilder) Queue(queue bool) *TransactionChargeAuthorizationRequestBuilder {
	b.request.Queue = queue

	return b
}

func (b *TransactionChargeAuthorizationRequestBuilder) Build() *TransactionChargeAuthorizationRequest {
	return &b.request
}

type ChargeAuthorizationResponse = types.Response[types.Transaction]

func (c *Client) ChargeAuthorization(ctx context.Context, builder *TransactionChargeAuthorizationRequestBuilder) (*ChargeAuthorizationResponse, error) {
	return net.Post[TransactionChargeAuthorizationRequest, types.Transaction](ctx, c.Client, c.Secret, fmt.Sprintf("%s%s", basePath, transactionChargeAuthorizationPath), builder.Build(), c.BaseURL)
}
