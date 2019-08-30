package mathutil_test

import (
	"testing"
	"time"

	"github.com/gookit/goutil/mathutil"
	"github.com/stretchr/testify/assert"
)

func TestPercent(t *testing.T) {
	assert.Equal(t, float64(34), mathutil.Percent(34, 100))
	assert.Equal(t, float64(0), mathutil.Percent(34, 0))
	assert.Equal(t, float64(-100), mathutil.Percent(34, -34))
}

func TestElapsedTime(t *testing.T) {
	ct := time.Now()
	time.Sleep(time.Second * 1)
	assert.Contains(t, mathutil.ElapsedTime(ct), "100")
}

func TestDataSize(t *testing.T) {
	assert.Equal(t, "3.38K", mathutil.DataSize(3456))
}

func TestHowLongAgo(t *testing.T) {
	assert.Equal(t, "57 mins", mathutil.HowLongAgo(3456))
}
