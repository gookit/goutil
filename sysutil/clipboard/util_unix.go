//go:build !windows && !darwin
// +build !windows,!darwin

package clipboard

// GetWriterBin program name
func GetWriterBin() string {
	return WriterOnLin
}

// GetReaderBin program name
func GetReaderBin() string {
	return ReaderOnLin
}
