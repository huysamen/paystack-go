package settlements

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
)

// ListTransactions retrieves transactions for a specific settlement using a builder (fluent interface)
func (c *Client) ListTransactions(ctx context.Context, settlementID string, builder *SettlementTransactionListRequestBuilder) (*SettlementTransactionListResponse, error) {
	if settlementID == "" {
		return nil, fmt.Errorf("settlement_id is required")
	}

	req := builder.Build()
	params := url.Values{}

	if req != nil {
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

	endpoint := fmt.Sprintf("%s/%s/transactions", settlementBasePath, settlementID)
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	resp, err := net.Get[SettlementTransactionListResponse](ctx, c.client, c.secret, endpoint, c.baseURL)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
