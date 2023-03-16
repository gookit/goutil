package httpctype_test

import (
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
}
