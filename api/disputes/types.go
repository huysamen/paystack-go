package disputes

import (
	"time"

	"github.com/huysamen/paystack-go/types"
)

// DisputeStatus represents the status of a dispute
type DisputeStatus string

const (
	DisputeStatusAwaitingMerchantFeedback DisputeStatus = "awaiting-merchant-feedback"
	DisputeStatusAwaitingBankFeedback     DisputeStatus = "awaiting-bank-feedback"
	DisputeStatusPending                  DisputeStatus = "pending"
	DisputeStatusResolved                 DisputeStatus = "resolved"
	DisputeStatusArchived                 DisputeStatus = "archived"
)

// String returns the string representation of DisputeStatus
func (ds DisputeStatus) String() string {
	return string(ds)
}

// DisputeResolution represents the resolution of a dispute
type DisputeResolution string

const (
	DisputeResolutionMerchantAccepted DisputeResolution = "merchant-accepted"
	DisputeResolutionDeclined         DisputeResolution = "declined"
)

// String returns the string representation of DisputeResolution
func (dr DisputeResolution) String() string {
	return string(dr)
}

// DisputeSource represents the source of a dispute
type DisputeSource string

const (
	DisputeSourceBank DisputeSource = "bank"
	DisputeSourceCard DisputeSource = "card"
)

// String returns the string representation of DisputeSource
func (ds DisputeSource) String() string {
	return string(ds)
}

// DisputeCategory represents the category of a dispute
type DisputeCategory string

const (
	DisputeCategoryGeneral          DisputeCategory = "general"
	DisputeCategoryFraud            DisputeCategory = "fraud"
	DisputeCategoryAuthorization    DisputeCategory = "authorization"
	DisputeCategoryProcessingErrors DisputeCategory = "processing_errors"
	DisputeCategoryConsumerDispute  DisputeCategory = "consumer_dispute"
)

// String returns the string representation of DisputeCategory
func (dc DisputeCategory) String() string {
	return string(dc)
}

// Transaction represents transaction information in a dispute
type Transaction struct {
	ID              int                    `json:"id"`
	Domain          string                 `json:"domain"`
	Status          string                 `json:"status"`
	Reference       string                 `json:"reference"`
	Amount          int                    `json:"amount"`
	Message         *string                `json:"message"`
	GatewayResponse string                 `json:"gateway_response"`
	PaidAt          *types.DateTime        `json:"paid_at"`
	CreatedAt       types.DateTime         `json:"created_at"`
	Channel         string                 `json:"channel"`
	Currency        string                 `json:"currency"`
	IPAddress       *string                `json:"ip_address"`
	Metadata        map[string]interface{} `json:"metadata"`
	Log             *TransactionLog        `json:"log"`
	Fees            *int                   `json:"fees"`
	FeesSplit       interface{}            `json:"fees_split"`
	Authorization   *Authorization         `json:"authorization"`
	Customer        *Customer              `json:"customer"`
	Plan            interface{}            `json:"plan"`
	Subaccount      interface{}            `json:"subaccount"`
	Split           interface{}            `json:"split"`
	OrderID         *string                `json:"order_id"`
	PaidBy          *string                `json:"paid_by"`
}

// TransactionLog represents transaction log information
type TransactionLog struct {
	StartTime int      `json:"start_time"`
	TimeSpent int      `json:"time_spent"`
	Attempts  int      `json:"attempts"`
	Errors    int      `json:"errors"`
	Success   bool     `json:"success"`
	Mobile    bool     `json:"mobile"`
	Input     []string `json:"input"`
	History   []struct {
		Type    string `json:"type"`
		Message string `json:"message"`
		Time    int    `json:"time"`
	} `json:"history"`
}

// Authorization represents authorization information
type Authorization struct {
	AuthorizationCode string `json:"authorization_code"`
	Bin               string `json:"bin"`
	Last4             string `json:"last4"`
	ExpMonth          string `json:"exp_month"`
	ExpYear           string `json:"exp_year"`
	Channel           string `json:"channel"`
	CardType          string `json:"card_type"`
	Bank              string `json:"bank"`
	CountryCode       string `json:"country_code"`
	Brand             string `json:"brand"`
	Reusable          bool   `json:"reusable"`
	Signature         string `json:"signature"`
	AccountName       string `json:"account_name"`
}

// Customer represents customer information
type Customer struct {
	ID            int                    `json:"id"`
	FirstName     *string                `json:"first_name"`
	LastName      *string                `json:"last_name"`
	Email         string                 `json:"email"`
	CustomerCode  string                 `json:"customer_code"`
	Phone         *string                `json:"phone"`
	Metadata      map[string]interface{} `json:"metadata"`
	RiskAction    string                 `json:"risk_action"`
	International bool                   `json:"international"`
}

// Dispute represents a dispute
type Dispute struct {
	ID                           int                `json:"id"`
	RefundAmount                 *int               `json:"refund_amount"`
	Currency                     *string            `json:"currency"`
	Status                       DisputeStatus      `json:"status"`
	Resolution                   *DisputeResolution `json:"resolution"`
	Domain                       string             `json:"domain"`
	Transaction                  *Transaction       `json:"transaction"`
	TransactionReference         *string            `json:"transaction_reference"`
	MerchantTransactionReference *string            `json:"merchant_transaction_reference"`
	Source                       DisputeSource      `json:"source"`
	Category                     DisputeCategory    `json:"category"`
	Note                         *string            `json:"note"`
	Attachments                  *string            `json:"attachments"`
	LastFour                     *string            `json:"last4"`
	BIN                          *string            `json:"bin"`
	CreatedAt                    types.DateTime     `json:"created_at"`
	UpdatedAt                    types.DateTime     `json:"updated_at"`
	DueAt                        *types.DateTime    `json:"due_at"`
	ResolvedAt                   *types.DateTime    `json:"resolved_at"`
	Evidence                     *Evidence          `json:"evidence"`
	Messages                     []DisputeMessage   `json:"messages"`
	History                      []DisputeHistory   `json:"history"`
}

