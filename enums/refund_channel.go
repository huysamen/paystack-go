package enums

import (
	"encoding/json"
	"fmt"
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

// UnmarshalJSON implements json.Unmarshaler
func (rc *RefundChannel) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	channel := RefundChannel(s)
	switch channel {
	case RefundChannelCard, RefundChannelBank, RefundChannelMIGS:
		*rc = channel
		return nil
	default:
		return fmt.Errorf("invalid RefundChannel value: %s", s)
	}
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
