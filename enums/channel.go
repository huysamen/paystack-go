package enums

import (
	"encoding/json"
	"fmt"
)

// Channel represents payment channels supported by Paystack
type Channel string

const (
	ChannelCard         Channel = "card"
	ChannelBank         Channel = "bank"
	ChannelUSSD         Channel = "ussd"
	ChannelQR           Channel = "qr"
	ChannelMobileMoney  Channel = "mobile_money"
	ChannelBankTransfer Channel = "bank_transfer"
	ChannelEFT          Channel = "eft"
)

// String returns the string representation of the channel
func (c Channel) String() string {
	return string(c)
}

// MarshalJSON implements json.Marshaler
func (c Channel) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(c))
}

// UnmarshalJSON implements json.Unmarshaler
func (c *Channel) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	channel := Channel(s)
	switch channel {
	case ChannelCard, ChannelBank, ChannelUSSD, ChannelQR,
		ChannelMobileMoney, ChannelBankTransfer, ChannelEFT:
		*c = channel
		return nil
	default:
		return fmt.Errorf("invalid Channel value: %s", s)
	}
}

// IsValid returns true if the channel is a valid known value
func (c Channel) IsValid() bool {
	switch c {
	case ChannelCard, ChannelBank, ChannelUSSD, ChannelQR,
		ChannelMobileMoney, ChannelBankTransfer, ChannelEFT:
		return true
	default:
		return false
	}
}

// AllChannels returns all valid Channel values
func AllChannels() []Channel {
	return []Channel{
		ChannelCard,
		ChannelBank,
		ChannelUSSD,
		ChannelQR,
		ChannelMobileMoney,
		ChannelBankTransfer,
		ChannelEFT,
	}
}
