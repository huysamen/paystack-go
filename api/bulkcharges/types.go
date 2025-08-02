package bulkcharges

import (
	"github.com/huysamen/paystack-go/types"
)

// BulkChargeItem represents a single charge in a bulk charge request
type BulkChargeItem struct {
	Authorization string `json:"authorization"`
	Amount        int64  `json:"amount"`
	Reference     string `json:"reference"`
}

// BulkChargeBatch represents a bulk charge batch
type BulkChargeBatch struct {
	ID             int    `json:"id"`
	BatchCode      string `json:"batch_code"`
	Reference      string `json:"reference,omitempty"`
	Integration    int    `json:"integration,omitempty"`
	Domain         string `json:"domain"`
	Status         string `json:"status"`
	TotalCharges   int    `json:"total_charges"`
	PendingCharges int    `json:"pending_charges"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
}

// BulkChargeCharge represents a single charge within a bulk charge batch
type BulkChargeCharge struct {
	Integration   int                 `json:"integration"`
	BulkCharge    int                 `json:"bulkcharge"`
	Customer      types.Customer      `json:"customer"`
	Authorization types.Authorization `json:"authorization"`
	Transaction   any                 `json:"transaction"`
	Domain        string              `json:"domain"`
	Amount        int64               `json:"amount"`
	Currency      string              `json:"currency"`
	Reference     string              `json:"reference"`
	Status        string              `json:"status"`
	Message       string              `json:"message"`
	PaidAt        string              `json:"paid_at"`
	CreatedAt     string              `json:"createdAt"`
	UpdatedAt     string              `json:"updatedAt"`
}
