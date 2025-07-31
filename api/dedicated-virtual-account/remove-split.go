package dedicatedvirtualaccount

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/huysamen/paystack-go/types"
)

// RemoveSplit removes split payment setup from a dedicated virtual account
// Note: This uses a custom implementation because the Paystack API requires DELETE with body
func (c *Client) RemoveSplit(ctx context.Context, req *RemoveSplitFromDedicatedAccountRequest) (*types.Response[DedicatedVirtualAccount], error) {
	if err := validateRemoveSplitRequest(req); err != nil {
		return nil, err
	}

	// Construct the full URL
	endpoint := dedicatedVirtualAccountBasePath + "/split"
	baseURL := c.baseURL
	if baseURL == "" {
		baseURL = "https://api.paystack.co"
	}
	fullURL := baseURL + endpoint

	// Marshal the request body
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// Create the HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodDelete, fullURL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	// Set headers
	httpReq.Header.Set("Authorization", "Bearer "+c.secret)
	httpReq.Header.Set("Content-Type", "application/json")

	// Execute the request
	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	var buf bytes.Buffer
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	body := buf.Bytes()

	// Check for HTTP errors
	if resp.StatusCode >= 400 {
		var paystackErr map[string]interface{}
		if err := json.Unmarshal(body, &paystackErr); err == nil {
			if msg, ok := paystackErr["message"].(string); ok {
				return nil, fmt.Errorf("paystack api error (status %d): %s", resp.StatusCode, msg)
			}
		}
		return nil, fmt.Errorf("paystack api error (status %d)", resp.StatusCode)
	}

	// Parse response
	result := new(types.Response[DedicatedVirtualAccount])
	if len(body) > 0 {
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
