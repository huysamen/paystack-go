package types

import (
	"encoding/json"
)

type Bearer int

const (
	BearerUnknown Bearer = iota
	BearerAccount
	BearerSubaccount
)

func (b Bearer) String() string {
	switch b {
	case BearerAccount:
		return "account"
	case BearerSubaccount:
		return "subaccount"
	default:
		return ""
	}
}

func (b Bearer) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *Bearer) UnmarshalJSON(data []byte) error {
	var str string

	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	switch str {
	case "account":
		*b = BearerAccount
	case "subaccount":
		*b = BearerSubaccount
	default:
		*b = BearerUnknown
	}

	return nil
}
