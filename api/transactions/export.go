package transactions

import (
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
	query := ""

	queryAppend(query, "perPage", r.PerPage)
	queryAppend(query, "page", r.Page)
	queryAppend(query, "from", r.From)
	queryAppend(query, "to", r.To)
	queryAppend(query, "customer", r.Customer)
	queryAppend(query, "status", r.Status)
	queryAppend(query, "currency", r.Currency.String())
	queryAppend(query, "amount", r.Amount)
	queryAppend(query, "settled", r.Settled)
	queryAppend(query, "settlement", r.Settlement)
	queryAppend(query, "payment_page", r.PaymentPage)

	return query
}

type TransactionExportResponse struct {
	Path      string         `json:"path"`
	ExpiresAt types.DateTime `json:"expiresAt"`
}

func (c *Client) Export(req *TransactionExportRequest) (*types.Response[TransactionExportResponse], error) {
	query := ""

	if req != nil {
		query = "?" + req.toQuery()

		if query == "?" {
			query = ""
		}
	}

	return net.Get[TransactionExportResponse](
		c.client,
		c.secret,
		fmt.Sprintf("%s%s%s", transactionBasePath, transactionExportPath, query),
	)
}

func queryAppend[T *string | string | *int | *uint64 | *bool | *time.Time](query, name string, value T) string {
	switch v := any(value).(type) {
	case *string:
		if v != nil {
			return fmt.Sprintf("%s&%s=%s", query, url.QueryEscape(name), url.QueryEscape(*v))
		}
	case string:
		return fmt.Sprintf("%s&%s=%s", query, url.QueryEscape(name), url.QueryEscape(v))
	case *int:
		if v != nil {
			return fmt.Sprintf("%s&%s=%s", query, url.QueryEscape(name), url.QueryEscape(strconv.Itoa(*v)))
		}
	case *uint64:
		if v != nil {
			return fmt.Sprintf("%s&%s=%s", query, url.QueryEscape(name), url.QueryEscape(strconv.FormatUint(*v, 10)))
		}
	case *bool:
		if v != nil {
			b := "false"

			if *v {
				b = "true"
			}

			return fmt.Sprintf("%s&%s=%s", query, url.QueryEscape(name), url.QueryEscape(b))
		}
	case *time.Time:
		if v != nil {
			t := *v

			return fmt.Sprintf("%s&%s=%s", query, url.QueryEscape(name), url.QueryEscape(t.Format("2006-01-02T15:04:05.999Z")))
		}
	}

	return query
}
