package timex_test

import (
	"testing"

	"github.com/gookit/goutil/timex"
	"github.com/stretchr/testify/assert"
)

func TestTimeX_basic(t *testing.T) {
	tx := timex.Now()
	assert.NotEmpty(t, tx.String())
	assert.NotEmpty(t, tx.Datetime())
}
