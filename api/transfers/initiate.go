package transfers

import (
	"context"

	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/optional"
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
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

// Initiate returns recipient as an ID in fixtures; define a narrow response type for this endpoint
type InitiateResponseData struct {
	ID            data.Int        `json:"id"`
	Integration   data.Int        `json:"integration"`
	Domain        data.String     `json:"domain"`
	Amount        data.Int        `json:"amount"`
	Currency      enums.Currency  `json:"currency"`
	Source        data.String     `json:"source"`
	SourceDetails types.Metadata  `json:"source_details"`
	Reason        data.String     `json:"reason"`
	Status        data.String     `json:"status"`
	Failures      types.Metadata  `json:"failures"`
	TransferCode  data.String     `json:"transfer_code"`
	TitanCode     data.NullString `json:"titan_code"`
	TransferredAt data.NullTime   `json:"transferred_at"`
	Reference     data.String     `json:"reference"`
	Recipient     data.Int        `json:"recipient"` // ID here
	CreatedAt     data.Time       `json:"createdAt"`
	UpdatedAt     data.Time       `json:"updatedAt"`
}
type InitiateResponse = types.Response[InitiateResponseData]

func (c *Client) Initiate(ctx context.Context, builder InitiateRequestBuilder) (*InitiateResponse, error) {
	return net.Post[initiateRequest, InitiateResponseData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
