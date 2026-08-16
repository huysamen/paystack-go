package enums

import (
	"encoding/json"
)

// RefundChannel represents the channel through which a refund is processed
type RefundChannel string

const (
	RefundChannelCard RefundChannel = "card"
	RefundChannelBank RefundChannel = "bank"
	RefundChannelMIGS RefundChannel = "migs"
)

// String returns the string representation of RefundChannel
func (rc RefundChannel) String() string {
	return string(rc)
}

// MarshalJSON implements json.Marshaler
func (rc RefundChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(rc))
}

// UnmarshalJSON implements json.Unmarshaler.
//
// Paystack extends this field without notice, so an unrecognised value is
// accepted and preserved verbatim rather than rejected: failing here would
// fail the decode of the entire payload over a single field. Use IsValid to
// test whether the value is one this version of the library models.
func (rc *RefundChannel) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*rc = RefundChannel(s)

	return nil
}

// IsValid returns true if the refund channel is a valid known value
func (rc RefundChannel) IsValid() bool {
	switch rc {
	case RefundChannelCard, RefundChannelBank, RefundChannelMIGS:
		return true
	default:
		return false
	}
}

// AllRefundChannels returns all valid RefundChannel values
func AllRefundChannels() []RefundChannel {
	return []RefundChannel{
		RefundChannelCard,
		RefundChannelBank,
		RefundChannelMIGS,
	}
}
