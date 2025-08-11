package paymentrequests

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFinalizeResponse_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "paymentrequests", "finalize_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var response FinalizeResponse
	err = json.Unmarshal(b, &response)
	require.NoError(t, err)
	assert.True(t, response.Status.Bool())
}
