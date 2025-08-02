package applepay

import (
	"github.com/huysamen/paystack-go/api"
)

const (
	basePath       = "/apple-pay"
	listPath       = basePath + "/domain"
	registerPath   = basePath + "/domain"
	unregisterPath = basePath + "/domain"
)

// Client is the Apple Pay API client
type Client api.API
