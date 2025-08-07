package types

import "github.com/huysamen/paystack-go/enums"

// Transfer represents a Paystack transfer
type Transfer struct {
	ID            int            `json:"id"`
	Integration   int            `json:"integration"`
	Domain        string         `json:"domain"`
	Amount        int            `json:"amount"`
	Currency      enums.Currency `json:"currency"`
	Source        string         `json:"source"`
	SourceDetails *Metadata      `json:"source_details"`
	Reason        string         `json:"reason"`
	Status        string         `json:"status"`
	Failures      *Metadata      `json:"failures"`
	TransferCode  string         `json:"transfer_code"`
	TitanCode     *string        `json:"titan_code"`
	TransferredAt *DateTime      `json:"transferred_at,omitempty"`
	Reference     string         `json:"reference"`
	Recipient     Recipient      `json:"recipient"`
	CreatedAt     DateTime       `json:"createdAt"`
	UpdatedAt     DateTime       `json:"updatedAt"`
}

// Recipient represents a transfer recipient
type Recipient struct {
	ID            int              `json:"id"`
	Integration   int              `json:"integration"`
	Domain        string           `json:"domain"`
	Type          string           `json:"type"`
	Currency      enums.Currency   `json:"currency"`
	Name          string           `json:"name"`
	Details       RecipientDetails `json:"details"`
	Description   string           `json:"description"`
	Metadata      *Metadata        `json:"metadata"`
	RecipientCode string           `json:"recipient_code"`
	Active        bool             `json:"active"`
	Email         *string          `json:"email"`
	IsDeleted     bool             `json:"is_deleted"`
	CreatedAt     DateTime         `json:"createdAt"`
	UpdatedAt     DateTime         `json:"updatedAt"`
}

// RecipientDetails represents recipient account details
type RecipientDetails struct {
	AuthorizationCode *string `json:"authorization_code"`
	AccountNumber     string  `json:"account_number"`
	AccountName       string  `json:"account_name"`
	BankCode          string  `json:"bank_code"`
	BankName          string  `json:"bank_name"`
}

// Balance represents account balance information
type Balance struct {
	Currency string `json:"currency"`
	Balance  int64  `json:"balance"`
}

// BalanceLedger represents a balance ledger entry
type BalanceLedger struct {
	Integration      int      `json:"integration"`
	Domain           string   `json:"domain"`
	Balance          int64    `json:"balance"`
	Currency         string   `json:"currency"`
	Difference       int64    `json:"difference"`
	Reason           string   `json:"reason"`
	ModelResponsible string   `json:"model_responsible"`
	ModelRow         int      `json:"model_row"`
	ID               int      `json:"id"`
	CreatedAt        DateTime `json:"createdAt"`
	UpdatedAt        DateTime `json:"updatedAt"`
}

// BulkRecipientItem represents a recipient item for bulk creation
type BulkRecipientItem struct {
	Type          string         `json:"type"`
	Name          string         `json:"name"`
	AccountNumber string         `json:"account_number"`
	BankCode      string         `json:"bank_code"`
	Currency      enums.Currency `json:"currency"`
	Description   string         `json:"description,omitempty"`
	Email         string         `json:"email,omitempty"`
	Metadata      *Metadata      `json:"metadata,omitempty"`
}

// BulkCreateResult represents the result of bulk recipient creation
type BulkCreateResult struct {
	Success []Recipient `json:"success"`
	Errors  []struct {
		Error   string            `json:"error"`
		Payload BulkRecipientItem `json:"payload"`
	} `json:"errors"`
}
