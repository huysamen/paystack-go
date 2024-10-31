package types

import (
	"encoding/json"
)

type MoMo int

const (
	MoMoUnknown MoMo = iota
	MoMoMTN
	MoMoVodafone
	MoMoAirtelTigo
)

func (c MoMo) String() string {
	switch c {
	case MoMoMTN:
		return "mtn"
	case MoMoVodafone:
		return "vod"
	case MoMoAirtelTigo:
		return "atl"
	default:
		return ""
	}
}

func (c MoMo) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *MoMo) UnmarshalJSON(data []byte) error {
	var str string

	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	switch str {
	case "mtn":
		*c = MoMoMTN
	case "vod":
		*c = MoMoVodafone
	case "atl":
		*c = MoMoAirtelTigo
	default:
		*c = MoMoUnknown
	}

	return nil
}
