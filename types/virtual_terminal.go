package types
import "github.com/huysamen/paystack-go/types/data"
// VirtualTerminalDestination represents a notification destination for a virtual terminal
type VirtualTerminalDestination struct {
	ID        int      `json:"id,omitempty"`
	Target    string   `json:"target"`
	Name      string   `json:"name"`
	Type      string   `json:"type,omitempty"`
	CreatedAt data.MultiDateTime `json:"created_at,omitempty"`
	UpdatedAt data.MultiDateTime `json:"updated_at,omitempty"`
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
	CreatedAt      data.MultiDateTime                     `json:"created_at,omitempty"`
	Metadata       *Metadata                    `json:"metadata,omitempty"`
	Destinations   []VirtualTerminalDestination `json:"destinations,omitempty"`
	Currency       string                       `json:"currency"`
	CustomFields   []VirtualTerminalCustomField `json:"custom_fields,omitempty"`
	ConnectAccount *int                         `json:"connect_account_id,omitempty"`
}

// VirtualTerminalCustomField represents a custom field for the virtual terminal form
type VirtualTerminalCustomField struct {
	DisplayName  string `json:"display_name"`
	VariableName string `json:"variable_name"`
}
