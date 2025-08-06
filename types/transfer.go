package types

// TransferRecipientType represents the type of transfer recipient
type TransferRecipientType string

const (
	RecipientTypeNuban       TransferRecipientType = "nuban"        // Nigerian bank account
	RecipientTypeGhipss      TransferRecipientType = "ghipss"       // Ghana bank account
	RecipientTypeMobileMoney TransferRecipientType = "mobile_money" // Mobile money account
	RecipientTypeBasa        TransferRecipientType = "basa"         // South African bank account
)

// String returns the string representation of RecipientType
func (r TransferRecipientType) String() string {
	return string(r)
}

// Transfer represents a Paystack transfer
type Transfer struct {
	ID            int            `json:"id"`
	Integration   int            `json:"integration"`
	Domain        string         `json:"domain"`
	Amount        int            `json:"amount"`
	Currency      Currency       `json:"currency"`
	Source        string         `json:"source"`
	SourceDetails map[string]any `json:"source_details"`
	Reason        string         `json:"reason"`
	Status        string         `json:"status"`
	Failures      any            `json:"failures"`
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
	Domain        string           `json:"domain"`
	Type          string           `json:"type"`
	Currency      Currency         `json:"currency"`
	Name          string           `json:"name"`
	Details       RecipientDetails `json:"details"`
	Description   string           `json:"description"`
	Metadata      map[string]any   `json:"metadata"`
	RecipientCode string           `json:"recipient_code"`
	Active        bool             `json:"active"`
	Email         *string          `json:"email"`
	CreatedAt     DateTime         `json:"createdAt"`
	UpdatedAt     DateTime         `json:"updatedAt"`
}

// RecipientDetails represents recipient account details
type RecipientDetails struct {
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
	BankCode      string `json:"bank_code"`
	BankName      string `json:"bank_name"`
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

// TransferRecipient represents a transfer recipient
type TransferRecipient struct {
	ID            uint64                `json:"id"`
	Domain        string                `json:"domain"`
	Type          TransferRecipientType `json:"type"`
	Currency      string                `json:"currency"`
	Name          string                `json:"name"`
	Description   string                `json:"description"`
	Details       RecipientDetails      `json:"details"`
	Metadata      map[string]any        `json:"metadata"`
	RecipientCode string                `json:"recipient_code"`
	Active        bool                  `json:"active"`
	IsDeleted     bool                  `json:"is_deleted"`
	CreatedAt     DateTime              `json:"createdAt"`
	UpdatedAt     DateTime              `json:"updatedAt"`
	Integration   uint64                `json:"integration"`
	Email         *string               `json:"email,omitempty"`
}

// BulkRecipientItem represents a single recipient in a bulk create request
type BulkRecipientItem struct {
	Type              TransferRecipientType `json:"type"`                         // Required: nuban, ghipss, mobile_money, basa
	Name              string                `json:"name"`                         // Required: recipient's name
	AccountNumber     string                `json:"account_number"`               // Required for all types except authorization
	BankCode          string                `json:"bank_code"`                    // Required for all types except authorization
	Description       *string               `json:"description,omitempty"`        // Optional: description
	Currency          *string               `json:"currency,omitempty"`           // Optional: currency
	AuthorizationCode *string               `json:"authorization_code,omitempty"` // Optional: authorization code
	Metadata          map[string]any        `json:"metadata,omitempty"`           // Optional: additional data
}

// BulkCreateResult represents the result of a bulk create operation
type BulkCreateResult struct {
	Success []TransferRecipient `json:"success"` // Successfully created recipients
	Errors  []struct {
		Type          string `json:"type"`
		Name          string `json:"name"`
		AccountNumber string `json:"account_number"`
		BankCode      string `json:"bank_code"`
		Message       string `json:"message"`
	} `json:"errors"` // Failed recipient creations
}
