package paystack

import (
	"net/http"
	"time"
)

// Environment represents the Paystack API environment
type Environment string

const (
	// EnvironmentProduction uses the live Paystack API
	EnvironmentProduction Environment = "production"
	// EnvironmentSandbox uses the test Paystack API (not officially supported by Paystack, but useful for testing)
	EnvironmentSandbox Environment = "sandbox"
)

// Config holds configuration for the Paystack client
type Config struct {
	// SecretKey is the Paystack secret key
	SecretKey string

	// Environment determines which API endpoint to use
	Environment Environment

	// HTTPClient is the HTTP client to use for requests
	// If nil, a default client will be created
	HTTPClient *http.Client

	// Timeout for HTTP requests
	Timeout time.Duration

	// BaseURL allows overriding the API base URL (useful for testing)
	BaseURL string
}

// NewConfig creates a new configuration with sensible defaults
func NewConfig(secretKey string) *Config {
	return &Config{
		SecretKey:   secretKey,
		Environment: EnvironmentProduction,
		Timeout:     60 * time.Second,
	}
}

// WithEnvironment sets the environment
func (c *Config) WithEnvironment(env Environment) *Config {
	c.Environment = env
	return c
}

// WithHTTPClient sets a custom HTTP client
func (c *Config) WithHTTPClient(client *http.Client) *Config {
	c.HTTPClient = client
	return c
}

// WithTimeout sets the request timeout
func (c *Config) WithTimeout(timeout time.Duration) *Config {
	c.Timeout = timeout
	return c
}

// WithBaseURL sets a custom base URL
func (c *Config) WithBaseURL(baseURL string) *Config {
	c.BaseURL = baseURL
	return c
}

// GetBaseURL returns the appropriate base URL for the environment
func (c *Config) GetBaseURL() string {
	if c.BaseURL != "" {
		return c.BaseURL
	}

	switch c.Environment {
	case EnvironmentSandbox:
		return "https://api.paystack.co" // Paystack doesn't have a separate sandbox URL
	default:
		return "https://api.paystack.co"
	}
}
