package cliutil

import "github.com/gookit/color"

/*************************************************************
 * quick use color print message
 *************************************************************/

// Redp print message with Red color
func Redp(a ...any) { color.Red.Print(a...) }

// Redf print message with Red color
func Redf(format string, a ...any) { color.Red.Printf(format, a...) }

// Redln print message line with Red color
func Redln(a ...any) { color.Red.Println(a...) }

// Bluep print message with Blue color
func Bluep(a ...any) { color.Blue.Print(a...) }

// Bluef print message with Blue color
func Bluef(format string, a ...any) { color.Blue.Printf(format, a...) }

// Blueln print message line with Blue color
func Blueln(a ...any) { color.Blue.Println(a...) }

// Cyanp print message with Cyan color
func Cyanp(a ...any) { color.Cyan.Print(a...) }

// Cyanf print message with Cyan color
func Cyanf(format string, a ...any) { color.Cyan.Printf(format, a...) }

// Cyanln print message line with Cyan color
func Cyanln(a ...any) { color.Cyan.Println(a...) }

// Grayp print message with gray color
func Grayp(a ...any) { color.Gray.Print(a...) }

// Grayf print message with gray color
func Grayf(format string, a ...any) { color.Gray.Printf(format, a...) }

// Grayln print message line with gray color
func Grayln(a ...any) { color.Gray.Println(a...) }

// Greenp print message with green color
func Greenp(a ...any) { color.Green.Print(a...) }

// Greenf print message with green color
func Greenf(format string, a ...any) { color.Green.Printf(format, a...) }

// Greenln print message line with green color
func Greenln(a ...any) { color.Green.Println(a...) }

// Yellowp print message with yellow color
func Yellowp(a ...any) { color.Yellow.Print(a...) }

// Yellowf print message with yellow color
func Yellowf(format string, a ...any) { color.Yellow.Printf(format, a...) }

// Yellowln print message line with yellow color
func Yellowln(a ...any) { color.Yellow.Println(a...) }

// Magentap print message with magenta color
func Magentap(a ...any) { color.Magenta.Print(a...) }

// Magentaf print message with magenta color
func Magentaf(format string, a ...any) { color.Magenta.Printf(format, a...) }

// Magentaln print message line with magenta color
func Magentaln(a ...any) { color.Magenta.Println(a...) }

/*************************************************************
 * quick use style print message
 *************************************************************/

// Infop print message with info color
func Infop(a ...any) { color.Info.Print(a...) }

// Infof print message with info style
func Infof(format string, a ...any) { color.Info.Printf(format, a...) }

// Infoln print message with info style
func Infoln(a ...any) { color.Info.Println(a...) }

// Successp print message with success color
func Successp(a ...any) { color.Success.Print(a...) }

// Successf print message with success style
func Successf(format string, a ...any) { color.Success.Printf(format, a...) }

// Successln print message with success style
func Successln(a ...any) { color.Success.Println(a...) }

// Errorp print message with error color
func Errorp(a ...any) { color.Error.Print(a...) }

// Errorf print message with error style
func Errorf(format string, a ...any) { color.Error.Printf(format, a...) }

// Errorln print message with error style
func Errorln(a ...any) { color.Error.Println(a...) }

// Warnp print message with warn color
func Warnp(a ...any) { color.Warn.Print(a...) }

// Warnf print message with warn style
func Warnf(format string, a ...any) { color.Warn.Printf(format, a...) }

// Warnln print message with warn style
func Warnln(a ...any) { color.Warn.Println(a...) }
