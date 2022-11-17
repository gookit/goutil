// Package clipboard provide a simple clipboard read and write operations.
package clipboard

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/sysutil"
)

// Clipboard struct
type Clipboard struct {
	// TODO add event on write, read
	buf *bytes.Buffer

	// available - bin file exist on the OS.
	available bool
	readerBin string
	writerBin string
	// add slashes. eg: '\' -> '\\'
	// addSlashes bool
	readArgs []string
}

// New instance
func New() *Clipboard {
	var readArgs []string

	// special handle on Windows
	reader := GetReaderBin()
	if strings.Contains(reader, " ") {
		args := strings.Split(reader, " ")
		reader, readArgs = args[0], args[1:]
	}

	return &Clipboard{
		readerBin: reader,
		readArgs:  readArgs,
		writerBin: GetWriterBin(),
		available: sysutil.HasExecutable(reader),
	}
}

// WithSlashes for the contents
// func (c *Clipboard) WithSlashes() *Clipboard {
// 	c.addSlashes = true
// 	return c
// }

// Reset and clean the clipboard
func (c *Clipboard) Reset() error {
	if c.buf != nil {
		c.buf.Reset()
	}

	// run: echo '' | pbcopy
	// echo empty string for clean clipboard.
	cmd := exec.Command(c.writerBin)
	cmd.Stdin = strings.NewReader("")
	return cmd.Run()
}

//
// -------------------- write --------------------
//

// Write bytes data to clipboard
func (c *Clipboard) Write(p []byte) (int, error) {
	return c.WriteString(string(p))
}

// WriteString data to clipboard
func (c *Clipboard) WriteString(s string) (int, error) {
	// if c.addSlashes {
	// 	s = strutil.AddSlashes(s)
	// }
	return c.buffer().WriteString(s)
}

// Flush buffer contents to clipboard
func (c *Clipboard) Flush() error {
	if c.buf == nil || c.buf.Len() == 0 {
		return errors.New("not write contents")
	}

	// linux:
	//   # Copy input to clipboard
	// 	 echo -n "$input" | xclip -selection c
	// Mac:
	//   echo hello | pbcopy
	//   pbcopy < tempfile.txt
	cmd := exec.Command(c.writerBin)
	cmd.Stdin = c.buf

	defer c.buf.Reset()
	return cmd.Run()
}

// WriteFromFile contents to clipboard
func (c *Clipboard) WriteFromFile(filepath string) error {
	// eg:
	// 	Mac: pbcopy < tempfile.txt
	// return exec.Command(c.writerBin, "<", filepath).Run()
	file, err := fsutil.OpenReadFile(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	cmd := exec.Command(c.writerBin)
	cmd.Stdin = file

	return cmd.Run()
}

//
// -------------------- read --------------------
//

// Read contents from clipboard
func (c *Clipboard) Read() ([]byte, error) {
	return exec.Command(c.readerBin, c.readArgs...).Output()
}

// ReadString contents as string from clipboard
func (c *Clipboard) ReadString() (string, error) {
	bts, err := c.Read()
	if err != nil {
		return "", err
	}

	// fix: at Windows will always return end of the "\r\n"
	if sysutil.IsWindows() {
		return strings.TrimRight(string(bts), "\r\n"), err
	}
	return string(bts), err
}

// ReadToFile dump clipboard data to file
func (c *Clipboard) ReadToFile(filepath string) error {
	// eg:
	// 	Mac: pbpaste >> tasklist.txt
	// return exec.Command(c.readerBin, ">>", filepath).Run()
	file, err := fsutil.QuickOpenFile(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	cmd := exec.Command(c.readerBin, c.readArgs...)
	cmd.Stdout = file

	return cmd.Run()
}

// Available check
func (c *Clipboard) Available() bool {
	return c.available
}

func (c *Clipboard) buffer() *bytes.Buffer {
	if c.buf == nil {
		c.buf = new(bytes.Buffer)
	}
	return c.buf
}
