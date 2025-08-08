package types
import "github.com/huysamen/paystack-go/types/data"
import "github.com/huysamen/paystack-go/enums"

// Dispute represents a dispute
type Dispute struct {
	ID                           int                      `json:"id"`
	RefundAmount                 *int                     `json:"refund_amount"`
	Currency                     *enums.Currency          `json:"currency"`
	Status                       enums.DisputeStatus      `json:"status"`
	Resolution                   *enums.DisputeResolution `json:"resolution"`
	Domain                       string                   `json:"domain"`
	Transaction                  *Transaction             `json:"transaction"`
	TransactionReference         *string                  `json:"transaction_reference"`
	MerchantTransactionReference *string                  `json:"merchant_transaction_reference"`
	Source                       *enums.DisputeSource     `json:"source"`
	Category                     *enums.DisputeCategory   `json:"category"`
	Note                         *string                  `json:"note"`
	Attachments                  *string                  `json:"attachments"`
	LastFour                     *string                  `json:"last4"`
	BIN                          *string                  `json:"bin"`
	Customer                     *Customer                `json:"customer"`
	CreatedAt                    data.MultiDateTime                 `json:"createdAt"`
	UpdatedAt                    data.MultiDateTime                 `json:"updatedAt"`
	DueAt                        *data.MultiDateTime                `json:"dueAt"`
	ResolvedAt                   *data.MultiDateTime                `json:"resolvedAt"`
	Evidence                     *Evidence                `json:"evidence"`
	Messages                     []DisputeMessage         `json:"messages"`
	History                      []DisputeHistory         `json:"history"`
}

// Evidence represents dispute evidence
type Evidence struct {
	ID              int       `json:"id"`
	CustomerEmail   string    `json:"customer_email"`
	CustomerName    string    `json:"customer_name"`
	CustomerPhone   string    `json:"customer_phone"`
	ServiceDetails  string    `json:"service_details"`
	DeliveryAddress *string   `json:"delivery_address"`
	DeliveryDate    *data.MultiDateTime `json:"delivery_date"`
	Dispute         int       `json:"dispute"`
	CreatedAt       data.MultiDateTime  `json:"created_at"`
	UpdatedAt       data.MultiDateTime  `json:"updated_at"`
}

// DisputeMessage represents a dispute message
type DisputeMessage struct {
	ID        int      `json:"id"`
	Sender    string   `json:"sender"`
	Body      string   `json:"body"`
	Dispute   int      `json:"dispute"`
	IsDeleted bool     `json:"is_deleted"`
	CreatedAt data.MultiDateTime `json:"created_at"`
	UpdatedAt data.MultiDateTime `json:"updated_at"`
}

// DisputeHistory represents dispute history
type DisputeHistory struct {
	ID        int                 `json:"id"`
	Dispute   int                 `json:"dispute"`
	Status    enums.DisputeStatus `json:"status"`
	By        string              `json:"by"`
	CreatedAt data.MultiDateTime            `json:"created_at"`
	UpdatedAt data.MultiDateTime            `json:"updated_at"`
}
