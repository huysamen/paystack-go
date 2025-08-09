package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Plan represents a Paystack subscription plan
type Plan struct {
	ID                       data.Uint          `json:"id"`
	Domain                   data.String        `json:"domain"`
	Name                     data.String        `json:"name"`
	PlanCode                 data.String        `json:"plan_code"`
	Description              data.NullString    `json:"description"`
	Amount                   data.Int           `json:"amount"`
	Interval                 enums.Interval     `json:"interval"`
	InvoiceLimit             data.Int           `json:"invoice_limit"`
	SendInvoices             data.Bool          `json:"send_invoices"`
	SendSms                  data.Bool          `json:"send_sms"`
	HostedPage               data.Bool          `json:"hosted_page"`
	HostedPageURL            data.NullString    `json:"hosted_page_url,omitempty"`
	HostedPageSummary        data.NullString    `json:"hosted_page_summary,omitempty"`
	Currency                 enums.Currency     `json:"currency"`
	Integration              data.Int           `json:"integration"`
	Migrate                  data.Bool          `json:"migrate"`
	IsDeleted                data.Bool          `json:"is_deleted"`
	IsArchived               data.Bool          `json:"is_archived"`
	CreatedAt                data.MultiDateTime `json:"createdAt"`
	UpdatedAt                data.MultiDateTime `json:"updatedAt"`
	PagesCount               data.Int           `json:"pages_count"`
	SubscribersCount         data.Int           `json:"subscribers_count"`
	ActiveSubscriptionsCount data.Int           `json:"active_subscriptions_count"`
	TotalRevenue             data.Int           `json:"total_revenue"`
	Subscriptions            []Subscription     `json:"subscriptions,omitempty"`
	Pages                    []PaymentPage      `json:"pages,omitempty"` // Use PaymentPage instead of Page
}
