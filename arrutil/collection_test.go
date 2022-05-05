package arrutil_test

import (
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/stretchr/testify/assert"
)

func TestContainsItemStringEqualComparer(t *testing.T) {
	data := []string{
		"a",
		"b",
		"c",
	}
	assert.True(t, arrutil.ContainsItem(data, "a", arrutil.StringEqualCompareFunc))
	assert.False(t, arrutil.ContainsItem(data, "d", arrutil.StringEqualCompareFunc))
}

func TestContainsItemEqualComparer(t *testing.T) {
	ptrVal := 3
	data := []interface{}{
		1,
		2,
		&ptrVal,
	}
	assert.True(t, arrutil.ContainsItem(data, 1, arrutil.EqualCompareFunc))
	assert.False(t, arrutil.ContainsItem(data, 3, arrutil.EqualCompareFunc))
	assert.False(t, arrutil.ContainsItem(data, 4, arrutil.EqualCompareFunc))
}

func TestContainsItemElementTypeComparer(t *testing.T) {
	ptrVal := 3
	data := []interface{}{
		1,
		2,
		&ptrVal,
	}
	assert.True(t, arrutil.ContainsItem(data, 1, arrutil.ElemTypeEqualCompareFunc))
	assert.True(t, arrutil.ContainsItem(data, 3, arrutil.ElemTypeEqualCompareFunc))
	assert.False(t, arrutil.ContainsItem(data, "a", arrutil.ElemTypeEqualCompareFunc))
}
