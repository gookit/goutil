package dump

import (
	"io"
	"os"
)

// Options for dumper
type Options struct {
	// Output the output writer
	Output io.Writer
	// NoType don't show data type TODO
	NoType bool
	// NoColor don't with color
	NoColor bool
	// IndentLen width. default is 2
	IndentLen int
	// IndentChar default is one space
	IndentChar byte
	// MaxDepth for nested print
	MaxDepth int
	// ShowFlag for display caller position
	ShowFlag int
	// CallerSkip skip for call runtime.Caller()
	CallerSkip int
	// ColorTheme for print result.
	ColorTheme Theme
	// SkipNilField value dump on map, struct.
	SkipNilField bool
	// SkipPrivate field dump on struct.
	SkipPrivate bool
	// BytesAsString dump handle.
	BytesAsString bool
	// MoreLenNL array/slice elements length > MoreLenNL, will wrap new line
	// MoreLenNL int
}

// OptionFunc type
type OptionFunc func(opts *Options)

// NewDefaultOptions create.
func NewDefaultOptions(out io.Writer, skip int) *Options {
	if out == nil {
		out = os.Stdout
	}

	return &Options{
		Output: out,
		// ---
		MaxDepth: 5,
		ShowFlag: Ffunc | Ffname | Fline,
		// MoreLenNL: 8,
		// ---
		IndentLen:  2,
		IndentChar: ' ',
		CallerSkip: skip,
		ColorTheme: defaultTheme,
	}
}

// SkipNilField setting.
func SkipNilField() OptionFunc {
	return func(opt *Options) {
		opt.SkipNilField = true
	}
}

// SkipPrivate field dump on struct.
func SkipPrivate() OptionFunc {
	return func(opt *Options) {
		opt.SkipPrivate = true
	}
}

// BytesAsString setting.
func BytesAsString() OptionFunc {
	return func(opt *Options) {
		opt.BytesAsString = true
	}
}

// WithCallerSkip on print caller position information.
func WithCallerSkip(skip int) OptionFunc {
	return func(opt *Options) {
		opt.CallerSkip = skip
	}
}

// WithoutPosition dont print call dump position information.
func WithoutPosition() OptionFunc {
	return func(opt *Options) {
		opt.ShowFlag = Fnopos
	}
}

// WithoutOutput setting.
func WithoutOutput(out io.Writer) OptionFunc {
	return func(opt *Options) {
		opt.Output = out
	}
}

// WithoutColor setting.
func WithoutColor() OptionFunc {
	return func(opt *Options) {
		opt.NoColor = true
	}
}

// WithoutType setting.
func WithoutType() OptionFunc {
	return func(opt *Options) {
		opt.NoType = true
	}
}
