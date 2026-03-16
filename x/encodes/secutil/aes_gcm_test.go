package secutil_test

import (
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/x/encodes/secutil"
)

func TestEncryptGCM(t *testing.T) {
	t.Run("round trip for valid key sizes", func(t *testing.T) {
		cases := []struct {
			name string
			key  []byte
		}{
			{name: "aes-128", key: []byte("1234567890abcdef")},
			{name: "aes-192", key: []byte("1234567890abcdefghijklmn")},
			{name: "aes-256", key: []byte("1234567890abcdefghijklmnopqrstuv")},
		}

		plaintext := []byte("hello from secutil aes gcm")
		for _, tc := range cases {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				ciphertext, nonce, err := secutil.EncryptGCM(tc.key, plaintext)
				assert.NoErr(t, err)
				assert.NotEmpty(t, ciphertext)
				assert.Eq(t, 12, len(nonce))

				got, err := secutil.DecryptGCM(tc.key, nonce, ciphertext)
				assert.NoErr(t, err)
				assert.Eq(t, string(plaintext), string(got))
			})
		}
	})

	t.Run("allow empty plaintext", func(t *testing.T) {
		key := []byte("1234567890abcdef")

		ciphertext, nonce, err := secutil.EncryptGCM(key, nil)
		assert.NoErr(t, err)
		assert.NotEmpty(t, ciphertext)
		assert.Eq(t, 12, len(nonce))

		got, err := secutil.DecryptGCM(key, nonce, ciphertext)
		assert.NoErr(t, err)
		assert.Empty(t, got)
	})

	t.Run("reject invalid key size", func(t *testing.T) {
		_, _, err := secutil.EncryptGCM([]byte("short-key"), []byte("hello"))
		assert.ErrIs(t, err, secutil.ErrInvalidAESKeySize)
	})
}

func TestDecryptGCM(t *testing.T) {
	key := []byte("1234567890abcdefghijklmnopqrstuv")
	plaintext := []byte("hello from secutil aes gcm")

	ciphertext, nonce, err := secutil.EncryptGCM(key, plaintext)
	assert.NoErr(t, err)

	t.Run("reject wrong nonce size", func(t *testing.T) {
		_, err := secutil.DecryptGCM(key, nonce[:len(nonce)-1], ciphertext)
		assert.ErrIs(t, err, secutil.ErrInvalidGCMNonceSize)
	})

	t.Run("reject wrong key", func(t *testing.T) {
		wrongKey := []byte("abcdef1234567890abcdefghijklmnop")
		_, err := secutil.DecryptGCM(wrongKey, nonce, ciphertext)
		assert.Err(t, err)
	})

	t.Run("reject invalid key size", func(t *testing.T) {
		_, err := secutil.DecryptGCM([]byte("short-key"), nonce, ciphertext)
		assert.ErrIs(t, err, secutil.ErrInvalidAESKeySize)
	})
}
