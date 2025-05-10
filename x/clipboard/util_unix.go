//go:build !windows && !darwin

package clipboard

import "os"

// GetWriterBin program name
func GetWriterBin() string {
	return WriterOnLin
}

// GetReaderBin program name
func GetReaderBin() string {
	return ReaderOnLin
}

func available() bool {
	// X clipboard is unavailable when not under X.
	return os.Getenv("DISPLAY") != ""
}
