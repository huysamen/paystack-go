package transactionsplits

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// List retrieves a list of transaction splits
func (c *Client) List(ctx context.Context, builder *TransactionSplitListRequestBuilder) (*types.Response[[]TransactionSplit], error) {
	params := url.Values{}

	if builder != nil {
		req := builder.Build()
		if req.Name != nil {
			params.Set("name", *req.Name)
		}
		if req.Active != nil {
			params.Set("active", strconv.FormatBool(*req.Active))
		}
		if req.SortBy != nil {
			params.Set("sort_by", *req.SortBy)
		}
		if req.PerPage != nil {
			params.Set("perPage", strconv.Itoa(*req.PerPage))
		}
		if req.Page != nil {
			params.Set("page", strconv.Itoa(*req.Page))
		}
		if req.From != nil {
			params.Set("from", req.From.Format(time.RFC3339))
		}
		if req.To != nil {
			params.Set("to", req.To.Format(time.RFC3339))
		}
	}

	endpoint := transactionSplitBasePath
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	return net.Get[[]TransactionSplit](ctx, c.client, c.secret, endpoint, c.baseURL)
}
