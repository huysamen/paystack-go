package transactions

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/optional"
	"github.com/huysamen/paystack-go/types"
)

type TransactionTotalsRequest struct {
	PerPage *int
	Page    *int
	From    *time.Time
	To      *time.Time
}

type TransactionTotalsRequestBuilder struct {
	req *TransactionTotalsRequest
}

func NewTransactionTotalsRequest() *TransactionTotalsRequestBuilder {
	return &TransactionTotalsRequestBuilder{
		req: &TransactionTotalsRequest{},
	}
}

func (b *TransactionTotalsRequestBuilder) PerPage(perPage int) *TransactionTotalsRequestBuilder {
	b.req.PerPage = optional.Int(perPage)

	return b
}

func (b *TransactionTotalsRequestBuilder) Page(page int) *TransactionTotalsRequestBuilder {
	b.req.Page = optional.Int(page)

	return b
}

func (b *TransactionTotalsRequestBuilder) DateRange(from, to time.Time) *TransactionTotalsRequestBuilder {
	b.req.From = optional.Time(from)
	b.req.To = optional.Time(to)

	return b
}

func (b *TransactionTotalsRequestBuilder) From(from time.Time) *TransactionTotalsRequestBuilder {
	b.req.From = optional.Time(from)

	return b
}

func (b *TransactionTotalsRequestBuilder) To(to time.Time) *TransactionTotalsRequestBuilder {
	b.req.To = optional.Time(to)

	return b
}

func (b *TransactionTotalsRequestBuilder) Build() *TransactionTotalsRequest {
	return b.req
}

func (r *TransactionTotalsRequest) toQuery() string {
	params := url.Values{}

	if r.PerPage != nil {
		params.Add("perPage", fmt.Sprintf("%d", *r.PerPage))
	}
	if r.Page != nil {
		params.Add("page", fmt.Sprintf("%d", *r.Page))
	}
	if r.From != nil {
		params.Add("from", r.From.Format("2006-01-02T15:04:05.999Z"))
	}
	if r.To != nil {
		params.Add("to", r.To.Format("2006-01-02T15:04:05.999Z"))
	}

	return params.Encode()
}

type CurrencyTotal struct {
	Currency types.Currency `json:"currency"`
	Amount   int            `json:"amount"`
}

type TotalsResponseData struct {
	TotalTransactions                    int             `json:"total_transactions"`
	TotalVolume                          int             `json:"total_volume"`
	TotalVolumeByCurrency                []CurrencyTotal `json:"total_volume_by_currency"`
	PendingTransfers                     int             `json:"pending_transfers"`
	PendingTransfersByCurrency           []CurrencyTotal `json:"pending_transfers_by_currency"`
	UnsettledTransactionCount            int             `json:"unsettled_transaction_count"`
	UnsettledTransactionVolume           int             `json:"unsettled_transaction_volume"`
	UnsettledTransactionVolumeByCurrency []CurrencyTotal `json:"unsettled_transaction_volume_by_currency"`
}

type TotalsResponse = types.Response[TotalsResponseData]

func (c *Client) Totals(ctx context.Context, builder *TransactionTotalsRequestBuilder) (*TotalsResponse, error) {
	req := builder.Build()
	query := ""

	if req != nil {
		query = req.toQuery()
	}

	return net.Get[TotalsResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s/totals", basePath), query, c.BaseURL)
}
