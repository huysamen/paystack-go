package paystack

import "testing"

func TestClient_WiresWebhookEndpoints(t *testing.T) {
	c := DefaultClient("sk_test_wiring")

	if c.WebhookEndpoints == nil {
		t.Fatal("WebhookEndpoints is nil")
	}

	if c.WebhookEndpoints.Secret != "sk_test_wiring" {
		t.Fatalf("WebhookEndpoints.Secret = %q, want sk_test_wiring", c.WebhookEndpoints.Secret)
	}
}
