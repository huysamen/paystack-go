package types

import "github.com/huysamen/paystack-go/types/data"

// Country represents a country supported by Paystack
type Country struct {
	ID                           data.Int             `json:"id"`
	ActiveForDashboardOnboarding data.Bool            `json:"active_for_dashboard_onboarding"`
	Name                         data.String          `json:"name"`
	ISOCode                      data.String          `json:"iso_code"`
	DefaultCurrencyCode          data.String          `json:"default_currency_code"`
	IntegrationDefaults          Metadata             `json:"integration_defaults"`
	CallingCode                  data.String          `json:"calling_code"`
	PilotMode                    data.Bool            `json:"pilot_mode"`
	Relationships                CountryRelationships `json:"relationships"`
}

// CountryRelationships represents the relationships for a country
type CountryRelationships struct {
	Currency           CurrencyRelationship `json:"currency"`
	IntegrationFeature CountryRelationship  `json:"integration_feature"`
	IntegrationType    CountryRelationship  `json:"integration_type"`
	PaymentMethod      CountryRelationship  `json:"payment_method"`
}

// CountryRelationship represents basic relationship data
type CountryRelationship struct {
	Type data.String   `json:"type"`
	Data []data.String `json:"data"`
}

// CurrencyRelationship represents currency relationship data with supported currencies
type CurrencyRelationship struct {
	Type                data.String                      `json:"type"`
	Data                []data.String                    `json:"data"`
	SupportedCurrencies map[string]CurrencyConfiguration `json:"supported_currencies"`
}

// CurrencyConfiguration represents the configuration for a specific currency
type CurrencyConfiguration struct {
	Bank                *BankConfiguration                `json:"bank,omitempty"`
	MobileMoney         *MobileMoneyConfiguration         `json:"mobile_money,omitempty"`
	MobileMoneyBusiness *MobileMoneyBusinessConfiguration `json:"mobile_money_business,omitempty"`
	EFT                 *EFTConfiguration                 `json:"eft,omitempty"`
}

// BankConfiguration represents bank account configuration
type BankConfiguration struct {
	BankType                    data.String          `json:"bank_type"`
	RequiredFields              []data.String        `json:"required_fields,omitempty"`
	BranchCode                  data.Bool            `json:"branch_code"`
	BranchCodeType              data.String          `json:"branch_code_type"`
	AccountName                 data.Bool            `json:"account_name"`
	AccountVerificationRequired data.Bool            `json:"account_verification_required"`
	AccountNumberLabel          data.String          `json:"account_number_label"`
	AccountNumberPattern        AccountNumberPattern `json:"account_number_pattern"`
	Documents                   []data.String        `json:"documents"`
	Notices                     []data.String        `json:"notices,omitempty"`
	ShowAccountNumberTooltip    data.Bool            `json:"show_account_number_tooltip"`
}

// MobileMoneyConfiguration represents mobile money configuration
type MobileMoneyConfiguration struct {
	BankType             data.String          `json:"bank_type"`
	PhoneNumberLabel     data.String          `json:"phone_number_label"`
	Placeholder          data.String          `json:"placeholder"`
	AccountNumberPattern AccountNumberPattern `json:"account_number_pattern"`
}

// MobileMoneyBusinessConfiguration represents mobile money business configuration
type MobileMoneyBusinessConfiguration struct {
	BankType                    data.String          `json:"bank_type"`
	AccountVerificationRequired data.Bool            `json:"account_verification_required"`
	PhoneNumberLabel            data.String          `json:"phone_number_label"`
	AccountNumberPattern        AccountNumberPattern `json:"account_number_pattern"`
}

// EFTConfiguration represents EFT configuration
type EFTConfiguration struct {
	AccountNumberPattern AccountNumberPattern `json:"account_number_pattern"`
	Placeholder          data.String          `json:"placeholder"`
}

// AccountNumberPattern represents account number validation pattern
type AccountNumberPattern struct {
	ExactMatch data.Bool   `json:"exact_match"`
	Pattern    data.String `json:"pattern"`
}

// State represents a state for address verification
type State struct {
	Name         data.String `json:"name"`
	Slug         data.String `json:"slug"`
	Abbreviation data.String `json:"abbreviation"`
}
