package enums

import (
	"encoding/json"
	"fmt"
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

// UnmarshalJSON implements json.Unmarshaler
func (tea *TerminalEventAction) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	action := TerminalEventAction(s)
	switch action {
	case TerminalEventActionProcess, TerminalEventActionView, TerminalEventActionPrint:
		*tea = action
		return nil
	default:
		return fmt.Errorf("invalid TerminalEventAction value: %s", s)
	}
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
