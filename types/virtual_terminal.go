package types

import "github.com/huysamen/paystack-go/types/data"

// VirtualTerminalDestination represents a notification destination for a virtual terminal
type VirtualTerminalDestination struct {
	ID        data.Int    `json:"id,omitempty"`
	Target    data.String `json:"target"`
	Name      data.String `json:"name"`
	Type      data.String `json:"type,omitempty"`
	CreatedAt data.Time   `json:"created_at,omitempty"`
	UpdatedAt data.Time   `json:"updated_at,omitempty"`
}

// VirtualTerminal represents a virtual terminal
type VirtualTerminal struct {
	ID             data.Int                     `json:"id"`
	Code           data.String                  `json:"code"`
	Name           data.String                  `json:"name"`
	Integration    data.Int                     `json:"integration"`
	Domain         data.String                  `json:"domain"`
	PaymentMethods []data.String                `json:"paymentMethods"`
	Active         data.Bool                    `json:"active"`
	CreatedAt      data.NullTime                `json:"created_at,omitempty"` // Change to NullTime
	Metadata       Metadata                     `json:"metadata,omitempty"`
	Destinations   []VirtualTerminalDestination `json:"destinations,omitempty"`
	Currency       data.String                  `json:"currency"`
	CustomFields   []VirtualTerminalCustomField `json:"custom_fields,omitempty"`
	ConnectAccount data.NullInt                 `json:"connect_account_id,omitempty"`
}

// VirtualTerminalCustomField represents a custom field for the virtual terminal form
type VirtualTerminalCustomField struct {
	DisplayName  data.String `json:"display_name"`
	VariableName data.String `json:"variable_name"`
}
