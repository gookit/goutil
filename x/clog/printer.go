package clog

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gookit/goutil/x/ccolor"
	"github.com/gookit/goutil/x/fmtutil"
)

type WriteFn func(level string, data map[string]string)

// Printer log printer for console.
type Printer struct {
	Template   string
	TimeFormat string
	// Output log output. default: os.Stdout
	Output io.Writer
	// OnWriteFn log write callback. see formatOutput
	OnWriteFn WriteFn
}

// NewPrinter create a new printer
func NewPrinter(output io.Writer) *Printer {
	return &Printer{
		Template:   DefaultTemplate,
		Output:     output,
		TimeFormat: "15:04:05.000",
	}
}

// Configure the printer
func (p *Printer) Configure(fns ...func(p *Printer)) {
	for _, fn := range fns {
		fn(p)
	}
}

// SetTemplate sets a custom template
func (p *Printer) SetTemplate(template string) { p.Template = template }

// Print logs a message with the specified level
func (p *Printer) Print(level string, v ...any) {
	logStr := p.formatOutput(level, fmtutil.ArgsWithSpaces(v))
	ccolor.Fprint(p.Output, logStr, "\n")
}

// Println logs a message with the specified level. alias of Print
func (p *Printer) Println(level string, v ...any) { p.Print(level, v...) }

// Printf logs a message with the specified level and format
func (p *Printer) Printf(level, format string, v ...any) {
	logStr := p.formatOutput(level, fmt.Sprintf(format, v...))
	ccolor.Fprint(p.Output, logStr, "\n")
}

// formatOutput formats the output based on template
func (p *Printer) formatOutput(level, message string) string {
	data := map[string]string{
		"time":    "",
		"level":   level,
		"message": message,
		"emoji":   getLevelEmoji(level),
	}
	if p.TimeFormat != "" {
		data["time"] = time.Now().Format(p.TimeFormat)
	}

	// fire onWrite hook
	if p.OnWriteFn != nil {
		p.OnWriteFn(level, data)
	}

	message = wrapColor(level, message)
	levelStr := formatLevel(level)
	if ln := len(levelStr); ln < 5 {
		levelStr += strings.Repeat(" ", 5-ln)
	} else if ln > 5 {
		levelStr = levelStr[:5]
	}
	levelStr = wrapColor(level, levelStr)

	// Replace placeholders in template
	return strings.NewReplacer(
		"{level}", levelStr, "{message}", message,
		"{time}", data["time"], "{emoji}", data["emoji"],
	).Replace(p.Template)
}

// Debug logs a debug message
func (p *Printer) Debug(v ...any) { p.Print(DebugLevel, v...) }

// Debugf logs a debug message with format
func (p *Printer) Debugf(format string, v ...any) { p.Printf(DebugLevel, format, v...) }

// Info logs an info message
func (p *Printer) Info(v ...any) { p.Print(InfoLevel, v...) }

// Infof logs an info message with format
func (p *Printer) Infof(format string, v ...any) { p.Printf(InfoLevel, format, v...) }

// Warn logs a warning message
func (p *Printer) Warn(v ...any) { p.Print(WarnLevel, v...) }

// Warnf logs a warning message with format
func (p *Printer) Warnf(format string, v ...any) { p.Printf(WarnLevel, format, v...) }

// Error logs an error message
func (p *Printer) Error(v ...any) { p.Print(ErrorLevel, v...) }

// Errorf logs an error message with format
func (p *Printer) Errorf(format string, v ...any) { p.Printf(ErrorLevel, format, v...) }

// ErrorT logs an error type value
// func (p *Printer) ErrorT(err error) {
// 	if err != nil {
// 		p.Print(ErrorLevel, fmt.Sprint(err))
// 	}
// }

// Fatal logs a fatal message
func (p *Printer) Fatal(v ...any) { p.Print(FatalLevel, v...) }

// Fatalf logs a fatal message with format
func (p *Printer) Fatalf(format string, v ...any) { p.Printf(FatalLevel, format, v...) }

// Trace logs a trace message
func (p *Printer) Trace(v ...any) { p.Print(TraceLevel, v...) }

// Tracef logs a trace message with format
func (p *Printer) Tracef(format string, v ...any) { p.Printf(TraceLevel, format, v...) }

// Success logs a success message
func (p *Printer) Success(v ...any) { p.Print(SuccessLevel, v...) }

// Successf logs a success message with format
func (p *Printer) Successf(format string, v ...any) { p.Printf(SuccessLevel, format, v...) }
