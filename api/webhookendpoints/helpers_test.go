package webhookendpoints

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func decodeFixture(t *testing.T, name string, out any) {
	t.Helper()

	b, err := os.ReadFile(filepath.Join("..", "..", "resources", "examples", "responses", "webhookendpoints", name))
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(b, out))
}

// recorded is the last request a test server saw.
type recorded struct {
	method  string
	rawPath string
	query   string
	body    []byte
}

// recordingClient returns a Client against a server that records each request
// and answers answer.
func recordingClient(t *testing.T, answer string) (*Client, *recorded) {
	t.Helper()

	rec := &recorded{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec.method = r.Method
		rec.rawPath = r.URL.EscapedPath()
		rec.query = r.URL.RawQuery
		rec.body, _ = io.ReadAll(r.Body)

		_, _ = w.Write([]byte(answer))
	}))
	t.Cleanup(srv.Close)

	return &Client{Client: srv.Client(), Secret: "sk_test_x", BaseURL: srv.URL}, rec
}
