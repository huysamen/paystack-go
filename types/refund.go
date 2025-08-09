package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Refund represents a refund object
type Refund struct {
	ID             data.Int             `json:"id"`
	Integration    data.Int             `json:"integration"`
	Domain         data.String          `json:"domain"`
	Transaction    *Transaction         `json:"transaction,omitempty"`
	Dispute        data.NullInt         `json:"dispute"`
	Settlement     data.NullInt         `json:"settlement"`
	Amount         data.Int             `json:"amount"`
	DeductedAmount data.Int             `json:"deducted_amount"`
	Currency       enums.Currency       `json:"currency"`
	Channel        *enums.RefundChannel `json:"channel"`
	FullyDeducted  data.Bool            `json:"fully_deducted"`
	Status         enums.RefundStatus   `json:"status"`
	RefundedBy     data.String          `json:"refunded_by"`
	RefundedAt     *data.MultiDateTime  `json:"refunded_at"`
	ExpectedAt     *data.MultiDateTime  `json:"expected_at"`
	CreatedAt      data.MultiDateTime   `json:"createdAt"`
	UpdatedAt      data.MultiDateTime   `json:"updatedAt"`
	CustomerNote   data.NullString      `json:"customer_note"`
	MerchantNote   data.NullString      `json:"merchant_note"`
}
