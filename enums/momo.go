package enums

import (
	"encoding/json"
	"fmt"
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

// UnmarshalJSON implements json.Unmarshaler
func (m *MoMo) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	momo := MoMo(s)
	switch momo {
	case MoMoMTN, MoMoAirtimeUG, MoMoVodafone, MoMoAirtel, MoMoTigo:
		*m = momo
		return nil
	default:
		return fmt.Errorf("invalid MoMo value: %s", s)
	}
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
