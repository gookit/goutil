package json

import (
	"encoding/json"
	"github.com/json-iterator/go"
	"io/ioutil"
	"os"
)

var parser = jsoniter.ConfigCompatibleWithStandardLibrary

// WriteFile
func WriteFile(filePath string, data interface{}) error {
	jsonBytes, err := Encode(data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, jsonBytes, 0664)
}

// ReadFile
func ReadFile(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return Decode(content, v)
}

// JsonEncode encode data to json bytes. use it instead of json.Marshal
func Encode(v interface{}) ([]byte, error) {
	return parser.Marshal(v)
}

// JsonEncode decode json bytes to data. use it instead of json.Unmarshal
func Decode(json []byte, v interface{}) error {
	return parser.Unmarshal(json, v)
}

// Pretty get pretty JSON string
func Pretty(v interface{}) (string, error) {
	out, err := json.MarshalIndent(v, "", "    ")
	return string(out), err
}
