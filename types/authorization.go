package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Authorization represents a payment authorization
type Authorization struct {
	AuthorizationCode         string           `json:"authorization_code"`
	Bin                       string           `json:"bin"`
	Last4                     string           `json:"last4"`
	ExpMonth                  data.MultiString `json:"exp_month"`
	ExpYear                   data.MultiString `json:"exp_year"`
	Channel                   enums.Channel    `json:"channel"`
	CardType                  string           `json:"card_type"`
	Brand                     string           `json:"brand"`
	Bank                      string           `json:"bank"`
	CountryCode               string           `json:"country_code"`
	CountryName               data.NullString  `json:"country_name,omitempty"`
	Reusable                  bool             `json:"reusable"`
	Signature                 string           `json:"signature"`
	AccountName               data.NullString  `json:"account_name"`
	ReceiverBankAccountNumber data.NullString  `json:"receiver_bank_account_number,omitempty"`
	ReceiverBank              data.NullString  `json:"receiver_bank,omitempty"`
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
	CreatedAt         data.MultiDateTime               `json:"created_at"`
	UpdatedAt         data.MultiDateTime               `json:"updated_at"`
}
