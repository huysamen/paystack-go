package transfers

import (
	"time"

	"github.com/huysamen/paystack-go/types"
)

// Transfer represents a Paystack transfer
type Transfer struct {
	ID            int            `json:"id"`
	Integration   int            `json:"integration"`
	Domain        string         `json:"domain"`
	Amount        int            `json:"amount"`
	Currency      types.Currency `json:"currency"`
	Source        string         `json:"source"`
	SourceDetails map[string]any `json:"source_details"`
	Reason        string         `json:"reason"`
	Status        string         `json:"status"`
	Failures      any            `json:"failures"`
	TransferCode  string         `json:"transfer_code"`
	TitanCode     *string        `json:"titan_code"`
	TransferredAt *time.Time     `json:"transferred_at"`
	Reference     string         `json:"reference"`
	Recipient     Recipient      `json:"recipient"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
}

// Recipient represents a transfer recipient
type Recipient struct {
	ID            int              `json:"id"`
	Domain        string           `json:"domain"`
	Type          string           `json:"type"`
	Currency      types.Currency   `json:"currency"`
	Name          string           `json:"name"`
	Details       RecipientDetails `json:"details"`
	Description   string           `json:"description"`
	Metadata      map[string]any   `json:"metadata"`
	RecipientCode string           `json:"recipient_code"`
	Active        bool             `json:"active"`
	Email         *string          `json:"email"`
	CreatedAt     time.Time        `json:"createdAt"`
	UpdatedAt     time.Time        `json:"updatedAt"`
}

// RecipientDetails represents recipient account details
type RecipientDetails struct {
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
	BankCode      string `json:"bank_code"`
	BankName      string `json:"bank_name"`
}

// BulkTransferItem represents a single transfer in a bulk transfer request
type BulkTransferItem struct {
	Amount    int    `json:"amount"`
	Reference string `json:"reference"`
	Reason    string `json:"reason"`
	Recipient string `json:"recipient"`
}

// BulkTransferResponse represents a single transfer result in a bulk transfer response
type BulkTransferResponse struct {
	Reference    string         `json:"reference"`
	Recipient    string         `json:"recipient"`
	Amount       int            `json:"amount"`
	TransferCode string         `json:"transfer_code"`
	Currency     types.Currency `json:"currency"`
	Status       string         `json:"status"`
}
