package types

// AccountResolution represents the result of resolving an account number
type AccountResolution struct {
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
}

// AccountValidation represents the result of validating an account
type AccountValidation struct {
	Verified            bool   `json:"verified"`
	VerificationMessage string `json:"verificationMessage"`
}

// CardBINResolution represents the result of resolving a card BIN
type CardBINResolution struct {
	BIN          string `json:"bin"`
	Brand        string `json:"brand"`
	SubBrand     string `json:"sub_brand"`
	CountryCode  string `json:"country_code"`
	CountryName  string `json:"country_name"`
	CardType     string `json:"card_type"`
	Bank         string `json:"bank"`
	LinkedBankID int    `json:"linked_bank_id"`
}
