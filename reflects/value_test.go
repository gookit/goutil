package reflects_test

import (
	"reflect"
	"testing"

	"github.com/gookit/goutil/reflects"
	"github.com/stretchr/testify/assert"
)

func TestValueOf(t *testing.T) {
	rv := reflects.ValueOf(int64(23))

	assert.Equal(t, reflect.Int64, rv.Kind())
	assert.Equal(t, reflects.Int, rv.BaseKind())
}
