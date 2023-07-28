// Package clipboard provide a simple clipboard read and write operations.
package clipboard

import (
	"bytes"
	"errors"
	"io"
	"os/exec"
	"strings"

	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/sysutil"
)

// Clipboard struct
type Clipboard struct {
	// TODO add event on write, read
	// buffer for write
	buf *bytes.Buffer

	// print exec command line on run
	verbose bool
	// available - bin file exist on the OS.
	writeable, readable bool

	readerBin string
	readArgs  []string
	writerBin string
	writeArgs []string
}

// New instance
func New() *Clipboard {
	// special handle on with args
	reader, readArgs := parseLine(GetReaderBin())
	writer, writeArgs := parseLine(GetWriterBin())

	return &Clipboard{
		readerBin: reader,
		readArgs:  readArgs,
		writerBin: writer,
		writeArgs: writeArgs,
		readable:  sysutil.HasExecutable(reader),
		writeable: sysutil.HasExecutable(writer),
	}
}

// SetReader for handle clip
// func (c *Clipboard) SetReader(line string)  {
// }

// SetWriter for handle clip
// func (c *Clipboard) SetWriter(line string)  {
// }

// WithVerbose setting
func (c *Clipboard) WithVerbose(yn bool) *Clipboard {
	c.verbose = yn
	return c
}

// Clean the clipboard
func (c *Clipboard) Clean() error { return c.Reset() }

// Reset and clean the clipboard
func (c *Clipboard) Reset() error {
	if c.buf != nil {
		c.buf.Reset()
	}

	// echo empty string for clean clipboard.
	// run: echo '' | pbcopy
	return c.WriteFrom(strings.NewReader(""))
}

//
// ---------------------------------------- write ----------------------------------------
//

// Write bytes data to buffer. should call Flush() to write to clipboard
func (c *Clipboard) Write(p []byte) (int, error) {
	return c.WriteString(string(p))
}

// WriteString data to buffer. should call Flush() to write to clipboard
func (c *Clipboard) WriteString(s string) (int, error) {
	// if c.addSlashes {
	// 	s = strutil.AddSlashes(s)
	// }
	return c.buffer().WriteString(s)
}

// Flush buffer contents to clipboard
func (c *Clipboard) Flush() error {
	if c.buf == nil || c.buf.Len() == 0 {
		return errors.New("clipboard: empty contents for write")
	}

	defer c.buf.Reset()
	return c.WriteFrom(c.buf)
}

// WriteFromFile contents to clipboard
func (c *Clipboard) WriteFromFile(filepath string) error {
	file, err := fsutil.OpenReadFile(filepath)
	if err != nil {
		return err
	}

	defer file.Close()
	return c.WriteFrom(file)
}

// WriteFrom reader data to clipboard
func (c *Clipboard) WriteFrom(r io.Reader) error {
	if !c.writeable {
		return errorx.Rawf("clipboard: write driver %q not found on OS", c.writerBin)
	}

	cmd := exec.Command(c.writerBin, c.writeArgs...)
	cmd.Stdin = r

	if c.verbose {
		cliutil.Yellowf("clipboard> %s\n", cliutil.BuildLine(c.writerBin, c.writeArgs))
	}
	return cmd.Run()
}

//
// ---------------------------------------- read ----------------------------------------
//

// Read bytes contents from clipboard
func (c *Clipboard) Read() ([]byte, error) {
	buf, err := c.ReadToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ReadToBuffer read clipboard contents to new buffer.
func (c *Clipboard) ReadToBuffer() (*bytes.Buffer, error) {
	var buf bytes.Buffer
	if err := c.ReadTo(&buf); err != nil {
		return nil, err
	}
	return &buf, nil
}

// SafeString read contents as string from clipboard, will return empty string on error
func (c *Clipboard) SafeString() string {
	s, err := c.ReadString()
	if err != nil {
		return ""
	}
	return s
}

// ReadString contents as string from clipboard
func (c *Clipboard) ReadString() (string, error) {
	bts, err := c.Read()
	if err != nil {
		return "", err
	}

	// fix: at Windows will always return end of the "\r\n"
	if sysutil.IsWindows() {
		return string(bytes.TrimRight(bts, "\r\n")), nil
	}
	return string(bts), nil
}

// ReadToFile dump clipboard data to file
func (c *Clipboard) ReadToFile(filepath string) error {
	file, err := fsutil.QuickOpenFile(filepath)
	if err != nil {
		return err
	}

	defer file.Close()
	return c.ReadTo(file)
}

// ReadTo read clipboard contents to writer
func (c *Clipboard) ReadTo(w io.Writer) error {
	if !c.readable {
		return errorx.Rawf("clipboard: read driver %q not found on OS", c.readerBin)
	}

	cmd := exec.Command(c.readerBin, c.readArgs...)
	cmd.Stdout = w

	if c.verbose {
		cliutil.Yellowf("clipboard> %s\n", cliutil.BuildLine(c.readerBin, c.readArgs))
	}
	return cmd.Run()
}

//
// ---------------------------------------- help ----------------------------------------
//

// Available check
func (c *Clipboard) Available() bool {
	return c.writeable && c.readable && available()
}

// Writeable check
func (c *Clipboard) Writeable() bool {
	return c.writeable
}

// Readable check
func (c *Clipboard) Readable() bool {
	return c.readable
}

func (c *Clipboard) buffer() *bytes.Buffer {
	if c.buf == nil {
		c.buf = new(bytes.Buffer)
	}
	return c.buf
}
