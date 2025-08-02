package charges

import (
	"net/http"

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

// Client is the Charges API client
type Client api.API

// NewClient creates a new Charges API client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		Client:  httpClient,
		Secret:  secret,
		BaseURL: baseURL,
	}
}
