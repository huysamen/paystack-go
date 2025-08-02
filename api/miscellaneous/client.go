package miscellaneous

import (
	"github.com/huysamen/paystack-go/api"
)

const (
	bankPath    = "/bank"
	countryPath = "/country"
	statesPath  = "/address_verification/states"
)

// Client provides miscellaneous operations
type Client api.API
