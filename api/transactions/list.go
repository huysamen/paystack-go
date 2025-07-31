package transactions

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
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

func (c *Client) List(ctx context.Context, req *TransactionListRequest) (*types.Response[TransactionListResponse], error) {
	path := transactionBasePath

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[TransactionListResponse](
		ctx,
		c.client,
		c.secret,
		path,
		c.baseURL,
	)
}

// ListPage is a convenience method for simple pagination (deprecated, use List with TransactionListRequest)
func (c *Client) ListPage(ctx context.Context, perPage, page int) (*types.Response[TransactionListResponse], error) {
	req := &TransactionListRequest{
		PerPage: &perPage,
		Page:    &page,
	}
	return c.List(ctx, req)
}
