package maputil

// Data an map data type
type Data map[string]interface{}

// Get value from the data map
func (d Data) Get(key string) interface{} {
	return d[key]
}

// Set value to the data map
func (d Data) Set(key string, val interface{}) {
	d[key] = val
}

// Has value on the data map
func (d Data) Has(key string) bool {
	_, ok := d[key]
	return ok
}

// Default get value from the data map with default value
func (d Data) Default(key string, def interface{}) interface{} {
	val, ok := d[key]
	if ok {
		return val
	}

	return def
}
