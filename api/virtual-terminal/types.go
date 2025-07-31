package virtualterminal

import (
	"github.com/huysamen/paystack-go/types"
)

// VirtualTerminalDestination represents a notification destination for a virtual terminal
type VirtualTerminalDestination struct {
	ID        int    `json:"id,omitempty"`
	Target    string `json:"target"`
	Name      string `json:"name"`
	Type      string `json:"type,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// VirtualTerminal represents a virtual terminal
type VirtualTerminal struct {
	ID             int                          `json:"id"`
	Code           string                       `json:"code"`
	Name           string                       `json:"name"`
	Integration    int                          `json:"integration"`
	Domain         string                       `json:"domain"`
	PaymentMethods []string                     `json:"paymentMethods"`
	Active         bool                         `json:"active"`
	CreatedAt      string                       `json:"created_at,omitempty"`
	Metadata       *types.Metadata              `json:"metadata,omitempty"`
	Destinations   []VirtualTerminalDestination `json:"destinations,omitempty"`
	Currency       string                       `json:"currency"`
	CustomFields   []CustomField                `json:"custom_fields,omitempty"`
	ConnectAccount *int                         `json:"connect_account_id,omitempty"`
}

// CustomField represents a custom field for the virtual terminal form
type CustomField struct {
	DisplayName  string `json:"display_name"`
	VariableName string `json:"variable_name"`
}

// CreateVirtualTerminalRequest represents the request to create a virtual terminal
type CreateVirtualTerminalRequest struct {
	Name         string                       `json:"name"`
	Destinations []VirtualTerminalDestination `json:"destinations,omitempty"`
	Metadata     *types.Metadata              `json:"metadata,omitempty"`
	Currency     string                       `json:"currency,omitempty"`
	CustomFields []CustomField                `json:"custom_fields,omitempty"`
}

// ListVirtualTerminalsRequest represents the request to list virtual terminals
type ListVirtualTerminalsRequest struct {
	Status   string `json:"status,omitempty"`
	PerPage  int    `json:"perPage,omitempty"`
	Search   string `json:"search,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
}

// UpdateVirtualTerminalRequest represents the request to update a virtual terminal
type UpdateVirtualTerminalRequest struct {
	Name string `json:"name"`
}

// AssignDestinationRequest represents the request to assign destinations to a virtual terminal
type AssignDestinationRequest struct {
	Destinations []VirtualTerminalDestination `json:"destinations"`
}

// UnassignDestinationRequest represents the request to unassign destinations from a virtual terminal
type UnassignDestinationRequest struct {
	Targets []string `json:"targets"`
}

// AddSplitCodeRequest represents the request to add a split code to a virtual terminal
type AddSplitCodeRequest struct {
	SplitCode string `json:"split_code"`
}

// RemoveSplitCodeRequest represents the request to remove a split code from a virtual terminal
type RemoveSplitCodeRequest struct {
	SplitCode string `json:"split_code"`
}

// CreateVirtualTerminalResponse represents the response from creating a virtual terminal
type CreateVirtualTerminalResponse struct {
	Status  bool            `json:"status"`
	Message string          `json:"message"`
	Data    VirtualTerminal `json:"data"`
}

// ListVirtualTerminalsResponse represents the response from listing virtual terminals
type ListVirtualTerminalsResponse struct {
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Data    []VirtualTerminal `json:"data"`
	Meta    *types.Meta       `json:"meta,omitempty"`
}

// FetchVirtualTerminalResponse represents the response from fetching a virtual terminal
type FetchVirtualTerminalResponse struct {
	Status  bool            `json:"status"`
	Message string          `json:"message"`
	Data    VirtualTerminal `json:"data"`
}

// UpdateVirtualTerminalResponse represents the response from updating a virtual terminal
type UpdateVirtualTerminalResponse struct {
	Status  bool            `json:"status"`
	Message string          `json:"message"`
	Data    VirtualTerminal `json:"data"`
}

// DeactivateVirtualTerminalResponse represents the response from deactivating a virtual terminal
type DeactivateVirtualTerminalResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// AssignDestinationResponse represents the response from assigning destinations
type AssignDestinationResponse struct {
	Status  bool                         `json:"status"`
	Message string                       `json:"message"`
	Data    []VirtualTerminalDestination `json:"data"`
}

// UnassignDestinationResponse represents the response from unassigning destinations
type UnassignDestinationResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// AddSplitCodeResponse represents the response from adding a split code
type AddSplitCodeResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"` // Using interface{} as the split response structure can vary
}

// RemoveSplitCodeResponse represents the response from removing a split code
type RemoveSplitCodeResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
