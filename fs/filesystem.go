package fs

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path"
	"path/filepath"
)

// FileExists reports whether the named file or directory exists.
func FileExists(path string) bool {
	if path == "" {
		return false
	}

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// IsAbsPath is abs path.
func IsAbsPath(filename string) bool {
	return path.IsAbs(filename)
}

// IsZipFile check is zip file
// from https://blog.csdn.net/wangshubo1989/article/details/71743374
func IsZipFile(filepath string) bool {
	f, err := os.Open(filepath)
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}

	return bytes.Equal(buf, []byte("PK\x03\x04"))
}

// Unzip a zip archive
// from https://blog.csdn.net/wangshubo1989/article/details/71743374
func Unzip(archive, targetDir string) (err error) {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return
	}

	if err = os.MkdirAll(targetDir, 0755); err != nil {
		return
	}

	for _, file := range reader.File {
		fullPath := filepath.Join(targetDir, file.Name)
		if file.FileInfo().IsDir() {
			err = os.MkdirAll(fullPath, file.Mode())
			if err != nil {
				return err
			}

			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}

		targetFile, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			fileReader.Close()
			return err
		}

		_, err = io.Copy(targetFile, fileReader)

		// close all
		fileReader.Close()
		targetFile.Close()

		if err != nil {
			return err
		}
	}

	return
}
