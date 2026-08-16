package enums

import (
	"encoding/json"
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
	ChannelApplePay     Channel = "apple_pay"
	ChannelCapitecPay   Channel = "capitec_pay"
)

// String returns the string representation of the channel
func (c Channel) String() string {
	return string(c)
}

// MarshalJSON implements json.Marshaler
func (c Channel) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(c))
}

// UnmarshalJSON implements json.Unmarshaler.
//
// Paystack introduces new channels without notice, so an unrecognised value is
// accepted and preserved verbatim rather than rejected: failing here would fail
// the decode of the entire payload (a webhook event, a transaction) over a
// single informational field. Use IsValid to test whether the value is one this
// version of the library knows about.
func (c *Channel) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*c = Channel(s)

	return nil
}

// IsValid returns true if the channel is a valid known value
func (c Channel) IsValid() bool {
	switch c {
	case ChannelCard, ChannelBank, ChannelUSSD, ChannelQR,
		ChannelMobileMoney, ChannelBankTransfer, ChannelEFT, ChannelApplePay,
		ChannelCapitecPay, "":
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
		ChannelApplePay,
		ChannelCapitecPay,
	}
}
