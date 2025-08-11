package transactions

import "github.com/huysamen/paystack-go/api"

const (
	basePath                           = "/transaction"
	transactionInitializePath          = "/initialize"
	transactionVerifyPath              = "/verify"
	transactionChargeAuthorizationPath = "/charge_authorization"
	transactionViewTimelinePath        = "/timeline"
	transactionExportPath              = "/export"
	transactionPartialDebitPath        = "/partial_debit"
)

type Client api.API
