package types

import (
	"github.com/huysamen/paystack-go/enums"
)

// Authorization represents a payment authorization
type Authorization struct {
	AuthorizationCode         string          `json:"authorization_code"`
	Bin                       string          `json:"bin"`
	Last4                     string          `json:"last4"`
	ExpMonth                  string          `json:"exp_month"`
	ExpYear                   string          `json:"exp_year"`
	Channel                   enums.Channel   `json:"channel"`
	CardType                  string          `json:"card_type"`
	Brand                     enums.CardBrand `json:"brand"`
	Bank                      string          `json:"bank"`
	CountryCode               string          `json:"country_code"`
	CountryName               *string         `json:"country_name,omitempty"`
	Reusable                  bool            `json:"reusable"`
	Signature                 string          `json:"signature"`
	AccountName               *string         `json:"account_name"`
	ReceiverBankAccountNumber *string         `json:"receiver_bank_account_number,omitempty"`
	ReceiverBank              *string         `json:"receiver_bank,omitempty"`
}

// MandateAuthorization represents a mandate authorization
type MandateAuthorization struct {
	ID                int                              `json:"id"`
	Status            enums.MandateAuthorizationStatus `json:"status"`
	MandateID         int                              `json:"mandate_id"`
	AuthorizationID   int                              `json:"authorization_id"`
	AuthorizationCode string                           `json:"authorization_code"`
	IntegrationID     int                              `json:"integration_id"`
	AccountNumber     string                           `json:"account_number"`
	BankCode          string                           `json:"bank_code"`
	BankName          string                           `json:"bank_name"`
	CustomerCode      string                           `json:"customer_code"`
	CreatedAt         DateTime                         `json:"created_at"`
	UpdatedAt         DateTime                         `json:"updated_at"`
}
