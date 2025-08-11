package subscriptions

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubscriptions_Enable_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "subscriptions", "enable_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp EnableResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
	assert.True(t, rsp.Status.Bool())
}

func TestSubscriptions_Enable_Builder(t *testing.T) {
	r := NewEnableRequestBuilder("SUB_x", "token").Build()
	assert.Equal(t, "SUB_x", r.Code)
	assert.Equal(t, "token", r.Token)
}
