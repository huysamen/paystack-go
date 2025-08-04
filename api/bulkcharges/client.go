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

type Client api.API
