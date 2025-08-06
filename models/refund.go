package models

import "github.com/huysamen/paystack-go/enums"

// Refund represents a refund object
type Refund struct {
	ID             int                  `json:"id"`
	Integration    int                  `json:"integration"`
	Domain         string               `json:"domain"`
	Transaction    *Transaction         `json:"transaction,omitempty"`
	Dispute        *int                 `json:"dispute"`
	Settlement     *int                 `json:"settlement"`
	Amount         int                  `json:"amount"`
	DeductedAmount int                  `json:"deducted_amount"`
	Currency       enums.Currency       `json:"currency"`
	Channel        *enums.RefundChannel `json:"channel"`
	FullyDeducted  bool                 `json:"fully_deducted"`
	Status         enums.RefundStatus   `json:"status"`
	RefundedBy     string               `json:"refunded_by"`
	RefundedAt     *DateTime            `json:"refunded_at"`
	ExpectedAt     *DateTime            `json:"expected_at"`
	CreatedAt      DateTime             `json:"createdAt"`
	UpdatedAt      DateTime             `json:"updatedAt"`
	CustomerNote   *string              `json:"customer_note"`
	MerchantNote   *string              `json:"merchant_note"`
}
