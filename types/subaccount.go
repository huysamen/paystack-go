package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Subaccount represents a Paystack subaccount
type Subaccount struct {
	ID                   data.Uint       `json:"id"`
	Integration          data.Int        `json:"integration"`
	Bank                 data.NullInt    `json:"bank"`
	ManagedByIntegration data.Int        `json:"managed_by_integration"`
	Domain               data.String     `json:"domain"`
	SubaccountCode       data.String     `json:"subaccount_code"`
	BusinessName         data.String     `json:"business_name"`
	Description          data.NullString `json:"description"`
	PrimaryContactName   data.NullString `json:"primary_contact_name"`
	PrimaryContactEmail  data.NullString `json:"primary_contact_email"`
	PrimaryContactPhone  data.NullString `json:"primary_contact_phone"`
	Metadata             Metadata        `json:"metadata"`
	PercentageCharge     data.Float      `json:"percentage_charge"`
	IsVerified           data.Bool       `json:"is_verified"`
	SettlementBank       data.String     `json:"settlement_bank"`
	AccountNumber        data.String     `json:"account_number"`
	SettlementSchedule   data.String     `json:"settlement_schedule"`
	Active               data.Bool       `json:"active"`
	Migrate              data.Bool       `json:"migrate"`
	Currency             enums.Currency  `json:"currency"`
	AccountName          data.String     `json:"account_name"`
	Product              data.String     `json:"product"`
	CreatedAt            data.Time       `json:"createdAt"`
	UpdatedAt            data.Time       `json:"updatedAt"`
}

// TransactionSplit represents a Paystack transaction split
type TransactionSplit struct {
	ID               data.Uint         `json:"id"`
	Name             data.String       `json:"name"`
	Type             data.String       `json:"type"` // percentage, flat
	Currency         enums.Currency    `json:"currency"`
	Integration      data.Int          `json:"integration"`
	Domain           data.String       `json:"domain"`
	SplitCode        data.String       `json:"split_code"`
	Active           data.Bool         `json:"active"`
	BearerType       data.String       `json:"bearer_type"` // all, account, subaccount
	BearerSubaccount data.NullString   `json:"bearer_subaccount"`
	CreatedAt        data.Time         `json:"createdAt"`
	UpdatedAt        data.Time         `json:"updatedAt"`
	IsDynamic        data.Bool         `json:"is_dynamic"`
	Subaccounts      []SplitSubaccount `json:"subaccounts"`
	TotalSubaccounts data.Int          `json:"total_subaccounts"`
}

// SplitSubaccount represents a subaccount within a split
type SplitSubaccount struct {
	Subaccount Subaccount `json:"subaccount"`
	Share      data.Float `json:"share"`
}

// TransactionSplitSubaccount represents a subaccount for creating transaction splits
type TransactionSplitSubaccount struct {
	Subaccount data.String `json:"subaccount"`
	Share      data.Int    `json:"share"`
}
