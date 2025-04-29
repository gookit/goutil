package ccolor_test

import (
	"testing"

	"github.com/gookit/goutil/x/ccolor"
)

func TestColorPrint(t *testing.T) {
	// code gen by: kite gen parse ccolor/_demo/gen-code.tpl
	ccolor.Redp("p:red color message, ")
	ccolor.Redf("f:%s color message, ", "red")
	ccolor.Redln("ln:red color message print in cli.")
	ccolor.Bluep("p:blue color message, ")
	ccolor.Bluef("f:%s color message, ", "blue")
	ccolor.Blueln("ln:blue color message print in cli.")
	ccolor.Cyanp("p:cyan color message, ")
	ccolor.Cyanf("f:%s color message, ", "cyan")
	ccolor.Cyanln("ln:cyan color message print in cli.")
	ccolor.Grayp("p:gray color message, ")
	ccolor.Grayf("f:%s color message, ", "gray")
	ccolor.Grayln("ln:gray color message print in cli.")
	ccolor.Greenp("p:green color message, ")
	ccolor.Greenf("f:%s color message, ", "green")
	ccolor.Greenln("ln:green color message print in cli.")
	ccolor.Yellowp("p:yellow color message, ")
	ccolor.Yellowf("f:%s color message, ", "yellow")
	ccolor.Yellowln("ln:yellow color message print in cli.")
	ccolor.Magentap("p:magenta color message, ")
	ccolor.Magentaf("f:%s color message, ", "magenta")
	ccolor.Magentaln("ln:magenta color message print in cli.")

	ccolor.Infop("p:info color message, ")
	ccolor.Infof("f:%s color message, ", "info")
	ccolor.Infoln("ln:info color message print in cli.")
	ccolor.Successp("p:success color message, ")
	ccolor.Successf("f:%s color message, ", "success")
	ccolor.Successln("ln:success color message print in cli.")
	ccolor.Warnp("p:warn color message, ")
	ccolor.Warnf("f:%s color message, ", "warn")
	ccolor.Warnln("ln:warn color message print in cli.")
	ccolor.Errorp("p:error color message, ")
	ccolor.Errorf("f:%s color message, ", "error")
	ccolor.Errorln("ln:error color message print in cli.")
}
