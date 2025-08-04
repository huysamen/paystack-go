package types

import (
	"time"
)

// SettlementStatus represents the status of a settlement
type SettlementStatus string

const (
	SettlementStatusSuccess    SettlementStatus = "success"    // Successfully settled
	SettlementStatusProcessing SettlementStatus = "processing" // Currently being processed
	SettlementStatusPending    SettlementStatus = "pending"    // Pending settlement
	SettlementStatusFailed     SettlementStatus = "failed"     // Failed settlement
)

// String returns the string representation of SettlementStatus
func (s SettlementStatus) String() string {
	return string(s)
}

// Settlement represents a settlement record
type Settlement struct {
	ID              uint64           `json:"id"`
	Domain          string           `json:"domain"`
	Status          SettlementStatus `json:"status"`
	Currency        string           `json:"currency"`
	Integration     uint64           `json:"integration"`
	TotalAmount     int64            `json:"total_amount"`     // Amount after fees in kobo
	EffectiveAmount int64            `json:"effective_amount"` // Amount actually settled in kobo
	TotalFees       int64            `json:"total_fees"`       // Total fees charged in kobo
	TotalProcessed  int64            `json:"total_processed"`  // Total amount processed in kobo
	Deductions      map[string]any   `json:"deductions"`       // Any deductions applied
	SettlementDate  time.Time        `json:"settlement_date"`  // Date settlement was made
	SettledBy       *string          `json:"settled_by"`       // Who processed the settlement
	CreatedAt       time.Time        `json:"createdAt"`        // When settlement record was created
	UpdatedAt       time.Time        `json:"updatedAt"`        // When settlement was last updated
}

// SettlementTransaction represents a transaction within a settlement
type SettlementTransaction struct {
	ID              uint64         `json:"id"`
	Domain          string         `json:"domain"`
	Status          string         `json:"status"`
	Reference       string         `json:"reference"`
	Amount          int64          `json:"amount"` // Transaction amount in kobo
	Message         string         `json:"message"`
	GatewayResponse string         `json:"gateway_response"`
	PaidAt          *time.Time     `json:"paid_at"`
	CreatedAt       time.Time      `json:"created_at"`
	Channel         string         `json:"channel"`
	Currency        string         `json:"currency"`
	IPAddress       string         `json:"ip_address"`
	Metadata        map[string]any `json:"metadata"`
	Log             map[string]any `json:"log"`
	Fees            int64          `json:"fees"`          // Fees charged for this transaction
	FeesSplit       map[string]any `json:"fees_split"`    // Breakdown of fees
	Customer        map[string]any `json:"customer"`      // Customer information
	Authorization   map[string]any `json:"authorization"` // Authorization details
	Plan            map[string]any `json:"plan"`          // Plan details if subscription
	Subaccount      map[string]any `json:"subaccount"`    // Subaccount details if applicable
}
