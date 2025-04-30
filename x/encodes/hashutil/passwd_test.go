package hashutil_test

import (
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/x/encodes/hashutil"
)

func TestHashPasswd(t *testing.T) {
	tests := []struct {
		password string
		key      string
		expected string
	}{
		{"123456", "secret", "4a83854cf6f0112b4295bddd535a9b3fbe54a3f90e853b59d42e4bed553c55a4"},
		{"password", "key", "4d42fb9ffc8d7d0a245429438b4bc73db1007a167026a0a0c6a74fa58e8e86ca"},
	}

	for _, test := range tests {
		hashed := hashutil.HashPasswd(test.password, test.key)
		assert.Eq(t, test.expected, hashed)
	}
}

func TestVerifyPasswd(t *testing.T) {
	tests := []struct {
		wantPwd  string
		password string
		key      string
		expected bool
	}{
		{"4a83854cf6f0112b4295bddd535a9b3fbe54a3f90e853b59d42e4bed553c55a4", "123456", "secret", true},
		{"invalidHash", "123456", "secret", false},
		{"4d42fb9ffc8d7d0a245429438b4bc73db1007a167026a0a0c6a74fa58e8e86ca", "wrongPassword", "secret", false},
	}

	for _, test := range tests {
		result := hashutil.VerifyPasswd(test.wantPwd, test.password, test.key)
		assert.Eq(t, test.expected, result)
	}
}
