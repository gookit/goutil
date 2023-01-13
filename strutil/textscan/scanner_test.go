package textscan_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/strutil/textscan"
	"github.com/gookit/goutil/testutil/assert"
)

func ExampleNewScanner() {
	ts := textscan.NewScanner(`source code`)
	// add token matcher, can add your custom matcher
	ts.AddMatchers(
		&textscan.CommentsMatcher{
			InlineChars: []byte{'#'},
		},
		&textscan.KeyValueMatcher{
			MergeComments: true,
		},
	)

	// scan and parsing
	for ts.Scan() {
		tok := ts.Token()

		if !tok.IsValid() {
			continue
		}

		// Custom handle the parsed token
		if tok.Kind() == textscan.TokValue {
			vt := tok.(*textscan.ValueToken)
			fmt.Println(vt)
		}
	}

	if ts.Err() != nil {
		fmt.Println("ERROR:", ts.Err())
	}
}

func TestTextScanner_kvLine(t *testing.T) {
	ts := textscan.TextScanner{}
	ts.AddMatchers(
		&textscan.CommentsMatcher{},
		&textscan.KeyValueMatcher{},
	)

	ts.SetInput(`
# comments 1
name = inhere

// comments 2
age = 28

/*
multi line
comments 3
*/
desc = '''
a multi
line string
'''
`)

	data := make(map[string]string)
	err := ts.Each(func(t textscan.Token) {
		fmt.Println("====> Token kind:", t.Kind())
		fmt.Println(t.String())

		if t.Kind() == textscan.TokValue {
			v := t.(*textscan.ValueToken)
			data[v.Key()] = v.Value()
		}
	})

	fmt.Println("\n==== Collected data:")
	dump.P(data)
	assert.NoErr(t, err)
	assert.NotEmpty(t, data)
	assert.ContainsKeys(t, data, []string{"age", "name", "desc"})
}
