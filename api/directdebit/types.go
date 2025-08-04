package directdebit

import (
	"github.com/huysamen/paystack-go/types"
)

type MandateAuthorizationStatus string

const (
	MandateAuthorizationStatusPending MandateAuthorizationStatus = "pending"
	MandateAuthorizationStatusActive  MandateAuthorizationStatus = "active"
	MandateAuthorizationStatusRevoked MandateAuthorizationStatus = "revoked"
)

type MandateAuthorization struct {
	ID                int                        `json:"id"`
	Status            MandateAuthorizationStatus `json:"status"`
	MandateID         int                        `json:"mandate_id"`
	AuthorizationID   int                        `json:"authorization_id"`
	AuthorizationCode string                     `json:"authorization_code"`
	IntegrationID     int                        `json:"integration_id"`
	AccountNumber     string                     `json:"account_number"`
	BankCode          string                     `json:"bank_code"`
	BankName          string                     `json:"bank_name"`
	Customer          *types.Customer            `json:"customer"`
	CreatedAt         string                     `json:"created_at,omitempty"`
	UpdatedAt         string                     `json:"updated_at,omitempty"`
}
