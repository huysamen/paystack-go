package types
import "github.com/huysamen/paystack-go/types/data"
import "github.com/huysamen/paystack-go/enums"

// BulkCharge represents a single charge within a bulk charge batch
type BulkCharge struct {
	Integration   int            `json:"integration"`
	BulkCharge    int            `json:"bulkcharge"`
	Customer      Customer       `json:"customer"`
	Authorization Authorization  `json:"authorization"`
	Transaction   *Transaction   `json:"transaction"`
	Domain        string         `json:"domain"`
	Amount        int64          `json:"amount"`
	Currency      enums.Currency `json:"currency"`
	Reference     string         `json:"reference"`
	Status        string         `json:"status"`
	Message       string         `json:"message"`
	PaidAt        *data.MultiDateTime      `json:"paid_at,omitempty"`
	CreatedAt     data.MultiDateTime       `json:"createdAt"`
	UpdatedAt     data.MultiDateTime       `json:"updatedAt"`
}

// BulkChargeBatch represents a bulk charge batch
type BulkChargeBatch struct {
	ID             int      `json:"id"`
	BatchCode      string   `json:"batch_code"`
	Reference      string   `json:"reference,omitempty"`
	Integration    int      `json:"integration,omitempty"`
	Domain         string   `json:"domain"`
	Status         string   `json:"status"`
	TotalCharges   int      `json:"total_charges"`
	PendingCharges int      `json:"pending_charges"`
	CreatedAt      data.MultiDateTime `json:"createdAt"`
	UpdatedAt      data.MultiDateTime `json:"updatedAt"`
}
