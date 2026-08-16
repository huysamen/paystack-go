package enums

import (
	"encoding/json"
)

// Interval represents the billing interval for subscriptions and plans
type Interval string

const (
	IntervalHourly   Interval = "hourly"
	IntervalDaily    Interval = "daily"
	IntervalWeekly   Interval = "weekly"
	IntervalMonthly  Interval = "monthly"
	IntervalBiannual Interval = "biannually"
	IntervalAnnually Interval = "annually"
)

// String returns the string representation of Interval
func (i Interval) String() string {
	return string(i)
}

// MarshalJSON implements json.Marshaler
func (i Interval) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(i))
}

// UnmarshalJSON implements json.Unmarshaler.
//
// Paystack extends this field without notice, so an unrecognised value is
// accepted and preserved verbatim rather than rejected: failing here would
// fail the decode of the entire payload over a single field. Use IsValid to
// test whether the value is one this version of the library models.
func (i *Interval) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*i = Interval(s)

	return nil
}

// IsValid returns true if the interval is a valid known value
func (i Interval) IsValid() bool {
	switch i {
	case IntervalHourly, IntervalDaily, IntervalWeekly, IntervalMonthly,
		IntervalBiannual, IntervalAnnually:
		return true
	default:
		return false
	}
}

// AllIntervals returns all valid Interval values
func AllIntervals() []Interval {
	return []Interval{
		IntervalHourly,
		IntervalDaily,
		IntervalWeekly,
		IntervalMonthly,
		IntervalBiannual,
		IntervalAnnually,
	}
}
