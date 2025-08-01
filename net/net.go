package net

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/huysamen/paystack-go/types"
)

const apiURL = "https://api.paystack.co"

func getBaseURL(baseURL ...string) string {
	if len(baseURL) > 0 && baseURL[0] != "" {
		return baseURL[0]
	}
	return apiURL
}

// getHTTPErrorMessage generates a meaningful error message from HTTP status and response body
func getHTTPErrorMessage(statusCode int, body []byte) string {
	// Try to extract a meaningful message from the response body
	if len(body) > 0 {
		// Try to parse as JSON and extract message
		var response map[string]any
		if err := json.Unmarshal(body, &response); err == nil {
			if msg, ok := response["message"].(string); ok && msg != "" {
				return msg
			}
			if msg, ok := response["error"].(string); ok && msg != "" {
				return msg
			}
		}
		// If not JSON or no message field, use raw body if it's short and looks like text
		if len(body) < 200 && !json.Valid(body) {
			return strings.TrimSpace(string(body))
		}
	}

	// Fallback to standard HTTP status text
	switch statusCode {
	case http.StatusBadRequest:
		return "Bad request: The request was invalid or malformed"
	case http.StatusUnauthorized:
		return "Unauthorized: Invalid or missing API key"
	case http.StatusForbidden:
		return "Forbidden: Access denied or insufficient permissions"
	case http.StatusNotFound:
		return "Not found: The requested resource does not exist"
	case http.StatusUnprocessableEntity:
		return "Unprocessable entity: Validation failed"
	case http.StatusTooManyRequests:
		return "Rate limited: Too many requests, please try again later"
	case http.StatusConflict:
		return "Conflict: The request conflicts with the current state"
	case http.StatusPreconditionFailed:
		return "Precondition failed: Required conditions were not met"
	default:
		return fmt.Sprintf("HTTP %d: %s", statusCode, http.StatusText(statusCode))
	}
}

// Get makes a GET request with context support
func Get[O any](ctx context.Context, client *http.Client, secret, path string, baseURL ...string) (*types.Response[O], error) {
	url := getBaseURL(baseURL...)
	body, err := doReq(ctx, client, http.MethodGet, secret, url+path, nil)
	if err != nil {
		return nil, err
	}

	rsp := new(types.Response[O])

	if len(body) > 0 {
		err = json.Unmarshal(body, rsp)
		if err != nil {
			return nil, err
		}
	}

	return rsp, nil
}

// Post makes a POST request with context support
func Post[I any, O any](ctx context.Context, client *http.Client, secret, path string, payload *I, baseURL ...string) (*types.Response[O], error) {
	url := getBaseURL(baseURL...)
	return putOrPost[I, O](ctx, client, http.MethodPost, secret, url+path, payload)
}

// Put makes a PUT request with context support
func Put[I any, O any](ctx context.Context, client *http.Client, secret, path string, payload *I, baseURL ...string) (*types.Response[O], error) {
	url := getBaseURL(baseURL...)
	return putOrPost[I, O](ctx, client, http.MethodPut, secret, url+path, payload)
}

// Delete makes a DELETE request with context support
func Delete[O any](ctx context.Context, client *http.Client, secret, path string, baseURL ...string) (*types.Response[O], error) {
	url := getBaseURL(baseURL...)
	body, err := doReq(ctx, client, http.MethodDelete, secret, url+path, nil)
	if err != nil {
		return nil, err
	}

	rsp := new(types.Response[O])

	if len(body) > 0 {
		err = json.Unmarshal(body, rsp)
		if err != nil {
			return nil, err
		}
	}

	return rsp, nil
}

func putOrPost[I any, O any](ctx context.Context, client *http.Client, method, secret, fullURL string, payload *I) (*types.Response[O], error) {
	body, err := doReq(ctx, client, method, secret, fullURL, payload)
	if err != nil {
		return nil, err
	}

	rsp := new(types.Response[O])

	if len(body) > 0 {
		err = json.Unmarshal(body, rsp)
		if err != nil {
			return nil, err
		}
	}

	return rsp, nil
}

func doReq(ctx context.Context, client *http.Client, method, secret, fullURL string, data any) ([]byte, error) {
	var req *http.Request
	var err error

	if data != nil {
		d, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequestWithContext(ctx, method, fullURL, bytes.NewBuffer(d))
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequestWithContext(ctx, method, fullURL, nil)
		if err != nil {
			return nil, err
		}
	}

	req.Header.Add("Authorization", "Bearer "+secret)

	if data != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rsp.Body.Close() }()

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	// For server errors (5xx), return actual Go errors since these are system issues
	if rsp.StatusCode >= 500 {
		return nil, fmt.Errorf("paystack server error (HTTP %d): %s", rsp.StatusCode, getHTTPErrorMessage(rsp.StatusCode, body))
	}

	// For all other status codes (including 4xx client errors), return the body
	// The response will be parsed by the calling function and API errors will be
	// represented as Response objects with status: false
	return body, nil
}
