package transfers

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/optional"
	"github.com/huysamen/paystack-go/types"
)

type TransferInitiateRequest struct {
	Source           string  `json:"source"`                      // Only "balance" supported for now
	Amount           int     `json:"amount"`                      // Amount in kobo (NGN) or pesewas (GHS)
	Recipient        string  `json:"recipient"`                   // Transfer recipient code
	Reason           *string `json:"reason,omitempty"`            // Reason for transfer
	Currency         *string `json:"currency,omitempty"`          // Defaults to NGN
	AccountReference *string `json:"account_reference,omitempty"` // Required for MPESA in Kenya
	Reference        *string `json:"reference,omitempty"`         // Unique identifier for transfer
}

type TransferInitiateRequestBuilder struct {
	req *TransferInitiateRequest
}

func NewInitiateTransferRequest(source string, amount int, recipient string) *TransferInitiateRequestBuilder {
	return &TransferInitiateRequestBuilder{
		req: &TransferInitiateRequest{
			Source:    source,
			Amount:    amount,
			Recipient: recipient,
		},
	}
}

func (b *TransferInitiateRequestBuilder) Reason(reason string) *TransferInitiateRequestBuilder {
	b.req.Reason = optional.String(reason)

	return b
}

func (b *TransferInitiateRequestBuilder) Currency(currency string) *TransferInitiateRequestBuilder {
	b.req.Currency = optional.String(currency)

	return b
}

func (b *TransferInitiateRequestBuilder) AccountReference(accountReference string) *TransferInitiateRequestBuilder {
	b.req.AccountReference = optional.String(accountReference)

	return b
}

func (b *TransferInitiateRequestBuilder) Reference(reference string) *TransferInitiateRequestBuilder {
	b.req.Reference = optional.String(reference)

	return b
}

func (b *TransferInitiateRequestBuilder) Build() *TransferInitiateRequest {
	return b.req
}

type InitiateResponse = types.Response[types.Transfer]

func (c *Client) Initiate(ctx context.Context, builder *TransferInitiateRequestBuilder) (*InitiateResponse, error) {
	return net.Post[TransferInitiateRequest, types.Transfer](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
