package paymentrequests

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSendNotificationResponse_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "paymentrequests", "send_notification_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var response SendNotificationResponse
	err = json.Unmarshal(b, &response)
	require.NoError(t, err)
	assert.True(t, response.Status.Bool())
}
