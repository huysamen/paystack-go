package bulkcharges

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
)

// List retrieves all bulk charge batches created by the integration using a builder
func (c *Client) List(ctx context.Context, builder *ListBulkChargeBatchesRequestBuilder) (*ListBulkChargeBatchesResponse, error) {
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
			params.Set("from", *req.From)
		}
		if req.To != nil {
			params.Set("to", *req.To)
		}
	}

	path := bulkChargesBasePath
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	resp, err := net.Get[ListBulkChargeBatchesResponse](ctx, c.client, c.secret, path, c.baseURL)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
