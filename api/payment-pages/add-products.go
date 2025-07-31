package paymentpages

import (
	"context"
	"strconv"

	"github.com/huysamen/paystack-go/net"
)

// AddProducts adds products to a payment page
func (c *Client) AddProducts(ctx context.Context, pageID int, req *AddProductsToPageRequest) (*PaymentPage, error) {
	if err := ValidatePageID(pageID); err != nil {
		return nil, err
	}

	if err := ValidateAddProductsToPageRequest(req); err != nil {
		return nil, err
	}

	resp, err := net.Post[AddProductsToPageRequest, PaymentPage](
		ctx, c.client, c.secret, paymentPagesBasePath+"/"+strconv.Itoa(pageID)+"/product", req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
