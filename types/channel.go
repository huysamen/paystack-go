package types

// Channel represents payment channels supported by Paystack
type Channel string

const (
	ChannelUnknown      Channel = ""
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