// Evidence represents dispute evidence
type Evidence struct {
	ID              int             `json:"id"`
	CustomerEmail   string          `json:"customer_email"`
	CustomerName    string          `json:"customer_name"`
	CustomerPhone   string          `json:"customer_phone"`
	ServiceDetails  string          `json:"service_details"`
	DeliveryAddress *string         `json:"delivery_address"`
	DeliveryDate    *types.DateTime `json:"delivery_date"`
	Dispute         int             `json:"dispute"`
	CreatedAt       types.DateTime  `json:"created_at"`
	UpdatedAt       types.DateTime  `json:"updated_at"`
}

// DisputeMessage represents a dispute message
type DisputeMessage struct {
	ID        int    `json:"id"`
	Sender    string `json:"sender"`
	Body      string `json:"body"`
	Dispute   int    `json:"dispute"`
	IsDeleted bool   `json:"is_deleted"`
}

// DisputeHistory represents dispute history
type DisputeHistory struct {
	ID        int            `json:"id"`
	Dispute   int            `json:"dispute"`
	Status    DisputeStatus  `json:"status"`
	By        string         `json:"by"`
	CreatedAt types.DateTime `json:"created_at"`
	UpdatedAt types.DateTime `json:"updated_at"`
}

// DisputeListRequest represents the request to list disputes
type DisputeListRequest struct {
	From        *time.Time     `json:"from,omitempty"`
	To          *time.Time     `json:"to,omitempty"`
	PerPage     *int           `json:"per_page,omitempty"`
	Page        *int           `json:"page,omitempty"`
	Transaction *string        `json:"transaction,omitempty"`
	Status      *DisputeStatus `json:"status,omitempty"`
}

// DisputeListResponse represents the response from listing disputes
type DisputeListResponse = types.Response[[]Dispute]

// DisputeFetchResponse represents the response from fetching a dispute
type DisputeFetchResponse = types.Response[Dispute]

// TransactionDisputeResponse represents the response from listing transaction disputes
type TransactionDisputeResponse = types.Response[TransactionDisputeData]

// DisputeUpdateResponse represents the response from updating a dispute
type DisputeUpdateResponse = types.Response[[]Dispute]

// DisputeEvidenceResponse represents the response from adding evidence to a dispute
type DisputeEvidenceResponse = types.Response[Evidence]

// DisputeUploadURLResponse represents the response from getting upload URL
type DisputeUploadURLResponse = types.Response[UploadURLData]

// DisputeResolveResponse represents the response from resolving a dispute
type DisputeResolveResponse = types.Response[Dispute]

// DisputeExportResponse represents the response from exporting disputes
type DisputeExportResponse = types.Response[ExportData]

// TransactionDisputeData represents transaction dispute data
type TransactionDisputeData struct {
	History  []DisputeHistory `json:"history"`
	Messages []DisputeMessage `json:"messages"`
	Dispute  *Dispute         `json:"dispute,omitempty"`
}

// DisputeUpdateRequest represents the request to update a dispute
type DisputeUpdateRequest struct {
	RefundAmount     *int    `json:"refund_amount,omitempty"`
	UploadedFileName *string `json:"uploaded_filename,omitempty"`
}

// DisputeEvidenceRequest represents the request to add evidence to a dispute
type DisputeEvidenceRequest struct {
	CustomerEmail   string     `json:"customer_email"`
	CustomerName    string     `json:"customer_name"`
	CustomerPhone   string     `json:"customer_phone"`
	ServiceDetails  string     `json:"service_details"`
	DeliveryAddress *string    `json:"delivery_address,omitempty"`
	DeliveryDate    *time.Time `json:"delivery_date,omitempty"`
}

// DisputeUploadURLRequest represents the request to get upload URL for a dispute
type DisputeUploadURLRequest struct {
	UploadFileName string `json:"upload_filename"`
}

// UploadURLData represents upload URL data
type UploadURLData struct {
	SignedURL string `json:"signedUrl"`
	FileName  string `json:"fileName"`
	ExpiresIn int    `json:"expiresIn"`
}

// DisputeResolveRequest represents the request to resolve a dispute
type DisputeResolveRequest struct {
	Resolution       DisputeResolution `json:"resolution"`
	Message          string            `json:"message"`
	RefundAmount     int               `json:"refund_amount"`
	UploadedFileName string            `json:"uploaded_filename"`
	Evidence         *int              `json:"evidence,omitempty"`
}

// DisputeExportRequest represents the request to export disputes
type DisputeExportRequest struct {
	From        *time.Time     `json:"from,omitempty"`
	To          *time.Time     `json:"to,omitempty"`
	PerPage     *int           `json:"per_page,omitempty"`
	Page        *int           `json:"page,omitempty"`
	Transaction *string        `json:"transaction,omitempty"`
	Status      *DisputeStatus `json:"status,omitempty"`
}

// ExportData represents export data
type ExportData struct {
	Path      string          `json:"path"`
	ExpiresAt *types.DateTime `json:"expires_at,omitempty"`
}
