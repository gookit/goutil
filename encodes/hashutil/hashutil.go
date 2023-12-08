// Package hashutil provide some util for quickly generate hash
package hashutil

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"hash/crc32"
	"hash/crc64"
	"strings"

	"github.com/gookit/goutil/encodes"
)

// hash algorithm names
const (
	AlgoCRC32  = "crc32"
	AlgoCRC64  = "crc64"
	AlgoMD5    = "md5"
	AlgoSHA1   = "sha1"
	AlgoSHA224 = "sha224"
	AlgoSHA256 = "sha256"
	AlgoSHA384 = "sha384"
	AlgoSHA512 = "sha512"
)

// MD5 generate md5 string by given src
func MD5(src any) string {
	return string(HexBytes(AlgoMD5, src))
}

// ShortMD5 Generate a 16-bit md5 bytes.
// remove first 8 and last 8 bytes from 32-bit md5.
func ShortMD5(src any) string {
	return string(HexBytes(AlgoMD5, src)[8:24])
}

// Hash generate hex hash string by given algorithm
func Hash(algo string, src any) string {
	return string(HexBytes(algo, src))
}

// HexBytes generate hex hash bytes by given algorithm
func HexBytes(algo string, src any) []byte {
	bs := HashSum(algo, src)
	dst := make([]byte, hex.EncodedLen(len(bs)))
	hex.Encode(dst, bs)
	return dst
}

// Hash32 generate hash by given algorithm, then use base32 encode.
func Hash32(algo string, src any) string {
	return string(Base32Bytes(algo, src))
}

// Base32Bytes generate base32 hash bytes by given algorithm
func Base32Bytes(algo string, src any) []byte {
	bs := HashSum(algo, src)
	dst := make([]byte, encodes.B32Hex.EncodedLen(len(bs)))
	encodes.B32Hex.Encode(dst, bs)
	return dst
}

// Hash64 generate hash by given algorithm, then use base64 encode.
func Hash64(algo string, src any) string {
	return string(Base64Bytes(algo, src))
}

// Base64Bytes generate base64 hash bytes by given algorithm
func Base64Bytes(algo string, src any) []byte {
	bs := HashSum(algo, src)
	dst := make([]byte, encodes.B64Std.EncodedLen(len(bs)))
	encodes.B64Std.Encode(dst, bs)
	return dst
}

// HashSum generate hash sum bytes by given algorithm
func HashSum(algo string, src any) []byte {
	hh := NewHash(algo)
	switch val := src.(type) {
	case []byte:
		hh.Write(val)
	case string:
		hh.Write([]byte(val))
	default:
		hh.Write([]byte(fmt.Sprint(src)))
	}
	return hh.Sum(nil)
}

// NewHash create hash.Hash instance
//
// algo: crc32, crc64, md5, sha1, sha224, sha256, sha384, sha512, sha512_224, sha512_256
func NewHash(algo string) hash.Hash {
	switch strings.ToLower(algo) {
	case AlgoCRC32:
		return crc32.NewIEEE()
	case AlgoCRC64:
		return crc64.New(crc64.MakeTable(crc64.ISO))
	case AlgoMD5:
		return md5.New()
	case AlgoSHA1:
		return sha1.New()
	case AlgoSHA224:
		return sha256.New224()
	case AlgoSHA256:
		return sha256.New()
	case AlgoSHA384:
		return sha512.New384()
	case AlgoSHA512:
		return sha512.New()
	case "sha512_224":
		return sha512.New512_224()
	case "sha512_256":
		return sha512.New512_256()
	default:
		panic("invalid hash algorithm:" + algo)
	}
}
