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

	u1 := User{Age: 34, Name: "tom"}
	w1 := structs.Wrap(u1)
	// get value
	assert.Nil(t, w1.Get("NotExists"))
	assert.Eq(t, "tom", w1.Get("Name"))
	assert.ErrSubMsg(t, w1.Set("NotExists", "val"), "field NotExists not found")
	assert.ErrSubMsg(t, w1.Set("Name", "john"), "can not set value for field: Name")

	assert.Panics(t, func() {
		structs.NewWriter(u1)
	})
	assert.Panics(t, func() {
		structs.NewWriter("invalid")
	})
}
