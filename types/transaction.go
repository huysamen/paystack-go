package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Transaction represents a Paystack transaction with comprehensive field coverage
type Transaction struct {
	ID                 uint64              `json:"id"`
	Domain             string              `json:"domain"`
	Status             string              `json:"status"`
	Reference          string              `json:"reference"`
	Amount             int                 `json:"amount"`
	Message            *string             `json:"message"`
	GatewayResponse    string              `json:"gateway_response"`
	PaidAt             *data.MultiDateTime `json:"paid_at,omitempty"`
	CreatedAt          data.MultiDateTime  `json:"created_at"`
	Channel            enums.Channel       `json:"channel"`
	Currency           enums.Currency      `json:"currency"`
	IPAddress          string              `json:"ip_address"`
	Metadata           Metadata            `json:"metadata"`
	Log                *TransactionLog     `json:"log,omitempty"`
	Fees               int                 `json:"fees"`
	FeesSplit          *FeesSplit          `json:"fees_split,omitempty"`
	Customer           Customer            `json:"customer"`
	Authorization      Authorization       `json:"authorization"`
	Plan               *Plan               `json:"plan,omitempty"`
	Split              *TransactionSplit   `json:"split,omitempty"`
	Subaccount         *Subaccount         `json:"subaccount,omitempty"`
	OrderID            *string             `json:"order_id,omitempty"`
	RequestedAmount    int                 `json:"requested_amount"`
	Source             *TransactionSource  `json:"source,omitempty"`
	Connect            *ConnectData        `json:"connect,omitempty"`
	POSTransactionData *POSTransactionData `json:"pos_transaction_data,omitempty"`
}

// TransactionLog represents the transaction processing log
type TransactionLog struct {
	StartTime int                   `json:"start_time"`
	TimeSpent int                   `json:"time_spent"`
	Attempts  int                   `json:"attempts"`
	Errors    int                   `json:"errors"`
	Success   bool                  `json:"success"`
	Mobile    bool                  `json:"mobile"`
	Input     []any                 `json:"input"`
	History   []TransactionLogEntry `json:"history"`
}

// TransactionLogEntry represents an entry in the transaction log
type TransactionLogEntry struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Time    int    `json:"time"`
}

// FeesSplit represents the breakdown of transaction fees
type FeesSplit struct {
	Paystack    int              `json:"paystack"`
	Integration int              `json:"integration"`
	Subaccount  int              `json:"subaccount,omitempty"`
	Params      *FeesSplitParams `json:"params,omitempty"`
}

// FeesSplitParams represents the parameters used for fee calculation
type FeesSplitParams struct {
	Bearer            string `json:"bearer"`
	TransactionCharge string `json:"transaction_charge"`
	PercentageCharge  string `json:"percentage_charge"`
}

// TransactionSource represents the source of a transaction
type TransactionSource struct {
	Source     string  `json:"source"`
	Type       string  `json:"type"`
	Identifier *string `json:"identifier"`
	EntryPoint string  `json:"entry_point"`
}

// ConnectData represents connect-related transaction data
type ConnectData struct {
	ConnectAccountID *string `json:"connect_account_id,omitempty"`
	Provider         *string `json:"provider,omitempty"`
	ExternalID       *string `json:"external_id,omitempty"`
}

// POSTransactionData represents point-of-sale transaction data
type POSTransactionData struct {
	// Define POS-specific fields based on actual API responses
	TerminalID  *string `json:"terminal_id,omitempty"`
	ReceiptData *string `json:"receipt_data,omitempty"`
}

// TransactionInitializeRequest represents a request to initialize a transaction
type TransactionInitializeRequest struct {
	Email             string          `json:"email"`
	Amount            int             `json:"amount"`
	Currency          *enums.Currency `json:"currency,omitempty"`
	Reference         *string         `json:"reference,omitempty"`
	CallbackURL       *string         `json:"callback_url,omitempty"`
	Plan              *string         `json:"plan,omitempty"`
	InvoiceLimit      *int            `json:"invoice_limit,omitempty"`
	Metadata          Metadata        `json:"metadata,omitempty"`
	Channels          []enums.Channel `json:"channels,omitempty"`
	SplitCode         *string         `json:"split_code,omitempty"`
	SubaccountCode    *string         `json:"subaccount,omitempty"`
	TransactionCharge *int            `json:"transaction_charge,omitempty"`
	Bearer            *enums.Bearer   `json:"bearer,omitempty"`
}

// TransactionInitializeResponse represents the response from transaction initialization
type TransactionInitializeResponse struct {
	AuthorizationURL string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}
