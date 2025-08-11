package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
)

// FetchResponseData represents the response data for fetching a payment request
type FetchResponseData struct {
	ID           data.Int         `json:"id"`
	Description  data.String      `json:"description"`
	Amount       data.Int         `json:"amount"`
	SplitCode    data.NullString  `json:"split_code"`
	Customer     types.Customer   `json:"customer"`
	Status       data.String      `json:"status"`
	Currency     data.String      `json:"currency"`
	RequestCode  data.String      `json:"request_code"`
	DueDate      data.NullTime    `json:"due_date"`
	HasInvoice   data.Bool        `json:"has_invoice"`
	SendEmail    data.Bool        `json:"send_email"`
	SendSMS      data.Bool        `json:"send_sms"`
	Paid         data.Bool        `json:"paid"`
	PaidAt       data.NullTime    `json:"paid_at"`
	Metadata     types.Metadata   `json:"metadata"`
	Notifications []types.Notification `json:"notifications"`
	OfflineReference data.NullString `json:"offline_reference"`
	Source       data.String      `json:"source"`
	PaymentMethod data.NullString `json:"payment_method"`
	Note         data.NullString  `json:"note"`
	Invoice      *types.Invoice   `json:"invoice"`
	PendingAmount data.Int        `json:"pending_amount"`
	Integration  data.Int         `json:"integration"`
	Domain       data.String      `json:"domain"`
	CreatedAt    data.Time        `json:"createdAt"`
	UpdatedAt    data.Time        `json:"updatedAt"`
}

type FetchResponse = types.Response[FetchResponseData]

func (c *Client) Fetch(ctx context.Context, idOrCode string) (*FetchResponse, error) {
	return net.Get[FetchResponseData](ctx, c.Client, c.Secret, basePath+"/"+idOrCode, c.BaseURL)
}
