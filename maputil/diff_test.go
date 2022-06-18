package maputil_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/gookit/goutil/maputil"
	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	var (
		err   error
		b     = true
		tests = []struct {
			name  string
			dataa interface{}
			datab interface{}
			want  map[string]interface{}
		}{
			{
				name: "test1",
				dataa: map[string]interface{}{
					"company_id":  "aabbcc",
					"create_time": "2022-06-01T02:14:58.72Z",
					"description": "",
					"null":        nil,
					"arr":         []string{"aaa", "bbb", "ccc"},
					"type":        0,
					"settings": struct {
						Name string                   `json:"name"`
						Op   int                      `json:"op"`
						Arr  []string                 `json:"arr"`
						X    []map[string]interface{} `json:"x"`
					}{
						"name",
						1,
						[]string{},
						[]map[string]interface{}{
							{"tag": "tag1"},
							{"filter": map[string]interface{}{
								"tag": "tag2",
								"struct": struct {
									Name string `json:"name"`
								}{
									"name",
								},
							}},
						},
					},
					"ptr": new(bool),
				},
				datab: map[string]interface{}{
					"create_time": "2022-06-01T02:14:58.72Z",
					"description": "xxx",
					"null":        nil,
					"arr":         []string{"aaa", "bbb", "ccc"},
					"type":        0,
					"settings": struct {
						Name string                   `json:"name"`
						Op   int                      `json:"op"`
						Arr  []string                 `json:"arr"`
						X    []map[string]interface{} `json:"x"`
					}{
						"name",
						1,
						[]string{},
						[]map[string]interface{}{
							{"tag": "changed tag"},
							{"filter": map[string]interface{}{
								"tag": "changed tag",
								"struct": struct {
									Name string `json:"name"`
								}{
									"changed name",
								},
							}},
						},
					},
					"ptr": &b,
				},
				want: map[string]interface{}{
					"company_id":  nil,
					"description": "xxx",
					"arr":         []interface{}{"aaa", "bbb", "ccc"},
					"settings": map[string]interface{}{
						"x": []interface{}{
							map[string]interface{}{"tag": "changed tag"},
							map[string]interface{}{"filter": map[string]interface{}{
								"tag": "changed tag",
								"struct": map[string]interface{}{
									"name": "changed name",
								},
							},
							},
						},
					},
					"ptr": b,
				},
			},
			{
				name: "test2",
				dataa: map[string]interface{}{
					"a": "a",
					"b": "b",
				},
				datab: map[string]interface{}{
					"b": "b",
					"c": "c",
				},
				want: map[string]interface{}{
					"a": nil,
					"c": "c",
				},
			},
		}
	)
	for _, test := range tests {
		var (
			m1 map[string]interface{}
			m2 map[string]interface{}
		)

		r1, _ := json.Marshal(test.dataa)
		r2, _ := json.Marshal(test.datab)

		if err = json.Unmarshal(r1, &m1); err != nil {
			t.Fatal(err)
		}
		if err = json.Unmarshal(r2, &m2); err != nil {
			t.Fatal(err)
		}
		diff, err := maputil.Diff(m1, m2)
		assert.NoError(t, err)

		assert.NotNil(t, diff)
		assert.Equal(t, test.want, diff)
		c, _ := json.Marshal(diff)
		fmt.Printf("%s", c)
	}
}
