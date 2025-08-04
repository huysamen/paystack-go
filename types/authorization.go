package types

type MandateAuthorizationStatus string

const (
	MandateAuthorizationStatusPending MandateAuthorizationStatus = "pending"
	MandateAuthorizationStatusActive  MandateAuthorizationStatus = "active"
	MandateAuthorizationStatusRevoked MandateAuthorizationStatus = "revoked"
)

// Authorization represents a payment authorization
type Authorization struct {
	AuthorizationCode         string    `json:"authorization_code"`
	Bin                       string    `json:"bin"`
	Last4                     string    `json:"last4"`
	Description               string    `json:"description"`
	ExpMonth                  string    `json:"exp_month"`
	ExpYear                   string    `json:"exp_year"`
	Channel                   Channel   `json:"channel"`
	CardType                  string    `json:"card_type"`
	Brand                     CardBrand `json:"brand"`
	Bank                      string    `json:"bank"`
	Reusable                  bool      `json:"reusable"`
	Signature                 string    `json:"signature"`
	AccountName               string    `json:"account_name"`
	CountryCode               string    `json:"country_code"`
	CountryName               string    `json:"country_name"`
	ReceiverBankAccountNumber string    `json:"receiver_bank_account_number"`
	ReceiverBank              string    `json:"receiver_bank"`
}

type MandateAuthorization struct {
	ID                int                        `json:"id"`
	Status            MandateAuthorizationStatus `json:"status"`
	MandateID         int                        `json:"mandate_id"`
	AuthorizationID   int                        `json:"authorization_id"`
	AuthorizationCode string                     `json:"authorization_code"`
	IntegrationID     int                        `json:"integration_id"`
	AccountNumber     string                     `json:"account_number"`
	BankCode          string                     `json:"bank_code"`
	BankName          string                     `json:"bank_name"`
	Customer          *Customer                  `json:"customer"`
	CreatedAt         string                     `json:"created_at,omitempty"`
	UpdatedAt         string                     `json:"updated_at,omitempty"`
}
