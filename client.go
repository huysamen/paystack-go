// Package paystack provides a Go client library for the Paystack API.
//
// Paystack is a payment infrastructure for businesses in Africa. This library
// provides a simple, type-safe way to interact with the Paystack API.
package paystack

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/huysamen/paystack-go/api/applepay"
	"github.com/huysamen/paystack-go/api/bulkcharges"
	"github.com/huysamen/paystack-go/api/charges"
	"github.com/huysamen/paystack-go/api/customers"
	"github.com/huysamen/paystack-go/api/dedicatedvirtualaccount"
	"github.com/huysamen/paystack-go/api/directdebit"
	"github.com/huysamen/paystack-go/api/disputes"
	"github.com/huysamen/paystack-go/api/integration"
	"github.com/huysamen/paystack-go/api/miscellaneous"
	"github.com/huysamen/paystack-go/api/paymentpages"
	"github.com/huysamen/paystack-go/api/paymentrequests"
	"github.com/huysamen/paystack-go/api/plans"
	"github.com/huysamen/paystack-go/api/products"
	"github.com/huysamen/paystack-go/api/refunds"
	"github.com/huysamen/paystack-go/api/settlements"
	"github.com/huysamen/paystack-go/api/subaccounts"
	"github.com/huysamen/paystack-go/api/subscriptions"
	"github.com/huysamen/paystack-go/api/terminal"
	"github.com/huysamen/paystack-go/api/transactions"
	"github.com/huysamen/paystack-go/api/transactionsplits"
	"github.com/huysamen/paystack-go/api/transfercontrol"
	"github.com/huysamen/paystack-go/api/transferrecipients"
	"github.com/huysamen/paystack-go/api/transfers"
	"github.com/huysamen/paystack-go/api/verification"
	"github.com/huysamen/paystack-go/api/virtualterminal"
)

type client struct {
	Transactions            *transactions.Client
	Plans                   *plans.Client
	Products                *products.Client
	PaymentPages            *paymentpages.Client
	PaymentRequests         *paymentrequests.Client
	Customers               *customers.Client
	Subscriptions           *subscriptions.Client
	Transfers               *transfers.Client
	TransferControl         *transfercontrol.Client
	TransferRecipients      *transferrecipients.Client
	BulkCharges             *bulkcharges.Client
	Charges                 *charges.Client
	Disputes                *disputes.Client
	Refunds                 *refunds.Client
	Subaccounts             *subaccounts.Client
	Settlements             *settlements.Client
	Miscellaneous           *miscellaneous.Client
	Verification            *verification.Client
	TransactionSplits       *transactionsplits.Client
	Terminal                *terminal.Client
	VirtualTerminal         *virtualterminal.Client
	DirectDebit             *directdebit.Client
	DedicatedVirtualAccount *dedicatedvirtualaccount.Client
	ApplePay                *applepay.Client
	Integration             *integration.Client
}

// DefaultClient creates a new client with default configuration
func DefaultClient(secretKey string) *client {
	return NewClient(NewConfig(secretKey))
}

// NewClientWithHTTP creates a new client with a custom HTTP client (deprecated, use NewClient with Config)
func NewClientWithHTTP(secretKey string, httpClient *http.Client) *client {
	config := NewConfig(secretKey).WithHTTPClient(httpClient)
	return NewClient(config)
}

// NewClient creates a new Paystack client with the given configuration
func NewClient(config *Config) *client {
	if config == nil {
		panic("config cannot be nil")
	}

	httpClient := config.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: config.Timeout,
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:          100,
				MaxIdleConnsPerHost:   10,
				IdleConnTimeout:       60 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				TLSClientConfig: &tls.Config{
					MinVersion: tls.VersionTLS12,
				},
			},
		}
	}

	client := &client{
		Transactions:            transactions.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Plans:                   plans.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Products:                products.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		PaymentPages:            paymentpages.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		PaymentRequests:         paymentrequests.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Customers:               customers.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Subscriptions:           subscriptions.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Transfers:               transfers.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		TransferControl:         transfercontrol.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		TransferRecipients:      transferrecipients.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		BulkCharges:             bulkcharges.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Charges:                 charges.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Disputes:                disputes.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Refunds:                 refunds.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Subaccounts:             subaccounts.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Settlements:             settlements.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Miscellaneous:           miscellaneous.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Verification:            verification.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		TransactionSplits:       transactionsplits.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Terminal:                terminal.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		VirtualTerminal:         virtualterminal.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		DirectDebit:             directdebit.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		DedicatedVirtualAccount: dedicatedvirtualaccount.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		ApplePay:                applepay.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Integration:             integration.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
	}

	return client
}
