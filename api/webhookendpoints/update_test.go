package webhookendpoints

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdate_JSONDeserialization(t *testing.T) {
	var rsp UpdateResponse
	decodeFixture(t, "update_200.json", &rsp)

	require.True(t, rsp.Status.Bool())
	assert.False(t, rsp.Data.Enabled.Bool())
	assert.Equal(t, []string{"charge.success", "charge.failed"}, rsp.Data.Events)
}

func TestUpdate_OmitsUnsetFields(t *testing.T) {
	c, rec := recordingClient(t, `{"status":true,"message":"ok","data":{}}`)

	_, err := c.Update(context.Background(), "WHK_1", *NewUpdateRequestBuilder().Enabled(false))
	require.NoError(t, err)

	assert.Equal(t, "PATCH", rec.method)
	assert.Equal(t, "/integration/webhooks/endpoints/WHK_1", rec.rawPath)
	assert.JSONEq(t, `{"enabled":false}`, string(rec.body))
}

func TestUpdate_URLAndEvents(t *testing.T) {
	c, rec := recordingClient(t, `{"status":true,"message":"ok","data":{}}`)

	_, err := c.Update(context.Background(), "WHK_1", *NewUpdateRequestBuilder().URL("https://x").Events([]string{"charge.success"}))
	require.NoError(t, err)

	assert.JSONEq(t, `{"url":"https://x","events":["charge.success"]}`, string(rec.body))
}

func TestUpdate_Name(t *testing.T) {
	c, rec := recordingClient(t, `{"status":true,"message":"ok","data":{}}`)

	_, err := c.Update(context.Background(), "WHK_1", *NewUpdateRequestBuilder().Name("n"))
	require.NoError(t, err)

	assert.JSONEq(t, `{"name":"n"}`, string(rec.body))
}
