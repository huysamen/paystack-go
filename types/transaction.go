package types

import "time"

// Transaction represents a Paystack transaction
type Transaction struct {
	ID                 uint64        `json:"id"`
	Domain             string        `json:"domain"`
	Status             string        `json:"status"`
	Reference          string        `json:"reference"`
	Amount             int           `json:"amount"`
	Message            string        `json:"message"`
	GatewayResponse    string        `json:"gateway_response"`
	PaidAt             time.Time     `json:"paid_at"`
	CreatedAt          time.Time     `json:"created_at"`
	Channel            Channel       `json:"channel"`
	Currency           Currency      `json:"currency"`
	IPAddress          string        `json:"ip_address"`
	Metadata           Metadata      `json:"metadata"`
	Log                Log           `json:"log"`
	Fees               int           `json:"fees"`
	FeesSplit          any           `json:"fees_split"`
	Customer           Customer      `json:"customer"`
	Authorization      Authorization `json:"authorization"`
	Plan               Plan          `json:"plan"`
	Split              Split         `json:"split"`
	Subaccount         Subaccount    `json:"subsccount"`
	OrderID            any           `json:"order_id"`
	RequestedAmount    int           `json:"requested_amount"`
	Source             Source        `json:"source"`
	Connect            any           `json:"connect"`
	POSTransactionData any           `json:"pos_transaction_data"`
}

// Log represents transaction log information
type Log struct {
	StartTime int  `json:"start_time"`
	TimeSpent int  `json:"time_spent"`
	Attempts  int  `json:"attempts"`
	Errors    int  `json:"errors"`
	Success   bool `json:"success"`
	Mobile    bool `json:"mobile"`
	Input     any  `json:"input"`
	History   []struct {
		Type    string `json:"type"`
		Message string `json:"message"`
		Time    int    `json:"time"`
	} `json:"history"`
}

// Source represents the source of a transaction
type Source struct {
	Type       string `json:"type"`
	Source     string `json:"source"`
	EntryPoint string `json:"entry_point"`
	Identifier string `json:"identifier"`
}
