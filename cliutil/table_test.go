package cliutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/x/assert"
)

func TestFormatTableDefaultAlign(t *testing.T) {
	cases := []struct {
		name         string
		defaultAlign int
		align        []int
		want         string
	}{
		{
			name: "left",
			want: "| A     | B   |\n| alpha | one |\n| x     | y   |\n",
		},
		{
			name:         "center",
			defaultAlign: 1,
			want:         "|   A   |  B  |\n| alpha | one |\n|   x   |  y  |\n",
		},
		{
			name:         "right",
			defaultAlign: 2,
			want:         "|     A |   B |\n| alpha | one |\n|     x |   y |\n",
		},
		{
			name:         "empty overrides",
			defaultAlign: 2,
			align:        []int{},
			want:         "|     A |   B |\n| alpha | one |\n|     x |   y |\n",
		},
		{
			name:         "partial overrides",
			defaultAlign: 2,
			align:        []int{0},
			want:         "| A     |   B |\n| alpha | one |\n| x     |   y |\n",
		},
		{
			name:         "all columns overridden",
			defaultAlign: 2,
			align:        []int{1, 0},
			want:         "|   A   | B   |\n| alpha | one |\n|   x   | y   |\n",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			style := &cliutil.TableStyle{
				BorderLeft:    "|",
				BorderRight:   "|",
				ColumnSep:     "|",
				ColumnPadding: 1,
				DefaultAlign:  tc.defaultAlign,
				Align:         tc.align,
			}
			got := cliutil.FormatTable([]string{"A", "B"}, [][]any{
				{"alpha", "one"},
				{"x", "y"},
			}, style)
			assert.Eq(t, tc.want, got)
		})
	}
}

func TestShowTable(t *testing.T) {
	cols := []string{"ID", "Name", "Age", "Active"}
	rows := [][]any{
		{1, "Alice", 28, true},
		{2, "Bob", 32, false},
		{3, "Charlie", 24, true},
	}

	fmt.Println("Base on text/tabwriter")
	cliutil.SimpleTable(cols, rows)

	fmt.Println("Use custom render:")
	cliutil.ShowTable(cols, rows)
	fmt.Println("Use MinimalStyle:")
	cliutil.ShowTable(cols, rows, cliutil.MinimalStyle)
}
