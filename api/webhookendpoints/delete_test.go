package webhookendpoints

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDelete_JSONDeserialization(t *testing.T) {
	var rsp DeleteResponse
	decodeFixture(t, "delete_200.json", &rsp)

	require.True(t, rsp.Status.Bool())
	assert.Equal(t, "WHK_1a2b3c4d", rsp.Data.WebhookURLCode.String())
}

func TestDelete_RequestShape(t *testing.T) {
	c, rec := recordingClient(t, `{"status":true,"message":"ok","data":{}}`)

	_, err := c.Delete(context.Background(), "WHK_1")
	require.NoError(t, err)

	assert.Equal(t, "DELETE", rec.method)
	assert.Equal(t, "/integration/webhooks/endpoints/WHK_1", rec.rawPath)
}
