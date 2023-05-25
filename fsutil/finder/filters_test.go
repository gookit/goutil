package finder_test

import (
	"testing"

	"github.com/gookit/goutil/fsutil/finder"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestRegexFilterFunc(t *testing.T) {
	tests := []struct {
		filePath string
		pattern  string
		include  bool
		match    bool
	}{
		{"path/to/util.go", `\.go$`, true, true},
		{"path/to/util.go", `\.go$`, false, false},
		{"path/to/util.go", `\.py$`, true, false},
		{"path/to/util.go", `\.py$`, false, true},
	}

	ent := &testutil.DirEnt{}

	for _, tt := range tests {
		fn := finder.RegexFilterFunc(tt.pattern, tt.include)
		assert.Eq(t, tt.match, fn(finder.NewElem(tt.filePath, ent)))
	}
}
