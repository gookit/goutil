package textscan_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/strutil/textscan"
	"github.com/gookit/goutil/testutil/assert"
)

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
		fmt.Println("====> Token.String(), kind:", t.Kind())
		fmt.Println(t.String())

		if t.Kind() == textscan.TokValue {
			v := t.(*textscan.ValueToken)
			data[v.Key()] = v.Value()
		}
	})

	assert.NoErr(t, err)
	assert.NotEmpty(t, data)
	assert.ContainsKeys(t, data, []string{"age", "name", "desc"})
}

func TestParser_ParseText(t *testing.T) {
	p := textscan.NewParser(func(t textscan.Token) {
		dump.P(t)
	})

	err := p.ParseText(`
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
	assert.NoErr(t, err)

}
