package types

import "time"

// Subaccount represents a Paystack subaccount
type Subaccount struct {
	Integration          int       `json:"integration"`
	Bank                 int       `json:"bank"`
	ManagedByIntegration int       `json:"managed_by_integration"`
	Domain               string    `json:"domain"`
	SubaccountCode       string    `json:"subaccount_code"`
	BusinessName         string    `json:"business_name"`
	Description          string    `json:"description"`
	PrimaryContactName   string    `json:"primary_contact_name"`
	PrimaryContactEmail  string    `json:"primary_contact_email"`
	PrimaryContactPhone  string    `json:"primary_contact_phone"`
	Metadata             Metadata  `json:"metadata"`
	PercentageCharge     int       `json:"percentage_charge"`
	IsVerified           bool      `json:"is_verified"`
	SettlementBank       string    `json:"settlement_bank"`
	AccountNumber        string    `json:"account_number"`
	SettlementSchedule   string    `json:"settlement_schedule"`
	Active               bool      `json:"active"`
	Migrate              bool      `json:"migrate"`
	Currency             Currency  `json:"currency"`
	AccountName          string    `json:"account_name"`
	Product              string    `json:"product"`
	ID                   uint64    `json:"id"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}

// Split represents a Paystack transaction split
type Split struct {
	ID               uint64    `json:"id"`
	Name             string    `json:"name"`
	Type             string    `json:"type"`
	Currency         Currency  `json:"currency"`
	Integration      uint64    `json:"integration"`
	Domain           string    `json:"domain"`
	SplitCode        string    `json:"split_code"`
	Active           bool      `json:"active"`
	BearerType       string    `json:"bearer_type"`
	BearerSubaccount *string   `json:"bearer_subaccount"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	IsDynamic        bool      `json:"is_dynamic"`
	Subaccounts      []struct {
		Subaccount string `json:"subaccount"`
		Share      int    `json:"share"`
	} `json:"subaccounts"`
	TotalSubaccounts int `json:"total_subaccounts"`
}
