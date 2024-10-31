package types

import (
	"encoding/json"
)

type Channel int

const (
	ChannelUnknown Channel = iota
	ChannelCard
	ChannelBank
	ChannelUSSD
	ChannelQR
	ChannelMobileMoney
	ChannelBankTransfer
	ChannelEFT
)

func (c Channel) String() string {
	switch c {
	case ChannelCard:
		return "card"
	case ChannelBank:
		return "bank"
	case ChannelUSSD:
		return "ussd"
	case ChannelQR:
		return "qr"
	case ChannelMobileMoney:
		return "mobile_money"
	case ChannelBankTransfer:
		return "bank_transfer"
	case ChannelEFT:
		return "eft"
	default:
		return ""
	}
}

func (c Channel) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *Channel) UnmarshalJSON(data []byte) error {
	var str string

	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	switch str {
	case "card":
		*c = ChannelCard
	case "bank":
		*c = ChannelBank
	case "ussd":
		*c = ChannelUSSD
	case "qr":
		*c = ChannelQR
	case "mobile_money":
		*c = ChannelMobileMoney
	case "bank_transfer":
		*c = ChannelBankTransfer
	case "eft":
		*c = ChannelEFT
	default:
		*c = ChannelUnknown
	}

	return nil
}
