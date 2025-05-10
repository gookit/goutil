//go:build darwin

package clipboard

// GetWriterBin program name
func GetWriterBin() string {
	return WriterOnMac
}

// GetReaderBin program name
func GetReaderBin() string {
	return ReaderOnMac
}

func available() bool { return true }
