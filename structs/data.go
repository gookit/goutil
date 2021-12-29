package structs

import "sync"

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
		data: map[string]interface{}{},
	}
}

// EnableLock for data
func (md *MapDataStore) EnableLock() {
	md.enableLock = true
}

// Data get all
func (md *MapDataStore) Data() map[string]interface{} {
	return md.data
}

// SetData set all data
func (md *MapDataStore) SetData(data map[string]interface{}) {
	md.data = data
}

// Value get from data
func (md *MapDataStore) Value(key string) interface{} {
	if md.enableLock {
		md.RLock()
		defer md.RUnlock()
	}

	return md.data[key]
}

// SetValue to data
func (md *MapDataStore) SetValue(key string, val interface{}) {
	if md.enableLock {
		md.Lock()
		defer md.Unlock()
	}

	if md.data == nil {
		md.data = make(map[string]interface{})
	}
	md.data[key] = val
}

// ClearData all data
func (md *MapDataStore) ClearData() {
	md.data = nil
}
