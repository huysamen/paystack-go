package enums

import (
	"encoding/json"
	"fmt"
)

// TerminalEventType represents the type of terminal event
type TerminalEventType string

const (
	TerminalEventTypeInvoice     TerminalEventType = "invoice"
	TerminalEventTypePayment     TerminalEventType = "payment"
	TerminalEventTypeTransaction TerminalEventType = "transaction"
)

// String returns the string representation of TerminalEventType
func (tet TerminalEventType) String() string {
	return string(tet)
}

// MarshalJSON implements json.Marshaler
func (tet TerminalEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(tet))
}

// UnmarshalJSON implements json.Unmarshaler
func (tet *TerminalEventType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	eventType := TerminalEventType(s)
	switch eventType {
	case TerminalEventTypeInvoice, TerminalEventTypePayment, TerminalEventTypeTransaction:
		*tet = eventType
		return nil
	default:
		return fmt.Errorf("invalid TerminalEventType value: %s", s)
	}
}

// IsValid returns true if the terminal event type is a valid known value
func (tet TerminalEventType) IsValid() bool {
	switch tet {
	case TerminalEventTypeInvoice, TerminalEventTypePayment, TerminalEventTypeTransaction:
		return true
	default:
		return false
	}
}

// AllTerminalEventTypes returns all valid TerminalEventType values
func AllTerminalEventTypes() []TerminalEventType {
	return []TerminalEventType{
		TerminalEventTypeInvoice,
		TerminalEventTypePayment,
		TerminalEventTypeTransaction,
	}
}
