package types

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
	CreatedAt                    DateTime           `json:"created_at"`
	UpdatedAt                    DateTime           `json:"updated_at"`
	DueAt                        *DateTime          `json:"due_at"`
	ResolvedAt                   *DateTime          `json:"resolved_at"`
	Evidence                     *Evidence          `json:"evidence"`
	Messages                     []DisputeMessage   `json:"messages"`
	History                      []DisputeHistory   `json:"history"`
}

// Evidence represents dispute evidence
type Evidence struct {
	ID              int       `json:"id"`
	CustomerEmail   string    `json:"customer_email"`
	CustomerName    string    `json:"customer_name"`
	CustomerPhone   string    `json:"customer_phone"`
	ServiceDetails  string    `json:"service_details"`
	DeliveryAddress *string   `json:"delivery_address"`
	DeliveryDate    *DateTime `json:"delivery_date"`
	Dispute         int       `json:"dispute"`
	CreatedAt       DateTime  `json:"created_at"`
	UpdatedAt       DateTime  `json:"updated_at"`
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
	ID        int           `json:"id"`
	Dispute   int           `json:"dispute"`
	Status    DisputeStatus `json:"status"`
	By        string        `json:"by"`
	CreatedAt DateTime      `json:"created_at"`
	UpdatedAt DateTime      `json:"updated_at"`
}
