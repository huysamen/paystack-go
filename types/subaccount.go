package types

// Subaccount represents a Paystack subaccount
type Subaccount struct {
	Integration          int      `json:"integration"`
	Bank                 int      `json:"bank"`
	ManagedByIntegration int      `json:"managed_by_integration"`
	Domain               string   `json:"domain"`
	SubaccountCode       string   `json:"subaccount_code"`
	BusinessName         string   `json:"business_name"`
	Description          string   `json:"description"`
	PrimaryContactName   string   `json:"primary_contact_name"`
	PrimaryContactEmail  string   `json:"primary_contact_email"`
	PrimaryContactPhone  string   `json:"primary_contact_phone"`
	Metadata             Metadata `json:"metadata"`
	PercentageCharge     int      `json:"percentage_charge"`
	IsVerified           bool     `json:"is_verified"`
	SettlementBank       string   `json:"settlement_bank"`
	AccountNumber        string   `json:"account_number"`
	SettlementSchedule   string   `json:"settlement_schedule"`
	Active               bool     `json:"active"`
	Migrate              bool     `json:"migrate"`
	Currency             Currency `json:"currency"`
	AccountName          string   `json:"account_name"`
	Product              string   `json:"product"`
	ID                   uint64   `json:"id"`
	CreatedAt            DateTime `json:"createdAt"`
	UpdatedAt            DateTime `json:"updatedAt"`
}

// Split is an alias for TransactionSplit for backward compatibility
// Deprecated: Use TransactionSplit instead
type Split = TransactionSplit
