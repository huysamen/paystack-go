package bulkcharges

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
)

// ValidateChargeStatus validates the charge status parameter
func ValidateChargeStatus(status string) error {
	if status == "" {
		return nil // optional parameter
	}

	validStatuses := []string{"pending", "success", "failed"}
	for _, validStatus := range validStatuses {
		if status == validStatus {
			return nil
		}
	}

	return nil // allow any status for flexibility
}

// FetchChargesInBatch retrieves the charges associated with a specified batch code
func (c *Client) FetchChargesInBatch(ctx context.Context, idOrCode string, req *FetchChargesInBatchRequest) (*FetchChargesInBatchResponse, error) {

	params := url.Values{}

	if req != nil {
		if req.Status != nil {
			params.Set("status", *req.Status)
		}
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

	path := bulkChargesBasePath + "/" + idOrCode + "/charges"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	resp, err := net.Get[FetchChargesInBatchResponse](ctx, c.client, c.secret, path, c.baseURL)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
