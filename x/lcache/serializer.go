package lcache

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io"

	"github.com/gookit/goutil/comdef"
)

//
// ----- builtin serializers -----
//

type Serializer interface {
	comdef.Codec
	DecodeFrom(r io.Reader, dest any) error
	EncodeTo(w io.Writer, src any) error
}

var serializers = map[string]Serializer{
	"json": JSONSerializer{},
	"gob":  GOBSerializer{},
}

// SetSerializer Set up the serializer for the cache
func SetSerializer(name string, serializer Serializer) {
	if serializer != nil {
		serializers[name] = serializer
	}
}

// JSONSerializer builtin serializer: json, gob
type JSONSerializer struct{}

// Decode implements Serializer
func (j JSONSerializer) Decode(data []byte, dest any) error {
	return json.Unmarshal(data, dest)
}

// Encode implements Serializer
func (j JSONSerializer) Encode(data any) ([]byte, error) {
	return json.Marshal(data)
}

// DecodeFrom implements Serializer
func (j JSONSerializer) DecodeFrom(r io.Reader, dest any) error {
	return json.NewDecoder(r).Decode(dest)
}

// EncodeTo implements Serializer
func (j JSONSerializer) EncodeTo(w io.Writer, src any) error {
	return json.NewEncoder(w).Encode(src)
}

// GOBSerializer builtin serializer: json, gob
type GOBSerializer struct{}

// Decode implements Serializer
func (g GOBSerializer) Decode(data []byte, dest any) error {
	buf := bytes.NewBuffer(data)
	return gob.NewDecoder(buf).Decode(dest)
}

// Encode implements Serializer
func (g GOBSerializer) Encode(data any) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := gob.NewEncoder(buf).Encode(data)
	return buf.Bytes(), err
}

// DecodeFrom implements Serializer
func (g GOBSerializer) DecodeFrom(r io.Reader, dest any) error {
	return gob.NewDecoder(r).Decode(dest)
}

// EncodeTo implements Serializer
func (g GOBSerializer) EncodeTo(w io.Writer, src any) error {
	return gob.NewEncoder(w).Encode(src)
}
