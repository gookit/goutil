package httpctype_test

import (
	"strings"
	"testing"

	"github.com/gookit/goutil/netutil/httpctype"
	"github.com/gookit/goutil/testutil/assert"
)

func TestToKind(t *testing.T) {
	tests := []struct {
		cType       string
		defaultType string
		want        string
	}{
		{"", "abc", "abc"},
		{"not-match", "", ""},
		{"not-match", "abc", "abc"},
		{httpctype.JSON, "", httpctype.KindJSON},
		{httpctype.Form, "", httpctype.KindForm},
		{httpctype.FormData, "", httpctype.KindFormData},
		{httpctype.XML, "", httpctype.KindXML},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, httpctype.ToKind(tt.cType, tt.defaultType))
	}

	assert.Eq(t, "", httpctype.ToKindWithFunc("not-match", nil))
	assert.Eq(t, httpctype.KindYAML, httpctype.ToKindWithFunc(httpctype.YAML, func(cType string) string {
		if strings.Contains(cType, "/x-yaml") {
			return httpctype.KindYAML
		}
		return ""
	}))
}
