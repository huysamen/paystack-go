package paymentrequests

import (
	"github.com/huysamen/paystack-go/types"
)

// LineItem represents an item in a payment request
type LineItem struct {
	Name     string `json:"name"`
	Amount   int    `json:"amount"`
	Quantity int    `json:"quantity,omitempty"`
}

// Tax represents a tax item in a payment request
type Tax struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

// Customer represents customer information in a payment request
type Customer struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Email        string `json:"email"`
	CustomerCode string `json:"customer_code"`
	Phone        string `json:"phone,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}

// Transaction represents a transaction related to a payment request
type Transaction struct {
	ID            int            `json:"id"`
	Reference     string         `json:"reference"`
	Amount        int            `json:"amount"`
	Currency      string         `json:"currency"`
	Status        string         `json:"status"`
	Gateway       string         `json:"gateway"`
	Channel       string         `json:"channel"`
	PaidAt        string         `json:"paid_at,omitempty"`
	CreatedAt     string         `json:"created_at"`
	Authorization map[string]any `json:"authorization,omitempty"`
	Customer      *Customer      `json:"customer,omitempty"`
}

// PaymentRequest represents a payment request
type PaymentRequest struct {
	ID               int             `json:"id"`
	Domain           string          `json:"domain"`
	Amount           int             `json:"amount"`
	Currency         string          `json:"currency"`
	DueDate          string          `json:"due_date,omitempty"`
	HasInvoice       bool            `json:"has_invoice"`
	InvoiceNumber    *int            `json:"invoice_number,omitempty"`
	Description      string          `json:"description"`
	PDFURL           *string         `json:"pdf_url,omitempty"`
	LineItems        []LineItem      `json:"line_items,omitempty"`
	Tax              []Tax           `json:"tax,omitempty"`
	RequestCode      string          `json:"request_code"`
	Status           string          `json:"status"`
	Paid             bool            `json:"paid"`
	PaidAt           *string         `json:"paid_at,omitempty"`
	Metadata         *types.Metadata `json:"metadata,omitempty"`
	Notifications    []any           `json:"notifications,omitempty"`
	OfflineReference *string         `json:"offline_reference,omitempty"`
	Customer         *Customer       `json:"customer,omitempty"`
	CreatedAt        string          `json:"created_at"`
	UpdatedAt        string          `json:"updated_at,omitempty"`
	Transactions     []Transaction   `json:"transactions,omitempty"`
	SplitCode        *string         `json:"split_code,omitempty"`
	Integration      int             `json:"integration,omitempty"`
	Archived         bool            `json:"archived,omitempty"`
}

// CreatePaymentRequestRequest represents the request to create a payment request
type CreatePaymentRequestRequest struct {
	Customer         string     `json:"customer"`
	Amount           *int       `json:"amount,omitempty"`
	DueDate          string     `json:"due_date,omitempty"`
	Description      string     `json:"description,omitempty"`
	LineItems        []LineItem `json:"line_items,omitempty"`
	Tax              []Tax      `json:"tax,omitempty"`
	Currency         string     `json:"currency,omitempty"`
	SendNotification *bool      `json:"send_notification,omitempty"`
	Draft            *bool      `json:"draft,omitempty"`
	HasInvoice       *bool      `json:"has_invoice,omitempty"`
	InvoiceNumber    *int       `json:"invoice_number,omitempty"`
	SplitCode        string     `json:"split_code,omitempty"`
}

// CreatePaymentRequestRequestBuilder provides a fluent interface for building CreatePaymentRequestRequest
type CreatePaymentRequestRequestBuilder struct {
	req *CreatePaymentRequestRequest
}

// NewCreatePaymentRequestRequest creates a new builder for CreatePaymentRequestRequest
func NewCreatePaymentRequestRequest(customer string) *CreatePaymentRequestRequestBuilder {
	return &CreatePaymentRequestRequestBuilder{
		req: &CreatePaymentRequestRequest{
			Customer: customer,
		},
	}
}

// Amount sets the payment amount (optional, for flexible amount)
func (b *CreatePaymentRequestRequestBuilder) Amount(amount int) *CreatePaymentRequestRequestBuilder {
	b.req.Amount = &amount
	return b
}

// DueDate sets the due date
func (b *CreatePaymentRequestRequestBuilder) DueDate(dueDate string) *CreatePaymentRequestRequestBuilder {
	b.req.DueDate = dueDate
	return b
}

// Description sets the payment description
func (b *CreatePaymentRequestRequestBuilder) Description(description string) *CreatePaymentRequestRequestBuilder {
	b.req.Description = description
	return b
}

// LineItems sets the line items
func (b *CreatePaymentRequestRequestBuilder) LineItems(lineItems []LineItem) *CreatePaymentRequestRequestBuilder {
	b.req.LineItems = lineItems
	return b
}

// Tax sets the tax information
func (b *CreatePaymentRequestRequestBuilder) Tax(tax []Tax) *CreatePaymentRequestRequestBuilder {
	b.req.Tax = tax
	return b
}

// Currency sets the currency
func (b *CreatePaymentRequestRequestBuilder) Currency(currency string) *CreatePaymentRequestRequestBuilder {
	b.req.Currency = currency
	return b
}

// SendNotification sets whether to send notification
func (b *CreatePaymentRequestRequestBuilder) SendNotification(sendNotification bool) *CreatePaymentRequestRequestBuilder {
	b.req.SendNotification = &sendNotification
	return b
}

// Draft sets whether this is a draft
func (b *CreatePaymentRequestRequestBuilder) Draft(draft bool) *CreatePaymentRequestRequestBuilder {
	b.req.Draft = &draft
	return b
}

// HasInvoice sets whether to generate an invoice
func (b *CreatePaymentRequestRequestBuilder) HasInvoice(hasInvoice bool) *CreatePaymentRequestRequestBuilder {
	b.req.HasInvoice = &hasInvoice
	return b
}

// InvoiceNumber sets the invoice number
func (b *CreatePaymentRequestRequestBuilder) InvoiceNumber(invoiceNumber int) *CreatePaymentRequestRequestBuilder {
	b.req.InvoiceNumber = &invoiceNumber
	return b
}

// SplitCode sets the split code
func (b *CreatePaymentRequestRequestBuilder) SplitCode(splitCode string) *CreatePaymentRequestRequestBuilder {
	b.req.SplitCode = splitCode
	return b
}

// Build returns the constructed CreatePaymentRequestRequest
func (b *CreatePaymentRequestRequestBuilder) Build() *CreatePaymentRequestRequest {
	return b.req
}

// UpdatePaymentRequestRequest represents the request to update a payment request
type UpdatePaymentRequestRequest struct {
	Customer         string     `json:"customer,omitempty"`
	Amount           *int       `json:"amount,omitempty"`
	DueDate          string     `json:"due_date,omitempty"`
	Description      string     `json:"description,omitempty"`
	LineItems        []LineItem `json:"line_items,omitempty"`
	Tax              []Tax      `json:"tax,omitempty"`
	Currency         string     `json:"currency,omitempty"`
	SendNotification *bool      `json:"send_notification,omitempty"`
	Draft            *bool      `json:"draft,omitempty"`
	InvoiceNumber    *int       `json:"invoice_number,omitempty"`
	SplitCode        string     `json:"split_code,omitempty"`
}

// UpdatePaymentRequestRequestBuilder provides a fluent interface for building UpdatePaymentRequestRequest
type UpdatePaymentRequestRequestBuilder struct {
	req *UpdatePaymentRequestRequest
}

// NewUpdatePaymentRequestRequest creates a new builder for UpdatePaymentRequestRequest
func NewUpdatePaymentRequestRequest() *UpdatePaymentRequestRequestBuilder {
	return &UpdatePaymentRequestRequestBuilder{
		req: &UpdatePaymentRequestRequest{},
	}
}

// Customer sets the customer
func (b *UpdatePaymentRequestRequestBuilder) Customer(customer string) *UpdatePaymentRequestRequestBuilder {
	b.req.Customer = customer
	return b
}

// Amount sets the payment amount
func (b *UpdatePaymentRequestRequestBuilder) Amount(amount int) *UpdatePaymentRequestRequestBuilder {
	b.req.Amount = &amount
	return b
}

// DueDate sets the due date
func (b *UpdatePaymentRequestRequestBuilder) DueDate(dueDate string) *UpdatePaymentRequestRequestBuilder {
	b.req.DueDate = dueDate
	return b
}

// Description sets the payment description
func (b *UpdatePaymentRequestRequestBuilder) Description(description string) *UpdatePaymentRequestRequestBuilder {
	b.req.Description = description
	return b
}

// LineItems sets the line items
func (b *UpdatePaymentRequestRequestBuilder) LineItems(lineItems []LineItem) *UpdatePaymentRequestRequestBuilder {
	b.req.LineItems = lineItems
	return b
}

// Tax sets the tax information
func (b *UpdatePaymentRequestRequestBuilder) Tax(tax []Tax) *UpdatePaymentRequestRequestBuilder {
	b.req.Tax = tax
	return b
}

// Currency sets the currency
func (b *UpdatePaymentRequestRequestBuilder) Currency(currency string) *UpdatePaymentRequestRequestBuilder {
	b.req.Currency = currency
	return b
}

// SendNotification sets whether to send notification
func (b *UpdatePaymentRequestRequestBuilder) SendNotification(sendNotification bool) *UpdatePaymentRequestRequestBuilder {
	b.req.SendNotification = &sendNotification
	return b
}

// Draft sets whether this is a draft
func (b *UpdatePaymentRequestRequestBuilder) Draft(draft bool) *UpdatePaymentRequestRequestBuilder {
	b.req.Draft = &draft
	return b
}

// InvoiceNumber sets the invoice number
func (b *UpdatePaymentRequestRequestBuilder) InvoiceNumber(invoiceNumber int) *UpdatePaymentRequestRequestBuilder {
	b.req.InvoiceNumber = &invoiceNumber
	return b
}

// SplitCode sets the split code
func (b *UpdatePaymentRequestRequestBuilder) SplitCode(splitCode string) *UpdatePaymentRequestRequestBuilder {
	b.req.SplitCode = splitCode
	return b
}

// Build returns the constructed UpdatePaymentRequestRequest
func (b *UpdatePaymentRequestRequestBuilder) Build() *UpdatePaymentRequestRequest {
	return b.req
}

// ListPaymentRequestsRequest represents the request to list payment requests
type ListPaymentRequestsRequest struct {
	PerPage        int    `json:"perPage,omitempty"`
	Page           int    `json:"page,omitempty"`
	Customer       string `json:"customer,omitempty"`
	Status         string `json:"status,omitempty"`
	Currency       string `json:"currency,omitempty"`
	IncludeArchive string `json:"include_archive,omitempty"`
	From           string `json:"from,omitempty"`
	To             string `json:"to,omitempty"`
}

// ListPaymentRequestsRequestBuilder provides a fluent interface for building ListPaymentRequestsRequest
type ListPaymentRequestsRequestBuilder struct {
	req *ListPaymentRequestsRequest
}

// NewListPaymentRequestsRequest creates a new builder for ListPaymentRequestsRequest
func NewListPaymentRequestsRequest() *ListPaymentRequestsRequestBuilder {
	return &ListPaymentRequestsRequestBuilder{
		req: &ListPaymentRequestsRequest{},
	}
}

// PerPage sets the number of payment requests per page
func (b *ListPaymentRequestsRequestBuilder) PerPage(perPage int) *ListPaymentRequestsRequestBuilder {
	b.req.PerPage = perPage
	return b
}

// Page sets the page number
func (b *ListPaymentRequestsRequestBuilder) Page(page int) *ListPaymentRequestsRequestBuilder {
	b.req.Page = page
	return b
}

// Customer sets the customer filter
func (b *ListPaymentRequestsRequestBuilder) Customer(customer string) *ListPaymentRequestsRequestBuilder {
	b.req.Customer = customer
	return b
}

// Status sets the status filter
func (b *ListPaymentRequestsRequestBuilder) Status(status string) *ListPaymentRequestsRequestBuilder {
	b.req.Status = status
	return b
}

// Currency sets the currency filter
func (b *ListPaymentRequestsRequestBuilder) Currency(currency string) *ListPaymentRequestsRequestBuilder {
	b.req.Currency = currency
	return b
}

// IncludeArchive sets whether to include archived requests
func (b *ListPaymentRequestsRequestBuilder) IncludeArchive(includeArchive string) *ListPaymentRequestsRequestBuilder {
	b.req.IncludeArchive = includeArchive
	return b
}

// From sets the start date filter
func (b *ListPaymentRequestsRequestBuilder) From(from string) *ListPaymentRequestsRequestBuilder {
	b.req.From = from
	return b
}

// To sets the end date filter
func (b *ListPaymentRequestsRequestBuilder) To(to string) *ListPaymentRequestsRequestBuilder {
	b.req.To = to
	return b
}

// DateRange sets both from and to dates for convenience
func (b *ListPaymentRequestsRequestBuilder) DateRange(from, to string) *ListPaymentRequestsRequestBuilder {
	b.req.From = from
	b.req.To = to
	return b
}

// Build returns the constructed ListPaymentRequestsRequest
func (b *ListPaymentRequestsRequestBuilder) Build() *ListPaymentRequestsRequest {
	return b.req
}

// FinalizePaymentRequestRequest represents the request to finalize a payment request
type FinalizePaymentRequestRequest struct {
	SendNotification *bool `json:"send_notification,omitempty"`
}

// FinalizePaymentRequestRequestBuilder provides a fluent interface for building FinalizePaymentRequestRequest
type FinalizePaymentRequestRequestBuilder struct {
	req *FinalizePaymentRequestRequest
}

// NewFinalizePaymentRequestRequest creates a new builder for FinalizePaymentRequestRequest
func NewFinalizePaymentRequestRequest() *FinalizePaymentRequestRequestBuilder {
	return &FinalizePaymentRequestRequestBuilder{
		req: &FinalizePaymentRequestRequest{},
	}
}

// SendNotification sets whether to send notification
func (b *FinalizePaymentRequestRequestBuilder) SendNotification(sendNotification bool) *FinalizePaymentRequestRequestBuilder {
	b.req.SendNotification = &sendNotification
	return b
}

// Build returns the constructed FinalizePaymentRequestRequest
func (b *FinalizePaymentRequestRequestBuilder) Build() *FinalizePaymentRequestRequest {
	return b.req
}

// PaymentRequestTotal represents payment request totals by status and currency
type PaymentRequestTotal struct {
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}

// PaymentRequestTotals represents the totals data structure
type PaymentRequestTotals struct {
	Pending    []PaymentRequestTotal `json:"pending"`
	Successful []PaymentRequestTotal `json:"successful"`
	Total      []PaymentRequestTotal `json:"total"`
}

// CreatePaymentRequestResponse represents the response from creating a payment request
type CreatePaymentRequestResponse struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    PaymentRequest `json:"data"`
}

// ListPaymentRequestsResponse represents the response from listing payment requests
type ListPaymentRequestsResponse struct {
	Status  bool             `json:"status"`
	Message string           `json:"message"`
	Data    []PaymentRequest `json:"data"`
	Meta    *types.Meta      `json:"meta,omitempty"`
}

// FetchPaymentRequestResponse represents the response from fetching a payment request
type FetchPaymentRequestResponse struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    PaymentRequest `json:"data"`
}

// VerifyPaymentRequestResponse represents the response from verifying a payment request
type VerifyPaymentRequestResponse struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    PaymentRequest `json:"data"`
}

// UpdatePaymentRequestResponse represents the response from updating a payment request
type UpdatePaymentRequestResponse struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    PaymentRequest `json:"data"`
}

// FinalizePaymentRequestResponse represents the response from finalizing a payment request
type FinalizePaymentRequestResponse struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    PaymentRequest `json:"data"`
}

// ArchivePaymentRequestResponse represents the response from archiving a payment request
type ArchivePaymentRequestResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// SendNotificationResponse represents the response from sending a notification
type SendNotificationResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// PaymentRequestTotalsResponse represents the response from getting payment request totals
type PaymentRequestTotalsResponse struct {
	Status  bool                 `json:"status"`
	Message string               `json:"message"`
	Data    PaymentRequestTotals `json:"data"`
}
