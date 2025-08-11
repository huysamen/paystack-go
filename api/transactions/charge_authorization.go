package transactions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type chargeAuthorizationRequest struct {
	Amount            int    `json:"amount"`
	Email             string `json:"email"`
	AuthorizationCode string `json:"authorization_code"`

	Reference         string          `json:"reference,omitempty"`
	Currency          enums.Currency  `json:"currency,omitempty"`
	Metadata          types.Metadata  `json:"metadata,omitempty"`
	Channels          []enums.Channel `json:"channels,omitempty"`
	Subaccount        string          `json:"subaccount,omitempty"`
	TransactionCharge int             `json:"transaction_charge,omitempty"`
	Bearer            enums.Bearer    `json:"bearer,omitempty"`
	Queue             bool            `json:"queue,omitempty"`
}

type ChargeAuthorizationRequestBuilder struct {
	request chargeAuthorizationRequest
}

func NewChargeAuthorizationRequestBuilder() *ChargeAuthorizationRequestBuilder {
	return &ChargeAuthorizationRequestBuilder{}
}

func (b *ChargeAuthorizationRequestBuilder) Amount(amount int) *ChargeAuthorizationRequestBuilder {
	b.request.Amount = amount

	return b
}

func (b *ChargeAuthorizationRequestBuilder) Email(email string) *ChargeAuthorizationRequestBuilder {
	b.request.Email = email

	return b
}

func (b *ChargeAuthorizationRequestBuilder) AuthorizationCode(authorizationCode string) *ChargeAuthorizationRequestBuilder {
	b.request.AuthorizationCode = authorizationCode

	return b
}

func (b *ChargeAuthorizationRequestBuilder) Reference(reference string) *ChargeAuthorizationRequestBuilder {
	b.request.Reference = reference

	return b
}

func (b *ChargeAuthorizationRequestBuilder) Currency(currency enums.Currency) *ChargeAuthorizationRequestBuilder {
	b.request.Currency = currency

	return b
}

func (b *ChargeAuthorizationRequestBuilder) Metadata(metadata types.Metadata) *ChargeAuthorizationRequestBuilder {
	b.request.Metadata = metadata

	return b
}

func (b *ChargeAuthorizationRequestBuilder) Channels(channels []enums.Channel) *ChargeAuthorizationRequestBuilder {
	b.request.Channels = channels

	return b
}

func (b *ChargeAuthorizationRequestBuilder) Subaccount(subaccount string) *ChargeAuthorizationRequestBuilder {
	b.request.Subaccount = subaccount

	return b
}

func (b *ChargeAuthorizationRequestBuilder) TransactionCharge(transactionCharge int) *ChargeAuthorizationRequestBuilder {
	b.request.TransactionCharge = transactionCharge

	return b
}

func (b *ChargeAuthorizationRequestBuilder) Bearer(bearer enums.Bearer) *ChargeAuthorizationRequestBuilder {
	b.request.Bearer = bearer

	return b
}

func (b *ChargeAuthorizationRequestBuilder) Queue(queue bool) *ChargeAuthorizationRequestBuilder {
	b.request.Queue = queue

	return b
}

func (b *ChargeAuthorizationRequestBuilder) Build() *chargeAuthorizationRequest {
	return &b.request
}

// Charge authorization responses can include plan as null; the standard Transaction now has *Plan
type ChargeAuthorizationResponseData = types.Transaction
type ChargeAuthorizationResponse = types.Response[ChargeAuthorizationResponseData]

func (c *Client) ChargeAuthorization(ctx context.Context, builder ChargeAuthorizationRequestBuilder) (*ChargeAuthorizationResponse, error) {
	return net.Post[chargeAuthorizationRequest, ChargeAuthorizationResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s%s", basePath, transactionChargeAuthorizationPath), builder.Build(), c.BaseURL)
}
