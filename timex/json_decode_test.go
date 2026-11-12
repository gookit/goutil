package timex_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/timex"
)

func TestTimeUnmarshalJSONEscapes(t *testing.T) {
	for _, tc := range []struct {
		name, input, text string
	}{
		{"plain", `"2018-10-16 12:34:01"`, "2018-10-16 12:34:01"},
		{"space", `"2018-10-16\u002012:34:01"`, "2018-10-16 12:34:01"},
		{"digits", `"\u0032018-10-16 12:34:01"`, "2018-10-16 12:34:01"},
		{"separator", `"2018\u002f10\u002f16 12:34:01"`, "2018/10/16 12:34:01"},
		{"offset", `"2018-10-16T12:34:01\u002b02:00"`, "2018-10-16T12:34:01+02:00"},
		{"whitespace", " \n\t\"2018-10-16 12:34:01\" \n", "2018-10-16 12:34:01"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			want, err := strutil.ToTime(tc.text)
			assert.NoErr(t, err)
			got := timex.Time{Layout: "custom"}
			assert.NoErr(t, json.Unmarshal([]byte(tc.input), &got))
			assert.True(t, want.Equal(got.Time))
			assert.Eq(t, "custom", got.Layout)
		})
	}
}

func TestTimeUnmarshalJSONInvalid(t *testing.T) {
	for _, input := range []string{`1201810160`, `0`, `1`, `true`, `false`, `[]`, `{}`, `""`, `"invalid"`, `"2018-10-16\q12:34:01"`, ``, `"`, `"2018-10-16`} {
		t.Run(input, func(t *testing.T) {
			before := time.Date(2020, 3, 4, 5, 6, 7, 8, time.UTC)
			got := timex.Time{Time: before, Layout: "custom"}
			defer func() {
				if recovered := recover(); recovered != nil {
					t.Errorf("UnmarshalJSON panicked: %v", recovered)
				}
			}()
			assert.Err(t, got.UnmarshalJSON([]byte(input)))
			assert.Eq(t, before, got.Time)
			assert.Eq(t, "custom", got.Layout)
		})
	}
}

func TestTimeUnmarshalJSONNull(t *testing.T) {
	for _, input := range []string{"null", " \n null\t "} {
		t.Run(input, func(t *testing.T) {
			before := time.Date(2020, 3, 4, 5, 6, 7, 8, time.UTC)
			got := timex.Time{Time: before, Layout: "custom"}
			assert.NoErr(t, got.UnmarshalJSON([]byte(input)))
			assert.Eq(t, before, got.Time)
			assert.Eq(t, "custom", got.Layout)
		})
	}
}

func TestTimeUnmarshalJSONRequest(t *testing.T) {
	for _, input := range []string{`{"time":0}`, `{"time":1}`, `{"time":1201810160}`} {
		t.Run(input, func(t *testing.T) {
			var request struct {
				Time timex.Time `json:"time"`
			}
			assert.NotPanics(t, func() {
				assert.Err(t, json.Unmarshal([]byte(input), &request))
			})
		})
	}
}
