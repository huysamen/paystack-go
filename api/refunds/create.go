package refunds

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Create initiates a refund on a transaction
func (c *Client) Create(ctx context.Context, req *RefundCreateRequest) (*RefundCreateResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("refund create request is required")
	}

	if err := validateRefundCreateRequest(req); err != nil {
		return nil, err
	}

	url := c.baseURL + refundsBasePath
	return net.Post[RefundCreateRequest, RefundCreateData](ctx, c.client, c.secret, url, req)
}

// validateRefundCreateRequest validates the refund create request
func validateRefundCreateRequest(req *RefundCreateRequest) error {
	if req.Transaction == "" {
		return fmt.Errorf("transaction is required")
	}

	if req.Amount != nil && *req.Amount <= 0 {
		return fmt.Errorf("amount must be greater than 0")
	}

	if req.Currency != nil && *req.Currency == "" {
		return fmt.Errorf("currency cannot be empty")
	}

	return nil
}
