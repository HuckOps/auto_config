package yaml

import (
	yamlv2 "github.com/ghodss/yaml"
	"github.com/huckops/auto_config/encoder"
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
