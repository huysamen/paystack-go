package types

import (
	"github.com/huysamen/paystack-go/enums"
)

// Transaction represents a Paystack transaction
type Transaction struct {
	ID                 uint64              `json:"id"`
	Domain             string              `json:"domain"`
	Status             string              `json:"status"`
	Reference          string              `json:"reference"`
	Amount             int                 `json:"amount"`
	Message            string              `json:"message"`
	GatewayResponse    string              `json:"gateway_response"`
	PaidAt             *DateTime           `json:"paid_at,omitempty"`
	CreatedAt          DateTime            `json:"created_at"`
	Channel            enums.Channel       `json:"channel"`
	Currency           enums.Currency      `json:"currency"`
	IPAddress          string              `json:"ip_address"`
	Metadata           Metadata            `json:"metadata"`
	Log                Log                 `json:"log"`
	Fees               int                 `json:"fees"`
	FeesSplit          *FeesSplit          `json:"fees_split,omitempty"`
	Customer           Customer            `json:"customer"`
	Authorization      Authorization       `json:"authorization"`
	Plan               *Plan               `json:"plan,omitempty"`
	Split              *TransactionSplit   `json:"split,omitempty"`
	Subaccount         *Subaccount         `json:"subaccount,omitempty"`
	OrderID            *string             `json:"order_id,omitempty"`
	RequestedAmount    int                 `json:"requested_amount"`
	Source             Source              `json:"source"`
	Connect            *ConnectData        `json:"connect,omitempty"`
	POSTransactionData *POSTransactionData `json:"pos_transaction_data,omitempty"`
}

// FeesSplit represents the breakdown of transaction fees
type FeesSplit struct {
	Paystack    int `json:"paystack"`
	Integration int `json:"integration"`
	Subaccount  int `json:"subaccount,omitempty"`
	Params      struct {
		Bearer            string `json:"bearer"`
		TransactionCharge string `json:"transaction_charge"`
		PercentageCharge  string `json:"percentage_charge"`
	} `json:"params,omitempty"`
}

// ConnectData represents connect-related transaction data
type ConnectData struct {
	ConnectAccountID *string `json:"connect_account_id,omitempty"`
	Provider         *string `json:"provider,omitempty"`
	ExternalID       *string `json:"external_id,omitempty"`
}

// POSTransactionData represents point-of-sale transaction data
type POSTransactionData struct {
	TerminalID  *string `json:"terminal_id,omitempty"`
	Location    *string `json:"location,omitempty"`
	MerchantID  *string `json:"merchant_id,omitempty"`
	ReceiptData *string `json:"receipt_data,omitempty"`
}

// LogHistoryEntry represents a single entry in transaction log history
type LogHistoryEntry struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Time    int    `json:"time"`
}

// LogInput represents the input data for a transaction log
type LogInput struct {
	Email       *string        `json:"email,omitempty"`
	Amount      *int           `json:"amount,omitempty"`
	Currency    *string        `json:"currency,omitempty"`
	Reference   *string        `json:"reference,omitempty"`
	CallbackURL *string        `json:"callback_url,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
	Channels    []string       `json:"channels,omitempty"`
	Custom      map[string]any `json:"custom,omitempty"`
}

// Log represents transaction log information
type Log struct {
	StartTime int               `json:"start_time"`
	TimeSpent int               `json:"time_spent"`
	Attempts  int               `json:"attempts"`
	Errors    int               `json:"errors"`
	Success   bool              `json:"success"`
	Mobile    bool              `json:"mobile"`
	Input     *LogInput         `json:"input,omitempty"`
	History   []LogHistoryEntry `json:"history"`
}

// Source represents the source of a transaction
type Source struct {
	Type       string `json:"type"`
	Source     string `json:"source"`
	EntryPoint string `json:"entry_point"`
	Identifier string `json:"identifier"`
}
