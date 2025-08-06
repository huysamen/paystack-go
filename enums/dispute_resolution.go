package enums

import (
	"encoding/json"
	"fmt"
)

// DisputeResolution represents the resolution of a dispute
type DisputeResolution string

const (
	DisputeResolutionMerchantAccepted DisputeResolution = "merchant-accepted"
	DisputeResolutionDeclined         DisputeResolution = "declined"
)

// String returns the string representation of DisputeResolution
func (dr DisputeResolution) String() string {
	return string(dr)
}

// MarshalJSON implements json.Marshaler
func (dr DisputeResolution) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(dr))
}

// UnmarshalJSON implements json.Unmarshaler
func (dr *DisputeResolution) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	resolution := DisputeResolution(s)
	switch resolution {
	case DisputeResolutionMerchantAccepted, DisputeResolutionDeclined:
		*dr = resolution
		return nil
	default:
		return fmt.Errorf("invalid DisputeResolution value: %s", s)
	}
}

// IsValid returns true if the dispute resolution is a valid known value
func (dr DisputeResolution) IsValid() bool {
	switch dr {
	case DisputeResolutionMerchantAccepted, DisputeResolutionDeclined:
		return true
	default:
		return false
	}
}

// AllDisputeResolutions returns all valid DisputeResolution values
func AllDisputeResolutions() []DisputeResolution {
	return []DisputeResolution{
		DisputeResolutionMerchantAccepted,
		DisputeResolutionDeclined,
	}
}
