package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// BulkCharge represents a single charge within a bulk charge batch
type BulkCharge struct {
	Integration   data.Int       `json:"integration"`
	BulkCharge    data.Int       `json:"bulkcharge"`
	Customer      Customer       `json:"customer"`
	Authorization Authorization  `json:"authorization"`
	Transaction   *Transaction   `json:"transaction"`
	Domain        data.String    `json:"domain"`
	Amount        data.Int       `json:"amount"`
	Currency      enums.Currency `json:"currency"`
	Reference     data.String    `json:"reference"`
	Status        data.String    `json:"status"`
	Message       data.String    `json:"message"`
	PaidAt        data.NullTime  `json:"paid_at,omitempty"`
	CreatedAt     data.Time      `json:"createdAt"`
	UpdatedAt     data.Time      `json:"updatedAt"`
}

// BulkChargeBatch represents a bulk charge batch
type BulkChargeBatch struct {
	ID             data.Int    `json:"id"`
	BatchCode      data.String `json:"batch_code"`
	Reference      data.String `json:"reference,omitempty"`
	Integration    data.Int    `json:"integration,omitempty"`
	Domain         data.String `json:"domain"`
	Status         data.String `json:"status"`
	TotalCharges   data.Int    `json:"total_charges"`
	PendingCharges data.Int    `json:"pending_charges"`
	CreatedAt      data.Time   `json:"createdAt"`
	UpdatedAt      data.Time   `json:"updatedAt"`
}
