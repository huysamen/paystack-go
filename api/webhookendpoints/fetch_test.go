package webhookendpoints

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetch_JSONDeserialization(t *testing.T) {
	var rsp FetchResponse
	decodeFixture(t, "fetch_200.json", &rsp)

	require.True(t, rsp.Status.Bool())
	assert.Equal(t, "WHK_1a2b3c4d", rsp.Data.WebhookURLCode.String())
	assert.Equal(t, []string{"charge.failed", "charge.success"}, rsp.Data.EventNames())
}

func TestFetch_EventNamesFlattensSortsAndDedupes(t *testing.T) {
	d := FetchResponseData{Events: map[string][]string{
		"transactions": {"charge.success", "charge.failed"},
		"refunds":      {"refund.processed", "charge.success"},
	}}

	assert.Equal(t, []string{"charge.failed", "charge.success", "refund.processed"}, d.EventNames())
}

func TestFetch_404Deserializes(t *testing.T) {
	var rsp FetchResponse
	decodeFixture(t, "fetch_404.json", &rsp)

	assert.False(t, rsp.Status.Bool())
	assert.Equal(t, "Webhook endpoint not found", rsp.Message)
}

func TestFetch_EscapesCode(t *testing.T) {
	c, rec := recordingClient(t, `{"status":true,"message":"ok","data":{}}`)

	_, err := c.Fetch(context.Background(), "WHK/x")
	require.NoError(t, err)

	assert.Equal(t, "GET", rec.method)
	assert.Equal(t, "/integration/webhooks/endpoints/WHK%2Fx", rec.rawPath)
}
