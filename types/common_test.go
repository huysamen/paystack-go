package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMeta_UnmarshalJSON(t *testing.T) {
	t.Run("deserializes perPage field", func(t *testing.T) {
		jsonData := `{"perPage": 20, "total": 100}`

		var meta Meta
		err := json.Unmarshal([]byte(jsonData), &meta)
		require.NoError(t, err)

		assert.Equal(t, 20, meta.PerPage, "should parse perPage field")
		assert.Equal(t, int64(100), meta.Total.Int, "should parse total field")
		assert.True(t, meta.Total.Valid, "total should be valid")
	})

	t.Run("deserializes per_page field", func(t *testing.T) {
		jsonData := `{"per_page": 25, "total": 150}`

		var meta Meta
		err := json.Unmarshal([]byte(jsonData), &meta)
		require.NoError(t, err)

		assert.Equal(t, 25, meta.PerPage, "should parse per_page field into PerPage")
		assert.Equal(t, int64(150), meta.Total.Int, "should parse total field")
		assert.True(t, meta.Total.Valid, "total should be valid")
	})

	t.Run("prioritizes perPage over per_page when both are present", func(t *testing.T) {
		jsonData := `{"perPage": 30, "per_page": 25, "total": 200}`

		var meta Meta
		err := json.Unmarshal([]byte(jsonData), &meta)
		require.NoError(t, err)

		assert.Equal(t, 30, meta.PerPage, "should prioritize perPage over per_page")
		assert.Equal(t, int64(200), meta.Total.Int, "should parse total field")
		assert.True(t, meta.Total.Valid, "total should be valid")
	})

	t.Run("uses per_page when perPage is not present", func(t *testing.T) {
		jsonData := `{"per_page": 35, "total": 250, "next": "cursor123"}`

		var meta Meta
		err := json.Unmarshal([]byte(jsonData), &meta)
		require.NoError(t, err)

		assert.Equal(t, 35, meta.PerPage, "should use per_page when perPage is not present")
		assert.Equal(t, int64(250), meta.Total.Int, "should parse total field")
		assert.True(t, meta.Total.Valid, "total should be valid")
		assert.Equal(t, "cursor123", meta.Next.Str, "should parse next field")
		assert.True(t, meta.Next.Valid, "next should be valid")
	})

	t.Run("handles empty JSON", func(t *testing.T) {
		jsonData := `{}`

		var meta Meta
		err := json.Unmarshal([]byte(jsonData), &meta)
		require.NoError(t, err)

		assert.Equal(t, 0, meta.PerPage, "PerPage should be zero for empty JSON")
		assert.False(t, meta.Total.Valid, "Total should be invalid for empty JSON")
		assert.False(t, meta.Next.Valid, "Next should be invalid for empty JSON")
	})
}
