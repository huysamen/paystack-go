package bulkcharges

import (
	"context"
	"errors"
	"strings"

	"github.com/huysamen/paystack-go/net"
)

// ValidateBatchCode validates the batch code parameter
func ValidateBatchCode(batchCode string) error {
	if strings.TrimSpace(batchCode) == "" {
		return errors.New("batch code is required")
	}
	return nil
}

// Pause pauses processing of a bulk charge batch
func (c *Client) Pause(ctx context.Context, batchCode string) (*PauseBulkChargeBatchResponse, error) {
	if err := ValidateBatchCode(batchCode); err != nil {
		return nil, err
	}

	resp, err := net.Get[PauseBulkChargeBatchResponse](
		ctx, c.client, c.secret, bulkChargesBasePath+"/pause/"+batchCode, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
