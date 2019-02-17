package fmtutil

import "fmt"

// DataSize format bytes number friendly.
// Usage:
// 	file, err := os.Open(path)
// 	fl, err := file.Stat()
// 	fmtSize := DataSize(fl.Size())
func DataSize(bytes uint64) string {
	switch {
	case bytes < 1024:
		return fmt.Sprintf("%dB", bytes)
	case bytes < 1024*1024:
		return fmt.Sprintf("%.2fK", float64(bytes)/1024)
	case bytes < 1024*1024*1024:
		return fmt.Sprintf("%.2fM", float64(bytes)/1024/1024)
	default:
		return fmt.Sprintf("%.2fG", float64(bytes)/1024/1024/1024)
	}
}
