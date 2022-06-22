package cflag

import (
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/strutil"
)

// OptCheckFn define
type OptCheckFn func(val interface{}) error

// FlagOpt struct
type FlagOpt struct {
	// Shortcuts short names. eg: ["o", "a"]
	Shortcuts []string
	// Required option
	Required bool
	// Validator for check option value
	Validator OptCheckFn
}

// HelpName string
func (o *FlagOpt) HelpName(name string) string {
	return AddPrefixes(name, o.Shortcuts)
}

// FlagArg struct
type FlagArg struct {
	// Value for the flag argument
	*structs.Value
	// defaultVal string
	defaultVal string
	// Name of the argument
	Name string
	// Desc arg description
	Desc string
	// Index of the argument
	Index int
	// Required argument
	Required bool
	// Validator for check value
	Validator func(val string) error
}

// NewArg create instance
func NewArg(name, desc string, required bool) *FlagArg {
	return &FlagArg{Name: name, Desc: desc, Required: required}
}

// check arg config and init
func (a *FlagArg) check() error {
	if a.Name == "" {
		return errorx.Rawf("cflag: arg#%d name cannot be empty", a.Index)
	}

	if a.Required && a.V != nil {
		return errorx.Rawf("cflag: cannot set default value for 'required' arg: %s", a.Name)
	}

	if a.Desc == "" {
		a.Desc = "no description"
	}

	a.defaultVal = a.String()
	return nil
}

// HelpDesc string
func (a *FlagArg) HelpDesc() string {
	desc := strutil.UpperFirst(a.Desc)
	if a.Required {
		desc = "<red>*</>" + desc
	}

	if a.defaultVal != "" {
		desc += "(Default:" + a.defaultVal + ")"
	}
	return desc
}
