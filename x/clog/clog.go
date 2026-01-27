package clog

import (
	"io"
	"os"
	"strings"

	"github.com/gookit/goutil/x/ccolor"
)

const (
	SimpleTemplate = `{emoji} [{level}] | {message}`
	DefaultTemplate = `{time} [{level}] | {emoji} {message}`
)

const (
	DebugLevel = "debug"
	InfoLevel = "info"
	WarnLevel = "warn"
	ErrorLevel = "error"
	FatalLevel = "fatal"
	TraceLevel = "trace"
	SuccessLevel = "success"
)

// LevelColorMap 定义日志级别对应的颜色
var LevelColorMap = map[string]string{
	DebugLevel: "cyan",
	InfoLevel: "blue",
	WarnLevel: "yellow",
	ErrorLevel: "red",
	FatalLevel: "red",
	TraceLevel: "gray",
	SuccessLevel: "green",
}

// LevelEmojiMap 定义日志级别对应的 emoji ⚠️💡
var LevelEmojiMap = map[string]string{
	DebugLevel: "🐛",
	InfoLevel: "ℹ️",
	WarnLevel: "💡",
	ErrorLevel: "❌",
	FatalLevel: "🚨",
	TraceLevel: "🔍",
	SuccessLevel: "🎉",
}

func wrapColor(level, s string) string {
	if color, ok := LevelColorMap[level]; ok {
		return ccolor.WrapTag(s, color)
	}
	return s
}

// formatLevel formats the level string
func formatLevel(level string) string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	case TraceLevel:
		return "TRACE"
	case SuccessLevel:
		return "SUCCESS"
	default:
		return strings.ToUpper(level)
	}
}

// getLevelEmoji returns the emoji for the given level
func getLevelEmoji(level string) string {
	if emoji, ok := LevelEmojiMap[level]; ok {
		return emoji
	}
	return "📝" // default emoji
}

var std = NewPrinter(os.Stdout)

// Configure the standard log printer
func Configure(optFns ...func(p *Printer)) { std.Configure(optFns...) }

// SetOnWrite sets a custom write function for the standard logger
func SetOnWrite(fn WriteFn) { std.OnWriteFn = fn }

// SetOutput sets the output for the standard logger
func SetOutput(w io.Writer) { std.Output = w }

// SetTemplate sets a custom template for the standard logger
func SetTemplate(template string) { std.SetTemplate(template) }

// Print logs a message with the specified level using the standard logger
func Print(level string, v ...any) { std.Print(level, v...) }

// Println logs a message with the specified level using the standard logger
func Println(level string, v ...any) { std.Println(level, v...) }

// Printf logs a message with the specified level and format using the standard logger
func Printf(level, format string, v ...any) { std.Printf(level, format, v...) }

// Debug logs a debug message using the standard logger
func Debug(v ...any) { std.Debug(v...) }

// Debugf logs a debug message with format using the standard logger
func Debugf(format string, v ...any) { std.Debugf(format, v...) }

// Info logs an info message using the standard logger
func Info(v ...any) { std.Info(v...) }

// Infof logs an info message with format using the standard logger
func Infof(format string, v ...any) { std.Infof(format, v...) }

// Warn logs a warning message using the standard logger
func Warn(v ...any) { std.Warn(v...) }

// Warnf logs a warning message with format using the standard logger
func Warnf(format string, v ...any) { std.Warnf(format, v...) }

// Error logs a message using the standard logger
func Error(v ...any) { std.Error(v...) }

// Errorf logs a message with format using the standard logger
func Errorf(format string, v ...any) { std.Errorf(format, v...) }

// Fatal logs a fatal message using the standard logger
func Fatal(v ...any) { std.Fatal(v...) }

// Fatalf logs a fatal message with format using the standard logger
func Fatalf(format string, v ...any) { std.Fatalf(format, v...) }

// Trace logs a trace message using the standard logger
func Trace(v ...any) { std.Trace(v...) }

// Tracef logs a trace message with format using the standard logger
func Tracef(format string, v ...any) { std.Tracef(format, v...) }

// Success logs a success message using the standard logger
func Success(v ...any) { std.Success(v...) }

// Successf logs a success message with format using the standard logger
func Successf(format string, v ...any) { std.Successf(format, v...) }
