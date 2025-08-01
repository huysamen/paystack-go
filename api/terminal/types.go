package terminal

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

// TerminalSendEventResponse represents the response from sending an event
type TerminalSendEventResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID string `json:"id"` // Event ID
	} `json:"data"`
}

// Terminal Event Status

// TerminalEventStatusResponse represents the response from fetching event status
type TerminalEventStatusResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Delivered bool `json:"delivered"` // Whether event was delivered to terminal
	} `json:"data"`
}

// Terminal Presence/Status

// TerminalPresenceResponse represents the response from checking terminal presence
type TerminalPresenceResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Online    bool `json:"online"`    // Whether terminal is online
		Available bool `json:"available"` // Whether terminal is available for events
	} `json:"data"`
}

// Terminal List

// TerminalListRequest represents the request to list terminals
type TerminalListRequest struct {
	PerPage  *int    `json:"perPage,omitempty"`  // Number of terminals per page (default: 50)
	Next     *string `json:"next,omitempty"`     // Cursor for next page (optional)
	Previous *string `json:"previous,omitempty"` // Cursor for previous page (optional)
}

// TerminalListResponse represents the response from listing terminals
type TerminalListResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    []Terminal `json:"data"`
	Meta    struct {
		Next     *string `json:"next"`     // Cursor for next page
		Previous *string `json:"previous"` // Cursor for previous page
		PerPage  int     `json:"perPage"`  // Records per page
	} `json:"meta"`
}

// Terminal Fetch

// TerminalFetchResponse represents the response from fetching a terminal
type TerminalFetchResponse struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Data    Terminal `json:"data"`
}

// Terminal Update

// TerminalUpdateRequest represents the request to update a terminal
type TerminalUpdateRequest struct {
	Name    *string `json:"name,omitempty"`    // Name of the terminal (optional)
	Address *string `json:"address,omitempty"` // Address of the terminal (optional)
}

// TerminalUpdateResponse represents the response from updating a terminal
type TerminalUpdateResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// Terminal Commission

// TerminalCommissionRequest represents the request to commission a terminal
type TerminalCommissionRequest struct {
	SerialNumber string `json:"serial_number"` // Device serial number
}

// TerminalCommissionResponse represents the response from commissioning a terminal
type TerminalCommissionResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// Terminal Decommission

// TerminalDecommissionRequest represents the request to decommission a terminal
type TerminalDecommissionRequest struct {
	SerialNumber string `json:"serial_number"` // Device serial number
}

// TerminalDecommissionResponse represents the response from decommissioning a terminal
type TerminalDecommissionResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
