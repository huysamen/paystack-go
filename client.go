package paystack

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/huysamen/paystack-go/api"
	"github.com/huysamen/paystack-go/api/applepay"
	"github.com/huysamen/paystack-go/api/bulkcharges"
	"github.com/huysamen/paystack-go/api/charge"
	"github.com/huysamen/paystack-go/api/customers"
	"github.com/huysamen/paystack-go/api/dedicatedvirtualaccounts"
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
	"github.com/huysamen/paystack-go/api/transferrecipients"
	"github.com/huysamen/paystack-go/api/transfers"
	"github.com/huysamen/paystack-go/api/transferscontrol"
	"github.com/huysamen/paystack-go/api/verification"
	"github.com/huysamen/paystack-go/api/virtualterminal"
	pnet "github.com/huysamen/paystack-go/net"
)

type Client struct {
	Transactions            *transactions.Client
	Plans                   *plans.Client
	Products                *products.Client
	PaymentPages            *paymentpages.Client
	PaymentRequests         *paymentrequests.Client
	Customers               *customers.Client
	Subscriptions           *subscriptions.Client
	Transfers               *transfers.Client
	TransferControl         *transferscontrol.Client
	TransferRecipients      *transferrecipients.Client
	BulkCharges             *bulkcharges.Client
	Charges                 *charge.Client
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
	DedicatedVirtualAccount *dedicatedvirtualaccounts.Client
	ApplePay                *applepay.Client
	Integration             *integration.Client
}

// DefaultClient creates a new client with default configuration
func DefaultClient(secretKey string) *Client {
	return NewClient(NewConfig(secretKey))
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

	// Wrap transport to add default headers and optional UA suffix
	httpClient.Transport = pnet.NewHeaderRoundTripper(httpClient.Transport, config.DefaultHeaders, config.UserAgentSuffix)

	client := &Client{
		Transactions:            (*transactions.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		Plans:                   (*plans.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		Products:                (*products.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		PaymentPages:            (*paymentpages.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		PaymentRequests:         (*paymentrequests.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		Customers:               (*customers.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		Subscriptions:           (*subscriptions.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		Transfers:               (*transfers.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		TransferControl:         (*transferscontrol.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		TransferRecipients:      (*transferrecipients.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		BulkCharges:             (*bulkcharges.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		Charges:                 (*charge.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		Disputes:                (*disputes.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		Refunds:                 (*refunds.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		Subaccounts:             (*subaccounts.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		Settlements:             (*settlements.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		Miscellaneous:           (*miscellaneous.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		Verification:            (*verification.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		TransactionSplits:       (*transactionsplits.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		Terminal:                (*terminal.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		VirtualTerminal:         (*virtualterminal.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		DirectDebit:             (*directdebit.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		DedicatedVirtualAccount: (*dedicatedvirtualaccounts.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		ApplePay:                (*applepay.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
		Integration:             (*integration.Client)(&api.API{Client: httpClient, Secret: config.SecretKey, BaseURL: config.GetBaseURL(), Headers: config.DefaultHeaders}),
	}

	return client
}
