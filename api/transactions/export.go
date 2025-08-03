package transactions

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/optional"
	"github.com/huysamen/paystack-go/types"
)

type TransactionExportRequest struct {
	// Optional
	PerPage     *int
	Page        *int
	From        *time.Time
	To          *time.Time
	Customer    *uint64
	Status      *string
	Currency    *types.Currency
	Amount      *int
	Settled     *bool
	Settlement  *uint64
	PaymentPage *uint64
}

// TransactionExportRequestBuilder provides a fluent interface for building TransactionExportRequest
type TransactionExportRequestBuilder struct {
	req *TransactionExportRequest
}

// NewTransactionExportRequest creates a new builder for TransactionExportRequest
func NewTransactionExportRequest() *TransactionExportRequestBuilder {
	return &TransactionExportRequestBuilder{
		req: &TransactionExportRequest{},
	}
}

// PerPage sets the number of records per page
func (b *TransactionExportRequestBuilder) PerPage(perPage int) *TransactionExportRequestBuilder {
	b.req.PerPage = optional.Int(perPage)
	return b
}

// Page sets the page number
func (b *TransactionExportRequestBuilder) Page(page int) *TransactionExportRequestBuilder {
	b.req.Page = optional.Int(page)
	return b
}

// DateRange sets both start and end date filters
func (b *TransactionExportRequestBuilder) DateRange(from, to time.Time) *TransactionExportRequestBuilder {
	b.req.From = optional.Time(from)
	b.req.To = optional.Time(to)
	return b
}

// From sets the start date filter
func (b *TransactionExportRequestBuilder) From(from time.Time) *TransactionExportRequestBuilder {
	b.req.From = optional.Time(from)
	return b
}

// To sets the end date filter
func (b *TransactionExportRequestBuilder) To(to time.Time) *TransactionExportRequestBuilder {
	b.req.To = optional.Time(to)
	return b
}

// Customer filters by customer ID
func (b *TransactionExportRequestBuilder) Customer(customer uint64) *TransactionExportRequestBuilder {
	b.req.Customer = optional.Uint64(customer)
	return b
}

// Status filters by transaction status
func (b *TransactionExportRequestBuilder) Status(status string) *TransactionExportRequestBuilder {
	b.req.Status = optional.String(status)
	return b
}

// Currency filters by currency
func (b *TransactionExportRequestBuilder) Currency(currency types.Currency) *TransactionExportRequestBuilder {
	b.req.Currency = &currency
	return b
}

// Amount filters by transaction amount
func (b *TransactionExportRequestBuilder) Amount(amount int) *TransactionExportRequestBuilder {
	b.req.Amount = optional.Int(amount)
	return b
}

// Settled filters by settlement status
func (b *TransactionExportRequestBuilder) Settled(settled bool) *TransactionExportRequestBuilder {
	b.req.Settled = optional.Bool(settled)
	return b
}

// Settlement filters by settlement ID
func (b *TransactionExportRequestBuilder) Settlement(settlement uint64) *TransactionExportRequestBuilder {
	b.req.Settlement = optional.Uint64(settlement)
	return b
}

// PaymentPage filters by payment page ID
func (b *TransactionExportRequestBuilder) PaymentPage(paymentPage uint64) *TransactionExportRequestBuilder {
	b.req.PaymentPage = optional.Uint64(paymentPage)
	return b
}

// Build returns the constructed TransactionExportRequest
func (b *TransactionExportRequestBuilder) Build() *TransactionExportRequest {
	return b.req
}

func (r *TransactionExportRequest) toQuery() string {
	params := url.Values{}

	if r.PerPage != nil {
		params.Add("perPage", strconv.Itoa(*r.PerPage))
	}
	if r.Page != nil {
		params.Add("page", strconv.Itoa(*r.Page))
	}
	if r.From != nil {
		params.Add("from", r.From.Format("2006-01-02T15:04:05.999Z"))
	}
	if r.To != nil {
		params.Add("to", r.To.Format("2006-01-02T15:04:05.999Z"))
	}
	if r.Customer != nil {
		params.Add("customer", strconv.FormatUint(*r.Customer, 10))
	}
	if r.Status != nil {
		params.Add("status", *r.Status)
	}
	if r.Currency != nil {
		params.Add("currency", r.Currency.String())
	}
	if r.Amount != nil {
		params.Add("amount", strconv.Itoa(*r.Amount))
	}
	if r.Settled != nil {
		params.Add("settled", strconv.FormatBool(*r.Settled))
	}
	if r.Settlement != nil {
		params.Add("settlement", strconv.FormatUint(*r.Settlement, 10))
	}
	if r.PaymentPage != nil {
		params.Add("payment_page", strconv.FormatUint(*r.PaymentPage, 10))
	}

	return params.Encode()
}

type TransactionExportResponse struct {
	Path      string         `json:"path"`
	ExpiresAt types.DateTime `json:"expiresAt"`
}

// Response type alias
type TransactionExportApiResponse = types.Response[TransactionExportResponse]

// Export exports transactions using a builder (fluent interface)
func (c *Client) Export(ctx context.Context, builder *TransactionExportRequestBuilder) (*TransactionExportApiResponse, error) {
	req := builder.Build()
	query := ""
	if req != nil {
		query = req.toQuery()
	}
	return net.Get[TransactionExportResponse](ctx, c.Client, c.Secret, fmt.Sprintf("%s%s", basePath, transactionExportPath), query, c.BaseURL)
}
