package models

import "github.com/huysamen/paystack-go/enums"

// DedicatedVirtualAccountBank represents a bank provider for dedicated virtual accounts
type DedicatedVirtualAccountBank struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// DedicatedVirtualAccountAssignment represents the assignment details of a dedicated virtual account
type DedicatedVirtualAccountAssignment struct {
	Integration  int      `json:"integration"`
	AssigneeID   int      `json:"assignee_id"`
	AssigneeType string   `json:"assignee_type"`
	Expired      bool     `json:"expired"`
	AccountType  string   `json:"account_type"`
	AssignedAt   DateTime `json:"assigned_at"`
}

// DedicatedVirtualAccount represents a dedicated virtual account
type DedicatedVirtualAccount struct {
	ID            int                                `json:"id"`
	AccountName   string                             `json:"account_name"`
	AccountNumber string                             `json:"account_number"`
	Assigned      bool                               `json:"assigned"`
	Currency      enums.Currency                     `json:"currency"`
	Metadata      *Metadata                          `json:"metadata,omitempty"`
	Active        bool                               `json:"active"`
	Bank          DedicatedVirtualAccountBank        `json:"bank"`
	Customer      *Customer                          `json:"customer,omitempty"`
	Assignment    *DedicatedVirtualAccountAssignment `json:"assignment,omitempty"`
	CreatedAt     DateTime                           `json:"created_at,omitempty"`
	UpdatedAt     DateTime                           `json:"updated_at,omitempty"`
	SplitConfig   *Metadata                          `json:"split_config,omitempty"`
}
