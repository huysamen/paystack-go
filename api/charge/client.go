package charges

import (
	"github.com/huysamen/paystack-go/api"
)

const (
	basePath           = "/charge"
	submitPinPath      = basePath + "/submit_pin"
	submitOtpPath      = basePath + "/submit_otp"
	submitPhonePath    = basePath + "/submit_phone"
	submitBirthdayPath = basePath + "/submit_birthday"
	submitAddressPath  = basePath + "/submit_address"
	checkPendingPath   = basePath + "/check_pending"
)

type Client api.API
