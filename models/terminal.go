package models

import "github.com/huysamen/paystack-go/enums"

// Terminal represents a Paystack terminal device
type Terminal struct {
	ID           uint64               `json:"id"`
	SerialNumber string               `json:"serial_number"`
	DeviceMake   *string              `json:"device_make"`
	TerminalID   string               `json:"terminal_id"`
	Integration  uint64               `json:"integration"`
	Domain       string               `json:"domain"`
	Name         string               `json:"name"`
	Address      *string              `json:"address"`
	Status       enums.TerminalStatus `json:"status"`
}

// TerminalEventData represents the data payload for terminal events
type TerminalEventData map[string]any

// TerminalEventResult represents the result of sending an event
type TerminalEventResult struct {
	ID string `json:"id"` // Event ID
}

// TerminalEventStatus represents the status of a terminal event
type TerminalEventStatus struct {
	Delivered bool `json:"delivered"` // Whether event was delivered to terminal
}

// TerminalPresenceStatus represents the presence status of a terminal
type TerminalPresenceStatus struct {
	Online    bool `json:"online"`    // Whether terminal is online
	Available bool `json:"available"` // Whether terminal is available for events
}
