package transactions

import (
	"net/http"
)

const (
	transactionBasePath                = "/transaction"
	transactionInitializePath          = "/initialize"
	transactionVerifyPath              = "/verify"
	transactionChargeAuthorizationPath = "/charge_authorization"
	transactionViewTimelinePath        = "/timeline"
	transactionExportPath              = "/export"
	transactionPartialDebitPath        = "/partial_debit"
)

type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	if baseURL == "" {
		baseURL = "https://api.paystack.co"
	}

	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
