package yaml

import (
	"auto_config/encoder"
	yamlv2 "github.com/ghodss/yaml"
)

type yaml struct{}

func (y yaml) Encode(v interface{}) ([]byte, error) {
	return yamlv2.Marshal(v)
}

func (y yaml) Decode(b []byte, v interface{}) error {
	return yamlv2.Unmarshal(b, v)
}

func (y yaml) String(v interface{}) (string, error) {
	yamlBytes, err := yamlv2.Marshal(&v)
	if err != nil {
		return "", err
	}
	return string(yamlBytes), nil
}

func (y yaml) Type() string {
	return "yaml"
}

func NewEncoder() encoder.Encoder {
	return yaml{}
}
