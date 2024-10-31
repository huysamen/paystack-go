package transactions

import (
	"net/http"
)

const (
	transactionBasePath               = "/transaction"
	transactionInitializePath         = "/initialize"
	transactionVerifyPath             = "/verify"
	transactionChargeAutorizationPath = "/charge_authorization"
	transactionViewTimelinePath       = "/timeline"
	transactionExportPath             = "/export"
	transactionPartialDebitPath       = "/partial_debit"
)

type Client struct {
	client *http.Client
	secret string
}

func NewClient(secret string, client *http.Client) *Client {
	return &Client{
		secret: secret,
		client: client,
	}
}
