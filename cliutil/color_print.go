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

// Grayp print message with Gray color
func Grayp(a ...interface{}) { color.Gray.Print(a...) }

// Grayf print message with Gray color
func Grayf(format string, a ...interface{}) { color.Gray.Printf(format, a...) }

// Grayln print message line with Gray color
func Grayln(a ...interface{}) { color.Gray.Println(a...) }

// Greenp print message with Green color
func Greenp(a ...interface{}) { color.Green.Print(a...) }

// Greenf print message with Green color
func Greenf(format string, a ...interface{}) { color.Green.Printf(format, a...) }

// Greenln print message line with Green color
func Greenln(a ...interface{}) { color.Green.Println(a...) }

// Yellowp print message with Yellow color
func Yellowp(a ...interface{}) { color.Yellow.Print(a...) }

// Yellowf print message with Yellow color
func Yellowf(format string, a ...interface{}) { color.Yellow.Printf(format, a...) }

// Yellowln print message line with Yellow color
func Yellowln(a ...interface{}) { color.Yellow.Println(a...) }

// Magentap print message with Magenta color
func Magentap(a ...interface{}) { color.Magenta.Print(a...) }

// Magentaf print message with Magenta color
func Magentaf(format string, a ...interface{}) { color.Magenta.Printf(format, a...) }

// Magentaln print message line with Magenta color
func Magentaln(a ...interface{}) { color.Magenta.Println(a...) }

/*************************************************************
 * quick use style print message
 *************************************************************/

// Infop print message with Info color
func Infop(a ...interface{}) { color.Info.Print(a...) }

// Infof print message with Info style
func Infof(format string, a ...interface{}) { color.Info.Printf(format, a...) }

// Infoln print message with Info style
func Infoln(a ...interface{}) { color.Info.Println(a...) }

// Errorp print message with Error color
func Errorp(a ...interface{}) { color.Error.Print(a...) }

// Errorf print message with Error style
func Errorf(format string, a ...interface{}) { color.Error.Printf(format, a...) }

// Errorln print message with Error style
func Errorln(a ...interface{}) { color.Error.Println(a...) }

// Warnp print message with Warn color
func Warnp(a ...interface{}) { color.Warn.Print(a...) }

// Warnf print message with Warn style
func Warnf(format string, a ...interface{}) { color.Warn.Printf(format, a...) }

// Warnln print message with Warn style
func Warnln(a ...interface{}) { color.Warn.Println(a...) }
