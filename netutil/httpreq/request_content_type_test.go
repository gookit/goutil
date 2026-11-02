package httpreq_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gookit/goutil/netutil/httpctype"
	"github.com/gookit/goutil/netutil/httpreq"
)

func TestRequestContentTypeOverridesDefaultEncoding(t *testing.T) {
	cases := []struct{ name, defaultType, optionType, want string }{
		{"JSON over form", httpctype.Form, httpctype.JSON, httpctype.JSON},
		{"form over JSON", httpctype.JSON, httpctype.Form, httpctype.Form},
		{"default JSON", httpctype.JSON, "", httpctype.JSON},
		{"option JSON", "", httpctype.JSON, httpctype.JSON},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				data, err := io.ReadAll(r.Body)
				if err != nil {
					t.Error(err)
				}
				if r.Header.Get("Content-Type") != tc.want {
					t.Errorf("content type %q want%q", r.Header.Get("Content-Type"), tc.want)
				}
				if tc.want == httpctype.JSON {
					var v map[string]string
					err = json.Unmarshal(data, &v)
					if err != nil || v["name"] != "Ada" {
						t.Errorf("invalid JSON body %q: %v", data, err)
					}
				} else {
					v, err := url.ParseQuery(string(data))
					if err != nil || v.Get("name") != "Ada" {
						t.Errorf("invalid form %q: %v", data, err)
					}
				}
			}))
			defer srv.Close()
			opt := httpreq.New().ContentType(tc.defaultType).WithData(map[string]string{"name": "Ada"}).WithContentType(tc.optionType)
			response, err := opt.Send(http.MethodPost, srv.URL)
			if err != nil {
				t.Fatal(err)
			}
			_ = response.Body.Close()
		})
	}
}
