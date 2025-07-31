package applepay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// UnregisterDomain unregisters a top-level domain or subdomain previously used for Apple Pay integration
// Note: This uses a custom implementation because the Paystack API requires DELETE with body
func (c *Client) UnregisterDomain(ctx context.Context, req *UnregisterDomainRequest) (*UnregisterDomainResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.DomainName == "" {
		return nil, fmt.Errorf("domainName is required")
	}

	// Construct the full URL
	endpoint := applePayBasePath + "/domain"
	baseURL := c.baseURL
	if baseURL == "" {
		baseURL = "https://api.paystack.co"
	}
	fullURL := baseURL + endpoint

	// Marshal the request body
	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create the HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, "DELETE", fullURL, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Authorization", "Bearer "+c.secret)
	httpReq.Header.Set("Content-Type", "application/json")

	// Execute the request
	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Parse the response
	var apiResp types.Response[UnregisterDomainResponse]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Check for API errors
	if !apiResp.Status {
		return nil, &net.PaystackError{
			StatusCode: resp.StatusCode,
			Message:    apiResp.Message,
			Status:     apiResp.Status,
		}
	}

	return &apiResp.Data, nil
}
