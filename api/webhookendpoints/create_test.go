package webhookendpoints

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate_JSONDeserialization(t *testing.T) {
	var rsp CreateResponse
	decodeFixture(t, "create_201.json", &rsp)

	require.True(t, rsp.Status.Bool())
	assert.Equal(t, "WHK_1a2b3c4d", rsp.Data.WebhookURLCode.String())
	assert.Equal(t, []string{"charge.success", "charge.failed"}, rsp.Data.Events)
	assert.True(t, rsp.Data.Enabled.Bool())
}

func TestCreate_409Deserializes(t *testing.T) {
	var rsp CreateResponse
	decodeFixture(t, "create_409.json", &rsp)

	assert.False(t, rsp.Status.Bool())
	assert.NotEmpty(t, rsp.Message)
}

func TestCreate_RequestShape(t *testing.T) {
	c, rec := recordingClient(t, `{"status":true,"message":"ok","data":{}}`)

	_, err := c.Create(context.Background(), *NewCreateRequestBuilder("n", "https://x", []string{"charge.success"}))
	require.NoError(t, err)

	assert.Equal(t, "POST", rec.method)
	assert.Equal(t, "/integration/webhooks/endpoints", rec.rawPath)
	assert.JSONEq(t, `{"name":"n","url":"https://x","events":["charge.success"]}`, string(rec.body))
}
