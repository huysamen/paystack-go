package types

// Interval represents subscription billing intervals
type Interval string

const (
	IntervalUnknown    Interval = ""
	IntervalHourly     Interval = "hourly"
	IntervalDaily      Interval = "daily"
	IntervalWeekly     Interval = "weekly"
	IntervalMonthly    Interval = "monthly"
	IntervalQuarterly  Interval = "quarterly"
	IntervalBiannually Interval = "biannually"
	IntervalAnnually   Interval = "annually"
)

// String returns the string representation of the interval
func (i Interval) String() string {
	return string(i)
}

// IsValid returns true if the interval is a valid known value
func (i Interval) IsValid() bool {
	switch i {
	case IntervalHourly, IntervalDaily, IntervalWeekly,
		IntervalMonthly, IntervalQuarterly, IntervalBiannually, IntervalAnnually:
		return true
	default:
		return false
	}
}
