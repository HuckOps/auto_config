package encoder

import (
	"github.com/huckops/auto_config/encoder/yaml"
	"io/ioutil"
	"testing"
)

func TestYamlRead(t *testing.T) {
	content, err := ioutil.ReadFile("./test/test.yaml")
	if err != nil {
		t.Error(err)
	}
	var dest interface{}
	encoder := yaml.NewEncoder()
	if err := encoder.Decode(content, &dest); err != nil {
		t.Error(err)
	}
	if _, err := encoder.Encode(&dest); err != nil {
		t.Error(err)
	}
}

func TestJsonRead(t *testing.T) {
	content, err := ioutil.ReadFile("./test/test.json")
	if err != nil {
		t.Error(err)
	}
	var dest map[string]interface{}
	encoder := yaml.NewEncoder()
	if err := encoder.Decode(content, &dest); err != nil {
		t.Error(err)
	}
	if _, err := encoder.Encode(&dest); err != nil {
		t.Error(err)
	}

}
