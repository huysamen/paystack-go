package transactionsplits

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

// TransactionSplitCreateRequestBuilder provides a fluent interface for building TransactionSplitCreateRequest
type TransactionSplitCreateRequestBuilder struct {
	name             string
	splitType        TransactionSplitType
	currency         types.Currency
	subaccounts      []TransactionSplitSubaccount
	bearerType       *TransactionSplitBearerType
	bearerSubaccount *string
}

// NewTransactionSplitCreateRequest creates a new builder for creating a transaction split
func NewTransactionSplitCreateRequest(name string, splitType TransactionSplitType, currency types.Currency) *TransactionSplitCreateRequestBuilder {
	return &TransactionSplitCreateRequestBuilder{
		name:        name,
		splitType:   splitType,
		currency:    currency,
		subaccounts: make([]TransactionSplitSubaccount, 0),
	}
}

// AddSubaccount adds a subaccount to the split
func (b *TransactionSplitCreateRequestBuilder) AddSubaccount(subaccount string, share int) *TransactionSplitCreateRequestBuilder {
	b.subaccounts = append(b.subaccounts, TransactionSplitSubaccount{
		Subaccount: subaccount,
		Share:      share,
	})
	return b
}

// Subaccounts sets all subaccounts at once
func (b *TransactionSplitCreateRequestBuilder) Subaccounts(subaccounts []TransactionSplitSubaccount) *TransactionSplitCreateRequestBuilder {
	b.subaccounts = subaccounts
	return b
}

// BearerType sets who bears the charges
func (b *TransactionSplitCreateRequestBuilder) BearerType(bearerType TransactionSplitBearerType) *TransactionSplitCreateRequestBuilder {
	b.bearerType = &bearerType
	return b
}

// BearerSubaccount sets the subaccount that bears the charges (when bearer_type is subaccount)
func (b *TransactionSplitCreateRequestBuilder) BearerSubaccount(subaccount string) *TransactionSplitCreateRequestBuilder {
	b.bearerSubaccount = &subaccount
	return b
}

// Build creates the TransactionSplitCreateRequest
func (b *TransactionSplitCreateRequestBuilder) Build() *TransactionSplitCreateRequest {
	return &TransactionSplitCreateRequest{
		Name:             b.name,
		Type:             b.splitType,
		Currency:         b.currency,
		Subaccounts:      b.subaccounts,
		BearerType:       b.bearerType,
		BearerSubaccount: b.bearerSubaccount,
	}
}

// TransactionSplitCreateResponse represents the response from creating a split
type TransactionSplitCreateResponse = types.Response[TransactionSplit]

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

// TransactionSplitListRequestBuilder provides a fluent interface for building TransactionSplitListRequest
type TransactionSplitListRequestBuilder struct {
	name    *string
	active  *bool
	sortBy  *string
	perPage *int
	page    *int
	from    *time.Time
	to      *time.Time
}

// NewTransactionSplitListRequest creates a new builder for listing transaction splits
func NewTransactionSplitListRequest() *TransactionSplitListRequestBuilder {
	return &TransactionSplitListRequestBuilder{}
}

// Name filters by split name
func (b *TransactionSplitListRequestBuilder) Name(name string) *TransactionSplitListRequestBuilder {
	b.name = &name
	return b
}

// Active filters by active status
func (b *TransactionSplitListRequestBuilder) Active(active bool) *TransactionSplitListRequestBuilder {
	b.active = &active
	return b
}

// SortBy sets the sort field
func (b *TransactionSplitListRequestBuilder) SortBy(sortBy string) *TransactionSplitListRequestBuilder {
	b.sortBy = &sortBy
	return b
}

// PerPage sets the number of records per page
func (b *TransactionSplitListRequestBuilder) PerPage(perPage int) *TransactionSplitListRequestBuilder {
	b.perPage = &perPage
	return b
}

// Page sets the page number
func (b *TransactionSplitListRequestBuilder) Page(page int) *TransactionSplitListRequestBuilder {
	b.page = &page
	return b
}

// DateRange sets both from and to dates
func (b *TransactionSplitListRequestBuilder) DateRange(from, to time.Time) *TransactionSplitListRequestBuilder {
	b.from = &from
	b.to = &to
	return b
}

// From sets the start date filter
func (b *TransactionSplitListRequestBuilder) From(from time.Time) *TransactionSplitListRequestBuilder {
	b.from = &from
	return b
}

// To sets the end date filter
func (b *TransactionSplitListRequestBuilder) To(to time.Time) *TransactionSplitListRequestBuilder {
	b.to = &to
	return b
}

// Build creates the TransactionSplitListRequest
func (b *TransactionSplitListRequestBuilder) Build() *TransactionSplitListRequest {
	return &TransactionSplitListRequest{
		Name:    b.name,
		Active:  b.active,
		SortBy:  b.sortBy,
		PerPage: b.perPage,
		Page:    b.page,
		From:    b.from,
		To:      b.to,
	}
}

// TransactionSplitListResponse represents the response from listing splits
type TransactionSplitListResponse = types.Response[[]TransactionSplit]

// TransactionSplit Fetch

// TransactionSplitFetchResponse represents the response from fetching a split
type TransactionSplitFetchResponse = types.Response[TransactionSplit]

// TransactionSplit Update

