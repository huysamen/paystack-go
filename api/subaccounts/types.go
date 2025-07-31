package subaccounts

import (
	"time"

	"github.com/huysamen/paystack-go/types"
)

// SettlementSchedule represents the settlement schedule for a subaccount
type SettlementSchedule string

const (
	SettlementScheduleAuto    SettlementSchedule = "auto"    // T+1 settlement
	SettlementScheduleWeekly  SettlementSchedule = "weekly"  // Weekly settlement
	SettlementScheduleMonthly SettlementSchedule = "monthly" // Monthly settlement
	SettlementScheduleManual  SettlementSchedule = "manual"  // Manual settlement only
)

// String returns the string representation of SettlementSchedule
func (s SettlementSchedule) String() string {
	return string(s)
}

// Subaccount represents a subaccount
type Subaccount struct {
	ID                   uint64         `json:"id"`
	SubaccountCode       string         `json:"subaccount_code"`
	BusinessName         string         `json:"business_name"`
	Description          *string        `json:"description"`
	PrimaryContactName   *string        `json:"primary_contact_name"`
	PrimaryContactEmail  *string        `json:"primary_contact_email"`
	PrimaryContactPhone  *string        `json:"primary_contact_phone"`
	Metadata             map[string]any `json:"metadata"`
	PercentageCharge     float64        `json:"percentage_charge"`
	IsVerified           bool           `json:"is_verified"`
	SettlementBank       string         `json:"settlement_bank"`
	AccountNumber        string         `json:"account_number"`
	AccountName          *string        `json:"account_name"`
	SettlementSchedule   *string        `json:"settlement_schedule"`
	Active               bool           `json:"active"`
	Currency             string         `json:"currency"`
	Domain               string         `json:"domain"`
	Integration          uint64         `json:"integration"`
	BankID               uint64         `json:"bank_id"`
	Bank                 uint64         `json:"bank"`
	ManagedByIntegration uint64         `json:"managed_by_integration"`
	Product              string         `json:"product"`
	CreatedAt            time.Time      `json:"createdAt"`
	UpdatedAt            time.Time      `json:"updatedAt"`
}

// Subaccount Create

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
type SubaccountCreateResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    Subaccount `json:"data"`
}

// Subaccount List

// SubaccountListRequest represents the request to list subaccounts
type SubaccountListRequest struct {
	PerPage *int       `json:"perPage,omitempty"` // Optional: records per page (default: 50)
	Page    *int       `json:"page,omitempty"`    // Optional: page number (default: 1)
	From    *time.Time `json:"from,omitempty"`    // Optional: start date filter
	To      *time.Time `json:"to,omitempty"`      // Optional: end date filter
}

// SubaccountListResponse represents the response from listing subaccounts
type SubaccountListResponse struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Data    []Subaccount `json:"data"`
	Meta    types.Meta   `json:"meta"`
}

// Subaccount Fetch

// SubaccountFetchResponse represents the response from fetching a subaccount
type SubaccountFetchResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    Subaccount `json:"data"`
}

// Subaccount Update

// SubaccountUpdateRequest represents the request to update a subaccount
type SubaccountUpdateRequest struct {
	BusinessName        string              `json:"business_name"`                   // Required: Name of business
	Description         string              `json:"description"`                     // Required: Description
	BankCode            *string             `json:"settlement_bank,omitempty"`       // Optional: Bank Code
	AccountNumber       *string             `json:"account_number,omitempty"`        // Optional: Bank Account Number
	Active              *bool               `json:"active,omitempty"`                // Optional: Activate/deactivate
	PercentageCharge    *float64            `json:"percentage_charge,omitempty"`     // Optional: Percentage charge
	PrimaryContactEmail *string             `json:"primary_contact_email,omitempty"` // Optional: Contact email
	PrimaryContactName  *string             `json:"primary_contact_name,omitempty"`  // Optional: Contact name
	PrimaryContactPhone *string             `json:"primary_contact_phone,omitempty"` // Optional: Contact phone
	SettlementSchedule  *SettlementSchedule `json:"settlement_schedule,omitempty"`   // Optional: Settlement schedule
	Metadata            map[string]any      `json:"metadata,omitempty"`              // Optional: Additional data
}

// SubaccountUpdateResponse represents the response from updating a subaccount
type SubaccountUpdateResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    Subaccount `json:"data"`
}
