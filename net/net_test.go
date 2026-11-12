package net

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPatch_SendsMethodBodyAndAuth(t *testing.T) {
	var (
		method, path, auth, contentType string
		body                            []byte
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method, path = r.Method, r.URL.Path
		auth, contentType = r.Header.Get("Authorization"), r.Header.Get("Content-Type")
		body, _ = io.ReadAll(r.Body)

		_, _ = w.Write([]byte(`{"status":true,"message":"ok","data":{}}`))
	}))
	defer srv.Close()

	payload := map[string]any{"enabled": false}

	rsp, err := Patch[map[string]any, map[string]any](context.Background(), srv.Client(), "sk_test_x", "/thing/1", &payload, srv.URL)
	require.NoError(t, err)

	assert.Equal(t, http.MethodPatch, method)
	assert.Equal(t, "/thing/1", path)
	assert.Equal(t, "Bearer sk_test_x", auth)
	assert.Equal(t, "application/json", contentType)
	assert.JSONEq(t, `{"enabled":false}`, string(body))
	assert.True(t, rsp.Status.Bool())
}
