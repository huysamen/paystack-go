package models

import "github.com/huysamen/paystack-go/enums"

// Subaccount represents a Paystack subaccount
type Subaccount struct {
	ID                   uint64         `json:"id"`
	Integration          int            `json:"integration"`
	Bank                 *int           `json:"bank"`
	ManagedByIntegration int            `json:"managed_by_integration"`
	Domain               string         `json:"domain"`
	SubaccountCode       string         `json:"subaccount_code"`
	BusinessName         string         `json:"business_name"`
	Description          *string        `json:"description"`
	PrimaryContactName   *string        `json:"primary_contact_name"`
	PrimaryContactEmail  *string        `json:"primary_contact_email"`
	PrimaryContactPhone  *string        `json:"primary_contact_phone"`
	Metadata             *Metadata      `json:"metadata"`
	PercentageCharge     float64        `json:"percentage_charge"`
	IsVerified           bool           `json:"is_verified"`
	SettlementBank       string         `json:"settlement_bank"`
	AccountNumber        string         `json:"account_number"`
	SettlementSchedule   string         `json:"settlement_schedule"`
	Active               bool           `json:"active"`
	Migrate              bool           `json:"migrate"`
	Currency             enums.Currency `json:"currency"`
	AccountName          string         `json:"account_name"`
	Product              string         `json:"product"`
	CreatedAt            DateTime       `json:"createdAt"`
	UpdatedAt            DateTime       `json:"updatedAt"`
}

// TransactionSplit represents a Paystack transaction split
type TransactionSplit struct {
	ID               uint64            `json:"id"`
	Name             string            `json:"name"`
	Type             string            `json:"type"` // percentage, flat
	Currency         enums.Currency    `json:"currency"`
	Integration      int               `json:"integration"`
	Domain           string            `json:"domain"`
	SplitCode        string            `json:"split_code"`
	Active           bool              `json:"active"`
	BearerType       string            `json:"bearer_type"` // all, account, subaccount
	BearerSubaccount *string           `json:"bearer_subaccount"`
	CreatedAt        DateTime          `json:"createdAt"`
	UpdatedAt        DateTime          `json:"updatedAt"`
	IsDynamic        bool              `json:"is_dynamic"`
	Subaccounts      []SplitSubaccount `json:"subaccounts"`
	TotalSubaccounts int               `json:"total_subaccounts"`
}

// SplitSubaccount represents a subaccount within a split
type SplitSubaccount struct {
	Subaccount Subaccount `json:"subaccount"`
	Share      float64    `json:"share"`
}
