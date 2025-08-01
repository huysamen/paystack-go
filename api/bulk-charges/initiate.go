package bulkcharges

import (
	"context"
	"errors"
	"strings"

	"github.com/huysamen/paystack-go/net"
)

// ValidateInitiateBulkChargeRequest validates the initiate bulk charge request
func ValidateInitiateBulkChargeRequest(req InitiateBulkChargeRequest) error {
	if len(req) == 0 {
		return errors.New("at least one bulk charge item is required")
	}

	if len(req) > 200 {
		return errors.New("maximum of 200 charges allowed per bulk charge request")
	}

	for i, item := range req {
		if strings.TrimSpace(item.Authorization) == "" {
			return errors.New("authorization is required for bulk charge item at index " +
				string(rune(i)))
		}

		if item.Amount <= 0 {
			return errors.New("amount must be greater than 0 for bulk charge item at index " +
				string(rune(i)))
		}

		if strings.TrimSpace(item.Reference) == "" {
			return errors.New("reference is required for bulk charge item at index " +
				string(rune(i)))
		}
	}

	return nil
}

// Initiate sends an array of objects with authorization codes and amounts for batch processing
func (c *Client) Initiate(ctx context.Context, req InitiateBulkChargeRequest) (*InitiateBulkChargeResponse, error) {

	resp, err := net.Post[InitiateBulkChargeRequest, InitiateBulkChargeResponse](
		ctx, c.client, c.secret, bulkChargesBasePath, &req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
