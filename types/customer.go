package types

import "github.com/huysamen/paystack-go/types/data"

// Customer represents a Paystack customer
type Customer struct {
	ID           data.Uint       `json:"id"`
	Integration  data.NullInt    `json:"integration,omitempty"`
	FirstName    data.NullString `json:"first_name"`
	LastName     data.NullString `json:"last_name"`
	Email        data.String     `json:"email"`
	Phone        data.NullString `json:"phone"`
	Metadata     Metadata        `json:"metadata"`
	Domain       data.String     `json:"domain"`
	CustomerCode data.String     `json:"customer_code"`
	RiskAction   data.String     `json:"risk_action"`
	CreatedAt    data.Time       `json:"createdAt"`
	UpdatedAt    data.Time       `json:"updatedAt"`

	// Additional fields from detailed customer responses
	TotalTransactions     data.Int                 `json:"total_transactions,omitempty"`
	TotalTransactionValue []any                    `json:"total_transaction_value,omitempty"`
	DedicatedAccount      *DedicatedVirtualAccount `json:"dedicated_account,omitempty"`
	Identified            data.Bool                `json:"identified,omitempty"`
	Identifications       *Metadata                `json:"identifications,omitempty"`

	// Alternative field names that also appear in responses
	CreatedAtSnake data.NullTime `json:"created_at,omitempty"`
	UpdatedAtSnake data.NullTime `json:"updated_at,omitempty"`
}
