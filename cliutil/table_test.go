package cliutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/cliutil"
)

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
