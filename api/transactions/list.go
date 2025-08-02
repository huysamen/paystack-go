package transactions

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/optional"
	"github.com/huysamen/paystack-go/types"
)

type TransactionListRequest struct {
	// Optional
	PerPage    *int
	Page       *int
	Customer   *uint64
	TerminalID *string
	Status     *string
	From       *time.Time
	To         *time.Time
	Amount     *int
}

// TransactionListRequestBuilder provides a fluent interface for building TransactionListRequest
type TransactionListRequestBuilder struct {
	req *TransactionListRequest
}

// NewTransactionListRequest creates a new builder for TransactionListRequest
func NewTransactionListRequest() *TransactionListRequestBuilder {
	return &TransactionListRequestBuilder{
		req: &TransactionListRequest{},
	}
}

// PerPage sets the number of records per page
func (b *TransactionListRequestBuilder) PerPage(perPage int) *TransactionListRequestBuilder {
	b.req.PerPage = optional.Int(perPage)
	return b
}

// Page sets the page number
func (b *TransactionListRequestBuilder) Page(page int) *TransactionListRequestBuilder {
	b.req.Page = optional.Int(page)
	return b
}

// Customer filters by customer ID
func (b *TransactionListRequestBuilder) Customer(customer uint64) *TransactionListRequestBuilder {
	b.req.Customer = optional.Uint64(customer)
	return b
}

// TerminalID filters by terminal ID
func (b *TransactionListRequestBuilder) TerminalID(terminalID string) *TransactionListRequestBuilder {
	b.req.TerminalID = optional.String(terminalID)
	return b
}

// Status filters by transaction status
func (b *TransactionListRequestBuilder) Status(status string) *TransactionListRequestBuilder {
	b.req.Status = optional.String(status)
	return b
}

// DateRange sets both start and end date filters
func (b *TransactionListRequestBuilder) DateRange(from, to time.Time) *TransactionListRequestBuilder {
	b.req.From = optional.Time(from)
	b.req.To = optional.Time(to)
	return b
}

// From sets the start date filter
func (b *TransactionListRequestBuilder) From(from time.Time) *TransactionListRequestBuilder {
	b.req.From = optional.Time(from)
	return b
}

// To sets the end date filter
func (b *TransactionListRequestBuilder) To(to time.Time) *TransactionListRequestBuilder {
	b.req.To = optional.Time(to)
	return b
}

// Amount filters by transaction amount
func (b *TransactionListRequestBuilder) Amount(amount int) *TransactionListRequestBuilder {
	b.req.Amount = optional.Int(amount)
	return b
}

// Build returns the constructed TransactionListRequest
func (b *TransactionListRequestBuilder) Build() *TransactionListRequest {
	return b.req
}

func (r *TransactionListRequest) toQuery() string {
	params := url.Values{}

	if r.PerPage != nil {
		params.Add("perPage", strconv.Itoa(*r.PerPage))
	}
	if r.Page != nil {
		params.Add("page", strconv.Itoa(*r.Page))
	}
	if r.Customer != nil {
		params.Add("customer", strconv.FormatUint(*r.Customer, 10))
	}
	if r.TerminalID != nil {
		params.Add("terminalid", *r.TerminalID)
	}
	if r.Status != nil {
		params.Add("status", *r.Status)
	}
	if r.From != nil {
		params.Add("from", r.From.Format("2006-01-02T15:04:05.999Z"))
	}
	if r.To != nil {
		params.Add("to", r.To.Format("2006-01-02T15:04:05.999Z"))
	}
	if r.Amount != nil {
		params.Add("amount", strconv.Itoa(*r.Amount))
	}

	return params.Encode()
}

type TransactionListResponse []types.Transaction

// List lists transactions using a builder (fluent interface)
func (c *Client) List(ctx context.Context, builder *TransactionListRequestBuilder) (*types.Response[TransactionListResponse], error) {
	req := builder.Build()
	query := ""
	if req != nil {
		query = req.toQuery()
	}
	return net.Get[TransactionListResponse](ctx, c.Client, c.Secret, basePath, query, c.BaseURL)
}
