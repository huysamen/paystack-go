package types

import (
	"github.com/huysamen/paystack-go/enums"
)

// TransactionSplitSubaccount represents a subaccount in a split configuration
type TransactionSplitSubaccount struct {
	Subaccount string `json:"subaccount"` // Subaccount code
	Share      int    `json:"share"`      // Share amount (percentage or flat amount)
}

// TransactionSplit represents a transaction split
type TransactionSplit struct {
	ID               uint64                           `json:"id"`
	Name             string                           `json:"name"`
	Type             enums.TransactionSplitType       `json:"type"`
	Currency         enums.Currency                   `json:"currency"`
	Integration      uint64                           `json:"integration"`
	Domain           string                           `json:"domain"`
	SplitCode        string                           `json:"split_code"`
	Active           bool                             `json:"active"`
	BearerType       enums.TransactionSplitBearerType `json:"bearer_type"`
	BearerSubaccount *string                          `json:"bearer_subaccount"`
	CreatedAt        DateTime                         `json:"createdAt"`
	UpdatedAt        DateTime                         `json:"updatedAt"`
	IsDynamic        bool                             `json:"is_dynamic"`
	Subaccounts      []TransactionSplitSubaccount     `json:"subaccounts"`
	TotalSubaccounts int                              `json:"total_subaccounts"`
}
