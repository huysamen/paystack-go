package types

// MoMo represents mobile money providers supported by Paystack
type MoMo string

const (
	MoMoUnknown    MoMo = ""
	MoMoMTN        MoMo = "mtn"
	MoMoVodafone   MoMo = "vod"
	MoMoAirtelTigo MoMo = "atl"
)

// String returns the string representation of the mobile money provider
func (m MoMo) String() string {
	return string(m)
}

// IsValid returns true if the mobile money provider is a valid known value
func (m MoMo) IsValid() bool {
	switch m {
	case MoMoMTN, MoMoVodafone, MoMoAirtelTigo:
		return true
	default:
		return false
	}
}
