package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Plan represents a Paystack subscription plan
type Plan struct {
	ID                       uint64             `json:"id"`
	Domain                   string             `json:"domain"`
	Name                     string             `json:"name"`
	PlanCode                 string             `json:"plan_code"`
	Description              data.NullString    `json:"description"`
	Amount                   int                `json:"amount"`
	Interval                 enums.Interval     `json:"interval"`
	InvoiceLimit             int                `json:"invoice_limit"`
	SendInvoices             bool               `json:"send_invoices"`
	SendSms                  bool               `json:"send_sms"`
	HostedPage               bool               `json:"hosted_page"`
	HostedPageURL            data.NullString    `json:"hosted_page_url,omitempty"`
	HostedPageSummary        data.NullString    `json:"hosted_page_summary,omitempty"`
	Currency                 enums.Currency     `json:"currency"`
	Integration              int                `json:"integration"`
	Migrate                  bool               `json:"migrate"`
	IsDeleted                bool               `json:"is_deleted"`
	IsArchived               bool               `json:"is_archived"`
	CreatedAt                data.MultiDateTime `json:"createdAt"`
	UpdatedAt                data.MultiDateTime `json:"updatedAt"`
	PagesCount               int                `json:"pages_count"`
	SubscribersCount         int                `json:"subscribers_count"`
	ActiveSubscriptionsCount int                `json:"active_subscriptions_count"`
	TotalRevenue             int                `json:"total_revenue"`
	Subscriptions            []Subscription     `json:"subscriptions,omitempty"`
	Pages                    []PaymentPage      `json:"pages,omitempty"` // Use PaymentPage instead of Page
}
