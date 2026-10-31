package fsutil

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/gookit/goutil/x/basefn"
)

// ************************************************************
//	temp file or dir
// ************************************************************

// OSTempFile create a temp file on os.TempDir()
//
// Usage:
//
//	fsutil.OSTempFile("example.*.txt")
func OSTempFile(pattern string) (*os.File, error) {
	return os.CreateTemp(os.TempDir(), pattern)
}

// TempFile is like os.CreateTemp, but can custom temp dir.
//
// Usage:
//
//	// create temp file on os.TempDir()
//	fsutil.TempFile("", "example.*.txt")
//	// create temp file on "testdata" dir
//	fsutil.TempFile("testdata", "example.*.txt")
func TempFile(dir, pattern string) (*os.File, error) {
	return os.CreateTemp(dir, pattern)
}

// OSTempDir creates a new temp dir on os.TempDir and return the temp dir path
//
// Usage:
//
//	fsutil.OSTempDir("example.*")
func OSTempDir(pattern string) (string, error) {
	return os.MkdirTemp(os.TempDir(), pattern)
}

// TempDir creates a new temp dir and return the temp dir path
//
// Usage:
//
//	fsutil.TempDir("", "example.*")
//	fsutil.TempDir("testdata", "example.*")
func TempDir(dir, pattern string) (string, error) {
	return os.MkdirTemp(dir, pattern)
}

// ************************************************************
//	write, copy files
// ************************************************************

// MustSave create file and write contents to file, panic on error.
//
//   - data type allow: string, []byte, io.Reader
//
// default option see NewOpenOption()
func MustSave(filePath string, data any, optFns ...OpenOptionFunc) {
	basefn.MustOK(SaveFile(filePath, data, optFns...))
}

// SaveFile create file and write contents to file. will auto create dir.
//
//   - data type allow: string, []byte, io.Reader
//
// default option see NewOpenOption()
func SaveFile(filePath string, data any, optFns ...OpenOptionFunc) error {
	opt := NewOpenOption(optFns...)
	return WriteFile(filePath, data, opt.Perm, opt.Flag)
}

// WriteData Quick write any data to file, alias of PutContents
func WriteData(filePath string, data any, fileFlag ...int) (int, error) {
	return PutContents(filePath, data, fileFlag...)
}

// PutContents create file and write contents to file at once. Will auto create dir
//
// data type allows: string, []byte, io.Reader
//
// Tip: file flag default is FsCWTFlags (override write)
//
// Usage:
//
//	fsutil.PutContents(filePath, contents, fsutil.FsCWAFlags) // append write
//	fsutil.Must2(fsutil.PutContents(filePath, contents)) // panic on error
func PutContents(filePath string, data any, fileFlag ...int) (int, error) {
	f, err := QuickOpenFile(filePath, basefn.FirstOr(fileFlag, FsCWTFlags))
	if err != nil {
		return 0, err
	}
	return WriteOSFile(f, data)
}

// WriteFile create file and write contents to file, can set perm for a file.
//
// data type allows: string, []byte, io.Reader
//
// Tip: file flag default is FsCWTFlags (override write)
//
// Usage:
//
//	fsutil.WriteFile(filePath, contents, fsutil.DefaultFilePerm, fsutil.FsCWAFlags)
func WriteFile(filePath string, data any, perm os.FileMode, fileFlag ...int) error {
	flag := basefn.FirstOr(fileFlag, FsCWTFlags)
	f, err := OpenFile(filePath, flag, perm)
	if err != nil {
		return err
	}

	_, err = WriteOSFile(f, data)
	return err
}

// WriteOSFile write data to give os.File, then close file.
//
// data type allows: string, []byte, io.Reader
func WriteOSFile(f *os.File, data any) (n int, err error) {
	switch typData := data.(type) {
	case []byte:
		n, err = f.Write(typData)
	case string:
		n, err = f.WriteString(typData)
	case io.Reader: // eg: buffer
		var n64 int64
		n64, err = io.Copy(f, typData)
		n = int(n64)
	default:
		_ = f.Close()
		panic("WriteFile: data type only allow: []byte, string, io.Reader")
	}

	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return n, err
}

// CopyFile copy a file to another file path.
// It returns an error if both opened paths refer to the same file.
func CopyFile(srcPath, dstPath string) error {
	srcFile, err := os.OpenFile(srcPath, FsRFlags, 0)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	srcInfo, err := srcFile.Stat()
	if err != nil {
		return err
	}

	// Open without truncating so identity is checked against the opened file.
	dstFile, err := QuickOpenFile(dstPath, FsCWTFlags&^os.O_TRUNC)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	dstInfo, err := dstFile.Stat()
	if err != nil {
		return err
	}
	if os.SameFile(srcInfo, dstInfo) {
		return fmt.Errorf("fsutil: source and destination refer to the same file: %s", srcPath)
	}
	if dstInfo.Mode().IsRegular() {
		if err := dstFile.Truncate(0); err != nil {
			return err
		}
	}

	_, err = io.Copy(dstFile, srcFile)
	return err
}

// MustCopyFile copy file to another path.
func MustCopyFile(srcPath, dstPath string) {
	err := CopyFile(srcPath, dstPath)
	if err != nil {
		panic(err)
	}
}

// UpdateContents read the whole contents of the file, call handleFn(contents) to handle it,
// then write the updated contents back to the same file.
//
// Notes:
//   - the file must already exist, it will not be created
//   - the write is in-place and not atomic, a crash in the middle can leave a half updated file
//   - the returned contents will replace the file contents: a shorter result is truncated,
//     so returning nil or empty contents will empty the file
func UpdateContents(filePath string, handleFn func(bs []byte) []byte) error {
	if handleFn == nil {
		return errors.New("fsutil: UpdateContents: the handleFn is nil")
	}

	// tip: open without O_TRUNC, otherwise the original contents are lost before it can be read
	osFile, err := os.OpenFile(filePath, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer osFile.Close()

	bs, err := io.ReadAll(osFile)
	if err != nil {
		return err
	}

	// the file offset is at EOF after reading, so write from offset 0 explicitly.
	bs = handleFn(bs)
	if _, err = osFile.WriteAt(bs, 0); err != nil {
		return err
	}

	// truncate to the new length, avoid a stale suffix when the new contents are shorter
	return osFile.Truncate(int64(len(bs)))
}

// CreateSymlink creates a symbolic link
func CreateSymlink(target, linkPath string) error {
	// Check if the link already exists
	if IsFile(linkPath) {
		// Remove existing link/file
		if err := os.Remove(linkPath); err != nil {
			return fmt.Errorf("failed to remove existing symlink: %w", err)
		}
	}

	return os.Symlink(target, linkPath)
}
