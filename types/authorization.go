package types

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/huysamen/paystack-go/enums"
)

// FlexibleString handles both string and number JSON values
type FlexibleString string

// UnmarshalJSON implements json.Unmarshaler for FlexibleString
func (fs *FlexibleString) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		*fs = FlexibleString(s)
		return nil
	}

	var n float64
	if err := json.Unmarshal(data, &n); err == nil {
		*fs = FlexibleString(strconv.FormatFloat(n, 'f', 0, 64))
		return nil
	}

	return fmt.Errorf("cannot unmarshal %s into FlexibleString", string(data))
}

// String returns the string representation
func (fs FlexibleString) String() string {
	return string(fs)
}

// MarshalJSON implements json.Marshaler
func (fs FlexibleString) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(fs))
}

// Authorization represents a payment authorization
type Authorization struct {
	AuthorizationCode         string          `json:"authorization_code"`
	Bin                       string          `json:"bin"`
	Last4                     string          `json:"last4"`
	ExpMonth                  FlexibleString  `json:"exp_month"`
	ExpYear                   FlexibleString  `json:"exp_year"`
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
