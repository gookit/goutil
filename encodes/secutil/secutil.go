// Package secutil provide some security utils
package secutil

import (
	"bytes"
	"errors"
)

// ErrUnPadding error
var ErrUnPadding = errors.New("un-padding decrypted data fail")

// PKCS5Padding input data
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padText...)
}

// PKCS5UnPadding input data
func PKCS5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	delLen := int(origData[length-1])

	if delLen > length {
		return nil, ErrUnPadding
	}

	// fix: 检查删除的填充是否是一样的字符，不一样说明 delLen 值是有问题的，无法解码
	if delLen > 1 && origData[length-1] != origData[length-2] {
		return nil, ErrUnPadding
	}

	return origData[:length-delLen], nil
}

// PKCS7Padding input data
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	return PKCS5Padding(ciphertext, blockSize)
}

// PKCS7UnPadding input data
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	return PKCS5UnPadding(origData)
}
