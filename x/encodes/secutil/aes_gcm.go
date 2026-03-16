package secutil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

var (
	// ErrInvalidAESKeySize means the AES key length is not 16, 24 or 32 bytes.
	ErrInvalidAESKeySize = errors.New("secutil: invalid AES key size")
	// ErrInvalidGCMNonceSize means the GCM nonce length does not match AEAD requirements.
	ErrInvalidGCMNonceSize = errors.New("secutil: invalid GCM nonce size")
)

// EncryptGCM encrypts plaintext with AES-GCM.
//
// The key length must be 16, 24 or 32 bytes. A random nonce is generated and
// returned alongside the ciphertext.
func EncryptGCM(key, plaintext []byte) (ciphertext, nonce []byte, err error) {
	aead, err := newGCM(key)
	if err != nil {
		return nil, nil, err
	}

	nonce = make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

	ciphertext = aead.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nonce, nil
}

// DecryptGCM decrypts ciphertext with AES-GCM.
func DecryptGCM(key, nonce, ciphertext []byte) ([]byte, error) {
	aead, err := newGCM(key)
	if err != nil {
		return nil, err
	}
	if len(nonce) != aead.NonceSize() {
		return nil, ErrInvalidGCMNonceSize
	}

	return aead.Open(nil, nonce, ciphertext, nil)
}

func newGCM(key []byte) (cipher.AEAD, error) {
	if err := validateAESKeySize(len(key)); err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}

func validateAESKeySize(n int) error {
	switch n {
	case 16, 24, 32:
		return nil
	default:
		return ErrInvalidAESKeySize
	}
}
