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

// CreateVirtualTerminalRequestBuilder provides a fluent interface for building CreateVirtualTerminalRequest
type CreateVirtualTerminalRequestBuilder struct {
	req *CreateVirtualTerminalRequest
}

// NewCreateVirtualTerminalRequest creates a new builder for CreateVirtualTerminalRequest
func NewCreateVirtualTerminalRequest(name string) *CreateVirtualTerminalRequestBuilder {
	return &CreateVirtualTerminalRequestBuilder{
		req: &CreateVirtualTerminalRequest{
			Name: name,
		},
	}
}

// Build returns the constructed CreateVirtualTerminalRequest
func (b *CreateVirtualTerminalRequestBuilder) Build() *CreateVirtualTerminalRequest {
	return b.req
}

// ListVirtualTerminalsRequest represents the request to list virtual terminals
type ListVirtualTerminalsRequest struct {
	Status   string `json:"status,omitempty"`
	PerPage  int    `json:"perPage,omitempty"`
	Search   string `json:"search,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
}

// ListVirtualTerminalsRequestBuilder provides a fluent interface for building ListVirtualTerminalsRequest
type ListVirtualTerminalsRequestBuilder struct {
	req *ListVirtualTerminalsRequest
}

// NewListVirtualTerminalsRequest creates a new builder for ListVirtualTerminalsRequest
func NewListVirtualTerminalsRequest() *ListVirtualTerminalsRequestBuilder {
	return &ListVirtualTerminalsRequestBuilder{
		req: &ListVirtualTerminalsRequest{},
	}
}

// Build returns the constructed ListVirtualTerminalsRequest
func (b *ListVirtualTerminalsRequestBuilder) Build() *ListVirtualTerminalsRequest {
	return b.req
}

// UpdateVirtualTerminalRequest represents the request to update a virtual terminal
type UpdateVirtualTerminalRequest struct {
	Name string `json:"name"`
}

// UpdateVirtualTerminalRequestBuilder provides a fluent interface for building UpdateVirtualTerminalRequest
type UpdateVirtualTerminalRequestBuilder struct {
	name string
}

// NewUpdateVirtualTerminalRequest creates a new builder for updating a virtual terminal
func NewUpdateVirtualTerminalRequest(name string) *UpdateVirtualTerminalRequestBuilder {
	return &UpdateVirtualTerminalRequestBuilder{
		name: name,
	}
}

// Name sets the name of the virtual terminal
func (b *UpdateVirtualTerminalRequestBuilder) Name(name string) *UpdateVirtualTerminalRequestBuilder {
	b.name = name
	return b
}

// Build creates the UpdateVirtualTerminalRequest
func (b *UpdateVirtualTerminalRequestBuilder) Build() *UpdateVirtualTerminalRequest {
	return &UpdateVirtualTerminalRequest{
		Name: b.name,
	}
}

// AssignDestinationRequest represents the request to assign destinations to a virtual terminal
type AssignDestinationRequest struct {
	Destinations []VirtualTerminalDestination `json:"destinations"`
}

// AssignDestinationRequestBuilder provides a fluent interface for building AssignDestinationRequest
type AssignDestinationRequestBuilder struct {
	destinations []VirtualTerminalDestination
}

// NewAssignDestinationRequest creates a new builder for assigning destinations
func NewAssignDestinationRequest() *AssignDestinationRequestBuilder {
	return &AssignDestinationRequestBuilder{
		destinations: make([]VirtualTerminalDestination, 0),
	}
}

// AddDestination adds a destination to the assignment request
func (b *AssignDestinationRequestBuilder) AddDestination(destination VirtualTerminalDestination) *AssignDestinationRequestBuilder {
	b.destinations = append(b.destinations, destination)
	return b
}

// Destinations sets all destinations at once
func (b *AssignDestinationRequestBuilder) Destinations(destinations []VirtualTerminalDestination) *AssignDestinationRequestBuilder {
	b.destinations = destinations
	return b
}

// Build creates the AssignDestinationRequest
func (b *AssignDestinationRequestBuilder) Build() *AssignDestinationRequest {
	return &AssignDestinationRequest{
		Destinations: b.destinations,
	}
}

// UnassignDestinationRequest represents the request to unassign destinations from a virtual terminal
type UnassignDestinationRequest struct {
	Targets []string `json:"targets"`
}

// UnassignDestinationRequestBuilder provides a fluent interface for building UnassignDestinationRequest
type UnassignDestinationRequestBuilder struct {
	targets []string
}

// NewUnassignDestinationRequest creates a new builder for unassigning destinations
func NewUnassignDestinationRequest() *UnassignDestinationRequestBuilder {
	return &UnassignDestinationRequestBuilder{
		targets: make([]string, 0),
	}
}

// AddTarget adds a target to the unassignment request
func (b *UnassignDestinationRequestBuilder) AddTarget(target string) *UnassignDestinationRequestBuilder {
	b.targets = append(b.targets, target)
	return b
}

// Targets sets all targets at once
func (b *UnassignDestinationRequestBuilder) Targets(targets []string) *UnassignDestinationRequestBuilder {
	b.targets = targets
	return b
}

// Build creates the UnassignDestinationRequest
func (b *UnassignDestinationRequestBuilder) Build() *UnassignDestinationRequest {
	return &UnassignDestinationRequest{
		Targets: b.targets,
	}
}

// AddSplitCodeRequest represents the request to add a split code to a virtual terminal
type AddSplitCodeRequest struct {
	SplitCode string `json:"split_code"`
}

// AddSplitCodeRequestBuilder provides a fluent interface for building AddSplitCodeRequest
type AddSplitCodeRequestBuilder struct {
	splitCode string
}

// NewAddSplitCodeRequest creates a new builder for adding a split code
func NewAddSplitCodeRequest(splitCode string) *AddSplitCodeRequestBuilder {
	return &AddSplitCodeRequestBuilder{
		splitCode: splitCode,
	}
}

// SplitCode sets the split code
func (b *AddSplitCodeRequestBuilder) SplitCode(splitCode string) *AddSplitCodeRequestBuilder {
	b.splitCode = splitCode
	return b
}

// Build creates the AddSplitCodeRequest
func (b *AddSplitCodeRequestBuilder) Build() *AddSplitCodeRequest {
	return &AddSplitCodeRequest{
		SplitCode: b.splitCode,
	}
}

// RemoveSplitCodeRequest represents the request to remove a split code from a virtual terminal
type RemoveSplitCodeRequest struct {
	SplitCode string `json:"split_code"`
}

// RemoveSplitCodeRequestBuilder provides a fluent interface for building RemoveSplitCodeRequest
type RemoveSplitCodeRequestBuilder struct {
	splitCode string
}

// NewRemoveSplitCodeRequest creates a new builder for removing a split code
func NewRemoveSplitCodeRequest(splitCode string) *RemoveSplitCodeRequestBuilder {
	return &RemoveSplitCodeRequestBuilder{
		splitCode: splitCode,
	}
}

// SplitCode sets the split code to remove
func (b *RemoveSplitCodeRequestBuilder) SplitCode(splitCode string) *RemoveSplitCodeRequestBuilder {
	b.splitCode = splitCode
	return b
}

// Build creates the RemoveSplitCodeRequest
func (b *RemoveSplitCodeRequestBuilder) Build() *RemoveSplitCodeRequest {
	return &RemoveSplitCodeRequest{
		SplitCode: b.splitCode,
	}
}

// FetchVirtualTerminalResponse represents the response from fetching a virtual terminal
type FetchVirtualTerminalResponse = types.Response[VirtualTerminal]

// UpdateVirtualTerminalResponse represents the response from updating a virtual terminal
type UpdateVirtualTerminalResponse = types.Response[VirtualTerminal]

// DeactivateVirtualTerminalResponse represents the response from deactivating a virtual terminal
type DeactivateVirtualTerminalResponse = types.Response[any]

// AssignDestinationResponse represents the response from assigning destinations
type AssignDestinationResponse = types.Response[[]VirtualTerminalDestination]

// UnassignDestinationResponse represents the response from unassigning destinations
type UnassignDestinationResponse = types.Response[any]

// AddSplitCodeResponse represents the response from adding a split code
type AddSplitCodeResponse = types.Response[any]

// RemoveSplitCodeResponse represents the response from removing a split code
type RemoveSplitCodeResponse = types.Response[any]

// Response type aliases for backwards compatibility
type CreateVirtualTerminalResponse = types.Response[VirtualTerminal]
type ListVirtualTerminalsResponse = types.Response[[]VirtualTerminal]
