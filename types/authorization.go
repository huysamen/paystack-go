package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Authorization represents a payment authorization
type Authorization struct {
	AuthorizationCode         data.String     `json:"authorization_code"`
	Bin                       data.String     `json:"bin"`
	Last4                     data.String     `json:"last4"`
	ExpMonth                  data.String     `json:"exp_month"`
	ExpYear                   data.String     `json:"exp_year"`
	Channel                   enums.Channel   `json:"channel"`
	CardType                  data.String     `json:"card_type"`
	Brand                     data.String     `json:"brand"`
	Bank                      data.String     `json:"bank"`
	CountryCode               data.String     `json:"country_code"`
	CountryName               data.NullString `json:"country_name,omitempty"`
	Reusable                  data.Bool       `json:"reusable"`
	Signature                 data.String     `json:"signature"`
	AccountName               data.NullString `json:"account_name"`
	ReceiverBankAccountNumber data.NullString `json:"receiver_bank_account_number,omitempty"`
	ReceiverBank              data.NullString `json:"receiver_bank,omitempty"`
}

// MandateAuthorization represents a mandate authorization
type MandateAuthorization struct {
	ID                data.Int                         `json:"id"`
	Status            enums.MandateAuthorizationStatus `json:"status"`
	MandateID         data.Int                         `json:"mandate_id"`
	AuthorizationID   data.Int                         `json:"authorization_id"`
	AuthorizationCode data.String                      `json:"authorization_code"`
	IntegrationID     data.Int                         `json:"integration_id"`
	AccountNumber     data.String                      `json:"account_number"`
	BankCode          data.String                      `json:"bank_code"`
	BankName          data.String                      `json:"bank_name"`
	CustomerCode      data.String                      `json:"customer_code"`
	CreatedAt         data.Time                        `json:"created_at"`
	UpdatedAt         data.Time                        `json:"updated_at"`
}
