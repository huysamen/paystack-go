package enums

import (
	"encoding/json"
	"fmt"
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

// UnmarshalJSON implements json.Unmarshaler
func (i *Interval) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	interval := Interval(s)
	switch interval {
	case IntervalHourly, IntervalDaily, IntervalWeekly, IntervalMonthly,
		IntervalBiannual, IntervalAnnually:
		*i = interval
		return nil
	default:
		return fmt.Errorf("invalid Interval value: %s", s)
	}
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
