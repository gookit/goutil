package cliutil

import (
	"github.com/gookit/goutil/x/ccolor"
)

/*************************************************************
 * quick use color print message
 *************************************************************/

// Redp print message with Red color
func Redp(a ...any) { ccolor.Red.Print(a...) }

// Redf print message with Red color
func Redf(format string, a ...any) { ccolor.Red.Printf(format, a...) }

// Redln print message line with Red color
func Redln(a ...any) { ccolor.Red.Println(a...) }

// Bluep print message with Blue color
func Bluep(a ...any) { ccolor.Blue.Print(a...) }

// Bluef print message with Blue color
func Bluef(format string, a ...any) { ccolor.Blue.Printf(format, a...) }

// Blueln print message line with Blue color
func Blueln(a ...any) { ccolor.Blue.Println(a...) }

// Cyanp print message with Cyan color
func Cyanp(a ...any) { ccolor.Cyan.Print(a...) }

// Cyanf print message with Cyan color
func Cyanf(format string, a ...any) { ccolor.Cyan.Printf(format, a...) }

// Cyanln print message line with Cyan color
func Cyanln(a ...any) { ccolor.Cyan.Println(a...) }

// Grayp print message with gray color
func Grayp(a ...any) { ccolor.Gray.Print(a...) }

// Grayf print message with gray color
func Grayf(format string, a ...any) { ccolor.Gray.Printf(format, a...) }

// Grayln print message line with gray color
func Grayln(a ...any) { ccolor.Gray.Println(a...) }

// Greenp print message with green color
func Greenp(a ...any) { ccolor.Green.Print(a...) }

// Greenf print message with green color
func Greenf(format string, a ...any) { ccolor.Green.Printf(format, a...) }

// Greenln print message line with green color
func Greenln(a ...any) { ccolor.Green.Println(a...) }

// Yellowp print message with yellow color
func Yellowp(a ...any) { ccolor.Yellow.Print(a...) }

// Yellowf print message with yellow color
func Yellowf(format string, a ...any) { ccolor.Yellow.Printf(format, a...) }

// Yellowln print message line with yellow color
func Yellowln(a ...any) { ccolor.Yellow.Println(a...) }

// Magentap print message with magenta color
func Magentap(a ...any) { ccolor.Magenta.Print(a...) }

// Magentaf print message with magenta color
func Magentaf(format string, a ...any) { ccolor.Magenta.Printf(format, a...) }

// Magentaln print message line with magenta color
func Magentaln(a ...any) { ccolor.Magenta.Println(a...) }

/*************************************************************
 * quick use style print message
 *************************************************************/

// Infop print message with info color
func Infop(a ...any) { ccolor.Info.Print(a...) }

// Infof print message with info style
func Infof(format string, a ...any) { ccolor.Info.Printf(format, a...) }

// Infoln print message with info style
func Infoln(a ...any) { ccolor.Info.Println(a...) }

// Successp print message with success color
func Successp(a ...any) { ccolor.Success.Print(a...) }

// Successf print message with success style
func Successf(format string, a ...any) { ccolor.Success.Printf(format, a...) }

// Successln print message with success style
func Successln(a ...any) { ccolor.Success.Println(a...) }

// Errorp print message with error color
func Errorp(a ...any) { ccolor.Error.Print(a...) }

// Errorf print message with error style
func Errorf(format string, a ...any) { ccolor.Error.Printf(format, a...) }

// Errorln print message with error style
func Errorln(a ...any) { ccolor.Error.Println(a...) }

// Warnp print message with warn color
func Warnp(a ...any) { ccolor.Warn.Print(a...) }

// Warnf print message with warn style
func Warnf(format string, a ...any) { ccolor.Warn.Printf(format, a...) }

// Warnln print message with warn style
func Warnln(a ...any) { ccolor.Warn.Println(a...) }
