package structs

import (
	"sync"

	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/strutil"
)

// DataStore struct TODO
type DataStore struct {
	sync.RWMutex
	enableLock bool
	// data store
	data map[string]interface{}
}

// NewMapData create
func NewMapData() *DataStore {
	return &DataStore{
		data: make(map[string]interface{}),
	}
}

// EnableLock for operate data
func (d *DataStore) EnableLock() *DataStore {
	d.enableLock = true
	return d
}

// Data get all
func (d *DataStore) Data() map[string]interface{} {
	return d.data
}

// SetData set all data
func (d *DataStore) SetData(data map[string]interface{}) {
	if !d.enableLock {
		d.data = data
		return
	}

	d.RLock()
	d.data = data
	d.RUnlock()
}

// Set value to data
func (d *DataStore) Set(key string, val interface{}) {
	d.SetValue(key, val)
}

// SetValue to data
func (d *DataStore) SetValue(key string, val interface{}) {
	if d.enableLock {
		d.Lock()
		defer d.Unlock()
	}

	d.data[key] = val
}

// Len of data
func (d *DataStore) Len() int {
	return len(d.data)
}

// Reset all data
func (d *DataStore) Reset() {
	d.data = make(map[string]interface{})
}

// Value get from data
func (d *DataStore) Value(key string) (val interface{}, ok bool) {
	if d.enableLock {
		d.RLock()
		defer d.RUnlock()
	}

	val, ok = d.data[key]
	return
}

// Get val from data
func (d *DataStore) Get(key string) interface{} {
	return d.GetVal(key)
}

// GetVal get from data
func (d *DataStore) GetVal(key string) interface{} {
	if d.enableLock {
		d.RLock()
		defer d.RUnlock()
	}

	return d.data[key]
}

// StrVal get from data
func (d *DataStore) StrVal(key string) string {
	return strutil.QuietString(d.GetVal(key))
}

// IntVal get from data
func (d *DataStore) IntVal(key string) int {
	return mathutil.QuietInt(d.GetVal(key))
}

// BoolVal get from data
func (d *DataStore) BoolVal(key string) bool {
	val, ok := d.Value(key)
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
func (d *DataStore) String() string {
	return maputil.ToString(d.data)
}
