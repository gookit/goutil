package structs

import "github.com/gookit/goutil/internal/comfunc"

// ToMap simple convert structs to map by reflect
func ToMap(st interface{}) map[string]interface{} {
	mp, _ := comfunc.TryStructToMap(st)
	return mp
}

// TryToMap simple convert structs to map by reflect
func TryToMap(st interface{}) (map[string]interface{}, error) {
	return comfunc.TryStructToMap(st)
}

// MustToMap alis of TryToMap, but will panic on error
func MustToMap(st interface{}) map[string]interface{} {
	mp, err := comfunc.TryStructToMap(st)
	if err != nil {
		panic(err)
	}
	return mp
}
