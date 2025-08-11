package transferscontrol

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransfersControl_EnableOTP_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transferscontrol", "enable_otp_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp EnableOTPResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
	assert.True(t, rsp.Status.Bool())
}

// Moved disable OTP test into disable_otp_test.go
