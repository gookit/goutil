package cliutil

import "github.com/gookit/color"

/*************************************************************
 * quick use color print message
 *************************************************************/

// Redp print message with Red color
func Redp(a ...interface{}) { color.Red.Print(a...) }

// Redf print message with Red color
func Redf(format string, a ...interface{}) { color.Red.Printf(format, a...) }

// Redln print message line with Red color
func Redln(a ...interface{}) { color.Red.Println(a...) }

// Bluep print message with Blue color
func Bluep(a ...interface{}) { color.Blue.Print(a...) }

// Bluef print message with Blue color
func Bluef(format string, a ...interface{}) { color.Blue.Printf(format, a...) }

// Blueln print message line with Blue color
func Blueln(a ...interface{}) { color.Blue.Println(a...) }

// Cyanp print message with Cyan color
func Cyanp(a ...interface{}) { color.Cyan.Print(a...) }

// Cyanf print message with Cyan color
func Cyanf(format string, a ...interface{}) { color.Cyan.Printf(format, a...) }

// Cyanln print message line with Cyan color
func Cyanln(a ...interface{}) { color.Cyan.Println(a...) }

// Grayp print message with gray color
func Grayp(a ...interface{}) { color.Gray.Print(a...) }

// Grayf print message with gray color
func Grayf(format string, a ...interface{}) { color.Gray.Printf(format, a...) }

// Grayln print message line with gray color
func Grayln(a ...interface{}) { color.Gray.Println(a...) }

// Greenp print message with green color
func Greenp(a ...interface{}) { color.Green.Print(a...) }

// Greenf print message with green color
func Greenf(format string, a ...interface{}) { color.Green.Printf(format, a...) }

// Greenln print message line with green color
func Greenln(a ...interface{}) { color.Green.Println(a...) }

// Yellowp print message with yellow color
func Yellowp(a ...interface{}) { color.Yellow.Print(a...) }

// Yellowf print message with yellow color
func Yellowf(format string, a ...interface{}) { color.Yellow.Printf(format, a...) }

// Yellowln print message line with yellow color
func Yellowln(a ...interface{}) { color.Yellow.Println(a...) }

// Magentap print message with magenta color
func Magentap(a ...interface{}) { color.Magenta.Print(a...) }

// Magentaf print message with magenta color
func Magentaf(format string, a ...interface{}) { color.Magenta.Printf(format, a...) }

// Magentaln print message line with magenta color
func Magentaln(a ...interface{}) { color.Magenta.Println(a...) }

/*************************************************************
 * quick use style print message
 *************************************************************/

// Infop print message with info color
func Infop(a ...interface{}) { color.Info.Print(a...) }

// Infof print message with info style
func Infof(format string, a ...interface{}) { color.Info.Printf(format, a...) }

// Infoln print message with info style
func Infoln(a ...interface{}) { color.Info.Println(a...) }

// Successp print message with success color
func Successp(a ...interface{}) { color.Success.Print(a...) }

// Successf print message with success style
func Successf(format string, a ...interface{}) { color.Success.Printf(format, a...) }

// Successln print message with success style
func Successln(a ...interface{}) { color.Success.Println(a...) }

// Errorp print message with error color
func Errorp(a ...interface{}) { color.Error.Print(a...) }

// Errorf print message with error style
func Errorf(format string, a ...interface{}) { color.Error.Printf(format, a...) }

// Errorln print message with error style
func Errorln(a ...interface{}) { color.Error.Println(a...) }

// Warnp print message with warn color
func Warnp(a ...interface{}) { color.Warn.Print(a...) }

// Warnf print message with warn style
func Warnf(format string, a ...interface{}) { color.Warn.Printf(format, a...) }

// Warnln print message with warn style
func Warnln(a ...interface{}) { color.Warn.Println(a...) }
