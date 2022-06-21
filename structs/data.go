package structs

import (
	"sync"

	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/strutil"
)

// MapDataStore struct
type MapDataStore struct {
	sync.RWMutex
	enableLock bool
	// data store
	data map[string]interface{}
}

// NewMapData create
func NewMapData() *MapDataStore {
	return &MapDataStore{
		data: make(map[string]interface{}),
	}
}

// EnableLock for operate data
func (md *MapDataStore) EnableLock() *MapDataStore {
	md.enableLock = true
	return md
}

// Data get all
func (md *MapDataStore) Data() map[string]interface{} {
	return md.data
}

// SetData set all data
func (md *MapDataStore) SetData(data map[string]interface{}) {
	if !md.enableLock {
		md.data = data
		return
	}

	md.RLock()
	md.data = data
	md.RUnlock()
}

// Set value to data
func (md *MapDataStore) Set(key string, val interface{}) {
	md.SetValue(key, val)
}

// SetValue to data
func (md *MapDataStore) SetValue(key string, val interface{}) {
	if md.enableLock {
		md.Lock()
		defer md.Unlock()
	}

	md.data[key] = val
}

// Len of data
func (md *MapDataStore) Len() int {
	return len(md.data)
}

// Clear all data
func (md *MapDataStore) Clear() {
	md.data = make(map[string]interface{})
}

// Value get from data
func (md *MapDataStore) Value(key string) (val interface{}, ok bool) {
	if md.enableLock {
		md.RLock()
		defer md.RUnlock()
	}

	val, ok = md.data[key]
	return
}

// Get val from data
func (md *MapDataStore) Get(key string) interface{} {
	return md.GetVal(key)
}

// GetVal get from data
func (md *MapDataStore) GetVal(key string) interface{} {
	if md.enableLock {
		md.RLock()
		defer md.RUnlock()
	}

	return md.data[key]
}

// StrVal get from data
func (md *MapDataStore) StrVal(key string) string {
	return strutil.QuietString(md.GetVal(key))
}

// IntVal get from data
func (md *MapDataStore) IntVal(key string) int {
	return mathutil.QuietInt(md.GetVal(key))
}

// BoolVal get from data
func (md *MapDataStore) BoolVal(key string) bool {
	val, ok := md.Value(key)
	if !ok {
		return false
	}

	if bol, ok := val.(bool); ok {
		return bol
	}

	if str, ok := val.(string); ok {
		return strutil.QuietBool(str)
	}
	return false
}

// String format data
func (md *MapDataStore) String() string {
	return maputil.ToString(md.data)
}
