package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Refund represents a refund object
type Refund struct {
	ID             int                  `json:"id"`
	Integration    int                  `json:"integration"`
	Domain         string               `json:"domain"`
	Transaction    *Transaction         `json:"transaction,omitempty"`
	Dispute        data.NullInt         `json:"dispute"`
	Settlement     data.NullInt         `json:"settlement"`
	Amount         int                  `json:"amount"`
	DeductedAmount int                  `json:"deducted_amount"`
	Currency       enums.Currency       `json:"currency"`
	Channel        *enums.RefundChannel `json:"channel"`
	FullyDeducted  bool                 `json:"fully_deducted"`
	Status         enums.RefundStatus   `json:"status"`
	RefundedBy     string               `json:"refunded_by"`
	RefundedAt     *data.MultiDateTime  `json:"refunded_at"`
	ExpectedAt     *data.MultiDateTime  `json:"expected_at"`
	CreatedAt      data.MultiDateTime   `json:"createdAt"`
	UpdatedAt      data.MultiDateTime   `json:"updatedAt"`
	CustomerNote   data.NullString      `json:"customer_note"`
	MerchantNote   data.NullString      `json:"merchant_note"`
}
