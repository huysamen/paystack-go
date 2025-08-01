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

// SubaccountCreateRequestBuilder provides a fluent interface for building SubaccountCreateRequest
type SubaccountCreateRequestBuilder struct {
	req *SubaccountCreateRequest
}

// NewSubaccountCreateRequest creates a new builder for SubaccountCreateRequest
func NewSubaccountCreateRequest(businessName, bankCode, accountNumber string, percentageCharge float64) *SubaccountCreateRequestBuilder {
	return &SubaccountCreateRequestBuilder{
		req: &SubaccountCreateRequest{
			BusinessName:     businessName,
			BankCode:         bankCode,
			AccountNumber:    accountNumber,
			PercentageCharge: percentageCharge,
		},
	}
}

// Description sets the subaccount description
func (b *SubaccountCreateRequestBuilder) Description(description string) *SubaccountCreateRequestBuilder {
	b.req.Description = &description
	return b
}

// PrimaryContactEmail sets the primary contact email
func (b *SubaccountCreateRequestBuilder) PrimaryContactEmail(email string) *SubaccountCreateRequestBuilder {
	b.req.PrimaryContactEmail = &email
	return b
}

// PrimaryContactName sets the primary contact name
func (b *SubaccountCreateRequestBuilder) PrimaryContactName(name string) *SubaccountCreateRequestBuilder {
	b.req.PrimaryContactName = &name
	return b
}

// PrimaryContactPhone sets the primary contact phone
func (b *SubaccountCreateRequestBuilder) PrimaryContactPhone(phone string) *SubaccountCreateRequestBuilder {
	b.req.PrimaryContactPhone = &phone
	return b
}

// Metadata sets the subaccount metadata
func (b *SubaccountCreateRequestBuilder) Metadata(metadata map[string]any) *SubaccountCreateRequestBuilder {
	b.req.Metadata = metadata
	return b
}

// Build returns the constructed SubaccountCreateRequest
func (b *SubaccountCreateRequestBuilder) Build() *SubaccountCreateRequest {
	return b.req
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

// SubaccountListRequestBuilder provides a fluent interface for building SubaccountListRequest
type SubaccountListRequestBuilder struct {
	req *SubaccountListRequest
}

// NewSubaccountListRequest creates a new builder for SubaccountListRequest
func NewSubaccountListRequest() *SubaccountListRequestBuilder {
	return &SubaccountListRequestBuilder{
		req: &SubaccountListRequest{},
	}
}

// PerPage sets the number of subaccounts per page
func (b *SubaccountListRequestBuilder) PerPage(perPage int) *SubaccountListRequestBuilder {
	b.req.PerPage = &perPage
	return b
}

// Page sets the page number
func (b *SubaccountListRequestBuilder) Page(page int) *SubaccountListRequestBuilder {
	b.req.Page = &page
	return b
}

// From sets the start date filter
func (b *SubaccountListRequestBuilder) From(from time.Time) *SubaccountListRequestBuilder {
	b.req.From = &from
	return b
}

// To sets the end date filter
func (b *SubaccountListRequestBuilder) To(to time.Time) *SubaccountListRequestBuilder {
	b.req.To = &to
	return b
}

// DateRange sets both from and to dates for convenience
func (b *SubaccountListRequestBuilder) DateRange(from, to time.Time) *SubaccountListRequestBuilder {
	b.req.From = &from
	b.req.To = &to
	return b
}

// Build returns the constructed SubaccountListRequest
func (b *SubaccountListRequestBuilder) Build() *SubaccountListRequest {
	return b.req
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

// SubaccountUpdateRequestBuilder provides a fluent interface for building SubaccountUpdateRequest
type SubaccountUpdateRequestBuilder struct {
	req *SubaccountUpdateRequest
}

// NewSubaccountUpdateRequest creates a new builder for SubaccountUpdateRequest
func NewSubaccountUpdateRequest(businessName, description string) *SubaccountUpdateRequestBuilder {
	return &SubaccountUpdateRequestBuilder{
		req: &SubaccountUpdateRequest{
			BusinessName: businessName,
			Description:  description,
		},
	}
}

// BankCode sets the settlement bank code
func (b *SubaccountUpdateRequestBuilder) BankCode(bankCode string) *SubaccountUpdateRequestBuilder {
	b.req.BankCode = &bankCode
	return b
}

// AccountNumber sets the bank account number
func (b *SubaccountUpdateRequestBuilder) AccountNumber(accountNumber string) *SubaccountUpdateRequestBuilder {
	b.req.AccountNumber = &accountNumber
	return b
}

// Active sets whether the subaccount is active
func (b *SubaccountUpdateRequestBuilder) Active(active bool) *SubaccountUpdateRequestBuilder {
	b.req.Active = &active
	return b
}

// PercentageCharge sets the percentage charge
func (b *SubaccountUpdateRequestBuilder) PercentageCharge(charge float64) *SubaccountUpdateRequestBuilder {
	b.req.PercentageCharge = &charge
	return b
}

// PrimaryContactEmail sets the primary contact email
func (b *SubaccountUpdateRequestBuilder) PrimaryContactEmail(email string) *SubaccountUpdateRequestBuilder {
	b.req.PrimaryContactEmail = &email
	return b
}

// PrimaryContactName sets the primary contact name
func (b *SubaccountUpdateRequestBuilder) PrimaryContactName(name string) *SubaccountUpdateRequestBuilder {
	b.req.PrimaryContactName = &name
	return b
}

// PrimaryContactPhone sets the primary contact phone
func (b *SubaccountUpdateRequestBuilder) PrimaryContactPhone(phone string) *SubaccountUpdateRequestBuilder {
	b.req.PrimaryContactPhone = &phone
	return b
}

// SettlementSchedule sets the settlement schedule
func (b *SubaccountUpdateRequestBuilder) SettlementSchedule(schedule *SettlementSchedule) *SubaccountUpdateRequestBuilder {
	b.req.SettlementSchedule = schedule
	return b
}

// Metadata sets the subaccount metadata
func (b *SubaccountUpdateRequestBuilder) Metadata(metadata map[string]any) *SubaccountUpdateRequestBuilder {
	b.req.Metadata = metadata
	return b
}

// Build returns the constructed SubaccountUpdateRequest
func (b *SubaccountUpdateRequestBuilder) Build() *SubaccountUpdateRequest {
	return b.req
}

// SubaccountUpdateResponse represents the response from updating a subaccount
type SubaccountUpdateResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    Subaccount `json:"data"`
}
