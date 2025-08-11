package verification

import "github.com/huysamen/paystack-go/api"

const (
	accountResolveBasePath  = "/bank/resolve"
	accountValidateBasePath = "/bank/validate"
	cardBINResolveBasePath  = "/decision/bin"
)

type Client api.API
