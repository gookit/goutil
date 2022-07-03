// Package clipboard provide a simple clipboard read and write function.
package clipboard

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/strutil"
)

// Clipboard struct
type Clipboard struct {
	buf *bytes.Buffer

	// TODO add event on write, read
	readerBin string
	writerBin string
	// add slashes. eg: '\' -> '\\'
	addSlashes bool
}

// New instance
func New() *Clipboard {
	return &Clipboard{
		readerBin: GetReaderBin(),
		writerBin: GetWriterBin(),
	}
}

// WithSlashes for the contents
func (c *Clipboard) WithSlashes() *Clipboard {
	c.addSlashes = true
	return c
}

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

// Write bytes data to clipboard
func (c *Clipboard) Write(p []byte) (int, error) {
	return c.WriteString(string(p))
}

// WriteString data to clipboard
func (c *Clipboard) WriteString(s string) (int, error) {
	if c.addSlashes {
		s = strutil.AddSlashes(s)
	}
	return c.buffer().WriteString(s)
}

// Flush buffer contents to clipboard
func (c *Clipboard) Flush() error {
	if c.buf == nil || c.buf.Len() == 0 {
		return errorx.Raw("not write contents")
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
	file, err := fsutil.OpenReadFile(filepath)
	if err != nil {
		return err
	}

	// eg:
	// 	Mac: pbcopy < tempfile.txt
	return c.doExec(c.writerBin, "<", file.Name())
}

// Read contents from clipboard
func (c *Clipboard) Read() ([]byte, error) {
	return exec.Command(c.readerBin).Output()
}

// ReadString contents as string from clipboard
func (c *Clipboard) ReadString() (string, error) {
	bts, err := c.Read()
	return string(bts), err
}

// ReadToFile dump clipboard data to file
func (c *Clipboard) ReadToFile(filepath string) error {
	// eg:
	// 	Mac: pbpaste >> tasklist.txt
	return c.doExec(c.readerBin, ">>", filepath)

	// cmd := exec.Command(c.readerBin)
	// file, err := fsutil.QuickOpenFile(filepath)
	// if err != nil {
	// 	return err
	// }
	// cmd.Stdout = file
	// return cmd.Run()
}

func (c *Clipboard) buffer() *bytes.Buffer {
	if c.buf == nil {
		c.buf = new(bytes.Buffer)
	}
	return c.buf
}

func (c *Clipboard) doExec(binName string, args ...string) error {
	return exec.Command(binName, args...).Run()
}

func (c *Clipboard) checkBin(binName string) error {
	_, err := exec.LookPath(binName)
	return err
}
