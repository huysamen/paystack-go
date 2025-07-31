package transactions

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
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

func (c *Client) Export(ctx context.Context, req *TransactionExportRequest) (*types.Response[TransactionExportResponse], error) {
	path := fmt.Sprintf("%s%s", transactionBasePath, transactionExportPath)

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[TransactionExportResponse](
		ctx,
		c.client,
		c.secret,
		path,
		c.baseURL,
	)
}
