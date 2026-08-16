package enums

import (
	"encoding/json"
)

// DisputeCategory represents the category of a dispute
type DisputeCategory string

const (
	DisputeCategoryGeneral          DisputeCategory = "general"
	DisputeCategoryFraud            DisputeCategory = "fraud"
	DisputeCategoryAuthorization    DisputeCategory = "authorization"
	DisputeCategoryProcessingErrors DisputeCategory = "processing_errors"
	DisputeCategoryConsumerDispute  DisputeCategory = "consumer_dispute"
)

// String returns the string representation of DisputeCategory
func (dc DisputeCategory) String() string {
	return string(dc)
}

// MarshalJSON implements json.Marshaler
func (dc DisputeCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(dc))
}

// UnmarshalJSON implements json.Unmarshaler.
//
// Paystack extends this field without notice, so an unrecognised value is
// accepted and preserved verbatim rather than rejected: failing here would
// fail the decode of the entire payload over a single field. Use IsValid to
// test whether the value is one this version of the library models.
func (dc *DisputeCategory) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*dc = DisputeCategory(s)

	return nil
}

// IsValid returns true if the dispute category is a valid known value
func (dc DisputeCategory) IsValid() bool {
	switch dc {
	case DisputeCategoryGeneral, DisputeCategoryFraud, DisputeCategoryAuthorization,
		DisputeCategoryProcessingErrors, DisputeCategoryConsumerDispute:
		return true
	default:
		return false
	}
}

// AllDisputeCategories returns all valid DisputeCategory values
func AllDisputeCategories() []DisputeCategory {
	return []DisputeCategory{
		DisputeCategoryGeneral,
		DisputeCategoryFraud,
		DisputeCategoryAuthorization,
		DisputeCategoryProcessingErrors,
		DisputeCategoryConsumerDispute,
	}
}
