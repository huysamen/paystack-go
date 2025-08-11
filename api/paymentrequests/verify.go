package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
)

// Integration represents integration details in payment request responses
type Integration struct {
	Key               data.String   `json:"key"`
	Name              data.String   `json:"name"`
	Logo              data.String   `json:"logo"`
	AllowedCurrencies []data.String `json:"allowed_currencies"`
}

// VerifyResponseData represents the response data for verifying a payment request
type VerifyResponseData struct {
	ID               data.Int             `json:"id"`
	Domain           data.String          `json:"domain"`
	Amount           data.Int             `json:"amount"`
	Currency         data.String          `json:"currency"`
	DueDate          data.NullTime        `json:"due_date"`
	HasInvoice       data.Bool            `json:"has_invoice"`
	InvoiceNumber    data.NullInt         `json:"invoice_number"`
	Description      data.String          `json:"description"`
	PDFUrl           data.NullString      `json:"pdf_url"`
	LineItems        []types.LineItem     `json:"line_items"`
	Tax              []types.Tax          `json:"tax"`
	RequestCode      data.String          `json:"request_code"`
	Status           data.String          `json:"status"`
	Paid             data.Bool            `json:"paid"`
	PaidAt           data.NullTime        `json:"paid_at"`
	Metadata         types.Metadata       `json:"metadata"`
	Notifications    []types.Notification `json:"notifications"`
	OfflineReference data.NullString      `json:"offline_reference"`
	Customer         types.Customer       `json:"customer"`
	CreatedAt        data.Time            `json:"created_at"`
	Integration      Integration          `json:"integration"`
	PendingAmount    data.Int             `json:"pending_amount"`
}

type VerifyResponse = types.Response[VerifyResponseData]

func (c *Client) Verify(ctx context.Context, code string) (*VerifyResponse, error) {
	return net.Get[VerifyResponseData](ctx, c.Client, c.Secret, basePath+"/verify/"+code, c.BaseURL)
}
