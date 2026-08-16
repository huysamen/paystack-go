package enums

import (
	"encoding/json"
)

// TransferRecipientType represents the type of transfer recipient
type TransferRecipientType string

const (
	TransferRecipientTypeNuban            TransferRecipientType = "nuban"
	TransferRecipientTypeMobileMoney      TransferRecipientType = "mobile_money"
	TransferRecipientTypeAuthorization    TransferRecipientType = "authorization"
	TransferRecipientTypeGhipss           TransferRecipientType = "ghipss"
	TransferRecipientTypeDomesticCGI      TransferRecipientType = "domiciliary_cgi"
	TransferRecipientTypeBankAccount      TransferRecipientType = "bank_account"
	TransferRecipientTypeInternationalCGI TransferRecipientType = "international_cgi"
)

// String returns the string representation of TransferRecipientType
func (trt TransferRecipientType) String() string {
	return string(trt)
}

// MarshalJSON implements json.Marshaler
func (trt TransferRecipientType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(trt))
}

// UnmarshalJSON implements json.Unmarshaler.
//
// Paystack extends this field without notice, so an unrecognised value is
// accepted and preserved verbatim rather than rejected: failing here would
// fail the decode of the entire payload over a single field. Use IsValid to
// test whether the value is one this version of the library models.
func (trt *TransferRecipientType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*trt = TransferRecipientType(s)

	return nil
}

// IsValid returns true if the transfer recipient type is a valid known value
func (trt TransferRecipientType) IsValid() bool {
	switch trt {
	case TransferRecipientTypeNuban, TransferRecipientTypeMobileMoney, TransferRecipientTypeAuthorization,
		TransferRecipientTypeGhipss, TransferRecipientTypeDomesticCGI, TransferRecipientTypeBankAccount,
		TransferRecipientTypeInternationalCGI:
		return true
	default:
		return false
	}
}

// AllTransferRecipientTypes returns all valid TransferRecipientType values
func AllTransferRecipientTypes() []TransferRecipientType {
	return []TransferRecipientType{
		TransferRecipientTypeNuban,
		TransferRecipientTypeMobileMoney,
		TransferRecipientTypeAuthorization,
		TransferRecipientTypeGhipss,
		TransferRecipientTypeDomesticCGI,
		TransferRecipientTypeBankAccount,
		TransferRecipientTypeInternationalCGI,
	}
}
