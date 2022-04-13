package timex_test

import (
	"testing"
	"time"

	"github.com/gookit/goutil/timex"
	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	sec := timex.NowUnix()

	assert.NotEmpty(t, timex.FormatUnix(sec))
	assert.NotEmpty(t, timex.FormatUnixBy(sec, time.RFC3339))
}
