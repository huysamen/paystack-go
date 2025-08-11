package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Terminal represents a Paystack terminal device
type Terminal struct {
	ID           data.Uint            `json:"id"`
	SerialNumber data.String          `json:"serial_number"`
	DeviceMake   data.NullString      `json:"device_make"`
	TerminalID   data.String          `json:"terminal_id"`
	Integration  data.Uint            `json:"integration"`
	Domain       data.String          `json:"domain"`
	Name         data.String          `json:"name"`
	Address      data.NullString      `json:"address"`
	Status       enums.TerminalStatus `json:"status"`
}

// TerminalEventData represents the data payload for terminal events
type TerminalEventData map[string]any

// TerminalEventResult represents the result of sending an event
type TerminalEventResult struct {
	ID data.String `json:"id"` // Event ID
}

// TerminalEventStatus represents the status of a terminal event
type TerminalEventStatus struct {
	Delivered data.Bool `json:"delivered"` // Whether event was delivered to terminal
}

// TerminalPresenceStatus represents the presence status of a terminal
type TerminalPresenceStatus struct {
	Online    data.Bool `json:"online"`    // Whether terminal is online
	Available data.Bool `json:"available"` // Whether terminal is available for events
}
