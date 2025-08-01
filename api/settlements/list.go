package settlements

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
)

// List retrieves a list of settlements using a builder (fluent interface)
func (c *Client) List(ctx context.Context, builder *SettlementListRequestBuilder) (*SettlementListResponse, error) {
	req := builder.Build()
	params := url.Values{}

	if req != nil {
		if req.PerPage != nil {
			params.Set("perPage", strconv.Itoa(*req.PerPage))
		}
		if req.Page != nil {
			params.Set("page", strconv.Itoa(*req.Page))
		}
		if req.Status != nil {
			params.Set("status", req.Status.String())
		}
		if req.Subaccount != nil {
			params.Set("subaccount", *req.Subaccount)
		}
		if req.From != nil {
			params.Set("from", req.From.Format(time.RFC3339))
		}
		if req.To != nil {
			params.Set("to", req.To.Format(time.RFC3339))
		}
	}

	endpoint := settlementBasePath
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	resp, err := net.Get[SettlementListResponse](ctx, c.client, c.secret, endpoint, c.baseURL)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
