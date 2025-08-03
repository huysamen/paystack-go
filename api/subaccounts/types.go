package subaccounts

import (
	"time"

	"github.com/huysamen/paystack-go/types"
)

// SubaccountCreateRequest represents the request to create a subaccount
type SubaccountCreateRequest struct {
	BusinessName        string         `json:"business_name"`                   // Required: Name of business
	BankCode            string         `json:"settlement_bank"`                 // Required: Bank Code (use settlement_bank as per API docs)
	AccountNumber       string         `json:"account_number"`                  // Required: Bank Account Number
	PercentageCharge    float64        `json:"percentage_charge"`               // Required: Percentage the main account receives
	Description         *string        `json:"description,omitempty"`           // Optional: Description
	PrimaryContactEmail *string        `json:"primary_contact_email,omitempty"` // Optional: Contact email
	PrimaryContactName  *string        `json:"primary_contact_name,omitempty"`  // Optional: Contact name
	PrimaryContactPhone *string        `json:"primary_contact_phone,omitempty"` // Optional: Contact phone
	Metadata            map[string]any `json:"metadata,omitempty"`              // Optional: Additional data
}

// SubaccountCreateResponse represents the response from creating a subaccount
type SubaccountCreateResponse = types.Response[types.Subaccount]

// Subaccount List

// SubaccountListRequest represents the request to list subaccounts
type SubaccountListRequest struct {
	PerPage *int       `json:"perPage,omitempty"` // Optional: records per page (default: 50)
	Page    *int       `json:"page,omitempty"`    // Optional: page number (default: 1)
	From    *time.Time `json:"from,omitempty"`    // Optional: start date filter
	To      *time.Time `json:"to,omitempty"`      // Optional: end date filter
}

// SubaccountListResponse represents the response from listing subaccounts
type SubaccountListResponse = types.Response[[]types.Subaccount]

// Subaccount Fetch

// SubaccountFetchResponse represents the response from fetching a subaccount
type SubaccountFetchResponse = types.Response[types.Subaccount]

// Subaccount Update

// SubaccountUpdateRequest represents the request to update a subaccount
type SubaccountUpdateRequest struct {
	BusinessName        *string        `json:"business_name,omitempty"`         // Optional: Business name
	BankCode            *string        `json:"settlement_bank,omitempty"`       // Optional: Bank code
	AccountNumber       *string        `json:"account_number,omitempty"`        // Optional: Account number
	PercentageCharge    *float64       `json:"percentage_charge,omitempty"`     // Optional: Percentage charge
	Description         *string        `json:"description,omitempty"`           // Optional: Description
	PrimaryContactEmail *string        `json:"primary_contact_email,omitempty"` // Optional: Primary contact email
	PrimaryContactName  *string        `json:"primary_contact_name,omitempty"`  // Optional: Primary contact name
	PrimaryContactPhone *string        `json:"primary_contact_phone,omitempty"` // Optional: Primary contact phone
	Active              *bool          `json:"active,omitempty"`                // Optional: Active status
	Metadata            map[string]any `json:"metadata,omitempty"`              // Optional: Metadata
}

// SubaccountUpdateResponse represents the response from updating a subaccount
type SubaccountUpdateResponse = types.Response[types.Subaccount]
