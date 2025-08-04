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

type Validator struct {
	secretKey string
}

func NewValidator(secretKey string) *Validator {
	return &Validator{
		secretKey: secretKey,
	}
}

func (v *Validator) ValidateSignature(payload []byte, signature string) bool {
	mac := hmac.New(sha512.New, []byte(v.secretKey))
	mac.Write(payload)
	expectedSignature := hex.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(signature), []byte(expectedSignature))
}

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
