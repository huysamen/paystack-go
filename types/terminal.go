package types

// TerminalEventType represents the type of event to send to terminal
type TerminalEventType string

const (
	TerminalEventTypeInvoice     TerminalEventType = "invoice"     // Invoice event
	TerminalEventTypeTransaction TerminalEventType = "transaction" // Transaction event
)

// String returns the string representation of TerminalEventType
func (t TerminalEventType) String() string {
	return string(t)
}

// TerminalEventAction represents the action for terminal events
type TerminalEventAction string

const (
	TerminalEventActionProcess TerminalEventAction = "process" // Process action (for invoice and transaction)
	TerminalEventActionView    TerminalEventAction = "view"    // View action (for invoice only)
	TerminalEventActionPrint   TerminalEventAction = "print"   // Print action (for transaction only)
)

// String returns the string representation of TerminalEventAction
func (t TerminalEventAction) String() string {
	return string(t)
}

// TerminalStatus represents terminal status
type TerminalStatus string

const (
	TerminalStatusActive   TerminalStatus = "active"   // Terminal is active
	TerminalStatusInactive TerminalStatus = "inactive" // Terminal is inactive
)

// String returns the string representation of TerminalStatus
func (t TerminalStatus) String() string {
	return string(t)
}

// Terminal represents a Paystack terminal device
type Terminal struct {
	ID           uint64         `json:"id"`
	SerialNumber string         `json:"serial_number"`
	DeviceMake   *string        `json:"device_make"`
	TerminalID   string         `json:"terminal_id"`
	Integration  uint64         `json:"integration"`
	Domain       string         `json:"domain"`
	Name         string         `json:"name"`
	Address      *string        `json:"address"`
	Status       TerminalStatus `json:"status"`
}

// Terminal Event

// TerminalEventData represents the data payload for terminal events
type TerminalEventData map[string]any

// TerminalSendEventRequest represents the request to send an event to terminal
type TerminalSendEventRequest struct {
	Type   TerminalEventType   `json:"type"`   // Type of event (invoice or transaction)
	Action TerminalEventAction `json:"action"` // Action to perform
	Data   TerminalEventData   `json:"data"`   // Event data payload
}

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
