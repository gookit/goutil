package cflag_test

import (
	"testing"

	"github.com/gookit/goutil/cliutil/cflag"
	"github.com/stretchr/testify/assert"
)

func TestAddPrefix(t *testing.T) {
	assert.Equal(t, "-a", cflag.AddPrefix("a"))
	assert.Equal(t, "--long", cflag.AddPrefix("long"))

	assert.Equal(t, "--long", cflag.AddPrefixes("long", nil))
}
