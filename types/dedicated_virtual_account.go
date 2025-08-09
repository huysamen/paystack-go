package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// DedicatedVirtualAccountBank represents a bank provider for dedicated virtual accounts
type DedicatedVirtualAccountBank struct {
	ID   data.Int    `json:"id"`
	Name data.String `json:"name"`
	Slug data.String `json:"slug"`
}

// DedicatedVirtualAccountAssignment represents the assignment details of a dedicated virtual account
type DedicatedVirtualAccountAssignment struct {
	Integration  data.Int    `json:"integration"`
	AssigneeID   data.Int    `json:"assignee_id"`
	AssigneeType data.String `json:"assignee_type"`
	Expired      data.Bool   `json:"expired"`
	AccountType  data.String `json:"account_type"`
	AssignedAt   data.Time   `json:"assigned_at"`
}

// DedicatedVirtualAccount represents a dedicated virtual account
type DedicatedVirtualAccount struct {
	ID            data.Int                           `json:"id"`
	AccountName   data.String                        `json:"account_name"`
	AccountNumber data.String                        `json:"account_number"`
	Assigned      data.Bool                          `json:"assigned"`
	Currency      enums.Currency                     `json:"currency"`
	Metadata      *Metadata                          `json:"metadata,omitempty"`
	Active        data.Bool                          `json:"active"`
	Bank          DedicatedVirtualAccountBank        `json:"bank"`
	Customer      *Customer                          `json:"customer,omitempty"`
	Assignment    *DedicatedVirtualAccountAssignment `json:"assignment,omitempty"`
	CreatedAt     data.Time                          `json:"created_at,omitempty"`
	UpdatedAt     data.Time                          `json:"updated_at,omitempty"`
	SplitConfig   *Metadata                          `json:"split_config,omitempty"`
}
