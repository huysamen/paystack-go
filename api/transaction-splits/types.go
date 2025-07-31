package transaction_splits

import (
	"time"

	"github.com/huysamen/paystack-go/types"
)

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
	Currency         types.Currency               `json:"currency"`
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

// TransactionSplit Create

// TransactionSplitCreateRequest represents the request to create a split
type TransactionSplitCreateRequest struct {
	Name             string                       `json:"name"`                        // Name of the transaction split
	Type             TransactionSplitType         `json:"type"`                        // Type of split (percentage or flat)
	Currency         types.Currency               `json:"currency"`                    // Currency for the split
	Subaccounts      []TransactionSplitSubaccount `json:"subaccounts"`                 // List of subaccounts and their shares
	BearerType       *TransactionSplitBearerType  `json:"bearer_type,omitempty"`       // Who bears the charges (optional)
	BearerSubaccount *string                      `json:"bearer_subaccount,omitempty"` // Subaccount code if bearer_type is subaccount (optional)
}

// TransactionSplitCreateResponse represents the response from creating a split
type TransactionSplitCreateResponse struct {
	Status  bool             `json:"status"`
	Message string           `json:"message"`
	Data    TransactionSplit `json:"data"`
}

// TransactionSplit List

// TransactionSplitListRequest represents the request to list splits
type TransactionSplitListRequest struct {
	Name    *string    `json:"name,omitempty"`    // Filter by name (optional)
	Active  *bool      `json:"active,omitempty"`  // Filter by active status (optional)
	SortBy  *string    `json:"sort_by,omitempty"` // Sort by field, defaults to createdAt (optional)
	PerPage *int       `json:"perPage,omitempty"` // Number of splits per page (default: 50)
	Page    *int       `json:"page,omitempty"`    // Page number (default: 1)
	From    *time.Time `json:"from,omitempty"`    // Start date filter (optional)
	To      *time.Time `json:"to,omitempty"`      // End date filter (optional)
}

// TransactionSplitListResponse represents the response from listing splits
type TransactionSplitListResponse struct {
	Status  bool               `json:"status"`
	Message string             `json:"message"`
	Data    []TransactionSplit `json:"data"`
	Meta    types.Meta         `json:"meta"`
}

// TransactionSplit Fetch

// TransactionSplitFetchResponse represents the response from fetching a split
type TransactionSplitFetchResponse struct {
	Status  bool             `json:"status"`
	Message string           `json:"message"`
	Data    TransactionSplit `json:"data"`
}

// TransactionSplit Update

// TransactionSplitUpdateRequest represents the request to update a split
type TransactionSplitUpdateRequest struct {
	Name             *string                     `json:"name,omitempty"`              // Name of the transaction split (optional)
	Active           *bool                       `json:"active,omitempty"`            // Active status (optional)
	BearerType       *TransactionSplitBearerType `json:"bearer_type,omitempty"`       // Who bears the charges (optional)
	BearerSubaccount *string                     `json:"bearer_subaccount,omitempty"` // Subaccount code if bearer_type is subaccount (optional)
}

// TransactionSplitUpdateResponse represents the response from updating a split
type TransactionSplitUpdateResponse struct {
	Status  bool             `json:"status"`
	Message string           `json:"message"`
	Data    TransactionSplit `json:"data"`
}

// TransactionSplit Subaccount Management

// TransactionSplitSubaccountAddRequest represents the request to add/update a subaccount in a split
type TransactionSplitSubaccountAddRequest struct {
	Subaccount string `json:"subaccount"` // Subaccount code
	Share      int    `json:"share"`      // Share amount (percentage or flat amount)
}

// TransactionSplitSubaccountAddResponse represents the response from adding/updating a subaccount in a split
type TransactionSplitSubaccountAddResponse struct {
	Status  bool             `json:"status"`
	Message string           `json:"message"`
	Data    TransactionSplit `json:"data"`
}

// TransactionSplitSubaccountRemoveRequest represents the request to remove a subaccount from a split
type TransactionSplitSubaccountRemoveRequest struct {
	Subaccount string `json:"subaccount"` // Subaccount code
}

// TransactionSplitSubaccountRemoveResponse represents the response from removing a subaccount from a split
type TransactionSplitSubaccountRemoveResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
