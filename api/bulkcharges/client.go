package bulkcharges

import (
	"github.com/huysamen/paystack-go/api"
)

const (
	basePath         = "/bulkcharge"
	pausePath        = basePath + "/pause"
	resumePath       = basePath + "/resume"
	fetchChargesPath = "/charges"
)

// Client is the Bulk Charges API client
type Client api.API
