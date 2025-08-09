package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Settlement represents a settlement record
type Settlement struct {
	ID              uint64                 `json:"id"`
	Domain          string                 `json:"domain"`
	Status          enums.SettlementStatus `json:"status"`
	Currency        enums.Currency         `json:"currency"`
	Integration     uint64                 `json:"integration"`
	TotalAmount     int64                  `json:"total_amount"`     // Amount after fees in kobo
	EffectiveAmount int64                  `json:"effective_amount"` // Amount actually settled in kobo
	TotalFees       int64                  `json:"total_fees"`       // Total fees charged in kobo
	TotalProcessed  int64                  `json:"total_processed"`  // Total amount processed in kobo
	Deductions      *Metadata              `json:"deductions"`       // Any deductions applied
	SettlementDate  data.MultiDateTime     `json:"settlement_date"`  // Date settlement was made
	SettledBy       data.NullString        `json:"settled_by"`       // Who processed the settlement
	CreatedAt       data.MultiDateTime     `json:"createdAt"`        // When settlement record was created
	UpdatedAt       data.MultiDateTime     `json:"updatedAt"`        // When settlement was last updated
}

// SettlementTransaction represents a transaction within a settlement
type SettlementTransaction struct {
	ID              uint64              `json:"id"`
	Domain          string              `json:"domain"`
	Status          string              `json:"status"`
	Reference       string              `json:"reference"`
	Amount          int64               `json:"amount"` // Transaction amount in kobo
	Message         string              `json:"message"`
	GatewayResponse string              `json:"gateway_response"`
	PaidAt          *data.MultiDateTime `json:"paid_at"`
	CreatedAt       data.MultiDateTime  `json:"createdAt"`
	Channel         string              `json:"channel"`
	Currency        string              `json:"currency"`
	IPAddress       string              `json:"ip_address"`
	Metadata        *Metadata           `json:"metadata"`
	Log             *Metadata           `json:"log"`
	Fees            int64               `json:"fees"`          // Fees charged for this transaction
	FeesSplit       *Metadata           `json:"fees_split"`    // Breakdown of fees
	Customer        *Metadata           `json:"customer"`      // Customer information
	Authorization   *Metadata           `json:"authorization"` // Authorization details
	Plan            *Metadata           `json:"plan"`          // Plan details if subscription
	Subaccount      *Metadata           `json:"subaccount"`    // Subaccount details if applicable
}
