package types

import "time"

// TransactionSplitType represents the type of transaction split
type TransactionSplitType string

const (
	TransactionSplitTypePercentage TransactionSplitType = "percentage" // Percentage-based split
	TransactionSplitTypeFlat       TransactionSplitType = "flat"       // Fixed amount split
)

// String returns the string representation of TransactionSplitType
func (s TransactionSplitType) String() string {
	return string(s)
}

// TransactionSplitBearerType represents who bears the charges for the split
type TransactionSplitBearerType string

const (
	TransactionSplitBearerTypeSubaccount      TransactionSplitBearerType = "subaccount"       // A specific subaccount bears the charges
	TransactionSplitBearerTypeAccount         TransactionSplitBearerType = "account"          // Main account bears the charges
	TransactionSplitBearerTypeAllProportional TransactionSplitBearerType = "all-proportional" // All parties bear charges proportionally
	TransactionSplitBearerTypeAll             TransactionSplitBearerType = "all"              // All parties bear charges equally
)

// String returns the string representation of TransactionSplitBearerType
func (b TransactionSplitBearerType) String() string {
	return string(b)
}

// TransactionSplitSubaccount represents a subaccount in a split configuration
type TransactionSplitSubaccount struct {
	Subaccount string `json:"subaccount"` // Subaccount code
	Share      int    `json:"share"`      // Share amount (percentage or flat amount)
}

// TransactionSplit represents a transaction split
type TransactionSplit struct {
	ID               uint64                       `json:"id"`
	Name             string                       `json:"name"`
	Type             TransactionSplitType         `json:"type"`
	Currency         Currency                     `json:"currency"`
	Integration      uint64                       `json:"integration"`
	Domain           string                       `json:"domain"`
	SplitCode        string                       `json:"split_code"`
	Active           bool                         `json:"active"`
	BearerType       TransactionSplitBearerType   `json:"bearer_type"`
	BearerSubaccount *string                      `json:"bearer_subaccount"`
	CreatedAt        time.Time                    `json:"createdAt"`
	UpdatedAt        time.Time                    `json:"updatedAt"`
	IsDynamic        bool                         `json:"is_dynamic"`
	Subaccounts      []TransactionSplitSubaccount `json:"subaccounts"`
	TotalSubaccounts int                          `json:"total_subaccounts"`
}
