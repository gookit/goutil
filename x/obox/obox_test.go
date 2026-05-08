package obox

import (
	"sync"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
)

type testService struct {
	name string
}

func TestMain(m *testing.M) {
	m.Run()
	Reset()
}

func TestSetAndGet(t *testing.T) {
	Reset()

	svc := &testService{name: "test"}
	Set("svc", svc)

	got, err := Get[*testService]("svc")
	assert.NoErr(t, err)
	assert.Eq(t, "test", got.name)
}

func TestGetNotFound(t *testing.T) {
	Reset()

	_, err := Get[*testService]("not-exist")
	assert.Err(t, err)
	assert.Eq(t, ErrNotFound, err)
}

func TestGetTypeMismatch(t *testing.T) {
	Reset()

	Set("svc", &testService{name: "test"})

	_, err := Get[string]("svc")
	assert.Err(t, err)
	assert.Eq(t, ErrTypeMismatch, err)
}

func TestMustGet(t *testing.T) {
	Reset()

	svc := &testService{name: "test"}
	Set("svc", svc)

	got := MustGet[*testService]("svc")
	assert.Eq(t, "test", got.name)
}

func TestMustGetPanic(t *testing.T) {
	Reset()

	assert.Panics(t, func() {
		MustGet[*testService]("not-exist")
	})
}

func TestMustGetTypeMismatchPanic(t *testing.T) {
	Reset()

	Set("svc", &testService{name: "test"})

	assert.Panics(t, func() {
		MustGet[string]("svc")
	})
}

func TestSingleton(t *testing.T) {
	Reset()

	svc := &testService{name: "original"}
	Set("svc", svc)

	got1, err := Get[*testService]("svc")
	assert.NoErr(t, err)

	got2, err := Get[*testService]("svc")
	assert.NoErr(t, err)

	assert.True(t, got1 == got2, "singleton should return same instance")
}

func TestTransient(t *testing.T) {
	Reset()

	svc := &testService{name: "transient"}
	Set("svc", svc, Transient())

	got1, err := Get[*testService]("svc")
	assert.NoErr(t, err)

	got2, err := Get[*testService]("svc")
	assert.NoErr(t, err)

	assert.True(t, got1 == got2, "transient without factory returns same value")
}

func TestTransientWithFactory(t *testing.T) {
	Reset()

	callCount := 0
	factory := func() any {
		callCount++
		return &testService{name: "transient"}
	}

	Set[*testService]("svc", nil, TransientWithFactory(factory))

	got1, err := Get[*testService]("svc")
	assert.NoErr(t, err)
	assert.Eq(t, 1, callCount)

	got2, err := Get[*testService]("svc")
	assert.NoErr(t, err)
	assert.Eq(t, 2, callCount)

	assert.True(t, got1 != got2, "transient should return different instances")
}

func TestLazy(t *testing.T) {
	Reset()

	callCount := 0
	factory := func() any {
		callCount++
		return &testService{name: "lazy"}
	}

	Set[*testService]("svc", nil, Lazy(factory))

	assert.Eq(t, 0, callCount, "factory should not be called on Set")

	got1, err := Get[*testService]("svc")
	assert.NoErr(t, err)
	assert.Eq(t, 1, callCount, "factory should be called on first Get")
	assert.Eq(t, "lazy", got1.name)

	got2, err := Get[*testService]("svc")
	assert.NoErr(t, err)
	assert.Eq(t, 1, callCount, "factory should not be called again")
	assert.True(t, got1 == got2, "lazy should return same instance")
}

func TestHas(t *testing.T) {
	Reset()

	assert.False(t, Has("svc"))

	Set("svc", &testService{})
	assert.True(t, Has("svc"))
}

func TestDelete(t *testing.T) {
	Reset()

	Set("svc", &testService{})
	assert.True(t, Has("svc"))

	Delete("svc")
	assert.False(t, Has("svc"))
}

func TestReset(t *testing.T) {
	Reset()

	Set("svc1", &testService{})
	Set("svc2", &testService{})

	assert.True(t, Has("svc1"))
	assert.True(t, Has("svc2"))

	Reset()

	assert.False(t, Has("svc1"))
	assert.False(t, Has("svc2"))
}

func TestConcurrentAccess(t *testing.T) {
	Reset()

	var wg sync.WaitGroup
	count := 100

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			name := string(rune('a' + n%26))
			Set(name, &testService{name: name})
		}(i)
	}

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			name := string(rune('a' + n%26))
			_, _ = Get[*testService](name)
		}(i)
	}

	wg.Wait()
}

func TestNamedRegistration(t *testing.T) {
	Reset()

	svc1 := &testService{name: "svc1"}
	svc2 := &testService{name: "svc2"}

	Set("db1", svc1)
	Set("db2", svc2)

	got1, err := Get[*testService]("db1")
	assert.NoErr(t, err)
	assert.Eq(t, "svc1", got1.name)

	got2, err := Get[*testService]("db2")
	assert.NoErr(t, err)
	assert.Eq(t, "svc2", got2.name)

	assert.True(t, got1 != got2)
}
