package transfers

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/optional"
	"github.com/huysamen/paystack-go/types"
)

type initiateRequest struct {
	Source           string  `json:"source"`                      // Only "balance" supported for now
	Amount           int     `json:"amount"`                      // Amount in kobo (NGN) or pesewas (GHS)
	Recipient        string  `json:"recipient"`                   // Transfer recipient code
	Reason           *string `json:"reason,omitempty"`            // Reason for transfer
	Currency         *string `json:"currency,omitempty"`          // Defaults to NGN
	AccountReference *string `json:"account_reference,omitempty"` // Required for MPESA in Kenya
	Reference        *string `json:"reference,omitempty"`         // Unique identifier for transfer
}

type InitiateRequestBuilder struct {
	req *initiateRequest
}

func NewInitiateRequestBuilder(source string, amount int, recipient string) *InitiateRequestBuilder {
	return &InitiateRequestBuilder{
		req: &initiateRequest{
			Source:    source,
			Amount:    amount,
			Recipient: recipient,
		},
	}
}

func (b *InitiateRequestBuilder) Reason(reason string) *InitiateRequestBuilder {
	b.req.Reason = optional.String(reason)

	return b
}

func (b *InitiateRequestBuilder) Currency(currency string) *InitiateRequestBuilder {
	b.req.Currency = optional.String(currency)

	return b
}

func (b *InitiateRequestBuilder) AccountReference(accountReference string) *InitiateRequestBuilder {
	b.req.AccountReference = optional.String(accountReference)

	return b
}

func (b *InitiateRequestBuilder) Reference(reference string) *InitiateRequestBuilder {
	b.req.Reference = optional.String(reference)

	return b
}

func (b *InitiateRequestBuilder) Build() *initiateRequest {
	return b.req
}

type InitiateResponseData = types.Transfer
type InitiateResponse = types.Response[InitiateResponseData]

func (c *Client) Initiate(ctx context.Context, builder InitiateRequestBuilder) (*InitiateResponse, error) {
	return net.Post[initiateRequest, InitiateResponseData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
