package types

// Customer represents a Paystack customer
type Customer struct {
	ID           uint64   `json:"id"`
	Integration  *int     `json:"integration,omitempty"`
	FirstName    *string  `json:"first_name"`
	LastName     *string  `json:"last_name"`
	Email        string   `json:"email"`
	Phone        *string  `json:"phone"`
	Metadata     Metadata `json:"metadata"`
	Domain       string   `json:"domain"`
	CustomerCode string   `json:"customer_code"`
	RiskAction   string   `json:"risk_action"`
	CreatedAt    DateTime `json:"createdAt"`
	UpdatedAt    DateTime `json:"updatedAt"`

	// Additional fields from detailed customer responses
	TotalTransactions     int                      `json:"total_transactions,omitempty"`
	TotalTransactionValue []any                    `json:"total_transaction_value,omitempty"`
	DedicatedAccount      *DedicatedVirtualAccount `json:"dedicated_account,omitempty"`
	Identified            bool                     `json:"identified,omitempty"`
	Identifications       *Metadata                `json:"identifications,omitempty"`

	// Alternative field names that also appear in responses
	CreatedAtSnake DateTime `json:"created_at,omitempty"`
	UpdatedAtSnake DateTime `json:"updated_at,omitempty"`
}
