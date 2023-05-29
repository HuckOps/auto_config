package auto_config

import (
	"auto_config/source/file"
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	dest := map[string]interface{}{}
	//sources := []Option{
	//	WithSource(file.NewSource(file.WithPath("./test/test.yaml"))),
	//}
	//dests := []interface{}{
	//	dest,
	//}
	config, err := NewConfig(WithSource(file.NewSource(file.WithPath("./test/test.yaml"))), WithEntity(dest))
	if err != nil {
		t.Error(err)
		panic(err)
	}
	fmt.Println(config)
	select {}
}
