package enums

import (
	"encoding/json"
)

// TerminalEventAction represents the action for terminal events
type TerminalEventAction string

const (
	TerminalEventActionProcess TerminalEventAction = "process"
	TerminalEventActionView    TerminalEventAction = "view"
	TerminalEventActionPrint   TerminalEventAction = "print"
)

// String returns the string representation of TerminalEventAction
func (tea TerminalEventAction) String() string {
	return string(tea)
}

// MarshalJSON implements json.Marshaler
func (tea TerminalEventAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(tea))
}

// UnmarshalJSON implements json.Unmarshaler.
//
// Paystack extends this field without notice, so an unrecognised value is
// accepted and preserved verbatim rather than rejected: failing here would
// fail the decode of the entire payload over a single field. Use IsValid to
// test whether the value is one this version of the library models.
func (tea *TerminalEventAction) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*tea = TerminalEventAction(s)

	return nil
}

// IsValid returns true if the terminal event action is a valid known value
func (tea TerminalEventAction) IsValid() bool {
	switch tea {
	case TerminalEventActionProcess, TerminalEventActionView, TerminalEventActionPrint:
		return true
	default:
		return false
	}
}

// AllTerminalEventActions returns all valid TerminalEventAction values
func AllTerminalEventActions() []TerminalEventAction {
	return []TerminalEventAction{
		TerminalEventActionProcess,
		TerminalEventActionView,
		TerminalEventActionPrint,
	}
}
