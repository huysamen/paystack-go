package net

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/huysamen/paystack-go/types"
)

const apiURL = "https://api.paystack.co"

// PaystackError represents an error response from the Paystack API
type PaystackError struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
	Status     bool   `json:"status"`
	Code       string `json:"code,omitempty"`
	Type       string `json:"type,omitempty"`
	Cause      error  `json:"-"` // Underlying cause of the error
}

func (e *PaystackError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("paystack api error (status %d): %s, cause: %v", e.StatusCode, e.Message, e.Cause)
	}
	return fmt.Sprintf("paystack api error (status %d): %s", e.StatusCode, e.Message)
}

// Unwrap returns the underlying cause error for error wrapping support
func (e *PaystackError) Unwrap() error {
	return e.Cause
}

// Is allows error comparison using errors.Is
func (e *PaystackError) Is(target error) bool {
	if t, ok := target.(*PaystackError); ok {
		return e.StatusCode == t.StatusCode && e.Code == t.Code
	}
	return false
}

func getBaseURL(baseURL ...string) string {
	if len(baseURL) > 0 && baseURL[0] != "" {
		return baseURL[0]
	}
	return apiURL
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

	switch rsp.StatusCode {
	case http.StatusOK, http.StatusCreated:
		return body, nil
	case http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden,
		http.StatusNotFound, http.StatusUnprocessableEntity:
		// Try to parse error response
		var paystackErr PaystackError
		if err := json.Unmarshal(body, &paystackErr); err == nil {
			paystackErr.StatusCode = rsp.StatusCode
			return nil, &paystackErr
		}
		// Fallback to generic error
		return nil, &PaystackError{
			StatusCode: rsp.StatusCode,
			Message:    fmt.Sprintf("HTTP %d: %s", rsp.StatusCode, http.StatusText(rsp.StatusCode)),
		}
	default:
		return nil, &PaystackError{
			StatusCode: rsp.StatusCode,
			Message:    fmt.Sprintf("unexpected status code: %d", rsp.StatusCode),
		}
	}
}
