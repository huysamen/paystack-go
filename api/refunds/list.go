package refunds

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// List retrieves all refunds available on your integration
func (c *Client) List(ctx context.Context, req *RefundListRequest) (*RefundListResponse, error) {
	if req == nil {
		req = &RefundListRequest{}
	}

	if err := validateRefundListRequest(req); err != nil {
		return nil, err
	}

	url := c.baseURL + refundsBasePath
	return net.Get[[]Refund](ctx, c.client, c.secret, url)
}

// validateRefundListRequest validates the refund list request
func validateRefundListRequest(req *RefundListRequest) error {
	if req.PerPage != nil && (*req.PerPage < 1 || *req.PerPage > 100) {
		return fmt.Errorf("per_page must be between 1 and 100")
	}

	if req.Page != nil && *req.Page < 1 {
		return fmt.Errorf("page must be greater than 0")
	}

	if req.From != nil && req.To != nil && req.From.After(*req.To) {
		return fmt.Errorf("from date must be before to date")
	}

	if req.Transaction != nil && *req.Transaction == "" {
		return fmt.Errorf("transaction cannot be empty")
	}

	if req.Currency != nil && *req.Currency == "" {
		return fmt.Errorf("currency cannot be empty")
	}

	return nil
}
