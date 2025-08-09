package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Transfer represents a Paystack transfer
type Transfer struct {
	ID            data.Int            `json:"id"`
	Integration   data.Int            `json:"integration"`
	Domain        data.String         `json:"domain"`
	Amount        data.Int            `json:"amount"`
	Currency      enums.Currency      `json:"currency"`
	Source        data.String         `json:"source"`
	SourceDetails *Metadata           `json:"source_details"`
	Reason        data.String         `json:"reason"`
	Status        data.String         `json:"status"`
	Failures      *Metadata           `json:"failures"`
	TransferCode  data.String         `json:"transfer_code"`
	TitanCode     data.NullString     `json:"titan_code"`
	TransferredAt *data.MultiDateTime `json:"transferred_at,omitempty"`
	Reference     data.String         `json:"reference"`
	Recipient     Recipient           `json:"recipient"`
	CreatedAt     data.MultiDateTime  `json:"createdAt"`
	UpdatedAt     data.MultiDateTime  `json:"updatedAt"`
}

// Recipient represents a transfer recipient
type Recipient struct {
	ID            data.Int           `json:"id"`
	Integration   data.Int           `json:"integration"`
	Domain        data.String        `json:"domain"`
	Type          data.String        `json:"type"`
	Currency      enums.Currency     `json:"currency"`
	Name          data.String        `json:"name"`
	Details       RecipientDetails   `json:"details"`
	Description   data.String        `json:"description"`
	Metadata      *Metadata          `json:"metadata"`
	RecipientCode data.String        `json:"recipient_code"`
	Active        data.Bool          `json:"active"`
	Email         data.NullString    `json:"email"`
	IsDeleted     data.Bool          `json:"is_deleted"`
	CreatedAt     data.MultiDateTime `json:"createdAt"`
	UpdatedAt     data.MultiDateTime `json:"updatedAt"`
}

// RecipientDetails represents recipient account details
type RecipientDetails struct {
	AuthorizationCode data.NullString `json:"authorization_code"`
	AccountNumber     data.String     `json:"account_number"`
	AccountName       data.String     `json:"account_name"`
	BankCode          data.String     `json:"bank_code"`
	BankName          data.String     `json:"bank_name"`
}

// Balance represents account balance information
type Balance struct {
	Currency data.String `json:"currency"`
	Balance  data.Int    `json:"balance"`
}

// BalanceLedger represents a balance ledger entry
type BalanceLedger struct {
	Integration      data.Int           `json:"integration"`
	Domain           data.String        `json:"domain"`
	Balance          data.Int           `json:"balance"`
	Currency         data.String        `json:"currency"`
	Difference       data.Int           `json:"difference"`
	Reason           data.String        `json:"reason"`
	ModelResponsible data.String        `json:"model_responsible"`
	ModelRow         data.Int           `json:"model_row"`
	ID               data.Int           `json:"id"`
	CreatedAt        data.MultiDateTime `json:"createdAt"`
	UpdatedAt        data.MultiDateTime `json:"updatedAt"`
}

// BulkRecipientItem represents a recipient item for bulk creation
type BulkRecipientItem struct {
	Type          data.String    `json:"type"`
	Name          data.String    `json:"name"`
	AccountNumber data.String    `json:"account_number"`
	BankCode      data.String    `json:"bank_code"`
	Currency      enums.Currency `json:"currency"`
	Description   data.String    `json:"description,omitempty"`
	Email         data.String    `json:"email,omitempty"`
	Metadata      *Metadata      `json:"metadata,omitempty"`
}

// BulkCreateResult represents the result of bulk recipient creation
type BulkCreateResult struct {
	Success []Recipient `json:"success"`
	Errors  []struct {
		Error   data.String       `json:"error"`
		Payload BulkRecipientItem `json:"payload"`
	} `json:"errors"`
}
