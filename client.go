// Package paystack provides a Go client library for the Paystack API.
//
// Paystack is a payment infrastructure for businesses in Africa. This library
// provides a simple, type-safe way to interact with the Paystack API.
//
// # Basic Usage
//
//	client := paystack.DefaultClient("sk_test_your_secret_key")
//
//	req := &transactions.TransactionInitializeRequest{
//		Amount: 50000, // 500.00 NGN in kobo
//		Email:  "customer@example.com",
//	}
//
//	resp, err := client.Transactions.Initialize(req)
//	if err != nil {
//		// Handle error
//	}
//
// # Configuration
//
// For advanced usage, you can create a custom configuration:
//
//	config := paystack.NewConfig("sk_test_your_secret_key").
//		WithTimeout(30 * time.Second).
//		WithEnvironment(paystack.EnvironmentProduction)
//
//	client := paystack.NewClient(config)
//
// # Error Handling
//
// The library provides structured error handling:
//
//	if paystackErr, ok := err.(*net.PaystackError); ok {
//		fmt.Printf("API Error: %s (Status: %d)", paystackErr.Message, paystackErr.StatusCode)
//	}
package paystack

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	applepay "github.com/huysamen/paystack-go/api/apple-pay"
	"github.com/huysamen/paystack-go/api/customers"
	dedicated_virtual_account "github.com/huysamen/paystack-go/api/dedicated-virtual-account"
	direct_debit "github.com/huysamen/paystack-go/api/direct-debit"
	"github.com/huysamen/paystack-go/api/miscellaneous"
	"github.com/huysamen/paystack-go/api/plans"
	"github.com/huysamen/paystack-go/api/products"
	"github.com/huysamen/paystack-go/api/settlements"
	"github.com/huysamen/paystack-go/api/subaccounts"
	"github.com/huysamen/paystack-go/api/subscriptions"
	"github.com/huysamen/paystack-go/api/terminal"
	transaction_splits "github.com/huysamen/paystack-go/api/transaction-splits"
	"github.com/huysamen/paystack-go/api/transactions"
	transfer_recipients "github.com/huysamen/paystack-go/api/transfer-recipients"
	"github.com/huysamen/paystack-go/api/transfers"
	"github.com/huysamen/paystack-go/api/verification"
	virtual_terminal "github.com/huysamen/paystack-go/api/virtual-terminal"
)

type Client struct {
	Transactions            *transactions.Client
	Plans                   *plans.Client
	Products                *products.Client
	Customers               *customers.Client
	Subscriptions           *subscriptions.Client
	Transfers               *transfers.Client
	TransferRecipients      *transfer_recipients.Client
	Subaccounts             *subaccounts.Client
	Settlements             *settlements.Client
	Miscellaneous           *miscellaneous.Client
	Verification            *verification.Client
	TransactionSplits       *transaction_splits.Client
	Terminal                *terminal.Client
	VirtualTerminal         *virtual_terminal.Client
	DirectDebit             *direct_debit.Client
	DedicatedVirtualAccount *dedicated_virtual_account.Client
	ApplePay                *applepay.Client
}

// DefaultClient creates a new client with default configuration
func DefaultClient(secretKey string) *Client {
	return NewClient(NewConfig(secretKey))
}

// NewClientWithHTTP creates a new client with a custom HTTP client (deprecated, use NewClient with Config)
func NewClientWithHTTP(secretKey string, httpClient *http.Client) *Client {
	config := NewConfig(secretKey).WithHTTPClient(httpClient)
	return NewClient(config)
}

// NewClient creates a new Paystack client with the given configuration
func NewClient(config *Config) *Client {
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

	client := &Client{
		Transactions:            transactions.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Plans:                   plans.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Products:                products.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Customers:               customers.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Subscriptions:           subscriptions.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Transfers:               transfers.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		TransferRecipients:      transfer_recipients.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Subaccounts:             subaccounts.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Settlements:             settlements.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Miscellaneous:           miscellaneous.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Verification:            verification.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		TransactionSplits:       transaction_splits.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		Terminal:                terminal.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		VirtualTerminal:         virtual_terminal.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		DirectDebit:             direct_debit.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		DedicatedVirtualAccount: dedicated_virtual_account.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
		ApplePay:                applepay.NewClient(httpClient, config.SecretKey, config.GetBaseURL()),
	}

	return client
}
