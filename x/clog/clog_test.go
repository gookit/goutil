package clog_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gookit/goutil/x/clog"
)

func Example() {
	// 使用默认全局logger
	clog.Info("Application started")
	clog.Warn("This is a warning message")
	clog.Error("An error occurred")
	clog.Success("Operation completed successfully")

	// 使用自定义模板
	clog.SetTemplate("{emoji} [{time}}] {message}")
	clog.Debug("Debug information")
}

func ExampleNewPrinter() {
	// 创建自定义printer
	printer := clog.NewPrinter(os.Stderr)
	printer.SetTemplate("[CUSTOM] {level}} - {message}")
	printer.Info("Custom printer message")
}

func TestStd_usage(t *testing.T) {
	clog.Info("Application started")
	clog.Debug("Debug information")
	clog.Warn("This is a warning message")
	clog.Error("An error occurred")
	clog.Fatal("Fatal error message")
	clog.Trace("Trace message")
	clog.Success("Operation completed successfully")
	clog.Print("custom", "custom level message")
	clog.Println("custom", "custom level message")

	fmt.Println("---- Custom the printer ----")
	// 使用自定义模板
	clog.SetOutput(os.Stdout)
	clog.SetTemplate("{time} | {emoji} {message}")
	clog.Configure(func(p *clog.Printer) {
		p.TimeFormat = "2006-01-02T15:04:05.000"
	})
	clog.SetOnWrite(func(level string, data map[string]string) {
		// nothing
	})

	clog.Debugf("Debug information %s", "arg1")
	clog.Infof("This is a %s message", "info")
	clog.Warnf("This is a %s message", "warning")
	clog.Errorf("An error occurred %s", "arg2")
	clog.Fatalf("Fatal error message %s", "arg3")
	clog.Tracef("Trace message %s", "arg4")
	clog.Successf("Operation completed successfully %s", "arg5")
	clog.Printf("custom", "custom level message %s", "arg6")
}
