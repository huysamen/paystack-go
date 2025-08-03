package terminal

import (
	"github.com/huysamen/paystack-go/types"
)

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

// TerminalSendEventResponse represents the response from sending an event
type TerminalSendEventResponse = types.Response[TerminalEventResult]

// TerminalEventStatusResponse represents the response from fetching event status
type TerminalEventStatusResponse = types.Response[TerminalEventStatus]

// TerminalPresenceResponse represents the response from checking terminal presence
type TerminalPresenceResponse = types.Response[TerminalPresenceStatus]

// Terminal List

// TerminalListRequest represents the request to list terminals
type TerminalListRequest struct {
	PerPage  *int    `json:"perPage,omitempty"`  // Number of terminals per page (default: 50)
	Next     *string `json:"next,omitempty"`     // Cursor for next page (optional)
	Previous *string `json:"previous,omitempty"` // Cursor for previous page (optional)
}

// TerminalListRequestBuilder provides a fluent interface for building TerminalListRequest
type TerminalListRequestBuilder struct {
	req *TerminalListRequest
}

// NewTerminalListRequest creates a new builder for TerminalListRequest
func NewTerminalListRequest() *TerminalListRequestBuilder {
	return &TerminalListRequestBuilder{
		req: &TerminalListRequest{},
	}
}

// PerPage sets the number of terminals per page
func (b *TerminalListRequestBuilder) PerPage(perPage int) *TerminalListRequestBuilder {
	b.req.PerPage = &perPage
	return b
}

// Next sets the cursor for next page
func (b *TerminalListRequestBuilder) Next(next string) *TerminalListRequestBuilder {
	b.req.Next = &next
	return b
}

// Previous sets the cursor for previous page
func (b *TerminalListRequestBuilder) Previous(previous string) *TerminalListRequestBuilder {
	b.req.Previous = &previous
	return b
}

// Build returns the constructed TerminalListRequest
func (b *TerminalListRequestBuilder) Build() *TerminalListRequest {
	return b.req
}

// TerminalListResponse represents the response from listing terminals
type TerminalListResponse = types.Response[[]Terminal]

// Terminal Fetch

// TerminalFetchResponse represents the response from fetching a terminal
type TerminalFetchResponse = types.Response[Terminal]

// Terminal Update

// TerminalUpdateRequest represents the request to update a terminal
type TerminalUpdateRequest struct {
	Name    *string `json:"name,omitempty"`    // Name of the terminal (optional)
	Address *string `json:"address,omitempty"` // Address of the terminal (optional)
}

// TerminalUpdateRequestBuilder provides a fluent interface for building TerminalUpdateRequest
type TerminalUpdateRequestBuilder struct {
	req *TerminalUpdateRequest
}

// NewTerminalUpdateRequest creates a new builder for TerminalUpdateRequest
func NewTerminalUpdateRequest() *TerminalUpdateRequestBuilder {
	return &TerminalUpdateRequestBuilder{
		req: &TerminalUpdateRequest{},
	}
}

// Name sets the terminal name
func (b *TerminalUpdateRequestBuilder) Name(name string) *TerminalUpdateRequestBuilder {
	b.req.Name = &name
	return b
}

// Address sets the terminal address
func (b *TerminalUpdateRequestBuilder) Address(address string) *TerminalUpdateRequestBuilder {
	b.req.Address = &address
	return b
}

// Build returns the constructed TerminalUpdateRequest
func (b *TerminalUpdateRequestBuilder) Build() *TerminalUpdateRequest {
	return b.req
}

// TerminalUpdateResponse represents the response from updating a terminal
type TerminalUpdateResponse = types.Response[Terminal]

// Terminal Commission

// TerminalCommissionRequest represents the request to commission a terminal
type TerminalCommissionRequest struct {
	SerialNumber string `json:"serial_number"` // Device serial number
}

// TerminalCommissionResponse represents the response from commissioning a terminal
type TerminalCommissionResponse = types.Response[Terminal]

// Terminal Decommission

// TerminalDecommissionRequest represents the request to decommission a terminal
type TerminalDecommissionRequest struct {
	SerialNumber string `json:"serial_number"` // Device serial number
}

// TerminalDecommissionResponse represents the response from decommissioning a terminal
type TerminalDecommissionResponse = types.Response[any]
