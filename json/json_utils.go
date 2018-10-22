package json

import (
	"encoding/json"
	"github.com/json-iterator/go"
	"io/ioutil"
	"os"
)

// WriteJsonFile
func WriteJsonFile(filePath string, data interface{}) error {
	jsonBytes, err := JsonEncode(data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, jsonBytes, 0664)
}

// ReadJsonFile
func ReadJsonFile(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	return JsonDecode(content, v)
}

// JsonEncode encode data to json bytes. use it instead of json.Marshal
func JsonEncode(v interface{}) ([]byte, error) {
	var parser = jsoniter.ConfigCompatibleWithStandardLibrary

	return parser.Marshal(v)
}

// JsonEncode decode json bytes to data. use it instead of json.Unmarshal
func JsonDecode(json []byte, v interface{}) error {
	var parser = jsoniter.ConfigCompatibleWithStandardLibrary

	return parser.Unmarshal(json, v)
}

// PrettyJson get pretty Json string
func PrettyJson(v interface{}) (string, error) {
	out, err := json.MarshalIndent(v, "", "    ")

	return string(out), err
}
