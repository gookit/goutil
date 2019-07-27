package sysutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentShell(t *testing.T) {
	path := CurrentShell(true)
	assert.NotEmpty(t, path)

	if path != "" {
		path = CurrentShell(false)
		assert.NotEmpty(t, path)
	}
}
