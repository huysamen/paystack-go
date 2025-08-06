package types

import (
	"github.com/huysamen/paystack-go/enums"
)

// Refund represents a refund object
type Refund struct {
	ID             int                 `json:"id"`
	Integration    int                 `json:"integration"`
	Domain         string              `json:"domain"`
	Transaction    int                 `json:"transaction"`
	Dispute        *int                `json:"dispute"`
	Settlement     *int                `json:"settlement"`
	Amount         int                 `json:"amount"`
	DeductedAmount int                 `json:"deducted_amount"`
	Currency       string              `json:"currency"`
	Channel        enums.RefundChannel `json:"channel"`
	FullyDeducted  bool                `json:"fully_deducted"`
	Status         enums.RefundStatus  `json:"status"`
	RefundedBy     string              `json:"refunded_by"`
	RefundedAt     *DateTime           `json:"refunded_at"`
	ExpectedAt     *DateTime           `json:"expected_at"`
	CreatedAt      *DateTime           `json:"created_at"`
	UpdatedAt      *DateTime           `json:"updated_at"`
	CustomerNote   *string             `json:"customer_note"`
	MerchantNote   *string             `json:"merchant_note"`
}
