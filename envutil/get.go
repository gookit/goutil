package envutil

import "os"

// GetEnv get ENV value
func GetEnv(name string, def ...string) string {
	val := os.Getenv(name)
	if val == "" && len(def) > 0 {
		val = def[0]
	}

	return val
}
