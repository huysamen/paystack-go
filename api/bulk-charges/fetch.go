package bulkcharges

import (
	"context"
	"errors"
	"strings"

	"github.com/huysamen/paystack-go/net"
)

// ValidateIDOrCode validates the ID or code parameter
func ValidateIDOrCode(idOrCode string) error {
	if strings.TrimSpace(idOrCode) == "" {
		return errors.New("id or code is required")
	}
	return nil
}

// Fetch retrieves a specific bulk charge batch by ID or batch code
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*FetchBulkChargeBatchResponse, error) {
	if err := ValidateIDOrCode(idOrCode); err != nil {
		return nil, err
	}

	resp, err := net.Get[FetchBulkChargeBatchResponse](
		ctx, c.client, c.secret, bulkChargesBasePath+"/"+idOrCode, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
