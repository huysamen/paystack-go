package dedicatedvirtualaccount

import (
	"github.com/huysamen/paystack-go/types"
)

// Bank represents a bank provider for dedicated virtual accounts
type Bank struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// BankProvider represents a bank provider
type BankProvider struct {
	ID           int    `json:"id"`
	ProviderSlug string `json:"provider_slug"`
	BankID       int    `json:"bank_id"`
	BankName     string `json:"bank_name"`
}

// Customer represents customer information in dedicated virtual account
type Customer struct {
	ID                       int             `json:"id"`
	FirstName                string          `json:"first_name,omitempty"`
	LastName                 string          `json:"last_name,omitempty"`
	Email                    string          `json:"email"`
	CustomerCode             string          `json:"customer_code"`
	Phone                    string          `json:"phone,omitempty"`
	Metadata                 *types.Metadata `json:"metadata,omitempty"`
	RiskAction               string          `json:"risk_action,omitempty"`
	InternationalFormatPhone *string         `json:"international_format_phone,omitempty"`
	CreatedAt                string          `json:"created_at,omitempty"`
	UpdatedAt                string          `json:"updated_at,omitempty"`
}

// DedicatedVirtualAccount represents a dedicated virtual account
type DedicatedVirtualAccount struct {
	ID            int             `json:"id"`
	AccountName   string          `json:"account_name"`
	AccountNumber string          `json:"account_number"`
	Assigned      bool            `json:"assigned"`
	Currency      string          `json:"currency"`
	Metadata      *types.Metadata `json:"metadata,omitempty"`
	Active        bool            `json:"active"`
	Bank          Bank            `json:"bank"`
	Customer      *Customer       `json:"customer,omitempty"`
	CreatedAt     string          `json:"created_at,omitempty"`
	UpdatedAt     string          `json:"updated_at,omitempty"`
	SplitConfig   interface{}     `json:"split_config,omitempty"`
}

// CreateDedicatedVirtualAccountRequest represents the request to create a dedicated virtual account
type CreateDedicatedVirtualAccountRequest struct {
	Customer      string `json:"customer"`
	PreferredBank string `json:"preferred_bank,omitempty"`
	Subaccount    string `json:"subaccount,omitempty"`
	SplitCode     string `json:"split_code,omitempty"`
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	Phone         string `json:"phone,omitempty"`
}

// AssignDedicatedVirtualAccountRequest represents the request to assign a dedicated virtual account
type AssignDedicatedVirtualAccountRequest struct {
	Email         string `json:"email"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Phone         string `json:"phone"`
	PreferredBank string `json:"preferred_bank"`
	Country       string `json:"country"`
	AccountNumber string `json:"account_number,omitempty"`
	BVN           string `json:"bvn,omitempty"`
	BankCode      string `json:"bank_code,omitempty"`
	Subaccount    string `json:"subaccount,omitempty"`
	SplitCode     string `json:"split_code,omitempty"`
	MiddleName    string `json:"middle_name,omitempty"`
}

// ListDedicatedVirtualAccountsRequest represents the request to list dedicated virtual accounts
type ListDedicatedVirtualAccountsRequest struct {
	Active       *bool  `json:"active,omitempty"`
	Currency     string `json:"currency,omitempty"`
	ProviderSlug string `json:"provider_slug,omitempty"`
	BankID       string `json:"bank_id,omitempty"`
	Customer     string `json:"customer,omitempty"`
}

// RequeryDedicatedAccountRequest represents the request to requery a dedicated account
type RequeryDedicatedAccountRequest struct {
	AccountNumber string `json:"account_number"`
	ProviderSlug  string `json:"provider_slug"`
	Date          string `json:"date,omitempty"`
}

// SplitDedicatedAccountTransactionRequest represents the request to add split to dedicated account
type SplitDedicatedAccountTransactionRequest struct {
	Customer      string `json:"customer"`
	Subaccount    string `json:"subaccount,omitempty"`
	SplitCode     string `json:"split_code,omitempty"`
	PreferredBank string `json:"preferred_bank,omitempty"`
}

// RemoveSplitFromDedicatedAccountRequest represents the request to remove split from dedicated account
type RemoveSplitFromDedicatedAccountRequest struct {
	AccountNumber string `json:"account_number"`
}

// CreateDedicatedVirtualAccountResponse represents the response from creating a dedicated virtual account
type CreateDedicatedVirtualAccountResponse struct {
	Status  bool                    `json:"status"`
	Message string                  `json:"message"`
	Data    DedicatedVirtualAccount `json:"data"`
}

// AssignDedicatedVirtualAccountResponse represents the response from assigning a dedicated virtual account
type AssignDedicatedVirtualAccountResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// ListDedicatedVirtualAccountsResponse represents the response from listing dedicated virtual accounts
type ListDedicatedVirtualAccountsResponse struct {
	Status  bool                      `json:"status"`
	Message string                    `json:"message"`
	Data    []DedicatedVirtualAccount `json:"data"`
	Meta    *types.Meta               `json:"meta,omitempty"`
}

// FetchDedicatedVirtualAccountResponse represents the response from fetching a dedicated virtual account
type FetchDedicatedVirtualAccountResponse struct {
	Status  bool                    `json:"status"`
	Message string                  `json:"message"`
	Data    DedicatedVirtualAccount `json:"data"`
}

// RequeryDedicatedAccountResponse represents the response from requerying a dedicated account
type RequeryDedicatedAccountResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// DeactivateDedicatedAccountResponse represents the response from deactivating a dedicated account
type DeactivateDedicatedAccountResponse struct {
	Status  bool                    `json:"status"`
	Message string                  `json:"message"`
	Data    DedicatedVirtualAccount `json:"data"`
}

// SplitDedicatedAccountTransactionResponse represents the response from adding split to dedicated account
type SplitDedicatedAccountTransactionResponse struct {
	Status  bool                    `json:"status"`
	Message string                  `json:"message"`
	Data    DedicatedVirtualAccount `json:"data"`
}

// RemoveSplitFromDedicatedAccountResponse represents the response from removing split from dedicated account
type RemoveSplitFromDedicatedAccountResponse struct {
	Status  bool                    `json:"status"`
	Message string                  `json:"message"`
	Data    DedicatedVirtualAccount `json:"data"`
}

// FetchBankProvidersResponse represents the response from fetching bank providers
type FetchBankProvidersResponse struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    []BankProvider `json:"data"`
}
