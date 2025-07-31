package products

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
)

// List retrieves products available on your integration
func (c *Client) List(ctx context.Context, req *ListProductsRequest) (*ListProductsResponse, error) {
	if req == nil {
		req = &ListProductsRequest{}
	}

	// Build query parameters
	params := url.Values{}
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

	path := productsBasePath
	if len(params) > 0 {
		path = fmt.Sprintf("%s?%s", productsBasePath, params.Encode())
	}

	resp, err := net.Get[ListProductsResponse](
		ctx,
		c.client,
		c.secret,
		path,
		c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
