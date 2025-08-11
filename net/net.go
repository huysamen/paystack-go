package net

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
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

// headerRoundTripper injects default headers and optionally appends a user-agent suffix.
type headerRoundTripper struct {
	base            http.RoundTripper
	headers         map[string]string
	userAgentSuffix string
}

// NewHeaderRoundTripper wraps a base RoundTripper to inject headers. If base is nil,
// http.DefaultTransport is used. Headers provided here will override any same-named
// headers set earlier on the request. If userAgentSuffix is non-empty and no explicit
// User-Agent header override is provided in headers, it will be appended to the
// existing User-Agent header value.
func NewHeaderRoundTripper(base http.RoundTripper, headers map[string]string, userAgentSuffix string) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	return &headerRoundTripper{base: base, headers: headers, userAgentSuffix: userAgentSuffix}
}

func (rt *headerRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clone the request to avoid mutating the caller's instance
	r := req.Clone(req.Context())

	// Apply explicit header overrides first
	if len(rt.headers) > 0 {
		for k, v := range rt.headers {
			r.Header.Set(k, v)
		}
	}

	// Append UA suffix only if not explicitly overridden via headers
	if rt.userAgentSuffix != "" {
		if _, ok := rt.headers["User-Agent"]; !ok {
			ua := r.Header.Get("User-Agent")
			if ua == "" {
				ua = fmt.Sprintf("paystack-go/%s (+github.com/huysamen/paystack-go)", runtime.Version())
			}
			r.Header.Set("User-Agent", strings.TrimSpace(ua+" "+rt.userAgentSuffix))
		}
	}

	return rt.base.RoundTrip(r)
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

// DeleteWithBody makes a DELETE request with a request body
func DeleteWithBody[I any, O any](ctx context.Context, client *http.Client, secret, path string, payload *I, baseURL ...string) (*types.Response[O], error) {
	url := getBaseURL(baseURL...)
	body, err := doReq(ctx, client, http.MethodDelete, secret, url+path, payload)
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

// doReq performs the HTTP request. If the client's Transport is a header-injecting
// RoundTripper, it can add default headers; otherwise we add minimal defaults here.
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
	req.Header.Add("Accept", "application/json")
	// Include a helpful User-Agent for diagnostics
	req.Header.Add("User-Agent", fmt.Sprintf("paystack-go/%s (+github.com/huysamen/paystack-go)", runtime.Version()))

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
