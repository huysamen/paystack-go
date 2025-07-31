package integration

import "github.com/huysamen/paystack-go/types"

// FetchTimeoutResponse represents the response from fetching payment session timeout
type FetchTimeoutResponse = types.Response[FetchTimeoutData]

// FetchTimeoutData contains the payment session timeout information
type FetchTimeoutData struct {
	PaymentSessionTimeout int `json:"payment_session_timeout"`
}

// UpdateTimeoutRequest represents the request to update payment session timeout
type UpdateTimeoutRequest struct {
	Timeout int `json:"timeout"`
}

// UpdateTimeoutResponse represents the response from updating payment session timeout
type UpdateTimeoutResponse = types.Response[UpdateTimeoutData]

// UpdateTimeoutData contains the updated payment session timeout information
type UpdateTimeoutData struct {
	PaymentSessionTimeout int `json:"payment_session_timeout"`
}
