package paystack

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/huysamen/paystack-go/api/plans"
	"github.com/huysamen/paystack-go/api/transactions"
)

type Client struct {
	secret       string
	client       *http.Client
	Transactions *transactions.Client
	Plans        *plans.Client
}

func DefaultClient(secretKey string) *Client {
	return NewClient(secretKey, nil)
}

func NewClient(secretKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: time.Second * 60,
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				TLSClientConfig: &tls.Config{
					MinVersion: tls.VersionTLS12,
				},
			},
		}
	}

	client := &Client{
		secret: secretKey,
		client: httpClient,
	}

	client.Transactions = transactions.NewClient(secretKey, httpClient)
	client.Plans = plans.NewClient(secretKey, httpClient)

	return client
}
