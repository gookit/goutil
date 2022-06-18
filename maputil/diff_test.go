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
		m1    map[string]interface{}
		m2    map[string]interface{}
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
				},
				datab: map[string]interface{}{
					"create_time": "2022-06-01T02:14:58.72Z",
					"description": "xxx",
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
				},
			},
		}
	)
	for _, test := range tests {
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
