package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Dispute represents a dispute
type Dispute struct {
	ID                           data.Int                 `json:"id"`
	RefundAmount                 data.NullInt             `json:"refund_amount"`
	Currency                     *enums.Currency          `json:"currency"`
	Status                       enums.DisputeStatus      `json:"status"`
	Resolution                   *enums.DisputeResolution `json:"resolution"`
	Domain                       data.String              `json:"domain"`
	Transaction                  *Transaction             `json:"transaction"`
	TransactionReference         data.NullString          `json:"transaction_reference"`
	MerchantTransactionReference data.NullString          `json:"merchant_transaction_reference"`
	Source                       *enums.DisputeSource     `json:"source"`
	Category                     *enums.DisputeCategory   `json:"category"`
	Note                         data.NullString          `json:"note"`
	Attachments                  data.NullString          `json:"attachments"`
	LastFour                     data.NullString          `json:"last4"`
	BIN                          data.NullString          `json:"bin"`
	Customer                     *Customer                `json:"customer"`
	CreatedAt                    data.MultiDateTime       `json:"createdAt"`
	UpdatedAt                    data.MultiDateTime       `json:"updatedAt"`
	DueAt                        *data.MultiDateTime      `json:"dueAt"`
	ResolvedAt                   *data.MultiDateTime      `json:"resolvedAt"`
	Evidence                     *Evidence                `json:"evidence"`
	Messages                     []DisputeMessage         `json:"messages"`
	History                      []DisputeHistory         `json:"history"`
}

// Evidence represents dispute evidence
type Evidence struct {
	ID              data.Int            `json:"id"`
	CustomerEmail   data.String         `json:"customer_email"`
	CustomerName    data.String         `json:"customer_name"`
	CustomerPhone   data.String         `json:"customer_phone"`
	ServiceDetails  data.String         `json:"service_details"`
	DeliveryAddress data.NullString     `json:"delivery_address"`
	DeliveryDate    *data.MultiDateTime `json:"delivery_date"`
	Dispute         data.Int            `json:"dispute"`
	CreatedAt       data.MultiDateTime  `json:"created_at"`
	UpdatedAt       data.MultiDateTime  `json:"updated_at"`
}

// DisputeMessage represents a dispute message
type DisputeMessage struct {
	ID        data.Int           `json:"id"`
	Sender    data.String        `json:"sender"`
	Body      data.String        `json:"body"`
	Dispute   data.Int           `json:"dispute"`
	IsDeleted data.MultiBool     `json:"is_deleted"`
	CreatedAt data.MultiDateTime `json:"created_at"`
	UpdatedAt data.MultiDateTime `json:"updated_at"`
}

// DisputeHistory represents dispute history
type DisputeHistory struct {
	ID        data.Int            `json:"id"`
	Dispute   data.Int            `json:"dispute"`
	Status    enums.DisputeStatus `json:"status"`
	By        data.String         `json:"by"`
	CreatedAt data.MultiDateTime  `json:"created_at"`
	UpdatedAt data.MultiDateTime  `json:"updated_at"`
}
