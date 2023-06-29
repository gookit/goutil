package structs_test

import (
	"testing"

	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/testutil/assert"
)

func TestNewWrapper(t *testing.T) {
	type User struct {
		Name string
		Age  int
		City string
	}

	u := &User{Age: 23, Name: "inhere"}
	w := structs.Wrap(u)

	assert.Equal(t, "inhere", w.Get("Name"))
	assert.Equal(t, 23, w.Get("Age"))

	assert.NoErr(t, w.Set("Age", 129))
	assert.Equal(t, 129, w.Get("Age"))
	assert.Equal(t, 129, u.Age)

	assert.NoErr(t, w.Set("City", "CD"))
	assert.Equal(t, "CD", w.Get("City"))
	assert.Equal(t, "CD", u.City)

	assert.Nil(t, w.Get("NotExists"))
	assert.Panics(t, func() {
		structs.NewWrapper(nil)
	})
}
