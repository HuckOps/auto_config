package json

import (
	"auto_config/encoder"
	jsonEncoding "encoding/json"
)

type json struct{}

func (j json) Encode(v interface{}) ([]byte, error) {
	r, err := jsonEncoding.Marshal(v)
	if err != nil {
		panic(err)
	}
	return r, err
}

func (j json) Decode(b []byte, v interface{}) error {
	return jsonEncoding.Unmarshal(b, v)
}

func (j json) String(v interface{}) (string, error) {
	yamlBytes, err := jsonEncoding.Marshal(&v)
	if err != nil {
		return "", err
	}
	return string(yamlBytes), nil
}

func (j json) Type() string {
	return "json"
}

func NewEncoder() encoder.Encoder {
	return json{}
}
