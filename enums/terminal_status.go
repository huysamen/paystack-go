package enums

import (
	"encoding/json"
	"fmt"
)

// TerminalStatus represents the status of a terminal
type TerminalStatus string

const (
	TerminalStatusActive   TerminalStatus = "active"
	TerminalStatusInactive TerminalStatus = "inactive"
	TerminalStatusPending  TerminalStatus = "pending"
)

// String returns the string representation of TerminalStatus
func (ts TerminalStatus) String() string {
	return string(ts)
}

// MarshalJSON implements json.Marshaler
func (ts TerminalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(ts))
}

// UnmarshalJSON implements json.Unmarshaler
func (ts *TerminalStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	status := TerminalStatus(s)
	switch status {
	case TerminalStatusActive, TerminalStatusInactive, TerminalStatusPending:
		*ts = status
		return nil
	default:
		return fmt.Errorf("invalid TerminalStatus value: %s", s)
	}
}

// IsValid returns true if the terminal status is a valid known value
func (ts TerminalStatus) IsValid() bool {
	switch ts {
	case TerminalStatusActive, TerminalStatusInactive, TerminalStatusPending:
		return true
	default:
		return false
	}
}

// AllTerminalStatuses returns all valid TerminalStatus values
func AllTerminalStatuses() []TerminalStatus {
	return []TerminalStatus{
		TerminalStatusActive,
		TerminalStatusInactive,
		TerminalStatusPending,
	}
}
