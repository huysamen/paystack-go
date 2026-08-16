package enums

import (
	"encoding/json"
)

// MoMo represents the mobile money provider
type MoMo string

const (
	MoMoMTN       MoMo = "mtn"
	MoMoAirtimeUG MoMo = "airtime_ug"
	MoMoVodafone  MoMo = "vodafone"
	MoMoAirtel    MoMo = "airtel"
	MoMoTigo      MoMo = "tigo"
)

// String returns the string representation of MoMo
func (m MoMo) String() string {
	return string(m)
}

// MarshalJSON implements json.Marshaler
func (m MoMo) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(m))
}

// UnmarshalJSON implements json.Unmarshaler.
//
// Paystack extends this field without notice, so an unrecognised value is
// accepted and preserved verbatim rather than rejected: failing here would
// fail the decode of the entire payload over a single field. Use IsValid to
// test whether the value is one this version of the library models.
func (m *MoMo) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*m = MoMo(s)

	return nil
}

// IsValid returns true if the mobile money provider is a valid known value
func (m MoMo) IsValid() bool {
	switch m {
	case MoMoMTN, MoMoAirtimeUG, MoMoVodafone, MoMoAirtel, MoMoTigo:
		return true
	default:
		return false
	}
}

// AllMoMoProviders returns all valid MoMo values
func AllMoMoProviders() []MoMo {
	return []MoMo{
		MoMoMTN,
		MoMoAirtimeUG,
		MoMoVodafone,
		MoMoAirtel,
		MoMoTigo,
	}
}
