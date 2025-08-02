package transfercontrol

import (
	"github.com/huysamen/paystack-go/types"
)

// Balance represents account balance information
type Balance struct {
	Currency string `json:"currency"`
	Balance  int64  `json:"balance"`
}

// BalanceLedger represents a balance ledger entry
type BalanceLedger struct {
	Integration      int    `json:"integration"`
	Domain           string `json:"domain"`
	Balance          int64  `json:"balance"`
	Currency         string `json:"currency"`
	Difference       int64  `json:"difference"`
	Reason           string `json:"reason"`
	ModelResponsible string `json:"model_responsible"`
	ModelRow         int    `json:"model_row"`
	ID               int    `json:"id"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
}

// ResendOTPRequest represents the request to resend OTP
type ResendOTPRequest struct {
	TransferCode string `json:"transfer_code"`
	Reason       string `json:"reason"`
}

// FinalizeDisableOTPRequest represents the request to finalize disabling OTP
type FinalizeDisableOTPRequest struct {
	OTP string `json:"otp"`
}

// CheckBalanceResponse represents the response from checking balance
type CheckBalanceResponse struct {
	Status  bool      `json:"status"`
	Message string    `json:"message"`
	Data    []Balance `json:"data"`
}

// FetchBalanceLedgerResponse represents the response from fetching balance ledger
type FetchBalanceLedgerResponse struct {
	Status  bool            `json:"status"`
	Message string          `json:"message"`
	Data    []BalanceLedger `json:"data"`
	Meta    *types.Meta     `json:"meta,omitempty"`
}

// ResendOTPResponse represents the response from resending OTP
type ResendOTPResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// DisableOTPResponse represents the response from disabling OTP
type DisableOTPResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// FinalizeDisableOTPResponse represents the response from finalizing disable OTP
type FinalizeDisableOTPResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// EnableOTPResponse represents the response from enabling OTP
type EnableOTPResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
