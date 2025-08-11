package types

import "github.com/huysamen/paystack-go/types/data"

// AccountResolution represents the result of resolving an account number
type AccountResolution struct {
	AccountNumber data.String `json:"account_number"`
	AccountName   data.String `json:"account_name"`
}

// AccountValidation represents the result of validating an account
type AccountValidation struct {
	Verified            data.Bool   `json:"verified"`
	VerificationMessage data.String `json:"verificationMessage"`
}

// CardBINResolution represents the result of resolving a card BIN
type CardBINResolution struct {
	BIN          data.String `json:"bin"`
	Brand        data.String `json:"brand"`
	SubBrand     data.String `json:"sub_brand"`
	CountryCode  data.String `json:"country_code"`
	CountryName  data.String `json:"country_name"`
	CardType     data.String `json:"card_type"`
	Bank         data.String `json:"bank"`
	LinkedBankID data.Int    `json:"linked_bank_id"`
}
