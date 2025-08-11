package refunds

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
)

type FetchResponseData struct {
	ID             data.Int             `json:"id"`
	Integration    data.Int             `json:"integration"`
	Domain         data.String          `json:"domain"`
	Transaction    data.Int             `json:"transaction"` // Transaction ID, not object in fetch response
	Dispute        data.NullInt         `json:"dispute"`
	Settlement     data.NullInt         `json:"settlement"`
	Amount         data.Int             `json:"amount"`
	DeductedAmount data.Int             `json:"deducted_amount"`
	Currency       enums.Currency       `json:"currency"`
	Channel        *enums.RefundChannel `json:"channel"`
	FullyDeducted  data.Bool            `json:"fully_deducted"`
	Status         enums.RefundStatus   `json:"status"`
	RefundedBy     data.String          `json:"refunded_by"`
	RefundedAt     data.NullTime        `json:"refunded_at"`
	ExpectedAt     data.NullTime        `json:"expected_at"`
	CreatedAt      data.Time            `json:"createdAt"`
	UpdatedAt      data.Time            `json:"updatedAt"`
	CustomerNote   data.NullString      `json:"customer_note"`
	MerchantNote   data.NullString      `json:"merchant_note"`
}

type FetchResponse = types.Response[FetchResponseData]

func (c *Client) Fetch(ctx context.Context, refundID string) (*FetchResponse, error) {
	return net.Get[FetchResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/%s", basePath, refundID), c.BaseURL)
}
