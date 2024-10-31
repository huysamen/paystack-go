package types

import (
	"encoding/json"
)

type Interval int

const (
	IntervalUnknown Interval = iota
	IntervalHourly
	IntervalDaily
	IntervalWeekly
	IntervalMonthly
	IntervalQuarterly
	IntervalBiannually
	IntervalAnnually
)

func (i Interval) String() string {
	switch i {
	case IntervalHourly:
		return "hourly"
	case IntervalDaily:
		return "daily"
	case IntervalWeekly:
		return "weekly"
	case IntervalMonthly:
		return "monthly"
	case IntervalQuarterly:
		return "quarterly"
	case IntervalBiannually:
		return "biannually"
	case IntervalAnnually:
		return "annually"
	default:
		return ""
	}
}

func (i Interval) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *Interval) UnmarshalJSON(data []byte) error {
	var str string

	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	switch str {
	case "hourly":
		*i = IntervalHourly
	case "daily":
		*i = IntervalDaily
	case "weekly":
		*i = IntervalWeekly
	case "monthly":
		*i = IntervalMonthly
	case "quarterly":
		*i = IntervalQuarterly
	case "biannually":
		*i = IntervalBiannually
	case "annually":
		*i = IntervalAnnually
	default:
		*i = IntervalUnknown
	}

	return nil
}
