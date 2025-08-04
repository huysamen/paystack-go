package transfers

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/optional"
	"github.com/huysamen/paystack-go/types"
)

// Request type
type TransferInitiateRequest struct {
	Source           string  `json:"source"`                      // Only "balance" supported for now
	Amount           int     `json:"amount"`                      // Amount in kobo (NGN) or pesewas (GHS)
	Recipient        string  `json:"recipient"`                   // Transfer recipient code
	Reason           *string `json:"reason,omitempty"`            // Reason for transfer
	Currency         *string `json:"currency,omitempty"`          // Defaults to NGN
	AccountReference *string `json:"account_reference,omitempty"` // Required for MPESA in Kenya
	Reference        *string `json:"reference,omitempty"`         // Unique identifier for transfer
}

// Builder for creating TransferInitiateRequest
type TransferInitiateRequestBuilder struct {
	req *TransferInitiateRequest
}

// NewInitiateTransferRequest creates a new builder for transfer initiation
func NewInitiateTransferRequest(source string, amount int, recipient string) *TransferInitiateRequestBuilder {
	return &TransferInitiateRequestBuilder{
		req: &TransferInitiateRequest{
			Source:    source,
			Amount:    amount,
			Recipient: recipient,
		},
	}
}

// Reason sets the reason for the transfer
func (b *TransferInitiateRequestBuilder) Reason(reason string) *TransferInitiateRequestBuilder {
	b.req.Reason = optional.String(reason)

	return b
}

// Currency sets the currency
func (b *TransferInitiateRequestBuilder) Currency(currency string) *TransferInitiateRequestBuilder {
	b.req.Currency = optional.String(currency)

	return b
}

// AccountReference sets the account reference (required for MPESA in Kenya)
func (b *TransferInitiateRequestBuilder) AccountReference(accountReference string) *TransferInitiateRequestBuilder {
	b.req.AccountReference = optional.String(accountReference)

	return b
}

// Reference sets the unique identifier for the transfer
func (b *TransferInitiateRequestBuilder) Reference(reference string) *TransferInitiateRequestBuilder {
	b.req.Reference = optional.String(reference)

	return b
}

// Build creates the TransferInitiateRequest
func (b *TransferInitiateRequestBuilder) Build() *TransferInitiateRequest {
	return b.req
}

// InitiateResponse represents the response for initiating a transfer
type InitiateResponse = types.Response[types.Transfer]

// Initiate creates a new transfer with the provided builder
func (c *Client) Initiate(ctx context.Context, builder *TransferInitiateRequestBuilder) (*InitiateResponse, error) {
	return net.Post[TransferInitiateRequest, types.Transfer](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
