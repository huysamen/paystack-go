package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Transaction represents a Paystack transaction with comprehensive field coverage
type Transaction struct {
	ID                 data.Uint           `json:"id"`
	Domain             data.String         `json:"domain"`
	Status             data.String         `json:"status"`
	Reference          data.String         `json:"reference"`
	Amount             data.Int            `json:"amount"`
	Message            data.NullString     `json:"message"`
	GatewayResponse    data.String         `json:"gateway_response"`
	PaidAt             data.NullTime       `json:"paid_at,omitempty"`
	CreatedAt          data.Time           `json:"created_at"`
	Channel            enums.Channel       `json:"channel"`
	Currency           enums.Currency      `json:"currency"`
	IPAddress          data.String         `json:"ip_address"`
	Metadata           Metadata            `json:"metadata"`
	Log                *TransactionLog     `json:"log,omitempty"`
	Fees               data.Int            `json:"fees"`
	FeesSplit          *FeesSplit          `json:"fees_split,omitempty"`
	Customer           Customer            `json:"customer"`
	Authorization      Authorization       `json:"authorization"`
	Plan               *Plan               `json:"plan,omitempty"`
	Split              *TransactionSplit   `json:"split,omitempty"`
	Subaccount         *Subaccount         `json:"subaccount,omitempty"`
	OrderID            data.NullString     `json:"order_id,omitempty"`
	RequestedAmount    data.Int            `json:"requested_amount"`
	Source             *TransactionSource  `json:"source,omitempty"`
	Connect            *ConnectData        `json:"connect,omitempty"`
	POSTransactionData *POSTransactionData `json:"pos_transaction_data,omitempty"`
}

// TransactionLog represents the transaction processing log
type TransactionLog struct {
	StartTime data.Int              `json:"start_time"`
	TimeSpent data.Int              `json:"time_spent"`
	Attempts  data.Int              `json:"attempts"`
	Errors    data.Int              `json:"errors"`
	Success   data.Bool             `json:"success"`
	Mobile    data.Bool             `json:"mobile"`
	Input     any                   `json:"input"` // todo: this comes through as multiple data types
	History   []TransactionLogEntry `json:"history"`
}

// TransactionLogEntry represents an entry in the transaction log
type TransactionLogEntry struct {
	Type    data.String `json:"type"`
	Message data.String `json:"message"`
	Time    data.Int    `json:"time"`
}

// FeesSplit represents the breakdown of transaction fees
type FeesSplit struct {
	Paystack    data.Int         `json:"paystack"`
	Integration data.Int         `json:"integration"`
	Subaccount  data.Int         `json:"subaccount,omitempty"`
	Params      *FeesSplitParams `json:"params,omitempty"`
}

// FeesSplitParams represents the parameters used for fee calculation
type FeesSplitParams struct {
	Bearer            data.String `json:"bearer"`
	TransactionCharge data.String `json:"transaction_charge"`
	PercentageCharge  data.String `json:"percentage_charge"`
}

// TransactionSource represents the source of a transaction
type TransactionSource struct {
	Source     data.String     `json:"source"`
	Type       data.String     `json:"type"`
	Identifier data.NullString `json:"identifier"`
	EntryPoint data.String     `json:"entry_point"`
}

// ConnectData represents connect-related transaction data
type ConnectData struct {
	ConnectAccountID data.NullString `json:"connect_account_id,omitempty"`
	Provider         data.NullString `json:"provider,omitempty"`
	ExternalID       data.NullString `json:"external_id,omitempty"`
}

// POSTransactionData represents point-of-sale transaction data
type POSTransactionData struct {
	// Define POS-specific fields based on actual API responses
	TerminalID  data.NullString `json:"terminal_id,omitempty"`
	ReceiptData data.NullString `json:"receipt_data,omitempty"`
}

// TransactionInitializeRequest represents a request to initialize a transaction
type TransactionInitializeRequest struct {
	Email             data.String     `json:"email"`
	Amount            data.Int        `json:"amount"`
	Currency          *enums.Currency `json:"currency,omitempty"`
	Reference         data.NullString `json:"reference,omitempty"`
	CallbackURL       data.NullString `json:"callback_url,omitempty"`
	Plan              data.NullString `json:"plan,omitempty"`
	InvoiceLimit      data.NullInt    `json:"invoice_limit,omitempty"`
	Metadata          Metadata        `json:"metadata,omitempty"`
	Channels          []enums.Channel `json:"channels,omitempty"`
	SplitCode         data.NullString `json:"split_code,omitempty"`
	SubaccountCode    data.NullString `json:"subaccount,omitempty"`
	TransactionCharge data.NullInt    `json:"transaction_charge,omitempty"`
	Bearer            *enums.Bearer   `json:"bearer,omitempty"`
}

// TransactionInitializeResponse represents the response from transaction initialization
type TransactionInitializeResponse struct {
	AuthorizationURL data.String `json:"authorization_url"`
	AccessCode       data.String `json:"access_code"`
	Reference        data.String `json:"reference"`
}
