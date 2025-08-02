package transferrecipients

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// List retrieves a list of transfer recipients
func (c *Client) List(ctx context.Context, builder *TransferRecipientListRequestBuilder) (*types.Response[[]TransferRecipient], error) {
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

	queryParams := ""
	if len(params) > 0 {
		queryParams = params.Encode()
	}

	return net.Get[[]TransferRecipient](ctx, c.Client, c.Secret, basePath, queryParams, c.BaseURL)
}
