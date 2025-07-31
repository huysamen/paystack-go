package bulkcharges

import (
	"github.com/huysamen/paystack-go/types"
)

// BulkChargeItem represents a single charge in a bulk charge request
type BulkChargeItem struct {
	Authorization string `json:"authorization"`
	Amount        int64  `json:"amount"`
	Reference     string `json:"reference"`
}

// BulkChargeBatch represents a bulk charge batch
type BulkChargeBatch struct {
	ID             int    `json:"id"`
	BatchCode      string `json:"batch_code"`
	Reference      string `json:"reference,omitempty"`
	Integration    int    `json:"integration,omitempty"`
	Domain         string `json:"domain"`
	Status         string `json:"status"`
	TotalCharges   int    `json:"total_charges"`
	PendingCharges int    `json:"pending_charges"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
}

// Customer represents customer information in bulk charge
type Customer struct {
	ID           int         `json:"id"`
	FirstName    string      `json:"first_name"`
	LastName     string      `json:"last_name"`
	Email        string      `json:"email"`
	CustomerCode string      `json:"customer_code"`
	Phone        string      `json:"phone"`
	Metadata     interface{} `json:"metadata"`
	RiskAction   string      `json:"risk_action"`
}

// Authorization represents authorization information for bulk charge
type Authorization struct {
	AuthorizationCode string `json:"authorization_code"`
	Bin               string `json:"bin"`
	Last4             string `json:"last4"`
	ExpMonth          string `json:"exp_month"`
	ExpYear           string `json:"exp_year"`
	CardType          string `json:"card_type"`
	Bank              string `json:"bank"`
	CountryCode       string `json:"country_code"`
	Brand             string `json:"brand"`
	Reusable          bool   `json:"reusable"`
	Signature         string `json:"signature"`
	AccountName       string `json:"account_name"`
}

// BulkChargeCharge represents a single charge within a bulk charge batch
type BulkChargeCharge struct {
	Integration   int           `json:"integration"`
	BulkCharge    int           `json:"bulkcharge"`
	Customer      Customer      `json:"customer"`
	Authorization Authorization `json:"authorization"`
	Transaction   interface{}   `json:"transaction"`
	Domain        string        `json:"domain"`
	Amount        int64         `json:"amount"`
	Currency      string        `json:"currency"`
	Reference     string        `json:"reference"`
	Status        string        `json:"status"`
	Message       string        `json:"message"`
	PaidAt        string        `json:"paid_at"`
	CreatedAt     string        `json:"createdAt"`
	UpdatedAt     string        `json:"updatedAt"`
}

// InitiateBulkChargeRequest represents the request to initiate a bulk charge
type InitiateBulkChargeRequest []BulkChargeItem

// ListBulkChargeBatchesRequest represents the request to list bulk charge batches
type ListBulkChargeBatchesRequest struct {
	PerPage *int    `json:"perPage,omitempty"`
	Page    *int    `json:"page,omitempty"`
	From    *string `json:"from,omitempty"`
	To      *string `json:"to,omitempty"`
}

// FetchChargesInBatchRequest represents the request to fetch charges in a batch
type FetchChargesInBatchRequest struct {
	Status  *string `json:"status,omitempty"`
	PerPage *int    `json:"perPage,omitempty"`
	Page    *int    `json:"page,omitempty"`
	From    *string `json:"from,omitempty"`
	To      *string `json:"to,omitempty"`
}

// InitiateBulkChargeResponse represents the response from initiating a bulk charge
type InitiateBulkChargeResponse struct {
	Status  bool            `json:"status"`
	Message string          `json:"message"`
	Data    BulkChargeBatch `json:"data"`
}

// ListBulkChargeBatchesResponse represents the response from listing bulk charge batches
type ListBulkChargeBatchesResponse struct {
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Data    []BulkChargeBatch `json:"data"`
	Meta    *types.Meta       `json:"meta,omitempty"`
}

// FetchBulkChargeBatchResponse represents the response from fetching a bulk charge batch
type FetchBulkChargeBatchResponse struct {
	Status  bool            `json:"status"`
	Message string          `json:"message"`
	Data    BulkChargeBatch `json:"data"`
}

// FetchChargesInBatchResponse represents the response from fetching charges in a batch
type FetchChargesInBatchResponse struct {
	Status  bool               `json:"status"`
	Message string             `json:"message"`
	Data    []BulkChargeCharge `json:"data"`
	Meta    *types.Meta        `json:"meta,omitempty"`
}

// PauseBulkChargeBatchResponse represents the response from pausing a bulk charge batch
type PauseBulkChargeBatchResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// ResumeBulkChargeBatchResponse represents the response from resuming a bulk charge batch
type ResumeBulkChargeBatchResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