// TransactionSplitUpdateRequest represents the request to update a split
type TransactionSplitUpdateRequest struct {
	Name             *string                     `json:"name,omitempty"`              // Name of the transaction split (optional)
	Active           *bool                       `json:"active,omitempty"`            // Active status (optional)
	BearerType       *TransactionSplitBearerType `json:"bearer_type,omitempty"`       // Who bears the charges (optional)
	BearerSubaccount *string                     `json:"bearer_subaccount,omitempty"` // Subaccount code if bearer_type is subaccount (optional)
}

// TransactionSplitUpdateRequestBuilder provides a fluent interface for building TransactionSplitUpdateRequest
type TransactionSplitUpdateRequestBuilder struct {
	name             *string
	active           *bool
	bearerType       *TransactionSplitBearerType
	bearerSubaccount *string
}

// NewTransactionSplitUpdateRequest creates a new builder for updating a transaction split
func NewTransactionSplitUpdateRequest() *TransactionSplitUpdateRequestBuilder {
	return &TransactionSplitUpdateRequestBuilder{}
}

// Name sets the split name
func (b *TransactionSplitUpdateRequestBuilder) Name(name string) *TransactionSplitUpdateRequestBuilder {
	b.name = &name
	return b
}

// Active sets the active status
func (b *TransactionSplitUpdateRequestBuilder) Active(active bool) *TransactionSplitUpdateRequestBuilder {
	b.active = &active
	return b
}

// BearerType sets who bears the charges
func (b *TransactionSplitUpdateRequestBuilder) BearerType(bearerType TransactionSplitBearerType) *TransactionSplitUpdateRequestBuilder {
	b.bearerType = &bearerType
	return b
}

// BearerSubaccount sets the subaccount that bears the charges
func (b *TransactionSplitUpdateRequestBuilder) BearerSubaccount(subaccount string) *TransactionSplitUpdateRequestBuilder {
	b.bearerSubaccount = &subaccount
	return b
}

// Build creates the TransactionSplitUpdateRequest
func (b *TransactionSplitUpdateRequestBuilder) Build() *TransactionSplitUpdateRequest {
	return &TransactionSplitUpdateRequest{
		Name:             b.name,
		Active:           b.active,
		BearerType:       b.bearerType,
		BearerSubaccount: b.bearerSubaccount,
	}
}

// TransactionSplitUpdateResponse represents the response from updating a split
type TransactionSplitUpdateResponse = types.Response[TransactionSplit]

// TransactionSplit Subaccount Management

// TransactionSplitSubaccountAddRequest represents the request to add/update a subaccount in a split
type TransactionSplitSubaccountAddRequest struct {
	Subaccount string `json:"subaccount"` // Subaccount code
	Share      int    `json:"share"`      // Share amount (percentage or flat amount)
}

// TransactionSplitSubaccountAddRequestBuilder provides a fluent interface for building TransactionSplitSubaccountAddRequest
type TransactionSplitSubaccountAddRequestBuilder struct {
	subaccount string
	share      int
}

// NewTransactionSplitSubaccountAddRequest creates a new builder for adding a subaccount to a split
func NewTransactionSplitSubaccountAddRequest(subaccount string, share int) *TransactionSplitSubaccountAddRequestBuilder {
	return &TransactionSplitSubaccountAddRequestBuilder{
		subaccount: subaccount,
		share:      share,
	}
}

// Subaccount sets the subaccount code
func (b *TransactionSplitSubaccountAddRequestBuilder) Subaccount(subaccount string) *TransactionSplitSubaccountAddRequestBuilder {
	b.subaccount = subaccount
	return b
}

// Share sets the share amount
func (b *TransactionSplitSubaccountAddRequestBuilder) Share(share int) *TransactionSplitSubaccountAddRequestBuilder {
	b.share = share
	return b
}

// Build creates the TransactionSplitSubaccountAddRequest
func (b *TransactionSplitSubaccountAddRequestBuilder) Build() *TransactionSplitSubaccountAddRequest {
	return &TransactionSplitSubaccountAddRequest{
		Subaccount: b.subaccount,
		Share:      b.share,
	}
}

// TransactionSplitSubaccountRemoveRequest represents the request to remove a subaccount from a split
type TransactionSplitSubaccountRemoveRequest struct {
	Subaccount string `json:"subaccount"` // Subaccount code
}

// TransactionSplitSubaccountRemoveRequestBuilder provides a fluent interface for building TransactionSplitSubaccountRemoveRequest
type TransactionSplitSubaccountRemoveRequestBuilder struct {
	subaccount string
}

// NewTransactionSplitSubaccountRemoveRequest creates a new builder for removing a subaccount from a split
func NewTransactionSplitSubaccountRemoveRequest(subaccount string) *TransactionSplitSubaccountRemoveRequestBuilder {
	return &TransactionSplitSubaccountRemoveRequestBuilder{
		subaccount: subaccount,
	}
}

// Subaccount sets the subaccount code
func (b *TransactionSplitSubaccountRemoveRequestBuilder) Subaccount(subaccount string) *TransactionSplitSubaccountRemoveRequestBuilder {
	b.subaccount = subaccount
	return b
}

// Build creates the TransactionSplitSubaccountRemoveRequest
func (b *TransactionSplitSubaccountRemoveRequestBuilder) Build() *TransactionSplitSubaccountRemoveRequest {
	return &TransactionSplitSubaccountRemoveRequest{
		Subaccount: b.subaccount,
	}
}

// TransactionSplitSubaccountRemoveResponse represents the response from removing a subaccount from a split
type TransactionSplitSubaccountRemoveResponse = types.Response[any]
