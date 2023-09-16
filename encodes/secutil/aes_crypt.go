package secutil

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/gookit/goutil/byteutil"
)

// aes methods
const (
	CryptAes128CBC = "aes-128-cbc"
	CryptAes256CBC = "aes-256-cbc"
)

// padding type for encoding contents.
const (
	PadTypeNone uint8 = iota
	PadTypeZeros
	PadTypePKCS5
	PadTypePKCS7
)

// type BlockModeFunc func(b cipher.Block, iv []byte, isEnc bool) cipher.BlockMode

// CryptConfig struct
type CryptConfig struct {
	Key string
	// IV string, length must be equals to cipher.Block.BlockSize()
	IV string
	// Method name. eg: CryptAes256CBC
	Method string
	// PadType for padding string.
	PadType uint8
	Encoder byteutil.BytesEncoder

	// BlockModeFn for create encrypter or decrypter blockMode
	// BlockModeFn BlockModeFunc
}

// AesCrypt struct
type AesCrypt struct {
	CryptConfig

	// iv length must is 16 = aes.BlockSize
	iv  []byte
	key []byte

	init   bool
	keyLen int
	// padding string

	encryptType string
}

// NewAesCrypt instance
func NewAesCrypt() *AesCrypt {
	return &AesCrypt{
		CryptConfig: CryptConfig{
			Method:  CryptAes256CBC,
			PadType: PadTypePKCS5,
			Encoder: byteutil.B64Encoder,
		},
	}
}

// Config crypt instance
func (p *AesCrypt) Config(fn func(c *CryptConfig)) *AesCrypt {
	fn(&p.CryptConfig)
	return p
}

// Init crypt instance
func (p *AesCrypt) Init() error {
	if p.init {
		return nil
	}

	p.init = true
	switch p.Method {
	case CryptAes128CBC:
		p.keyLen = 16
		p.encryptType = "cbc"
	case CryptAes256CBC:
		p.keyLen = 32
		p.encryptType = "cbc"
	}

	// iv length must is 16 = aes.BlockSize
	p.iv = []byte(p.IV)

	// padding ASCII 0(NUL)
	p.key = make([]byte, p.keyLen)
	copy(p.key, p.Key)

	return nil
}

// Encrypt input source bytes. return error on fail.
func (p *AesCrypt) Encrypt(src []byte) ([]byte, error) {
	if err := p.Init(); err != nil {
		return nil, err
	}

	// TODO 优化: 在 Init() 创建好 block 和 blockMode
	block, err := aes.NewCipher(p.key)
	if err != nil {
		return nil, err
	}

	padSrc := PKCS5Padding(src, block.BlockSize())

	encrypted := make([]byte, len(padSrc))
	blockMode := cipher.NewCBCEncrypter(block, p.iv)
	blockMode.CryptBlocks(encrypted, padSrc)

	if p.Encoder == nil {
		return encrypted, nil
	}
	return p.Encoder.Encode(encrypted), nil
}

// EncryptString to encoded string. return error on fail.
func (p *AesCrypt) EncryptString(src string) (string, error) {
	bs, err := p.Encrypt([]byte(src))
	return string(bs), err
}

// Decrypt an encrypt to source data
func (p *AesCrypt) Decrypt(enc []byte) ([]byte, error) {
	if err := p.Init(); err != nil {
		return nil, err
	}

	if p.Encoder != nil {
		var err error
		enc, err = p.Encoder.Decode(enc)
		if err != nil {
			return nil, err
		}
	}

	// TODO 优化: 在 Init() 创建好 block 和 blockMode
	block, err := aes.NewCipher(p.key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, p.iv)

	srcBytes := make([]byte, len(enc))
	blockMode.CryptBlocks(srcBytes, enc)

	return PKCS5UnPadding(srcBytes)
}

// DecryptString to source string. return error on fail.
func (p *AesCrypt) DecryptString(enc string) (string, error) {
	bs, err := p.Decrypt([]byte(enc))
	return string(bs), err
}
