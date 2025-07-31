package webhook

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Validator provides utilities for validating Paystack webhooks
type Validator struct {
	secretKey string
}

// NewValidator creates a new webhook validator with your secret key
func NewValidator(secretKey string) *Validator {
	return &Validator{
		secretKey: secretKey,
	}
}

// ValidateSignature validates the webhook signature to ensure it came from Paystack
func (v *Validator) ValidateSignature(payload []byte, signature string) bool {
	mac := hmac.New(sha512.New, []byte(v.secretKey))
	mac.Write(payload)
	expectedSignature := hex.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(signature), []byte(expectedSignature))
}

// ValidateRequest validates an incoming webhook request
func (v *Validator) ValidateRequest(r *http.Request) (*Event, error) {
	signature := r.Header.Get("x-paystack-signature")
	if signature == "" {
		return nil, fmt.Errorf("missing x-paystack-signature header")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}

	if !v.ValidateSignature(body, signature) {
		return nil, fmt.Errorf("invalid webhook signature")
	}

	var event Event
	if err := json.Unmarshal(body, &event); err != nil {
		return nil, fmt.Errorf("failed to parse webhook payload: %w", err)
	}

	return &event, nil
}
