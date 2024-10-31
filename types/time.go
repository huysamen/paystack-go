package types

import "time"

type DateTime struct {
	time.Time
}

func (ct *DateTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1]

	parsedTime, err := time.Parse(time.DateTime, str)
	if err != nil {
		return err
	}

	ct.Time = parsedTime

	return nil
}
