package transactions

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionTotalsRequest struct {
	// Optional
	PerPage *int
	Page    *int
	From    *time.Time
	To      *time.Time
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

type TransactionTotalsResponse struct {
	TotalTransactions                    int             `json:"total_transactions"`
	TotalVolume                          int             `json:"total_volume"`
	TotalVolumeByCurrency                []CurrencyTotal `json:"total_volume_by_currency"`
	PendingTransfers                     int             `json:"pending_transfers"`
	PendingTransfersByCurrency           []CurrencyTotal `json:"pending_transfers_by_currency"`
	UnsettledTransactionCount            int             `json:"unsettled_transaction_count"`
	UnsettledTransactionVolume           int             `json:"unsettled_transaction_volume"`
	UnsettledTransactionVolumeByCurrency []CurrencyTotal `json:"unsettled_transaction_volume_by_currency"`
}

func (c *Client) Totals(ctx context.Context, req *TransactionTotalsRequest) (*types.Response[TransactionTotalsResponse], error) {
	path := fmt.Sprintf("%s/totals", transactionBasePath)

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[TransactionTotalsResponse](
		ctx,
		c.client,
		c.secret,
		path,
		c.baseURL,
	)
}
