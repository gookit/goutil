package fmtutil_test

import (
	"bytes"
	"testing"

	"github.com/gookit/goutil/fmtutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestStringOrJSON(t *testing.T) {
	t.Run("string input", func(t *testing.T) {
		input := "Hello, world!"
		expected := []byte(input)
		actual, err := fmtutil.StringOrJSON(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !bytes.Equal(expected, actual) {
			t.Errorf("expected %q, but got %q", expected, actual)
		}
	})

	t.Run("JSON input", func(t *testing.T) {
		input := map[string]any{
			"foo": "bar",
			"baz": 123,
		}

		actual, err := fmtutil.StringOrJSON(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assert.StrContains(t, string(actual), `"foo": "bar"`)
		assert.StrContains(t, string(actual), `"baz": 123`)
	})

	t.Run("invalid JSON input", func(t *testing.T) {
		input := make(chan int) // channel is not JSON-serializable
		_, err := fmtutil.StringOrJSON(input)
		if err == nil {
			t.Error("expected error, but got nil")
		}
	})
}
