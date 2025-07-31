package directdebit

import (
	"github.com/huysamen/paystack-go/types"
)

// MandateAuthorizationStatus represents the status of a mandate authorization
type MandateAuthorizationStatus string

const (
	MandateAuthorizationStatusPending MandateAuthorizationStatus = "pending"
	MandateAuthorizationStatusActive  MandateAuthorizationStatus = "active"
	MandateAuthorizationStatusRevoked MandateAuthorizationStatus = "revoked"
)

// Customer represents customer information in mandate authorization
type Customer struct {
	ID           int             `json:"id"`
	CustomerCode string          `json:"customer_code"`
	Email        string          `json:"email"`
	FirstName    string          `json:"first_name,omitempty"`
	LastName     string          `json:"last_name,omitempty"`
	Phone        string          `json:"phone,omitempty"`
	Metadata     *types.Metadata `json:"metadata,omitempty"`
	RiskAction   string          `json:"risk_action,omitempty"`
	CreatedAt    string          `json:"created_at,omitempty"`
	UpdatedAt    string          `json:"updated_at,omitempty"`
}

// MandateAuthorization represents a direct debit mandate authorization
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
	Customer          Customer                   `json:"customer"`
	CreatedAt         string                     `json:"created_at,omitempty"`
	UpdatedAt         string                     `json:"updated_at,omitempty"`
}

// TriggerActivationChargeRequest represents the request to trigger activation charge
type TriggerActivationChargeRequest struct {
	CustomerIDs []int `json:"customer_ids"`
}

// ListMandateAuthorizationsRequest represents the request to list mandate authorizations
type ListMandateAuthorizationsRequest struct {
	Cursor  string                     `json:"cursor,omitempty"`
	Status  MandateAuthorizationStatus `json:"status,omitempty"`
	PerPage int                        `json:"per_page,omitempty"`
}

// TriggerActivationChargeResponse represents the response from triggering activation charge
type TriggerActivationChargeResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// ListMandateAuthorizationsResponse represents the response from listing mandate authorizations
type ListMandateAuthorizationsResponse struct {
	Status  bool                   `json:"status"`
	Message string                 `json:"message"`
	Data    []MandateAuthorization `json:"data"`
	Meta    *types.Meta            `json:"meta,omitempty"`
}
