package fsutil

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// Mkdir alias of os.MkdirAll()
func Mkdir(dirPath string, perm os.FileMode) error {
	return os.MkdirAll(dirPath, perm)
}

// MkParentDir quick create parent dir
func MkParentDir(fpath string) error {
	dirPath := filepath.Dir(fpath)
	if !IsDir(dirPath) {
		return os.MkdirAll(dirPath, 0775)
	}
	return nil
}

// DiscardReader anything from the reader
func DiscardReader(src io.Reader) {
	_, _ = io.Copy(ioutil.Discard, src)
}

// MustReadFile read file contents, will panic on error
func MustReadFile(filePath string) []byte {
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return bs
}

// MustReadReader read contents from io.Reader, will panic on error
func MustReadReader(r io.Reader) []byte {
	// TODO go 1.16+ bs, err := io.ReadAll(r)
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return bs
}

// GetContents read contents from path or io.Reader, will panic on error
func GetContents(in interface{}) []byte {
	if fPath, ok := in.(string); ok {
		return MustReadFile(fPath)
	}

	if r, ok := in.(io.Reader); ok {
		return MustReadReader(r)
	}

	panic("invalid type of input")
}

// ReadExistFile read file contents if existed, will panic on error
func ReadExistFile(filePath string) []byte {
	if IsFile(filePath) {
		bs, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		return bs
	}
	return nil
}

// ************************************************************
//	open/create files
// ************************************************************

// OpenFile like os.OpenFile, but will auto create dir.
func OpenFile(filepath string, flag int, perm os.FileMode) (*os.File, error) {
	fileDir := path.Dir(filepath)

	// if err := os.Mkdir(dir, 0775); err != nil {
	if err := os.MkdirAll(fileDir, DefaultDirPerm); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(filepath, flag, perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

/* TODO MustOpenFile() */

// QuickOpenFile like os.OpenFile, open for write, if not exists, will create it.
func QuickOpenFile(filepath string) (*os.File, error) {
	return OpenFile(filepath, DefaultFileFlags, DefaultFilePerm)
}

// OpenReadFile like os.OpenFile, open file for read contents
func OpenReadFile(filepath string) (*os.File, error) {
	return os.OpenFile(filepath, OnlyReadFileFlags, OnlyReadFilePerm)
}

// CreateFile create file if not exists
//
// Usage:
// 	CreateFile("path/to/file.txt", 0664, 0666)
func CreateFile(fpath string, filePerm, dirPerm os.FileMode) (*os.File, error) {
	dirPath := path.Dir(fpath)
	if !IsDir(dirPath) {
		err := os.MkdirAll(dirPath, dirPerm)
		if err != nil {
			return nil, err
		}
	}

	return os.OpenFile(fpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, filePerm)
}

// MustCreateFile create file, will panic on error
func MustCreateFile(filePath string, filePerm, dirPerm os.FileMode) *os.File {
	file, err := CreateFile(filePath, filePerm, dirPerm)
	if err != nil {
		panic(err)
	}
	return file
}

// ************************************************************
//	write, copy files
// ************************************************************

// PutContents to file
func PutContents(filePath string, contents string) (int, error) {
	// create and open file
	dstFile, err := QuickOpenFile(filePath)
	if err != nil {
		return 0, err
	}

	defer dstFile.Close()
	return dstFile.WriteString(contents)
}

// CopyFile copy a file to another file path.
func CopyFile(srcPath string, dstPath string) error {
	srcFile, err := os.OpenFile(srcPath, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// create and open file
	dstFile, err := QuickOpenFile(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

// MustCopyFile copy file to another path.
func MustCopyFile(srcPath string, dstPath string) {
	err := CopyFile(srcPath, dstPath)
	if err != nil {
		panic(err)
	}
}

// ************************************************************
//	remove files
// ************************************************************

// alias methods
var (
	// MustRm removes the named file or (empty) directory.
	MustRm = MustRemove
	// QuietRm removes the named file or (empty) directory.
	QuietRm = QuietRemove
)

// Remove removes the named file or (empty) directory.
func Remove(fPath string) error {
	return os.Remove(fPath)
}

// MustRemove removes the named file or (empty) directory.
// NOTICE: will panic on error
func MustRemove(fPath string) {
	if err := os.Remove(fPath); err != nil {
		panic(err)
	}
}

// QuietRemove removes the named file or (empty) directory.
// NOTICE: will ignore error
func QuietRemove(fPath string) {
	_ = os.Remove(fPath)
}

// RmIfExist removes the named file or (empty) directory on exists.
func RmIfExist(fPath string) error { return DeleteIfExist(fPath) }

// DeleteIfExist removes the named file or (empty) directory on exists.
func DeleteIfExist(fPath string) error {
	if PathExists(fPath) {
		return os.Remove(fPath)
	}
	return nil
}

// RmFileIfExist removes the named file on exists.
func RmFileIfExist(fPath string) error { return DeleteIfFileExist(fPath) }

// DeleteIfFileExist removes the named file on exists.
func DeleteIfFileExist(fPath string) error {
	if IsFile(fPath) {
		return os.Remove(fPath)
	}
	return nil
}

// ************************************************************
//	other operates
// ************************************************************

// Unzip a zip archive
// from https://blog.csdn.net/wangshubo1989/article/details/71743374
func Unzip(archive, targetDir string) (err error) {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(targetDir, DefaultDirPerm); err != nil {
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
			_ = fileReader.Close()
			return err
		}

		_, err = io.Copy(targetFile, fileReader)

		// close all
		_ = fileReader.Close()
		targetFile.Close()

		if err != nil {
			return err
		}
	}

	return
}
