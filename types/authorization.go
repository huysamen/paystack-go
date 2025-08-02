package types

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
