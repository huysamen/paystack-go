package transactions

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/optional"
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
)

type exportRequest struct {
	PerPage     *int
	Page        *int
	From        *time.Time
	To          *time.Time
	Customer    *uint64
	Status      *string
	Currency    *enums.Currency
	Amount      *int
	Settled     *bool
	Settlement  *uint64
	PaymentPage *uint64
}

type ExportRequestBuilder struct {
	req *exportRequest
}

func NewExportRequestBuilder() *ExportRequestBuilder {
	return &ExportRequestBuilder{
		req: &exportRequest{},
	}
}

func (b *ExportRequestBuilder) PerPage(perPage int) *ExportRequestBuilder {
	b.req.PerPage = optional.Int(perPage)

	return b
}

func (b *ExportRequestBuilder) Page(page int) *ExportRequestBuilder {
	b.req.Page = optional.Int(page)

	return b
}

func (b *ExportRequestBuilder) DateRange(from, to time.Time) *ExportRequestBuilder {
	b.req.From = optional.Time(from)
	b.req.To = optional.Time(to)

	return b
}

func (b *ExportRequestBuilder) From(from time.Time) *ExportRequestBuilder {
	b.req.From = optional.Time(from)

	return b
}

func (b *ExportRequestBuilder) To(to time.Time) *ExportRequestBuilder {
	b.req.To = optional.Time(to)

	return b
}

func (b *ExportRequestBuilder) Customer(customer uint64) *ExportRequestBuilder {
	b.req.Customer = optional.Uint64(customer)

	return b
}

func (b *ExportRequestBuilder) Status(status string) *ExportRequestBuilder {
	b.req.Status = optional.String(status)

	return b
}

func (b *ExportRequestBuilder) Currency(currency enums.Currency) *ExportRequestBuilder {
	b.req.Currency = &currency

	return b
}

func (b *ExportRequestBuilder) Amount(amount int) *ExportRequestBuilder {
	b.req.Amount = optional.Int(amount)

	return b
}

func (b *ExportRequestBuilder) Settled(settled bool) *ExportRequestBuilder {
	b.req.Settled = optional.Bool(settled)

	return b
}

func (b *ExportRequestBuilder) Settlement(settlement uint64) *ExportRequestBuilder {
	b.req.Settlement = optional.Uint64(settlement)

	return b
}

func (b *ExportRequestBuilder) PaymentPage(paymentPage uint64) *ExportRequestBuilder {
	b.req.PaymentPage = optional.Uint64(paymentPage)

	return b
}

func (b *ExportRequestBuilder) Build() *exportRequest {
	return b.req
}

func (r *exportRequest) toQuery() string {
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

type ExportResponseData struct {
	Path      string             `json:"path"`
	ExpiresAt data.MultiDateTime `json:"expiresAt"`
}

type ExportResponse = types.Response[ExportResponseData]

func (c *Client) Export(ctx context.Context, builder ExportRequestBuilder) (*ExportResponse, error) {
	req := builder.Build()
	query := ""

	if req != nil {
		query = req.toQuery()
	}

	return net.Get[ExportResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s%s", basePath, transactionExportPath), query, c.BaseURL)
}
