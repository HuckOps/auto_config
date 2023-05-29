package loader

import (
	"auto_config/source/file"
	"testing"
)

func TestMemory_Load(t *testing.T) {
	source := file.NewSource(file.WithPath("./test/test.yaml"))
	loader := NewLoader()
	if err := loader.Load(source); err != nil {
		t.Error(err)
	}
	select {}

}
