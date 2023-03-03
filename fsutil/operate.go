package fsutil

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gookit/goutil/basefn"
)

// Mkdir alias of os.MkdirAll()
func Mkdir(dirPath string, perm os.FileMode) error {
	return os.MkdirAll(dirPath, perm)
}

// MkDirs batch make multi dirs at once
func MkDirs(perm os.FileMode, dirPaths ...string) error {
	for _, dirPath := range dirPaths {
		if err := os.MkdirAll(dirPath, perm); err != nil {
			return err
		}
	}
	return nil
}

// MkSubDirs batch make multi sub-dirs at once
func MkSubDirs(perm os.FileMode, parentDir string, subDirs ...string) error {
	for _, dirName := range subDirs {
		dirPath := parentDir + "/" + dirName
		if err := os.MkdirAll(dirPath, perm); err != nil {
			return err
		}
	}
	return nil
}

// MkParentDir quick create parent dir
func MkParentDir(fpath string) error {
	dirPath := filepath.Dir(fpath)
	if !IsDir(dirPath) {
		return os.MkdirAll(dirPath, 0775)
	}
	return nil
}

// ************************************************************
//	open/create files
// ************************************************************

// some flag consts for open file
const (
	FsCWAFlags = os.O_CREATE | os.O_WRONLY | os.O_APPEND // create, append write-only
	FsCWTFlags = os.O_CREATE | os.O_WRONLY | os.O_TRUNC  // create, override write-only
	FsCWFlags  = os.O_CREATE | os.O_WRONLY               // create, write-only
	FsRFlags   = os.O_RDONLY                             // read-only
)

// OpenFile like os.OpenFile, but will auto create dir.
func OpenFile(filepath string, flag int, perm os.FileMode) (*os.File, error) {
	fileDir := path.Dir(filepath)
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

// QuickOpenFile like os.OpenFile, open for append write. if not exists, will create it.
func QuickOpenFile(filepath string, fileFlag ...int) (*os.File, error) {
	flag := basefn.FirstOr(fileFlag, FsCWAFlags)
	return OpenFile(filepath, flag, DefaultFilePerm)
}

// OpenAppendFile like os.OpenFile, open for append write. if not exists, will create it.
func OpenAppendFile(filepath string) (*os.File, error) {
	return OpenFile(filepath, FsCWAFlags, DefaultFilePerm)
}

// OpenTruncFile like os.OpenFile, open for override write. if not exists, will create it.
func OpenTruncFile(filepath string) (*os.File, error) {
	return OpenFile(filepath, FsCWTFlags, DefaultFilePerm)
}

// OpenReadFile like os.OpenFile, open file for read contents
func OpenReadFile(filepath string) (*os.File, error) {
	return os.OpenFile(filepath, FsRFlags, OnlyReadFilePerm)
}

// CreateFile create file if not exists
//
// Usage:
//
//	CreateFile("path/to/file.txt", 0664, 0666)
func CreateFile(fpath string, filePerm, dirPerm os.FileMode, fileFlag ...int) (*os.File, error) {
	dirPath := path.Dir(fpath)
	if !IsDir(dirPath) {
		err := os.MkdirAll(dirPath, dirPerm)
		if err != nil {
			return nil, err
		}
	}

	flag := basefn.FirstOr(fileFlag, FsCWAFlags)
	return os.OpenFile(fpath, flag, filePerm)
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
//
// NOTICE: will ignore error
func QuietRemove(fPath string) { _ = os.Remove(fPath) }

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
		if strings.Contains(file.Name, "..") {
			return fmt.Errorf("illegal file path in zip: %v", file.Name)
		}

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
