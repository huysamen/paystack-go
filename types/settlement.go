package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Settlement represents a settlement record
type Settlement struct {
	ID              data.Uint              `json:"id"`
	Domain          data.String            `json:"domain"`
	Status          enums.SettlementStatus `json:"status"`
	Currency        enums.Currency         `json:"currency"`
	Integration     data.Uint              `json:"integration"`
	TotalAmount     data.Int               `json:"total_amount"`     // Amount after fees in kobo
	EffectiveAmount data.Int               `json:"effective_amount"` // Amount actually settled in kobo
	TotalFees       data.Int               `json:"total_fees"`       // Total fees charged in kobo
	TotalProcessed  data.Int               `json:"total_processed"`  // Total amount processed in kobo
	Deductions      *Metadata              `json:"deductions"`       // Any deductions applied
	SettlementDate  data.MultiDateTime     `json:"settlement_date"`  // Date settlement was made
	SettledBy       data.NullString        `json:"settled_by"`       // Who processed the settlement
	CreatedAt       data.MultiDateTime     `json:"createdAt"`        // When settlement record was created
	UpdatedAt       data.MultiDateTime     `json:"updatedAt"`        // When settlement was last updated
}

// SettlementTransaction represents a transaction within a settlement
type SettlementTransaction struct {
	ID              data.Uint           `json:"id"`
	Domain          data.String         `json:"domain"`
	Status          data.String         `json:"status"`
	Reference       data.String         `json:"reference"`
	Amount          data.Int            `json:"amount"` // Transaction amount in kobo
	Message         data.String         `json:"message"`
	GatewayResponse data.String         `json:"gateway_response"`
	PaidAt          *data.MultiDateTime `json:"paid_at"`
	CreatedAt       data.MultiDateTime  `json:"createdAt"`
	Channel         data.String         `json:"channel"`
	Currency        data.String         `json:"currency"`
	IPAddress       data.String         `json:"ip_address"`
	Metadata        *Metadata           `json:"metadata"`
	Log             *Metadata           `json:"log"`
	Fees            data.Int            `json:"fees"`          // Fees charged for this transaction
	FeesSplit       *Metadata           `json:"fees_split"`    // Breakdown of fees
	Customer        *Metadata           `json:"customer"`      // Customer information
	Authorization   *Metadata           `json:"authorization"` // Authorization details
	Plan            *Metadata           `json:"plan"`          // Plan details if subscription
	Subaccount      *Metadata           `json:"subaccount"`    // Subaccount details if applicable
}
