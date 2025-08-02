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
type CheckBalanceResponse = types.Response[[]Balance]

// FetchBalanceLedgerResponse represents the response from fetching balance ledger
type FetchBalanceLedgerResponse = types.Response[[]BalanceLedger]

// ResendOTPResponse represents the response from resending OTP
type ResendOTPResponse = types.Response[any]

// DisableOTPResponse represents the response from disabling OTP
type DisableOTPResponse = types.Response[any]

// FinalizeDisableOTPResponse represents the response from finalizing disable OTP
type FinalizeDisableOTPResponse = types.Response[any]

// EnableOTPResponse represents the response from enabling OTP
type EnableOTPResponse = types.Response[any]
